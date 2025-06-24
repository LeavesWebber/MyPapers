import http from '../utils/request'

// 导入NFT API
export * from './nft'

// 请求首页数据
export const getHomeData = () => {
    return http.get('/home')
}
export const register = (data) => {
    return http.post('/user/register', data)
}
export const login = (data) => {
    return http.post('/user/login', data).then(response => {
        if (response && response.data) {
            return response.data;
        }
        return Promise.reject('Invalid response format');
    });
}
export const getMenu = (data) => {
    return http.get('/menu/getMenu', { params: data }).then(response => {
        if (response && response.data) {
            return response.data;
        }
        return Promise.reject('Invalid response format');
    });
}

// 用户
export const getSelfInfo = () => {
    return http.get('/user/getSelfInfo').then(response => {
        if (response && response.data) {
            return response.data;
        }
        return Promise.reject(new Error('Invalid response format'));
    }).catch(error => {
        console.error('获取用户信息失败:', error);
        if (error.response && error.response.status === 401) {
            // 清除本地存储的token
            localStorage.removeItem('token');
            // 抛出401错误，让调用者处理重定向
            throw error;
        }
        throw error;
    });
}
// 修改头像
export const changeHeaderImg = (data) => {
    return http.post('/user/changeHeaderImg', data)
}
// 修改用户信息
export const setSelfInfo = (data) => {
    return http.put('/user/setSelfInfo', data)
}
export const SendMail = (data) =>{
    return http.post('/user/SendMail',data)
}
// 查询自己的投稿
export const getSelfPapers = (data) => {
    return http.get('/paper/selfList', { params: data })
}

export const getDetailPapers = (data) => {
    return http.get('/paper/detail', { params: data })
}

export const deletePaper = (data) => {
    return http.delete('/paper/delete', { params: data })
}

// 审稿
export const getReviews = (data) => {
    return http.get('/review/list', { params: data })
}

export const allotReviewers = (data) => {
    return http.post('/review/allot', data)
}

export const submitReview = (data) => {
    return http.post('/review/submit', data)
}

// 创建委员会
export const createCommittee = (data) => {
    return http.post('/committee/create', data)
}
// 获取委员会列表
export const getCommitteeList = () => {
    return http.get('/committee/list')
}
// 查询自己所在的委员会列表
export const getSelfCommitteeList = () => {
    return http.get('/committee/selfList')
}
// 获取委员会详情
export const getCommitteeDetail = (data) => {
    return http.get('/committee/detail', { params: data })
}
// 根据委员会查询会议列表
export const getConferencelistByCommittee = (data) => {
    return http.get('/conference/listByCommittee', { params: data })
}
// 根据委员会查询期刊列表
export const getJournallistByCommittee = (data) => {
    return http.get('/journal/listByCommittee', { params: data })
}
// 更新委员会信息
export const updateCommittee = (data) => {
    return http.put('/committee/update', data)
}
// 删除委员会
export const deleteCommittee = (data) => {
    return http.delete('/committee/delete', { params: data })
}

// 创建会议
export const createConference = (data) => {
    return http.post('/conference/create', data)
}
// 获取会议列表
export const getConferenceList = () => {
    return http.get('/conference/list')
}
// 查询自己所在的会议列表
export const getSelfConferenceList = () => {
    return http.get('/conference/selfList')
}
// 获取会议详情
export const getConferenceDetail = (data) => {
    return http.get('/conference/detail', { params: data })
}
// 更新会议信息
export const updateConference = (data) => {
    return http.put('/conference/update', data)
}
// 删除会议
export const deleteConference = (data) => {
    return http.delete('/conference/delete', { params: data })
}

// 创建期刊
export const createJournal = (data) => {
    return http.post('/journal/create', data)
}
// 获取期刊列表
export const getJournalList = () => {
    return http.get('/journal/list')
}

// 查询自己所在的期刊列表
export const getSelfJournalList = () => {
    return http.get('/journal/selfList')
}
// 获取期刊详情
export const getJournalDetail = (data) => {
    return http.get('/journal/detail', { params: data })
}
// 更新期刊信息
export const updateJournal = (data) => {
    return http.put('/journal/update', data)
}
// 删除期刊
export const deleteJournal = (data) => {
    return http.delete('/journal/delete', { params: data })
}

