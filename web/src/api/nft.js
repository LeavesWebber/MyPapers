import http from '../utils/request'

// NFT相关API接口

// 1. 铸造NFT
export const mintNFT = (data) => {
    return http.post('/nft/mint', data)
}

// 2. 获取我的NFT列表
export const getMyNFTs = () => {
    return http.get('/nft/my-nfts')
}

// 3. 根据Token ID获取NFT信息
export const getNFTByTokenId = (tokenId) => {
    return http.get(`/nft/token/${tokenId}`)
}

// 4. 获取NFT元数据
export const getNFTMetadata = (tokenId) => {
    return http.get(`/nft/metadata/${tokenId}`)
}

// 5. 设置NFT价格
export const setNFTPrice = (data) => {
    return http.put('/nft/set-price', data)
}

// 6. 获取NFT市场列表
export const getNFTMarketplace = (params) => {
    return http.get('/nft/marketplace', { params })
}

// 7. 购买NFT
export const buyNFT = (data) => {
    return http.post('/nft/buy', data)
}

// 8. 出售NFT
export const sellNFT = (data) => {
    return http.post('/nft/sell', data)
}

// 9. 取消NFT出售
export const cancelNFTSale = (tokenId) => {
    return http.delete(`/nft/cancel-sale/${tokenId}`)
}

// 10. 获取NFT交易历史
export const getNFTTransactionHistory = (tokenId) => {
    return http.get(`/nft/transactions/${tokenId}`)
}

// 11. 获取用户NFT余额
export const getNFTBalance = (address) => {
    return http.get(`/nft/balance/${address}`)
}

// 12. 转移NFT
export const transferNFT = (data) => {
    return http.post('/nft/transfer', data)
}

// 13. 获取NFT版税信息
export const getNFTRoyaltyInfo = (tokenId, salePrice) => {
    return http.get('/nft/royalty-info', { 
        params: { tokenId, salePrice } 
    })
}

// 14. 批量铸造NFT
export const batchMintNFT = (data) => {
    return http.post('/nft/batch-mint', data)
}

// 15. 获取NFT统计信息
export const getNFTStats = () => {
    return http.get('/nft/stats')
}

// 16. 搜索NFT
export const searchNFTs = (params) => {
    return http.get('/nft/search', { params })
}

// 17. 获取NFT分类
export const getNFTCategories = () => {
    return http.get('/nft/categories')
}

// 18. 根据分类获取NFT
export const getNFTsByCategory = (category) => {
    return http.get(`/nft/category/${category}`)
}

// 19. 获取热门NFT
export const getHotNFTs = () => {
    return http.get('/nft/hot')
}

// 20. 获取最新NFT
export const getLatestNFTs = () => {
    return http.get('/nft/latest')
}

// 21. 获取NFT创建者信息
export const getNFTCreator = (tokenId) => {
    return http.get(`/nft/creator/${tokenId}`)
}

// 22. 获取NFT所有者信息
export const getNFTOwner = (tokenId) => {
    return http.get(`/nft/owner/${tokenId}`)
}

// 23. 检查NFT是否已授权
export const checkNFTApproval = (data) => {
    return http.post('/nft/check-approval', data)
}

// 24. 授权NFT
export const approveNFT = (data) => {
    return http.post('/nft/approve', data)
}

// 25. 获取NFT事件
export const getNFTEvents = (tokenId) => {
    return http.get(`/nft/events/${tokenId}`)
}

// 26. 获取NFT价格历史
export const getNFTPriceHistory = (tokenId) => {
    return http.get(`/nft/price-history/${tokenId}`)
}

// 27. 获取NFT收藏夹
export const getNFTFavorites = () => {
    return http.get('/nft/favorites')
}

// 28. 添加NFT到收藏夹
export const addToFavorites = (tokenId) => {
    return http.post(`/nft/favorites/${tokenId}`)
}

// 29. 从收藏夹移除NFT
export const removeFromFavorites = (tokenId) => {
    return http.delete(`/nft/favorites/${tokenId}`)
}

// 30. 获取NFT评论
export const getNFTComments = (tokenId) => {
    return http.get(`/nft/comments/${tokenId}`)
}

// 31. 添加NFT评论
export const addNFTComment = (data) => {
    return http.post('/nft/comments', data)
}

// 32. 获取NFT排行榜
export const getNFTRanking = (type = 'volume') => {
    return http.get(`/nft/ranking/${type}`)
}

// 33. 获取NFT分析数据
export const getNFTAnalytics = (tokenId) => {
    return http.get(`/nft/analytics/${tokenId}`)
}

// 34. 获取NFT相似推荐
export const getNFTSimilar = (tokenId) => {
    return http.get(`/nft/similar/${tokenId}`)
}

// 35. 获取NFT创建者作品
export const getCreatorNFTs = (address) => {
    return http.get(`/nft/creator-works/${address}`)
}

// 36. 获取NFT收藏者
export const getNFTCollectors = (tokenId) => {
    return http.get(`/nft/collectors/${tokenId}`)
}

// 37. 获取NFT交易统计
export const getNFTTradeStats = (tokenId) => {
    return http.get(`/nft/trade-stats/${tokenId}`)
}

// 38. 获取NFT市场趋势
export const getNFTMarketTrends = () => {
    return http.get('/nft/market-trends')
}

// 39. 获取NFT稀有度
export const getNFTRarity = (tokenId) => {
    return http.get(`/nft/rarity/${tokenId}`)
}

// 40. 获取NFT属性
export const getNFTAttributes = (tokenId) => {
    return http.get(`/nft/attributes/${tokenId}`)
} 