# MyPapers 运维日志  
> 请在此处记录你在 `mypapers` 机子上的关键操作，模板可参照前文，创建新日志的时候，请直接在本文章开头书写。  

## 2025年2月27日16:04:52
### 操作人
叶文博
### 操作说明  
今天试图修复 mypapers 网站访问 404 的问题，试图跑前端进程时，系统 ram 吃满，导致 ssh 都连不上了。  
所以我针对 mypapers 这台机子写了一个 ssh 脚本，确保 ssh 在极端负载的情况下仍可用。  
### 操作步骤  
#### 1. 创建 SSH 资源保障脚本
我命名为 `ssh_resource_guard.sh`
``` bash
vim /usr/local/bin/ssh_resource_guard.sh
# 将脚本内容粘贴到文件中
# 按ESC键，然后输入:wq保存并退出
```
脚本内容：
``` bash
#!/bin/bash
# 脚本名: ssh_resource_guard.sh
# 功能: 为SSH服务保留网络带宽、CPU和内存资源
# 使用方法: 以root用户运行此脚本

# 检查root权限
if [ "$(id -u)" -ne 0 ]; then
    echo "此脚本需要root权限运行"
    exit 1
fi

# 确认网卡名称
NIC="eth0"
if ! ip link show $NIC &>/dev/null; then
    echo "网卡 $NIC 不存在，脚本将退出"
    exit 1
fi

echo "=== 开始配置资源保障 ==="

# ===== 网络带宽保障 =====
echo "正在配置网络优先级..."

# 清除现有的tc配置
tc qdisc del dev $NIC root 2>/dev/null

# 创建根队列
tc qdisc add dev $NIC root handle 1: htb default 9999

# 为SSH预留1Mbps带宽（优先级最高）
tc class add dev $NIC parent 1: classid 1:1 htb rate 1mbps ceil 1mbps prio 0

# 标记SSH流量（端口22）
iptables -F OUTPUT -t mangle 2>/dev/null
iptables -A OUTPUT -t mangle -p tcp --sport 22 -j MARK --set-mark 0x1
tc filter add dev $NIC parent 1: protocol ip handle 0x1 fw flowid 1:1

# 其他流量分配到低优先级类（总带宽3Mbps，预留1Mbps给SSH）
tc class add dev $NIC parent 1: classid 1:9999 htb rate 2mbps ceil 2mbps prio 7

# ===== CPU优先级保障 =====
echo "正在配置CPU优先级..."

# 检查cgroups v2是否可用
if [ -d "/sys/fs/cgroup" ] && [ ! -d "/sys/fs/cgroup/memory" ]; then
    echo "检测到使用cgroups v2"
    
    # 创建SSH专用cgroup (如果不存在)
    if [ ! -d "/sys/fs/cgroup/ssh_reserved" ]; then
        mkdir -p /sys/fs/cgroup/ssh_reserved
    fi
    
    # 配置CPU权重
    echo "100" > /sys/fs/cgroup/ssh_reserved/cpu.weight 2>/dev/null
    
    # 将sshd服务移入cgroup
    echo "正在配置systemd服务属性..."
    systemctl set-property sshd.service CPUWeight=100
    
    # 确保sshd进程在cgroup中
    for pid in $(pgrep -f "/usr/sbin/sshd"); do
        echo $pid > /sys/fs/cgroup/ssh_reserved/cgroup.procs 2>/dev/null
    done
else
    echo "检测到使用cgroups v1或混合模式"
    
    # 使用旧版cgroups
    if [ -d "/sys/fs/cgroup/cpu" ]; then
        # 创建SSH专用cgroup
        mkdir -p /sys/fs/cgroup/cpu/ssh_reserved
        echo 1024 > /sys/fs/cgroup/cpu/ssh_reserved/cpu.shares
        
        # 将sshd进程移入cgroup
        for pid in $(pgrep -f "/usr/sbin/sshd"); do
            echo $pid > /sys/fs/cgroup/cpu/ssh_reserved/cgroup.procs 2>/dev/null
        done
    fi
fi

# 设置sshd进程为实时优先级
for pid in $(pgrep -f "/usr/sbin/sshd"); do
    chrt -rr 50 -p $pid 2>/dev/null
done

# ===== 内存保护 =====
echo "正在配置内存保护..."

# 降低SSH进程的OOM score
for pid in $(pgrep -f "/usr/sbin/sshd"); do
    echo -1000 > /proc/$pid/oom_score_adj 2>/dev/null
done

# 检查cgroups版本并配置内存限制
if [ -d "/sys/fs/cgroup/memory" ]; then
    # cgroups v1
    echo "配置cgroups v1内存保护..."
    
    # 创建低优先级组
    mkdir -p /sys/fs/cgroup/memory/other_processes
    
    # 为其他进程设置内存限制 (例如限制为总内存的80%)
    MEM_TOTAL=$(free -b | grep "Mem:" | awk '{print $2}')
    MEM_LIMIT=$(echo "$MEM_TOTAL * 0.8" | bc | cut -d. -f1)
    echo $MEM_LIMIT > /sys/fs/cgroup/memory/other_processes/memory.limit_in_bytes
    
    # 将非关键进程移入该组
    for pid in $(ps -eo pid --no-headers); do
        if ! pgrep -f "sshd|systemd|init|journald|udevd|bash" | grep -q "$pid"; then
            echo $pid > /sys/fs/cgroup/memory/other_processes/cgroup.procs 2>/dev/null
        fi
    done
elif [ -d "/sys/fs/cgroup" ]; then
    # cgroups v2
    echo "配置cgroups v2内存保护..."
    
    # 创建低优先级组
    mkdir -p /sys/fs/cgroup/other_processes
    
    # 启用内存控制器
    echo "+memory" > /sys/fs/cgroup/cgroup.subtree_control 2>/dev/null
    
    # 设置内存限制
    MEM_TOTAL=$(free -b | grep "Mem:" | awk '{print $2}')
    MEM_LIMIT=$(echo "$MEM_TOTAL * 0.8" | bc | cut -d. -f1)
    echo $MEM_LIMIT > /sys/fs/cgroup/other_processes/memory.max 2>/dev/null
    
    # 将非关键进程移入该组
    for pid in $(ps -eo pid --no-headers); do
        if ! pgrep -f "sshd|systemd|init|journald|udevd|bash" | grep -q "$pid"; then
            echo $pid > /sys/fs/cgroup/other_processes/cgroup.procs 2>/dev/null
        fi
    done
fi

echo "=== 资源保障配置完成 ==="

# 打印当前状态
echo ""
echo "当前配置状态:"
echo "网络带宽配置:"
tc -s qdisc show dev $NIC
tc -s class show dev $NIC

echo ""
echo "SSH进程的OOM分数:"
for pid in $(pgrep -f "/usr/sbin/sshd"); do
    echo "PID $pid: $(cat /proc/$pid/oom_score_adj)"
done

echo ""
echo "SSH进程的调度策略:"
for pid in $(pgrep -f "/usr/sbin/sshd"); do
    chrt -p $pid
done

echo ""
echo "已完成所有配置。"
```
#### 2. 设置执行权限：
``` bash
chmod +x /usr/local/bin/ssh_resource_guard.sh
```

#### 3. 创建systemd服务，使脚本在系统启动时自动运行：
``` bash
cat > /etc/systemd/system/ssh-resource-guard.service << 'EOF'
[Unit]
Description=SSH Resource Guard Service
After=network.target sshd.service

[Service]
Type=oneshot
ExecStart=/usr/local/bin/ssh_resource_guard.sh
RemainAfterExit=yes
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
EOF
```
#### 4. 启用并启动服务：
``` bash
systemctl daemon-reload
systemctl enable ssh-resource-guard.service
systemctl start ssh-resource-guard.service
```
#### 5. 检查systemd服务状态：
``` bash
systemctl status ssh-resource-guard.service
```
#### 如果需要检查运行日志：
``` bash
journalctl -u ssh-resource-guard.service
```
执行成功后，可见  
![](https://pic.kiss1314.top/d/LeavesResource/webImage/20250227155732.png)

#### 如果需要手动启动脚本，可以直接执行：
``` bash
bash /usr/local/bin/ssh_resource_guard.sh
```
