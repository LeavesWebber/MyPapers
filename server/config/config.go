package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt"`
	Zap     Zap     `mapstructure:"zap"`
	System  System  `mapstructure:"system"`
	Captcha Captcha `mapstructure:"captcha"`
	Mysql   Mysql   `mapstructure:"mysql"`
	// 跨域配置
	Cors          CORS       `mapstructure:"cors" json:"cors" yaml:"cors"`
	ImagePath     string     `mapstructure:"image-path"`
	ImageIpfsPath string     `mapstructure:"image-ipfs-path"`
	IPFS          IPFS       `mapstructure:"ipfs"`
	Nginx         Nginx      `mapstructure:"nginx"`
	Blockchain    Blockchain `mapstructure:"blockchain"`
	WxPay         WxPay      `mapstructure:"wxpay"`
	AliPay        AliPay     `mapstructure:"alipay"`
	Business      Business   `mapstructure:"business"`
	Smtp          Smtp       `mapstructure:"smtp"`
	Redis         Redis      `mapstructure:"redis"`
	RabbitMQ      RabbitMQ   `mapstructure:"rabbitmq-config"`
}

// Blockchain 区块链配置
type Blockchain struct {
	EthNodeURL string `mapstructure:"eth-node-url"` // 以太坊节点URL
	ChainID    int64  `mapstructure:"chain-id"`     // 链ID
	GasLimit   uint64 `mapstructure:"gas-limit"`    // Gas限制
	Decimals   int64  `mapstructure:"decimals"`
	// 智能合约地址
	MPSContractAddress         string `mapstructure:"mps-contract-address"`         // MPS代币合约地址
	ERC20ContractAddress       string `mapstructure:"erc20-contract-address"`       // ERC20合约地址
	ERC721ContractAddress      string `mapstructure:"erc721-contract-address"`      // ERC721合约地址
	MarketplaceContractAddress string `mapstructure:"marketplace-contract-address"` // 市场合约地址

	// 管理员配置
	AdminPrivateKey string `mapstructure:"admin-private-key"` // 管理员私钥
	AdminAddress    string `mapstructure:"admin-address"`     // 管理员地址

}

// WxPay 微信支付配置
type WxPay struct {
	WxPayType string `mapstructure:"wxpay-type"`
	// 基础配置
	AppID string `mapstructure:"app-id"` // 公众号ID
	MchID string `mapstructure:"mch-id"` // 商户号
	Key   string `mapstructure:"key"`    // API密钥

	// 支付配置
	TradeType string `mapstructure:"trade-type"` // 交易类型
	SignType  string `mapstructure:"sign-type"`  // 签名类型

	// URL配置
	NotifyURL  string `mapstructure:"notify-url"`  // 支付回调通知地址
	SuccessURL string `mapstructure:"success-url"` // 支付成功跳转页面
	FailURL    string `mapstructure:"fail-url"`    // 支付失败跳转页面
}

// AliPay 支付宝支付配置
type AliPay struct {
	//支付标识
	AliPayType string `mapstructure:"ali-pay-type"`
	// 基础配置
	AppID      string `mapstructure:"app-id"`      // 公众号ID
	PrivateKey string `mapstructure:"private-key"` // API密钥
	Charset    string `mapstructure:"charset"`     // 请求使用的编码格式
	Format     string `mapstructure:"format"`      // 请求格式
	//证书配置
	PublicCert    string `mapstructure:"public-cert"`    // 应用公钥证书
	PayPublicCert string `mapstructure:"paypublic-cert"` //支付宝公钥证书
	PayRootCert   string `mapstructure:"payroot-cert"`   //支付宝根证书

	// 支付配置
	SignType    string `mapstructure:"sign-type"`    // 签名类型
	QrPayMode   string `mapstructure:"qrpay-mode"`   // PC扫码方式
	QrcodeWidth string `mapstructure:"qrcode_width"` // 自定义二维码宽度,qr_pay_mode=4时该参数有效
	IsProd      bool   `mapstructure:"is-prod"`      //是否是正式环境

	// URL配置
	NotifyURL string `mapstructure:"notify-url"` // 支付回调通知地址
}

// Business 系统业务配置
type Business struct {
	MPSExchangeRate   float64 `mapstructure:"mps-exchange-rate"`   // MPS兑换比率
	OrderTimeout      int     `mapstructure:"order-timeout"`       // 订单超时时间（分）
	MinRechargeAmount float64 `mapstructure:"min-recharge-amount"` // 最小充值金额
	MaxRechargeAmount float64 `mapstructure:"max-recharge-amount"` // 最大充值金额
	AccountType       int8    `mapstructure:"account-type"`        // 收款账户类型。 1：对公（在金融机构开设的公司账户） 2：对私（在金融机构开设的个人账户）
	InstName          string  `mapstructure:"inst-name"`           //#银行卡卡开户银行
	RegisterMPSAmount float64 `mapstructure:"register-mps-amount"` //充值赠送MPS数量
}
