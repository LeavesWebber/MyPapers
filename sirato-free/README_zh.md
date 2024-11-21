# Sirato

区块链浏览器，支持 Besu、Quorum、VMware Blockchain 和与以太坊兼容的区块链

![Sirato 仪表板](https://raw.githubusercontent.com/web3labs/sirato-free/master/_images/sirato-dashboard.png "Sirato 仪表板")

## 简介

Sirato 是一个用于以太坊兼容区块链的数据和分析平台。

它提供了丰富的 API 和易用的界面，用于展示区块链上各种资产的信息，例如代币和智能合约。

此仓库中提供了一个免费开发者版，同时我们还提供了托管计划，详细信息如下。

## 免费计划

此版本的 Sirato 是一个免费版本，设计用于查看公共和私有的以太坊网络。支持以下网络：  
[Quorum](https://github.com/ConsenSys/quorum)、[Hyperledger Besu](https://besu.hyperledger.org/en/stable/)、[VMware Blockchain for Ethereum](https://www.vmware.com/products/blockchain.html) 和 [Ethereum](https://github.com/ethereum/go-ethereum)。

![Sirato 免费版截图](https://raw.githubusercontent.com/web3labs/sirato-free/master/_images/sirato-free.png "Sirato 免费版")

## 托管计划

Web3 Labs 提供托管计划，包含以下额外功能：

- 定制品牌和自定义域名的托管
- 代币的专用视图
- 智能合约管理和源码上传
- OpenAPI 后端
- 与 Tableau、Microsoft PowerBI 和 Qlik 等商业智能工具的集成
- 生产级 SLA
- 大量交易量支持（超过 1 亿笔交易）

![Sirato 托管版截图](https://raw.githubusercontent.com/web3labs/sirato-free/master/_images/sirato-hosted.png "Sirato 托管版示例：带验证源代码的 Palm")

托管计划的优势是用户只需提供兼容的 web3 客户端端点，其他工作由我们完成。

可在[此处](https://www.web3labs.com/blockchain-explorer-sirato-plans)查看托管计划的更多信息，或直接通过 [hi@web3labs.com](mailto:hi@web3labs.com?subject=Sirato%20hosted%20plans) 联系 Web3 Labs。

### 部署说明

此仓库包含用于使用 Docker Compose 或 Kubernetes 运行免费版本的配置。

根据需要选择适当的指南，在本地运行 Sirato 并连接到 Ethereum、Quorum、Hyperledger Besu 或 VMware 区块链兼容网络。

- [Docker Compose 部署](docker-compose/README.md)
- [Kubernetes 部署](k8s/README.md)

### 系统要求

推荐最低系统要求：

| 组件  | 描述          |
| --- | ----------- |
| CPU | 1 核         |
| 内存  | 8 GB        |
| 硬盘  | 根据区块链大小动态调整 |

## 许可证

Sirato 仅供非商业用途和评估目的免费使用，详情请参阅 [LICENSE](LICENSE)。关于商业用途，请通过 `hi <at> web3labs.com` 联系我们，或通过[此处](https://pages.web3labs.com/sirato-enterprise)提交咨询。
