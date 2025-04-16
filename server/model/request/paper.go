package request

import "mime/multipart"

// SubmitPaper 投稿请求参数
type SubmitPaper struct {
	Id           uint   `form:"id"`
	ConferenceId uint   `form:"conference_id"`
	JournalId    uint   `form:"journal_id"`
	VersionId    uint   `form:"version_id"`
	PaperType    string `form:"paper_type" binding:"required"`
	Title        string `form:"title" binding:"required"`
	Abstract     string `form:"abstract" binding:"required"`
	//InformedConsent string `form:"informed_consent" binding:"required"`
	//AnimalSubjects     string `form:"animal_subjects" binding:"required"`
	CorAuthor string `form:"cor_author" binding:"required"`
	//ManuscriptType     string `form:"manuscript_type" binding:"required"`
	UniqueContribution      string `form:"unique_contribution" binding:"required"`
	Hash                    string `form:"hash" binding:"required"`
	BlockAddress            string `form:"block_address" binding:"required"`
	PaperTransactionAddress string `form:"paper_transaction_address" binding:"required"`
	Authors                 string `form:"authors" binding:"required"`
	//Authors         []string `form:"authors"`
	//Keywords        []string `form:"keywords" binding:"required"`
	//SubjectCategory []string `form:"subject_category" binding:"required"`
	Keywords        string `form:"keywords" binding:"required"`
	SubjectCategory string `form:"subject_category" binding:"required"`

	Data *multipart.FileHeader `form:"data" binding:"required"`
}

type UpdatePaper struct {
	PaperId   uint   `form:"paper_id" binding:"required"`
	PaperType string `form:"paper_type"`
	Title     string `form:"title"`
	Abstract  string `form:"abstract"`
	//InformedConsent string `form:"informed_consent"`
	//AnimalSubjects     string   `form:"animal_subjects"`
	CorAuthor               string `form:"cor_author"`
	ManuscriptType          string `form:"manuscript_type"`
	UniqueContribution      string `form:"unique_contribution"`
	Hash                    string `form:"hash" binding:"required"`
	BlockAddress            string `form:"block_address" binding:"required"`
	PaperTransactionAddress string `form:"paper_transaction_address"`
	Authors                 string `form:"authors" binding:"required"`
	//Keywords           []string              `form:"keywords"`
	//SubjectCategory    []string              `form:"subject_category"`
	Keywords        string                `form:"keywords" binding:"required"`
	SubjectCategory string                `form:"subject_category" binding:"required"`
	Data            *multipart.FileHeader `form:"data" binding:"required"`
}

type PublishPaper struct {
	PaperId               uint   `json:"paper_id" binding:"required"`
	DownloadPrice         uint   `json:"download_price" binding:"required"`
	CopyrightTradingTrice uint   `json:"copyright_trading_price" binding:"required"`
	TokenId               string `json:"token_id" binding:"required"`
	TransactionHash       string `json:"transaction_hash" binding:"required"`
}

// AddPaperViewer 设置投稿可查看者
type AddPaperViewer struct {
	PaperId  uint `json:"paper_id" binding:"required"`
	ViewerId uint `json:"viewer_id"`
}

// UpdatePrice 更新价格
type UpdatePrice struct {
	PaperId               uint `json:"paper_id" binding:"required"`
	DownloadPrice         uint `json:"download_price" binding:"required"`
	CopyrightTradingTrice uint `json:"copyright_trading_price" binding:"required"`
}

// UploadPublishedPaper 上传已发表论文请求参数
type UploadPublishedPaper struct {
	Title                   string                `form:"title" binding:"required"`
	Authors                 string                `form:"authors" binding:"required"`
	Keywords                string                `form:"keywords" binding:"required"`
	CorrespondingEmail      string                `form:"corresponding_email" binding:"required"`
	PaperType               string                `form:"paper_type" binding:"required"`
	JournalName             string                `form:"journal_name"`
	VolumeAndIssue          string                `form:"volume_and_issue"`
	PublicationDate         string                `form:"publication_date"`
	ConferenceName          string                `form:"conference_name"`
	ConferenceDate          string                `form:"conference_date"`
	ConferenceLocation      string                `form:"conference_location"`
	Pages                   string                `form:"pages"`
	Issn                    string                `form:"issn"`
	PaperLink               string                `form:"paper_link"`
	Hash                    string                `form:"hash" binding:"required"`
	BlockAddress            string                `form:"block_address" binding:"required"`
	PaperTransactionAddress string                `form:"paper_transaction_address" binding:"required"`
	Data                    *multipart.FileHeader `form:"data" binding:"required"`
}
