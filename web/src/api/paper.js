import request from '@/utils/request'

// 提交已出版论文
export function submitPublishedPaper(data) {
  return request({
    url: '/paper/published',
    method: 'post',
    data
  })
}

// 获取已出版论文列表
export function getPublishedPapers() {
  return request({
    url: '/paper/published',
    method: 'get'
  })
}

// 获取论文详情
export function getPaperDetail(paperId) {
  return request({
    url: `/paper/${paperId}`,
    method: 'get'
  })
}

// 获取我的已出版论文
export function getMyPublishedPapers() {
  return request({
    url: '/paper/published/my',
    method: 'get'
  })
}

// 获取荣誉证书
export function getHonoraryCertificate(data) {
  return request({
    url: '/paper/honorary-certificate',
    method: 'post',
    data
  })
}

// 发布论文
export function publishPaper(data) {
  return request({
    url: '/paper/publish',
    method: 'post',
    data
  })
} 