# MyPapers 开发手册

[English](README.md) | [中文简体](README_zh.md)  

> 这是一个基于 web3 的去中心化论文管理系统。  

> 这并**不是**一个最终的 README 文件，我仅在这里做基本简述，便于大家开发  

> 私有仓库，内含敏感凭证信息，**不要**操作公开此仓库  

---
> 新功能的后端文档请见 [已发表论文后端开发指引.md](PUBLISHED_PAPERS_BACKEND.md)

---
> 项目的运维日志请见 [DevOps.md](DevOps.md)

## 关于项目结构

> 新成员若需要 demo 项目，可参考 [项目 demo 视频.mp4](docs/新版Demo视频.mp4)  

### 前端：基于`VUE2`

- 位置  
`web` 文件夹： 基于 **vue2**  开发的**前端**项目文件，并使用`yarn`作为包管理器  

- 启动方法  
启动前先根据需要修改前端的配置，可见 [web/vue.config.js](web/vue.config.js)  
  
进入前端文件夹  

``` bash
cd web 
```

启动开发服务器  

```bash
yarn serve
```

### 后端：基于 `GO`

- 位置  
`server` 文件夹：基于 **go** 开发的**后端**项目文件  
- 启动方法  
**请先调整[后端配置文件](server/config-example.yaml)，尤其注意安装 mysql 环境并且先自己创建一个数据库。**  
  
在这之后，进入后端文件夹  

``` bash
cd server  
```

启动后端入口程序  

``` bash
go run main.go
```

### 区块链浏览器

> 已弃用。我们已经用自己公网上的区块链浏览器  

- `sirato-free` 文件夹： 一个区块链浏览器，用于查看区块链交易数据，可使用 docker 部署  

### 文档和小工具

- 位置：  
`docs` 文件夹。里面有一些项目开发会用到的信息和工具，有助于你了解和使用项目。  

除了文档外，其中还有一些测试用信息：  
比如当你在测试项目的时候可能需要填一些论文信息，需要一些论文文件，里面就有一个 `python` 写的小工具能生成一些随机的论文用于测试。  

目前可以用的工具：  

- [随机论文生成器 by.Leaves](docs\测试用附件\随机论文生成器\man.py)  
运行代码以后，它会往 [生成的随机论文](docs\测试用附件\随机论文生成器\生成的随机论文) 文件夹里生成 7 篇测试用的文件，文件内容是一串随机的 base58 编码的区块链地址。为啥要这样？因为项目有些地方不允许上传相同的论文。  

### 合约开发：基于 `Solidity`