// 获取用户在期刊的level
export const getJournalLevel = (data) => {
    return http.get('/journal/level', { params: data })
}
// 投稿
export const submitPaper = (data) => {
    return http.post('/paper/submit', data)
}
// 修改投稿
export const updatePaper = (data) => {
    return http.put('/paper/update', data)
}

// 查询已经审核通过的投稿列表
export const getAllAcceptPapers = () => {
    return http.get('/paper/acceptPaperList')
}

// 查询投稿的历史版本
export const getPaperVersions = (data) => {
    return http.get('/paper/version', { params: data })
}

// 分期
export const getJournalIssuePapers = (data) => {
    return http.get('/paper/acceptPaperListByJournalAndTime', { params: data })
}

// 分期
export const getConferenceIssuePapers = (data) => {
    return http.get('/paper/acceptPaperListByConferenceAndTime', { params: data })
}

// 查询期刊Issue列表
export const getJournalIssues = (data) => {
    return http.get('/journal/issue/list', { params: data })
}

// 创建期刊Issue
export const createJournalIssue = (data) => {
    return http.post('/journal/issue/create', data)
}

// 更新期刊Issue
export const updateJournalIssue = (data) => {
    return http.put('/journal/issue/update', data)
}

// 删除期刊Issue
export const deleteJournalIssue = (data) => {
    return http.delete('/journal/issue/delete', { params: data })
}

// 切换角色
export const changeAuthority = (data) => {
    return http.post('/authority/changeAuthority', data)
}

// 查询会议Issue列表
export const getConferenceIssues = (data) => {
    return http.get('/conference/issue/list', { params: data })
}

// 创建会议Issue
export const createConferenceIssue = (data) => {
    return http.post('/conference/issue/create', data)
}

// 更新会议Issue
export const updateConferenceIssue = (data) => {
    return http.put('/conference/issue/update', data)
}

// 删除会议Issue
export const deleteConferenceIssue = (data) => {
    return http.delete('/conference/issue/delete', { params: data })
}

// 获取用户在会议的level
export const getConferenceLevel = (data) => {
    return http.get('/conference/level', { params: data })
}

// 获取荣誉证书
export const getHonoraryCertificate = (data) => {
    return http.get('/paper/honoraryCertificate', { params: data })
}

// 发布论文
export const publishPaper = (data) => {
    return http.post('/paper/publish', data)
}

// 查看用户是否有权限查看投稿
export const checkPaperViewer = (data) => {
    return http.get('/paper/checkPaperViewer', { params: data })
}

// 增加投稿可查看者
export const addPaperViewer = (data) => {
    return http.post('/paper/addPaperViewer', data)
}

// 获取我的NFT
export const getMyNFTs = () => {
    return http.get('/paper/myNFTs')
}

// 修改价格
export const updatePrice = (data) => {
    return http.put('/paper/updatePrice', data)
}

// 查询所有用户
export const getUserList = () => {
    return http.get('/user/list')
}

// 设置用户信息
export const setUserInfo = (data) => {
    return http.put('/user/setUserInfo', data)
}

// 根据tokenId查询NFT信息
export const getNFTInfoByTokenId = (data) => {
    return http.get('/paper/getNFTInfo', { params: data })
}

// 修改投稿对应的user_id
export const updatePaperUserId = (data) => {
    return http.put('/paper/updatePaperUserId', data)
}

// MPS通证相关API
// 1. 查询用户MPS余额
export const getMPSBalance = () => {
    return http.get('/mps/balance')
}

// 2. 获取用户通证交易记录
export const getMPSTransactions = (data) => {
    return http.get('/mps/transactions', { params: data })
}

// 3. 法币购买MPS
export const buyMPSWithFiat = (data) => {
    return http.post('/mps/buy', data)
}

// 4. MPS卖出换取法币
export const sellMPSToFiat = (data) => {
    return http.post('/mps/sell', data)
}

// 5. 获取当前MPS兑换率
export const getMPSRate = () => {
    return http.get('/mps/rate')
}

// 6. 获取用户激励记录
export const getMPSRewards = () => {
    return http.get('/mps/rewards')
}