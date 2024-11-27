package mysql

import (
	"server/global"
	"server/model/request"
	"server/model/tables"
)

// GetReviews 获取各个审核人的审核结果
func GetReviews(paperID uint) (reviews []*tables.Review, err error) {
	err = global.MPS_DB.Where("paper_id = ?", paperID).Find(&reviews).Error
	return
}

// GetReviewIds 查询user的paperIds
func GetReviewIds(userId uint) (paperIds []uint, err error) {
	// 去reviews表找自己对应的paperId
	//err = global.MPS_DB.Table("reviews").Where("reviewer_id = ?", userId).Pluck("paper_id", &paperIds).Error
	// 软删除的paper不显示
	err = global.MPS_DB.Table("reviews").Where("reviewer_id = ? and deleted_at is null and old_version = ?", userId, false).Pluck("paper_id", &paperIds).Error
	return
}

// InsertReview 插入一条review
func InsertReview(paperId uint, reviewersId []uint) (err error) {
	// 不存在则插入
	tx := global.MPS_DB.Begin()
	for _, v := range reviewersId {
		if err = tx.Where("paper_id = ? and reviewer_id = ?", paperId, v).FirstOrCreate(&tables.Review{PaperId: paperId, ReviewerId: v}).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

// SubmitReview 更新reviews表
func SubmitReview(in *request.SubmitReview, currentId uint) (out *tables.Review, err error) {
	out = &tables.Review{
		Comment: in.Comment,
		Status:  in.Status,
	}
	return out, global.MPS_DB.Model(&tables.Review{}).Where("paper_id = ? and reviewer_id = ?", in.PaperId, currentId).Updates(out).First(out).Error
}