合约地址在 [paperschain](https://paperschain.io/tokens) 可见（这是我们项目组自己部署的链端网站，由于部署在国外，很可能需要 [科学上网](https://clashverge.net/other_tools/) 才能访问）  

合约的地址需要正确填进 [web\src\constant.js](web\src\constant.js) 和 [server\config.yaml](server\config.yaml) 里项目的链端逻辑才能正确工作。  

- `ERC20.sol` 代币合约  

- `ERC721.sol` NFT合约  

- `nft-market.sol` 交易市场合约  

## 关于 Gitee 分支

> 请创建并使用自己的分支，请**不要**直接向 `main` 分支推送代码  

> 每周例会结束后，组长会合并代码到 `main` 分支，请大家及时从 `main` 分支 pull 代码到自己的分支

- `main` 分支： 用于发布合并（Merge）后的代码到生产环境  
- `raw` 分支： 用于存放学长最初给的代码，分支只读，请不要向此推送  
- `Leaves` 分支： [22 网络 叶文博](https://oksanye.com) 工作的分支  
- `MRdada` 分支： 22 计科 林振超 工作的分支  
- 待添加  

## 关于项目部署

> 目前，项目组在国内和国外各有一台机子部署了 Mypapers 项目。  
> 我倡议负责项目运维的同学积极维护 docs 文件夹内的 [运维文档](docs\DevOps.md)，其中应当包含机子的运维日志和基本信息。  
> 敏感的凭证不应该出现在项目里的文档里，或者以任何形式同步进 VCS。运维的同学可以把环境变量和其他凭证放进服务器的项目根目录。

### 域名

- 新加坡：  
[MyPapers.io](https://www.mypapers.io)  

- 国内京东云：  
[mypapers.com.cn](https://www.mypapers.com.cn)

### 项目地址

私有仓库，目前仅项目组员可见，项目代码请不要外传。

- [MyPapers Project（Gitee）](https://gitee.com/leaveswebber/MyPapers)  

- [MyPapers Project（Github）](https://github.com/LeavesWebber/MyPapers)  
  
### Host

- 新加坡 Ucloud： `107.155.56.166`  

- 国内京东云： `117.72.49.185`  

### 端口

> 下面这些信息仅供参考，具体可能需要问运维的同学设置的什么

- `22`: ssh && sftp
- `3306`: mysql 数据库  
- `4001`: IPFS P2P 节点通信  
- `5001`: IPFS 的 webui && API  
  
   （web UI 请使用 `http://ip:5001/webui/` 访问）  
- `8081`: IPFS 的网关  
- `8080`: MyPapers web 前端  
- `8887`: MyPapers 后端  

---

文档更新者： **叶文博**  
最后更新时间： **2025-05-11**  

---

## 附录目录

1. [附录1： 项目待办事项](#附录1-项目待办事项)
2. [附录2： 后端数据库结构](#附录2-后端数据库结构)
3. [附录3： 学长最初的开发环境](#附录3学长最初的开发环境)  

<br/>

### 附录1： 项目待办事项  

- [ ] 后端部署 Dify  
- [ ] 调整已出版论文生成 NFT 逻辑  
- [ ] 重复注册-无法注册但未提醒, 包含用户名，邮箱，手机号码和钱包地址重复注册无警告？  
- [ ] 已投稿论文上传，通过 ai 识别邮箱，一起识别其他信息  
- [ ] 注册送币机制修改  
- [ ] 现在论文都是明文上传到 IPFS ，现在需要想一个论文加密授权解决方案

### 附录2： 后端数据库结构

> 我把后端数据库结构做成了 `Markdown` 表格，作为大家执行 `sql` 查询时的参考。表格在 [server/model/tables](server/model/tables) 里定义

#### 1. authorities

| Column Name    | Data Type |
| -------------- | --------- |
| created_at     | datetime  |
| updated_at     | datetime  |
| deleted_at     | datetime  |
| authority_id   | bigint    |
| parent_id      | bigint    |
| authority_name | varchar   |

#### 2. authority_menus

| Column Name            | Data Type |
| ---------------------- | --------- |
| base_menu_id           | bigint    |
| authority_authority_id | bigint    |

#### 3. base_menus

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| created_at  | datetime  |
| updated_at  | datetime  |
| deleted_at  | datetime  |
| parent_id   | bigint    |
| sort        | bigint    |
| path        | varchar   |
| name        | varchar   |
| url         | varchar   |
| title       | varchar   |
| icon        | varchar   |

#### 4. committees

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| created_at  | datetime  |
| updated_at  | datetime  |
| deleted_at  | datetime  |
| creator_id  | bigint    |
| name        | varchar   |
| description | text      |

#### 5. conference_issues

| Column Name           | Data Type |
| --------------------- | --------- |
| id                    | bigint    |
| created_at            | datetime  |
| updated_at            | datetime  |
| deleted_at            | datetime  |
| conference_id         | bigint    |
| name                  | varchar   |
| submission_start_time | datetime  |
| submission_end_time   | datetime  |
| description           | varchar   |
| year                  | bigint    |
| volume                | bigint    |

#### 6. conferences

| Column Name  | Data Type |
| ------------ | --------- |
| id           | bigint    |
| created_at   | datetime  |
| updated_at   | datetime  |
| deleted_at   | datetime  |
| committee_id | bigint    |
| creator_id   | bigint    |
| name         | varchar   |
| description  | text      |
| location     | varchar   |
| category     | varchar   |
| start_time   | datetime  |
| end_time     | datetime  |

#### 7. journal_issues

| Column Name           | Data Type |
| --------------------- | --------- |
| id                    | bigint    |
| created_at            | datetime  |
| updated_at            | datetime  |
| deleted_at            | datetime  |
| journal_id            | bigint    |
| name                  | varchar   |
| submission_start_time | datetime  |
| submission_end_time   | datetime  |
| description           | varchar   |
| year                  | bigint    |
| volume                | bigint    |

#### 8. journals

| Column Name  | Data Type |
| ------------ | --------- |
| id           | bigint    |
| created_at   | datetime  |
| updated_at   | datetime  |
| deleted_at   | datetime  |
| committee_id | bigint    |
| creator_id   | bigint    |
| name         | varchar   |
| description  | text      |
| category     | varchar   |

#### 9. paper_viewers

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| created_at  | datetime  |
| updated_at  | datetime  |
| deleted_at  | datetime  |
| paper_id    | bigint    |
| viewer_id   | bigint    |

#### 10. papers

| Column Name             | Data Type |
| ----------------------- | --------- |
| id                      | bigint    |
| created_at              | datetime  |
| updated_at              | datetime  |
| deleted_at              | datetime  |
| conference_id           | bigint    |
| journal_id              | bigint    |
| version_id              | bigint    |
| download_price          | bigint    |
| copyright_trading_price | bigint    |
| token_id                | varchar   |
| title                   | varchar   |
| authors                 | varchar   |
| paper_type              | varchar   |
| abstract                | text      |
| key_words               | text      |
| subject_category        | varchar   |
| manuscript_id           | varchar   |
| informed_consent        | varchar   |
| animal_subjects         | varchar   |
| cor_author              | varchar   |
| manuscript_type         | varchar   |
| unique_contribution     | text      |
| hash                    | varchar   |
| block_address           | varchar   |
| paper_transaction_hash  | varchar   |
| filepath                | varchar   |
| cid                     | varchar   |
| status                  | varchar   |
| image_uri               | varchar   |
| image_url               | varchar   |
| image_cid               | varchar   |
| json_uri                | varchar   |
| transaction_hash        | varchar   |
| journal_name            | varchar   |
| volume_and_issue        | varchar   |
| publication_date        | varchar   |
| conference_name         | varchar   |
| conference_date         | varchar   |
| conference_location     | varchar   |
| pages                   | varchar   |
| issn                    | varchar   |
| paper_link              | varchar   |

#### 11. reviews

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| created_at  | datetime  |
| updated_at  | datetime  |
| deleted_at  | datetime  |
| reviewer_id | bigint    |
| paper_id    | bigint    |
| comment     | text      |
| status      | varchar   |
| old_version | tinyint   |

#### 12. user_authority

| Column Name            | Data Type |
| ---------------------- | --------- |
| user_id                | bigint    |
| authority_authority_id | bigint    |

#### 13. user_committee

| Column Name  | Data Type |
| ------------ | --------- |
| user_id      | bigint    |
| committee_id | bigint    |
| start_time   | datetime  |
| end_time     | datetime  |
| position     | varchar   |
| level        | varchar   |

#### 14. user_conference

| Column Name   | Data Type |
| ------------- | --------- |
| user_id       | bigint    |
| conference_id | bigint    |
| start_time    | datetime  |
| end_time      | datetime  |
| position      | varchar   |
| level         | varchar   |

#### 15. user_journal

| Column Name | Data Type |
| ----------- | --------- |
| user_id     | bigint    |
| journal_id  | bigint    |
| start_time  | datetime  |
| end_time    | datetime  |
| position    | varchar   |
| level       | varchar   |

#### 16. user_paper

| Column Name | Data Type |
| ----------- | --------- |
| user_id     | bigint    |
| paper_id    | bigint    |
| old_version | tinyint   |

#### 17. users

| Column Name         | Data Type |
| ------------------- | --------- |
| id                  | bigint    |
| created_at          | datetime  |
| updated_at          | datetime  |
| deleted_at          | datetime  |
| uuid                | bigint    |
| authority_id        | bigint    |
| sex                 | bigint    |
| username            | varchar   |
| password            | varchar   |
| first_name          | varchar   |
| middle_name         | varchar   |
| last_name           | varchar   |
| header_img          | varchar   |
| email               | varchar   |
| department          | varchar   |
| phone               | varchar   |
| address             | varchar   |
| education           | varchar   |
| title               | varchar   |
| research            | varchar   |
| block_chain_address | varchar   |
| affiliation         | varchar   |
| affiliation_type    | varchar   |

#### 18. mps_recharge_orders

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| user_id     | bigint    |
| order_no    | varchar   |
| amount      | float     |
| mps_amount  | float     |
| status      | int       |
| wx_trade_no | varchar   |
| ali_trade_no| varchar   |
| wallet_addr | varchar   |
| created_at  | datetime  |
| updated_at  | datetime  |

#### 19. mps_transactions

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| user_id     | bigint    |
| type        | int       |
| mps_amount  | float     |
| tx_hash     | varchar   |
| order_no    | varchar   |
| description | varchar   |
| created_at  | datetime  |

#### 20. mps_business_transfer_orders

| Column Name           | Data Type |
| --------------------- | --------- |
| id                    | bigint    |
| created_at            | datetime  |
| updated_at            | datetime  |
| deleted_at            | datetime  |
| user_id               | bigint    |
| order_no              | varchar   |
| identity              | varchar   |
| identity_type         | varchar   |
| mps_amount            | float     |
| fait_amount           | float     |
| wx_trade_no           | varchar   |
| ali_pay_fund_order_id | varchar   |
| trans_date            | varchar   |
| status                | varchar   |

---

### 附录3：学长最初的开发环境

操作系统：centos7.9.2009  
开发平台：Visual Studio Code 1.88.0；Remix   0.47.0；Goland  2023.1.2  
开发语言：Solidity ^0.8.0 ；javascript ；css ；html；go 1.20.2

---
