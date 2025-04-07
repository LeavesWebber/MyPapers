# 已发表论文后端开发指引

## 目录

- [文档说明](#文档说明)
- [功能概述](#功能概述)
- [开发建议](#开发建议)
- [API 接口设计](#api-接口设计)
  - [1. 邮箱验证相关接口](#1-邮箱验证相关接口)
    - [1.1 发送验证码](#11-发送验证码)
    - [1.2 验证验证码](#12-验证验证码)
  - [2. 论文管理接口](#2-论文管理接口)
    - [2.1 上传论文](#21-上传论文)
    - [2.2 获取论文列表](#22-获取论文列表)
    - [2.3 下载论文](#23-下载论文)
    - [2.4 批量下载论文](#24-批量下载论文)
    - [2.5 查看NFT详情](#25-查看nft详情)
- [数据库设计](#数据库设计)
  - [1. published_papers 表](#1-published_papers-表)
  - [2. email_verifications 表](#2-email_verifications-表)
- [业务逻辑实现](#业务逻辑实现)
  - [1. 邮箱验证流程](#1-邮箱验证流程)
  - [2. 论文上传流程](#2-论文上传流程)
  - [3. NFT生成流程](#3-nft生成流程)
- [安全考虑](#安全考虑)
- [注意事项](#注意事项)

## 文档说明

更新日期：  2025年3月5日00:55:54  
commit id：  `a670f8b3`  

我在基本完成新功能前端部分后撰写了这篇文章，旨在帮助后端同学更高效的完成开发和协同。  

## 功能概述

已发表论文（Published Papers）功能的后端开发要求。该功能允许用户上传已发表的学术论文，进行邮箱验证，并生成NFT版权证书。

## 开发建议

1. 遵循RESTful API设计规范
2. 错误处理机制要做好点
3. 使用 [postman](https://www.postman.com/downloads/) 进行接口测试  

## API 接口设计

### 1. 邮箱验证相关接口

#### 1.1 发送验证码

- **接口**: `POST /api/published-papers/verify-email`

- **请求体**:
  
  ```json
  {
    "email": "string"  // 作者邮箱地址
  }
  ```

- **响应**:
  
  ```json
  {
    "code": 1000,      // 成功状态码
    "msg": "string",   // 响应消息
    "data": null
  }
  ```

- **说明**:
  
  - 验证码有效期建议设置为10分钟
  - 同一邮箱60秒内只能请求一次验证码
  - 验证码应为6位数字

#### 1.2 验证验证码

- **接口**: `POST /api/published-papers/verify-code`

- **请求体**:
  
  ```json
  {
    "email": "string",  // 作者邮箱地址
    "code": "string"    // 6位验证码
  }
  ```

- **响应**:
  
  ```json
  {
    "code": 1000,
    "msg": "string",
    "data": null
  }
  ```

### 2. 论文管理接口

#### 2.1 上传论文

- **接口**: `POST /api/published-papers/upload`

- **请求体**: `multipart/form-data`
  
  ```
  title: string            // 论文标题
  authors: string          // 作者列表（逗号分隔）
  abstract: string         // 摘要
  correspondingEmail: string  // 通讯作者邮箱
  venueType: string       // 发表类型（journal/conference）
  venueName: string       // 期刊/会议名称
  publicationDate: string // 发表日期
  
  ```

- **响应**:
  
  ```json
  {
    "code": 1000,
    "msg": "string",
    "data": {
      "paperId": "string",
      "title": "string",
      "status": "string"
    }
  }
  ```

- **说明**:
  
  - 文件大小限制：50MB
  - 仅支持PDF格式
  - 需要验证用户登录状态
  - 需要验证邮箱已通过验证

#### 2.2 获取论文列表

- **接口**: `GET /api/published-papers`

- **参数**:
  
  ```
  page: number          // 页码，从1开始
  pageSize: number      // 每页数量
  query: string         // 搜索关键词（可选）
  ```

- **响应**:
  
  ```json
  {
    "code": 1000,
    "msg": "string",
    "data": {
      "total": number,
      "papers": [
        {
          "id": "string",
          "title": "string",
          "authors": "string",
          "venueName": "string",
          "venueType": "string",
          "publicationDate": "string",
          "status": "string",
          "nftGenerated": boolean,
          "createdAt": "string"
        }
      ]
    }
  }
  ```

#### 2.3 下载论文

- **接口**: `GET /api/published-papers/:id/download`
- **响应**: `application/pdf`
- **说明**:
  - 需要验证用户权限
  - 返回PDF文件流

#### 2.4 批量下载论文

- **接口**: `POST /api/published-papers/batch-download`

- **请求体**:
  
  ```json
  {
    "paperIds": ["string"]
  }
  ```

- **响应**: `application/zip`

- **说明**:
  
  - 将多个PDF打包成zip文件返回
  - 需要验证用户权限

#### 2.5 查看NFT详情

- **接口**: `GET /api/published-papers/:id/nft`

- **响应**:
  
  ```json
  {
    "code": 1000,
    "msg": "string",
    "data": {
      "tokenId": "string",
      "contractAddress": "string",
      "blockchain": "string",
      "creationDate": "string",
      "imageUrl": "string",
      "blockchainUrl": "string",
      "certificateUrl": "string"
    }
  }
  ```

## 数据库设计

### 1. published_papers 表

```sql
CREATE TABLE published_papers (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    title VARCHAR(500) NOT NULL,
    authors TEXT NOT NULL,
    abstract TEXT,
    corresponding_email VARCHAR(255) NOT NULL,
    venue_type ENUM('journal', 'conference') NOT NULL,
    venue_name VARCHAR(255) NOT NULL,
    publication_date DATE NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    status ENUM('pending', 'verified', 'rejected') DEFAULT 'pending',
    nft_generated BOOLEAN DEFAULT FALSE,
    nft_token_id VARCHAR(100),
    nft_contract_address VARCHAR(100),
    nft_blockchain VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user_id (user_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
);
```

### 2. email_verifications 表

```sql
CREATE TABLE email_verifications (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    code VARCHAR(6) NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_email_code (email, code),
    INDEX idx_expires_at (expires_at)
);
```

## 业务逻辑实现

### 1. 邮箱验证流程

1. 生成6位随机数字验证码
2. 保存验证码到数据库，设置过期时间
3. 发送验证码邮件
4. 验证时检查验证码是否正确且未过期
5. 验证成功后标记验证状态

### 2. 论文上传流程

1. 验证用户登录状态和权限
2. 验证邮箱是否已通过验证
3. 验证并保存PDF文件
4. 创建论文记录
5. 触发NFT生成任务

### 3. NFT生成流程

1. 创建异步任务处理NFT生成
2. 生成NFT元数据（包含论文信息）
3. 调用区块链接口铸造NFT
4. 更新论文记录的NFT信息
5. 生成可下载的证书文件

## 安全考虑

1. **用户认证**
   
   - 实现JWT或Session based认证
   - 验证用户权限和所有权

2. **文件上传安全**
   
   - 严格限制文件类型和大小
   - 使用安全的文件存储路径

3. **API访问控制**
   
   - 实现请求频率限制
   - 验证请求来源
   - 实现CORS策略

4. **数据安全**
   
   - 加密敏感信息

## 注意事项

1. 确保邮件服务的稳定性和送达率
2. 注意区块链交易的gas费用控制
3. 实现适当的重试机制
4. 考虑系统的可扩展性 