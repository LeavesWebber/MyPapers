# MyPapers 开发手册

[English](README.md) | [中文简体](README_zh.md)  

> 这是一个基于 web3 的去中心化论文管理系统。  

> 这并**不是**一个最终的 README 文件，我仅在这里做基本简述，便于大家开发  

> 私有仓库，内含敏感凭证信息，**不要**操作公开此仓库   

---
> 新功能的后端文档请见 [已发表论文后端开发指引.md](PUBLISHED_PAPERS_BACKEND.md)

---
> 项目的运维日志请见 [DevOps.md](DevOps.md)



## 关于URL

### 域名：

[MyPapers.io](https://mypapers.io)  

### 项目地址：  
私有仓库，目前仅项目组员可见
- [MyPapers Project（Gitee）](https://gitee.com/leaveswebber/MyPapers)

- [MyPapers Project（Github）](https://github.com/LeavesWebber/MyPapers)  
  
### Host

- 内网： `10.35.54.29`  

- 公网： `107.155.56.166`  

## 关于项目结构
### 前端：基于`VUE` + `JS` + `HTML` + `CSS`  
- 位置  
`web` 文件夹： 基于 **vue** + **js** + **html** + **css** 开发的**前端**项目文件，并使用`yarn`作为包管理器  
- 启动方法  

先进入前端文件夹  
``` bash
cd web 
```
启动开发服务器  
```bash
yarn serve
```
前端配置可见 [web/vue.config.js](web/vue.config.js)
### 后端：基于 `GO`  
- 位置  
`server` 文件夹：基于 **go** 开发的**后端**项目文件  
- 启动方法  
先进入后端文件夹  
``` bash
cd server  
```
启动后端入口程序  
``` bash
go run main.go
```

### 区块链浏览器：
- `sirato-free` 文件夹： 一个区块链浏览器，用于查看区块链交易数据，可使用 docker 部署 

### 合约开发：基于 `Solidity`：

- `ERC20.sol` 代币合约  

- `ERC721.sol` NFT合约  

- `nft-market.sol` 交易市场合约  

## 关于 Gitee 分支

> 请创建并使用自己的分支，请**不要**直接向 `main` 分支推送代码  

> 每周例会结束后，组长会合并代码到 `main` 分支，请大家及时从 `main` 分支 pull 代码到自己的分支

- `main` 分支： 用于发布合并（Merge）后的代码到生产环境  
- `raw` 分支： 用于存放学长最初给的代码，分支只读，请不要向此推送  
- `Leaves` 分支： [22 网络 叶文博](https://kiss1314.top) 工作的分支  
- `MRdada` 分支： [22 计科 林振超]() 工作的分支  
- `le` 分支： [22 网络 尧家乐](http://lekami.cn/) 工作的分支
- 待添加  

## 关于项目部署  

### SSH：

- username: `ubuntu`

- password: `xmutBC2024`
  
### 数据库（mysql）:

- db-name: `MyPapers`  

- username: `leaves`  

- password: `leaves`  

- port: `3306`

### 端口

- `22`: ssh && sftp
- `3306`: mysql 数据库  
- `4001`: IPFS P2P 节点通信  
- `5001`: IPFS 的 webui && API  
  
   （web UI 请使用 http://107.155.56.166:5001/webui/ 访问）  
- `8080`: MyPapers web 前端  
- `8081`: IPFS 的网关  
- `8887`: MyPapers 后端  

---  

本次文档更新者： **叶文博**   
最后更新时间： **2025-03-05**  

---
## 附录目录：  
1. [附录1： 项目待办事项](#附录1-项目待办事项-)
2. [附录2： 后端数据库结构](#附录2-后端数据库结构)
3. [附录3： 学长最初的开发环境](#附录3学长最初的开发环境)  

<br>

### 附录1： 项目待办事项  
#### Avatar 显示异常
本地已经可显示、修改头像，但是服务器显示不了，有多个报错如下：  
![](https://kiss1314.top:5555/d/webImage/20250104205949.png)
经测试，图片上传是正常的。

#### 权限管理异常
> 这个问题在 **center** 里的 **User Management** 里能复现。
- 仅管理员才能有修改 role 的权限
  ![](https://kiss1314.top:5555/d/webImage/20250104224620.png)



#### ws 请求头问题
这个问题服务器有，本地没有，尝试过很多修改仍未解决
![](https://kiss1314.top:5555/d/webImage/20250104210137.png)

#### 属性值复定义的问题
这在打开 center 时会显示出来，一堆让人头疼的报错  
![](https://kiss1314.top:5555/d/webImage/20250104221345.png)


#### smtp 验证码机制
- [ ] 在咱们的 VPS 上搭建 smtp 服务器，或采用域名注册商那边的邮件服务。
- [ ] 用户注册时，用咱们的域名（`mypapers.io`）作为邮件发件方，发送邮箱验证码，认证后才准许注册。
- [ ] 有基本的域审查（仅像`@edu`这样受到认可的邮箱才准许注册）
  
#### 已完成
- [x] ~~Journey 显示异常~~

- [x] ~~修复侧边栏折叠按钮的显示异常~~

- [x] ~~Desr 加入编辑栏~~

### 附录2： 后端数据库结构

> 我把后端数据库结构做成了 `Markdown` 表格，作为大家执行 `sql` 查询时的参考

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

  
---
### 附录3：学长最初的开发环境

操作系统：centos7.9.2009  
开发平台：Visual Studio Code 1.88.0；Remix   0.47.0；Goland  2023.1.2  
开发语言：Solidity ^0.8.0 ；javascript ；css ；html；go 1.20.2

---  
