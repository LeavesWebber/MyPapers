
# MyPapers 项目运维交接文档

## 任务概述
- MyPapers 当前部署在 UCloud 服务器上，计划排除页面无法访问的问题后迁移至阿里云服务器。
---

## 当前部署环境（与myarxiv共用同一台服务器）
- **服务器**：UCloud  
- **操作系统**：Ubuntu  
- **平台账号密码**：  
  - 2904976636@qq.com@aliyun.com
  - xmut_block_5922
- **服务器账号密码**：  
  - ubuntu 
  - xmutBC2024

## 仍需部署环境（供mypapers单独使用）
- **服务器**：阿里云  
- **操作系统**：Ubuntu  
- **平台账号密码**：  
  - hi30179517@aliyun.com
  - XMUT_rchen_5924
- **服务器账号密码**：  
  - root
  - XMUT_block_5924
---

## 已解决问题

### 1. Nginx 无法启动问题
- Nginx 服务无法启动，日志显示80端口冲突。
- **当前状态**：  
 ```bash
     ubuntu@10-35-54-29:~$ sudo systemctl status nginx
     ● nginx.service - A high performance web server and a reverse proxy server
          Loaded: loaded (/lib/systemd/system/nginx.service; disabled; vendor preset: enabled)
          Active: active (running) since Mon 2025-03-17 21:01:25 +08; 1 day 22h ago
            Docs: man:nginx(8)
        Main PID: 4019974 (nginx)
           Tasks: 5 (limit: 9473)
          Memory: 5.8M
             CPU: 56ms
          CGroup: /system.slice/nginx.service
                  ├─4019974 "nginx: master process /usr/sbin/nginx -g daemon on; master_process on;"
                  ├─4019975 "nginx: worker process"
                  ├─4019976 "nginx: worker process"
                  ├─4019977 "nginx: worker process"
                  └─4019978 "nginx: worker process"
```  
- Nginx 服务已可以正常运行。


### 2. 更新仓库后页面无法访问问题(已排查部分问题)
- **问题描述**：更新代码仓库后，访问 `http://mypapers.io` 返回错误：  
  ```bash
    ubuntu@10-35-54-29:~$ curl -I -L http://mypapers.io
    curl: (7) Failed to connect to mypapers.io port 80 after 533 ms: No route to host
  ```  

- **已排查内容**：  

  1. **检查端口占用**：  
    ```bash
    ubuntu@10-35-54-29:~$ sudo netstat -tuln | grep 80
    tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN     
    tcp        0      0 0.0.0.0:6380            0.0.0.0:*               LISTEN     
    tcp        0      0 0.0.0.0:8001            0.0.0.0:*               LISTEN     
    tcp        0      0 0.0.0.0:8090            0.0.0.0:*               LISTEN     
    ```  
     - 端口 80 已被 Nginx 占用，排除端口冲突问题。  

  3. **检查防火墙状态**：  
    ```bash
        ubuntu@10-35-54-29:~$ sudo ufw status
        Status: active

        To                         Action      From
        --                         ------      ----
        80/tcp                     ALLOW       Anywhere                  
        80/tcp (v6)                ALLOW       Anywhere (v6)             
    ```  

    - 防火墙已开放端口 80、443，排除防火墙限制的问题。  

  4. **检查 DNS 解析**：  
    ```bash
        ubuntu@10-35-54-29:~$ dig mypapers.io
        ;; ANSWER SECTION:
        mypapers.io.		599	IN	A	107.155.56.166
    ```  
    - DNS 解析正常，域名 `mypapers.io` 正确解析到 IP `107.155.56.166`。  

  5. **检查网络连通性**：  
    - **检查外部网络连通性**：  
    ```bash
       ubuntu@10-35-54-29:~$ ping google.com
       PING google.com (142.251.175.138) 56(84) bytes of data.
       64 bytes from sh-in-f138.1e100.net (142.251.175.138): icmp_seq=1 ttl=100 time=2.13 ms
    ```  
    - 服务器外部网络连通性正常。  

    - **检查内部网络连通性**：  
    ```bash
       ubuntu@10-35-54-29:~$ ping 10.35.0.0
       PING 10.35.0.0 (10.35.0.0) 56(84) bytes of data.
       From 10.35.54.29 icmp_seq=1 Destination Host Unreachable
       From 10.35.54.29 icmp_seq=2 Destination Host Unreachable
       From 10.35.54.29 icmp_seq=5 Destination Host Unreachable
       From 10.35.54.29 icmp_seq=8 Destination Host Unreachable
       From 10.35.54.29 icmp_seq=9 Destination Host Unreachable
       From 10.35.54.29 icmp_seq=10 Destination Host Unreachable
       ^C
       --- 10.35.0.0 ping statistics ---
       12 packets transmitted, 0 received, +6 errors, 100% packet loss, time 11212ms
       pipe 4
    ```  
    - 服务器内部网络（`10.35.0.0/16`）存在连通性问题。  

  6. **检查 Nginx 日志**：  
    ```bash
    ubuntu@10-35-54-29:~$ sudo tail -n 50 /var/log/nginx/error.log
    ```  

    - 日志未显示明显错误信息。  
   

---

## 待解决问题
### 1. 使项目可以通过mypapers.io以及www.mypapers.io正常访问
- **任务描述**：继续排查页面不能访问的原因
### 2. UCloud 服务器内容同步至阿里云新服务器
- **任务描述**：将 UCloud 服务器上的 MyPapers 项目完整迁移至阿里云新服务器。 

---

## 其他事项
1. **日志查看**：  
   - Nginx 日志：`/var/log/nginx/error.log`  
2. **配置文件**：
   - mypapers.io配置文件路径：`/etc/nginx/sites-enabled/mypapers.io`
3. **常用命令**：  
   - 重启 Nginx：`sudo systemctl restart nginx`  
   - 查看Nginx状态：`sudo systemctl status nginx`  
   - 项目目录：`/opt/MyPapers/` 
4. **关于 HTTPS 的部署**
- **SSL 证书配置**:
   - 已使用 Certbot 和 Let's Encrypt 获取 SSL 证书，并完成基础配置。  
   - 证书路径：`/etc/letsencrypt/live/mypapers.io`  
   - 证书有效期：90 天（Let's Encrypt 默认有效期）。
- **当前问题**:
   - 使用 HTTPS 访问 `https://mypapers.io` 时，页面显示空白，控制台报错显示 WebSocket（WS）请求头错误。  
   - 已排查内容：SSL 证书配置已生效，HTTPS连接正常。  
- **注意事项**:
   - Let's Encrypt 证书未设置自动续期，需手动续期，请提前续期以避免服务中断。
   ```bash
      sudo certbot renew --dry-run
   ``` 

5. **联系方式**：  
   - QQ:280289810  

---



**交接人**：王煜晶 
**日期**：2025年3月16日  


