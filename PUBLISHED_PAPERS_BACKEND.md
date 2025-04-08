## 文档说明
撰写人：叶文博
更新日期：  2025 年 4 月 8 日 22:39:00  


## 功能概述

已出版论文（Published Papers）功能的后端开发要求。该功能允许用户上传已发表的学术论文，进行邮箱验证，并生成NFT版权证书。  
1. 所谓已出版的论文是指已经在其他学术出版平台出版的论文，该模块的主要功能是接收这些论文并将这类论文上链存证和生成NFT版权证书。这些论文由版权所有者上传到mypapers平台。  
2. 这类已出版的论文上传mypapers功能大致参考mypapers的投稿系统，但不需要审稿和出版。  
3. 这类论文上传myapapers平台是需要增加邮箱验证，防止将别人的论文当成自己的上传我们平台。邮箱通常为该论文通讯作者的邮箱，如果没有特别指明通讯作者，那选第一作者的邮箱。(论文中就有邮箱信息)  
4. 论文上链和NFT版权证书生成和交易与项目之前的相关功能一样。

### 核心数据：  
#### 期刊论文 (Journal Paper)

|论文名称|期刊名称|论文期数|出版时间|论文页数|ISSN号|论文链接|
|---|---|---|---|---|---|---|
|Journal Paper|Journal Name|Volume & Issue|Date of Publication|Pages|ISSN|Paper Link|
比如这篇已出版在IEEE NETWORK期刊上面的论文，Endex: Degree of Mining Power Decentralization for Proof-of-Work Based Blockchain Systems。填写内容如下：  

