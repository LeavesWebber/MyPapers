## 使用方法

克隆仓库，导航到克隆的目录并运行实例：

```bash
cd docker-compose
docker-compose pull
NODE_ENDPOINT=http://<node_endpoint> docker-compose up
```

浏览器默认运行在 `80` 端口，如需自定义端口，例如 26000，请设置 `PORT` 环境变量。浏览器将可通过 http://localhost:26000 访问。

```bash
NODE_ENDPOINT=http://localhost:8585 PORT=26000 docker-compose up -d
```

请注意，如果将 `NODE_ENDPOINT` 设置为本地以太坊实例，则可能需要使用 Docker 桥接接口的 IP 地址。

在 Linux 上，桥接适配器的 IP 地址应为 172.16.239.1，具体见 docker-compose.yml。连接本地节点时，使用以下命令：

```bash
NODE_ENDPOINT=http://172.16.239.1:8545 docker-compose up
```

在 MacOS 和 Windows 上，由于 Docker 网络栈的限制，需要使用以下平台特定命令：

```bash
NODE_ENDPOINT=http://host.docker.internal:8545 docker-compose up
```

注意，无论哪种情况，本地的 geth 实例都必须使用 `--rpcaddr 0.0.0.0` 和 `--rpcvhosts="*"` 参数启动，否则 Sirato 将无法访问。

在 Windows 上，由于文件系统性能问题，Sirato 启动可能需要较长时间（有时长达 20 分钟）。

附加 `-d` 参数以在后台运行容器：

```bash
docker-compose up -d
```

您可以通过以下方式访问浏览器 UI：

- http://localhost/

停止容器使用以下命令：

```bash
docker-compose down
```

如需连接新网络，应删除与旧网络相关的存储卷：

```bash
docker-compose down -v
```

## Quorum 和 Hyperledger Besu

运行带有 Sirato 免费版的 Quorum 7 节点示例，请[参考此说明](examples/Quorum_Example.md)。

运行 Pantheon-quickstart 隐私网络示例，请[参考此说明](examples/Pantheon_Privacy_Example.md)。

## 限制

- 由于 Docker 的[限制](https://github.com/moby/moby/issues/1143)，同一时间只能运行一个浏览器实例。
