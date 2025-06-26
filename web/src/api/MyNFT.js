import http from '../utils/request'

// 铸造NFT
export const mintNFT = (data) => http.post('/mynft/mint', data)
// 批量铸造
export const bulkMintNFT = (data) => http.post('/mynft/bulk-mint', data)
// 设置合约元数据
export const setContractMetadata = (data) => http.post('/mynft/set-metadata', data)
// 更新版税
export const updateRoyalty = (data) => http.post('/mynft/update-royalty', data)
// 查询tokenURI
export const getTokenURI = (tokenId) => http.get(`/mynft/token-uri/${tokenId}`)
// 查询合约支持的接口
export const supportsInterface = (interfaceId) => http.get(`/mynft/supports-interface/${interfaceId}`)
// 查询owner
export const getOwner = () => http.get('/mynft/owner')
// 查询余额
export const getBalanceOf = (address) => http.get(`/mynft/balance-of/${address}`)
// 查询ownerOf
export const getOwnerOf = (tokenId) => http.get(`/mynft/owner-of/${tokenId}`)
// 查询总供应量
export const getTotalSupply = () => http.get('/mynft/total-supply') 