IEEE Network，Volume: 34, [Issue: 6](https://ieeexplore.ieee.org/xpl/tocresult.jsp?isnumber=9275414&punumber=65)，12 August 2020，266-271，https://ieeexplore.ieee.org/document/9165548  

#### 会议论文 (Conference Paper)

| 论文名称             | 会议名称            | 会议时间               | 会议地点                | 论文页数  | ISSN号 | 论文链接       |
| ---------------- | --------------- | ------------------ | ------------------- | ----- | ----- | ---------- |
| Conference Paper | Conference Name | Date of Conference | Conference Location | Pages | ISSN  | Paper Link |

比如这篇已出版在IEEE INFOCOM会议上面的论文，Privacy-Preserving Data Evaluation via Functional Encryption, Revisited。填写内容如下：  

[IEEE Conference on Computer Communications](https://ieeexplore.ieee.org/xpl/conhome/10621050/proceeding)(IEEE INFOCOM), 20-23 May 2024, Vancouver, BC, Canada, 233-243,  0743-166X ，https://ieeexplore.ieee.org/document/10621262

注：**会议和期刊论文的的后三个(论文页数，ISSN号 ，论文链接) 为可选**

## 开发建议

1. 遵循RESTful API设计规范
2. 错误处理机制要做好点
3. 使用 [postman](https://www.postman.com/downloads/) 进行接口测试  

## 要实现的接口
### 1. **邮箱验证相关接口**：

```go
// 发送验证码
POST /api/email/verify
请求体：
{
    "email": "corresponding@example.com"
}
响应：
{
    "code": 1000,
    "msg": "验证码已发送",
    "data": null
}

// 验证验证码
POST /api/email/verify-code
请求体：
{
    "email": "corresponding@example.com",
    "code": "123456"
}
响应：
{
    "code": 1000,
    "msg": "验证成功",
    "data": null
}
```

### 2. **已出版论文上传接口**：

```go
// 上传已出版论文
POST /api/paper/published
请求体：multipart/form-data
{
    // 基本信息
    "title": "论文标题",
    "authors": ["作者1", "作者2"],
    "keywords": ["关键词1", "关键词2"],
    "abstract": "论文摘要",
    "paper_type": "journal", // 或 "conference"
    
    // 期刊论文特有字段
    "journal_name": "期刊名称",
    "volume_issue": "Volume: 34, Issue: 6",
    "publication_date": "2020-08-12",
    
    // 会议论文特有字段
    "conference_name": "会议名称",
    "conference_date": ["2024-05-20", "2024-05-23"],
    "conference_location": "会议地点",
    
    // 可选字段
    "pages": "266-271",
    "issn": "0743-166X",
    "paper_link": "https://example.com/paper",
    
    // 验证信息
    "corresponding_email": "corresponding@example.com",
    
    // 区块链信息
    "hash": "文件哈希值",
    "block_address": "区块地址",
    "paper_transaction_address": "交易地址",
    
    // 文件
    "data": 文件二进制数据
}

响应：
{
    "code": 1000,
    "msg": "上传成功",
    "data": {
        "paper_id": 123,
        "title": "论文标题",
        "status": "published"
    }
}
```

### 3. **接口测试示例**：

```bash
# 1. 发送验证码
curl -X POST http://localhost:8080/api/email/verify \
  -H "Content-Type: application/json" \
  -d '{
    "email": "corresponding@example.com"
  }'

# 2. 验证验证码
curl -X POST http://localhost:8080/api/email/verify-code \
  -H "Content-Type: application/json" \
  -d '{
    "email": "corresponding@example.com",
    "code": "123456"
  }'

# 3. 上传期刊论文
curl -X POST http://localhost:8080/api/paper/published \
  -H "Content-Type: multipart/form-data" \
  -F "title=Endex: Degree of Mining Power Decentralization for Proof-of-Work Based Blockchain Systems" \
  -F "authors=Author1" \
  -F "authors=Author2" \
  -F "keywords=blockchain" \
  -F "keywords=mining" \
  -F "abstract=This paper proposes..." \
  -F "paper_type=journal" \
  -F "journal_name=IEEE Network" \
  -F "volume_issue=Volume: 34, Issue: 6" \
  -F "publication_date=2020-08-12" \
  -F "pages=266-271" \
  -F "issn=0743-166X" \
  -F "paper_link=https://ieeexplore.ieee.org/document/9165548" \
  -F "corresponding_email=corresponding@example.com" \
  -F "hash=abc123..." \
  -F "block_address=0x..." \
  -F "paper_transaction_address=0x..." \
  -F "data=@/path/to/paper.pdf"

# 4. 上传会议论文
curl -X POST http://localhost:8080/api/paper/published \
  -H "Content-Type: multipart/form-data" \
  -F "title=Privacy-Preserving Data Evaluation via Functional Encryption, Revisited" \
  -F "authors=Author1" \
  -F "authors=Author2" \
  -F "keywords=privacy" \
  -F "keywords=encryption" \
  -F "abstract=This paper revisits..." \
  -F "paper_type=conference" \
  -F "conference_name=IEEE Conference on Computer Communications" \
  -F "conference_date=2024-05-20" \
  -F "conference_date=2024-05-23" \
  -F "conference_location=Vancouver, BC, Canada" \
  -F "pages=233-243" \
  -F "issn=0743-166X" \
  -F "paper_link=https://ieeexplore.ieee.org/document/10621262" \
  -F "corresponding_email=corresponding@example.com" \
  -F "hash=def456..." \
  -F "block_address=0x..." \
  -F "paper_transaction_address=0x..." \
  -F "data=@/path/to/paper.pdf"
```

### 4. **错误响应示例**：

```json
// 邮箱格式错误
{
    "code": 1001,
    "msg": "邮箱格式不正确",
    "data": null
}

// 验证码错误
{
    "code": 1002,
    "msg": "验证码错误或已过期",
    "data": null
}

// 文件格式错误
{
    "code": 1003,
    "msg": "只支持PDF格式文件",
    "data": null
}

// 文件大小超限
{
    "code": 1004,
    "msg": "文件大小不能超过15MB",
    "data": null
}

// 区块链交易失败
{
    "code": 1005,
    "msg": "区块链交易失败",
    "data": null
}
```

## 数据库设计
### 复用 Papers 表，但需要添加的字段：
   - 已出版论文特有的字段：
     ```go
     type Paper struct {
         // ... 现有字段 ...
         
         // 新增字段
         VolumeIssue        string    `json:"volume_issue" gorm:"comment:期刊卷期号"`
         PublicationDate    time.Time `json:"publication_date" gorm:"comment:出版日期"`
         Pages             string    `json:"pages" gorm:"comment:页码"`
         ISSN              string    `json:"issn" gorm:"comment:ISSN号"`
         PaperLink         string    `json:"paper_link" gorm:"comment:论文链接"`
         CorrespondingEmail string    `json:"corresponding_email" gorm:"comment:通讯作者邮箱"`
         EmailVerified     bool      `json:"email_verified" gorm:"comment:邮箱是否已验证"`
     }
     ```

4. **需要新增的表**：
   ```go
   // 邮箱验证码表
   type EmailVerification struct {
       global.MPS_MODEL
       Email     string    `json:"email" gorm:"comment:邮箱地址;unique"`
       Code      string    `json:"code" gorm:"comment:验证码"`
       ExpiredAt time.Time `json:"expired_at" gorm:"comment:过期时间"`
   }
   ```

5. **数据库迁移建议**：
   ```sql
   -- 添加新字段到papers表
   ALTER TABLE papers
   ADD COLUMN volume_issue VARCHAR(255) COMMENT '期刊卷期号',
   ADD COLUMN publication_date DATETIME COMMENT '出版日期',
   ADD COLUMN pages VARCHAR(50) COMMENT '页码',
   ADD COLUMN issn VARCHAR(20) COMMENT 'ISSN号',
   ADD COLUMN paper_link VARCHAR(255) COMMENT '论文链接',
   ADD COLUMN corresponding_email VARCHAR(255) COMMENT '通讯作者邮箱',
   ADD COLUMN email_verified BOOLEAN DEFAULT FALSE COMMENT '邮箱是否已验证';

   -- 创建邮箱验证码表
   CREATE TABLE email_verifications (
       id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
       created_at DATETIME,
       updated_at DATETIME,
       deleted_at DATETIME,
       email VARCHAR(255) NOT NULL COMMENT '邮箱地址',
       code VARCHAR(6) NOT NULL COMMENT '验证码',
       expired_at DATETIME NOT NULL COMMENT '过期时间',
       UNIQUE KEY uk_email (email)
   );
   ```

### 一些小总结
   - 复用现有的`Paper`表，添加新字段
   - 使用`Status`字段区分论文类型（投稿/已出版）
   - 使用`EmailVerification`表管理邮箱验证
   - 保持与现有系统的兼容性


## 一些 Q & A

### 上传到 IPFS 的流程？

1. 首先将PDF文件上传到IPFS：
```go
func uploadToIPFS(fileData []byte) (string, error) {
    // 1. 连接到IPFS节点
    // 2. 上传文件
    // 3. 获取CID
    return cid, nil
}
```

2. 生成论文元数据JSON：
```go
func generateMetadata(paper *PaperMetadata, fileCID string) ([]byte, error) {
    metadata := struct {
        Title       string   `json:"title"`
        Authors     []string `json:"authors"`
        Keywords    []string `json:"keywords"`
        Abstract    string   `json:"abstract"`
        PaperType   string   `json:"paper_type"`
        // ... 其他字段
        FileCID     string   `json:"file_cid"`
        BlockAddress string  `json:"block_address"`
        TransactionHash string `json:"transaction_hash"`
    }{
        // 填充数据
    }
    
    return json.Marshal(metadata)
}
```

3. 将元数据JSON上传到IPFS：
```go
func uploadMetadata(metadata []byte) (string, error) {
    // 1. 连接到IPFS节点
    // 2. 上传元数据JSON
    // 3. 获取CID
    return metadataCID, nil
}
```

4. 生成NFT证书：
```go
func generateCertificate(paper *PaperMetadata, fileCID string, metadataCID string) (string, error) {
    // 1. 生成证书图片
    // 2. 将证书图片上传到IPFS
    // 3. 返回证书CID
    return certificateCID, nil
}
```

5. 更新数据库记录：
```go
func updatePaperRecord(paperID uint, fileCID, metadataCID, certificateCID string) error {
    paper := &tables.Paper{
        ID: paperID,
        Cid: fileCID,
        ImageCid: certificateCID,
        JsonUri: metadataCID,
        // ... 其他字段
    }
    return mysql.UpdatePaper(paper)
}
```

这样，一篇论文在IPFS上会有三个CID：
1. 论文PDF文件的CID
2. 论文元数据JSON的CID
3. NFT证书图片的CID

这些CID可以用于：
- 访问论文内容
- 验证论文元数据
- 显示NFT证书
- 进行版权交易

所有这些数据都需要在后端进行适当的验证和处理，确保数据的完整性和安全性。
