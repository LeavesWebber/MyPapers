# MyPapers Development Manual

[English](README.md) | [简体中文](README_zh.md)

> This is a decentralized paper management system based on web3.

> This is **not** a final README file; it is a basic overview for the development team.

> Private repository with sensitive credentials. **Do not** make this repository public.

## About URL

### Domain Name:

[MyPapers.io](https://mypapers.io)

### Project Repository:
Private repository, visible only to project members
- [MyPapers Project (Gitee)](https://gitee.com/leaveswebber/MyPapers)

- [MyPapers Project (GitHub)](https://github.com/LeavesWebber/MyPapers)

### Hosts:

- Internal Network: `10.35.54.29`

- Public Network: `107.155.56.166`

## About the Project Structure

### Based on `GO` + `VUE` + `JS` + `HTML` + `CSS`:

- `server` folder: Backend project files developed in **Go**

- `sirato-free` folder: A blockchain explorer for viewing blockchain transaction data

- `web` folder: Frontend project files developed in **Vue** + **JS** + **HTML** + **CSS**

### Based on `Solidity`:

- `ERC20.sol` Token contract

- `ERC721.sol` NFT contract

- `nft-market.sol` Marketplace contract

## About Gitee Branches

> Please create and use your own branch. **Do not** push code directly to the `main` branch.

> After each weekly meeting, the team leader will merge code into the `main` branch. Please pull the latest code from the `main` branch to your own branch.

- `main` branch: Used to merge the final code for production release.
- `raw` branch: Stores the initial code given by seniors, this branch is read-only, do not push to it.
- `Leaves` branch: [22 Network, Ye Wenbo](https://kiss1314.top)'s working branch.
- `MRdada` branch: [22 CS, Lin Zhenchao]'s working branch.
- To be added.

## About Configuration Information

### SSH:

- Username: `ubuntu`

- Password: `xmutBC2024`

### Database (MySQL):

- Database Name: `MyPapers`

- Username: `leaves`

- Password: `leaves`

- Port: `3306`

### Ports:

- `22`: SSH && SFTP
- `3306`: MySQL Database
- `4001`: IPFS P2P Node Communication
- `5001`: IPFS Web UI && API
  (Access the web UI via [http://107.155.56.166:5001/webui/](http://107.155.56.166:5001/webui/))
- `8080`: MyPapers Web Frontend
- `8081`: IPFS Gateway
- `8887`: MyPapers Backend

---

Document edited by: **Ye Wenbo**  
Last updated on: **2024-11-25**

---

## *Additional: Senior's Initial Development Environment*

Operating System: CentOS 7.9.2009  
Development Platforms: Visual Studio Code 1.88.0; Remix 0.47.0; Goland 2023.1.2  
Programming Languages: Solidity ^0.8.0; JavaScript; CSS; HTML; Go 1.20.2

---

## *Additional: Backend Database Structure*

> The backend database structure is provided as a `Markdown` table for reference when executing `SQL` queries.

### 1. authorities  


| Column Name    | Data Type |
| -------------- | --------- |
| created_at     | datetime  |
| updated_at     | datetime  |
| deleted_at     | datetime  |
| authority_id   | bigint    |
| parent_id      | bigint    |
| authority_name | varchar   |

### 2. authority_menus

| Column Name            | Data Type |
| ---------------------- | --------- |
| base_menu_id           | bigint    |
| authority_authority_id | bigint    |

### 3. base_menus

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

### 4. committees

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| created_at  | datetime  |
| updated_at  | datetime  |
| deleted_at  | datetime  |
| creator_id  | bigint    |
| name        | varchar   |
| description | text      |

### 5. conference_issues

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

### 6. conferences

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

### 7. journal_issues

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

### 8. journals

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

### 9. paper_viewers

| Column Name | Data Type |
| ----------- | --------- |
| id          | bigint    |
| created_at  | datetime  |
| updated_at  | datetime  |
| deleted_at  | datetime  |
| paper_id    | bigint    |
| viewer_id   | bigint    |

### 10. papers

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

### 11. reviews

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

### 12. user_authority

| Column Name            | Data Type |
| ---------------------- | --------- |
| user_id                | bigint    |
| authority_authority_id | bigint    |

### 13. user_committee

| Column Name  | Data Type |
| ------------ | --------- |
| user_id      | bigint    |
| committee_id | bigint    |
| start_time   | datetime  |
| end_time     | datetime  |
| position     | varchar   |
| level        | varchar   |

### 14. user_conference

| Column Name   | Data Type |
| ------------- | --------- |
| user_id       | bigint    |
| conference_id | bigint    |
| start_time    | datetime  |
| end_time      | datetime  |
| position      | varchar   |
| level         | varchar   |

### 15. user_journal

| Column Name | Data Type |
| ----------- | --------- |
| user_id     | bigint    |
| journal_id  | bigint    |
| start_time  | datetime  |
| end_time    | datetime  |
| position    | varchar   |
| level       | varchar   |

### 16. user_paper

| Column Name | Data Type |
| ----------- | --------- |
| user_id     | bigint    |
| paper_id    | bigint    |
| old_version | tinyint   |

### 17. users

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
