package response

import (
	"server/model/tables"
)

type GetPaper struct {
	Paper       *tables.Paper `json:"paper"`
	ReviewInfos []*ReviewInfo `json:"review_infos"`
	Level       string        `json:"level"`
}
type ReviewInfo struct {
	ReviewerName string `json:"reviewer_name"`
	Comment      string `json:"comment"`
	Status       string `json:"status"`
}

type HonoraryCertificateInfo struct {
	Cid         string `json:"cid"`
	Url         string `json:"url"`
	Uri         string `json:"image_uri"`
	MetadataUri string `json:"metadata_uri"`
}

type GetMyNFTs struct {
	PaperId               uint   `json:"paper_id"`
	TokenId               string `json:"token_id"`
	ImageUri              string `json:"image_uri"`
	ImageUrl              string `json:"image_url"`
	ImageCid              string `json:"image_cid"`
	CopyrightTradingPrice uint   `json:"copy_right_trading_price"`
	TransactionHash       string `json:"transaction_hash"`
}
