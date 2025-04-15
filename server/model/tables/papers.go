package tables

import (
	"mime/multipart"
	"server/global"
)

type Paper struct {
	global.MPS_MODEL
	ConferenceId          uint   `json:"conference_id" gorm:"comment:会议id"`
	JournalId             uint   `json:"journal_id" gorm:"comment:期刊id"`
	VersionId             uint   `json:"version_id" gorm:"comment:版本id;unique"`
	DownloadPrice         uint   `json:"download_price" gorm:"comment:下载价格" binding:"required"`
	CopyrightTradingPrice uint   `json:"copyright_trading_price" gorm:"comment:版权交易价格"`
	TokenId               string `json:"token_id" gorm:"comment:token id"`
	Title                 string `json:"title" gorm:"comment:标题"`
	Authors               string `json:"authors" gorm:"comment:作者"`
	PaperType             string `json:"paper_type" gorm:"comment:论文类型"`
	Abstract              string `json:"abstract" gorm:"comment:摘要;type:text"`
	KeyWords              string `json:"key_words" gorm:"comment:关键词;type:text"`
	SubjectCategory       string `json:"subject_category" gorm:"comment:论文所属学科分类"`
	ManuscriptID          string `json:"manuscript_id" gorm:"comment:投稿编号"`
	InformedConsent       string `json:"informed_consent" gorm:"comment:是否已获得知情同意"`
	AnimalSubjects        string `json:"animal_subjects" gorm:"comment:是否涉及动物试验"`
	CorAuthor             string `json:"cor_author" gorm:"comment:通讯作者"`
	ManuscriptType        string `json:"manuscript_type" gorm:"comment:投稿类型"`
	UniqueContribution    string `json:"unique_contribution" gorm:"comment:论文独特贡献;type:text"`
	Hash                  string `json:"hash" gorm:"comment:论文hash值"`
	BlockAddress          string `json:"block_address" gorm:"comment:投稿区块地址"`
	PaperTransactionHash  string `json:"paper_transaction_hash" gorm:"comment:存论文hash的transaction address"`
	Filepath              string `json:"filepath" gorm:"comment:论文文件路径"`
	Cid                   string `json:"cid" gorm:"comment:论文cid"`
	Status                string `json:"status" gorm:"comment:投稿状态"`
	ImageUri              string `json:"image_uri" gorm:"comment:证书图片IPFs路径"`
	ImageUrl              string `json:"image_url" gorm:"comment:证书图片路径"`
	ImageCid              string `json:"image_cid" gorm:"comment:证书图片cid"`
	JsonUri               string `json:"json_uri" gorm:"comment:证书nft元数据json路径"`
	TransactionHash       string `json:"transaction_hash" gorm:"comment:交易hash"`
	JournalName           string `json:"journal_name" gorm:"comment:期刊名称"`
	VolumeAndIssue        string `json:"volume_and_issue" gorm:"comment:卷号和期号"`
	PublicationDate       string `json:"publication_date" gorm:"comment:发表日期"`
	ConferenceName        string `json:"conference_name" gorm:"comment:会议名称"`
	ConferenceDate        string `json:"conference_date" gorm:"comment:会议日期"`
	ConferenceLocation    string `json:"conference_location" gorm:"comment:会议地点"`
	Pages                 string `json:"pages" gorm:"comment:页码"`
	Issn                  string `json:"issn" gorm:"comment:ISSN号"`
	PaperLink             string `json:"paper_link" gorm:"comment:论文链接"`
	//Justification      string    `json:"justification" comment:"投稿理由"`
	//Scope              string    `json:"scope" comment:"研究范围"`
	Users []User                `json:"user,omitempty" gorm:"many2many:user_paper;"`
	Data  *multipart.FileHeader `json:"-" gorm:"-" binding:"required"`
}
