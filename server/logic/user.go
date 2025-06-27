package logic

import (
	"context"
	"fmt"
	"path/filepath"
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"server/utils"

	"errors"
	"strconv"
	"time"

	"math/big" // 用于 ChainID
	"server/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // 用于 TransactOpts
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"

	//"crypto/ecdsa" // 用于私钥
	"github.com/ethereum/go-ethereum/crypto" // 用于 HexToECDSA
)

// Login 用户登录
func Login(in *request.Login) (out *response.Login, err error) {
	// 查询数据库
	if out, err = mysql.Login(in); err != nil {
		return nil, err
	}
	// 签发token
	token, err := utils.GenToken(out.UserInfo)
	if err != nil {
		return
	}
	out.Token = token
	// 返回信息
	return out, err
}

// Register 创建用户
func Register(in *request.Register) (err error, id int64) {
	// 判断用户名是已经否存在
	if _, err = mysql.UserNameExist(in.Username); err == global.ErrorUserExist {
		return
	}
	//判断邮箱是否存在
	// if _, err = mysql.UserEmailExist(in.Email); err == global.ErrorUserEmailExist {
	// 	return
	// }
	// 创建用户
	user := &tables.User{
		Sex:               in.Sex,
		Username:          in.Username,
		Password:          in.Password,
		FirstName:         in.FirstName,
		LastName:          in.LastName,
		HeaderImg:         in.HeaderImg,
		Email:             in.Email,
		Department:        in.Department,
		Phone:             in.Phone,
		Address:           in.Address,
		Education:         in.Education,
		Title:             in.Title,
		Research:          in.Research,
		BlockChainAddress: in.BlockChainAddress,
		Affiliation:       in.Affiliation,
		AffiliationType:   in.AffiliationType,
	}
	// 生成uuid
	user.UUID = utils.GenID()
	// 存入数据库
	return mysql.Register(user), user.UUID
}

// SendMail 发送邮件验证码
func SendMail(in *request.SendMail) error {
	// 检查Redis中邮箱验证码的TTL，以防止重复发送
	get := global.MPS_REDIS.Do(context.Background(), "TTL", global.REDIS_SMTP_PREFIX+in.Email)
	ttlResult, err := get.Result()
	if err != nil {
		global.MPS_LOG.Error("failed to check TTL in Redis", zap.String("email", in.Email), zap.Error(err))
		return fmt.Errorf("failed to check TTL in Redis: %v", err)
	}
	// 处理 TTL 返回值
	ttl, ok := ttlResult.(int64)
	if ok && ttl > 0 { // 键不存在
		if int64(global.SMTP_EXPIRED_TIME.Seconds())-ttl <= int64(global.SMTP_RETRY_TIME.Seconds()) {
			return global.ErrorInvalidEmailReSend{}
		}
	}

	// 生成6位数字验证码
	code, err := utils.GenerateRandomNumericCode(6)
	if err != nil {
		global.MPS_LOG.Error("failed to generate random numeric code", zap.Error(err))
		return fmt.Errorf("failed to generate random numeric code: %v", err)
	}

	// 输出验证码到控制台（仅在开发环境使用）
	global.MPS_LOG.Info("Email verification code",
		zap.String("email", in.Email),
		zap.String("code", code))

	// 构建邮件内容
	m := gomail.NewMessage()
	m.SetHeader("From", global.MPS_CONFIG.Smtp.Username)
	m.SetHeader("To", in.Email)
	m.SetHeader("Subject", "验证码")
	msg := fmt.Sprintf("您的验证码为: %s", code)
	m.SetBody("text/html", msg)

	// 连接SMTP服务器
	d := gomail.NewDialer(global.MPS_CONFIG.Smtp.Host, global.MPS_CONFIG.Smtp.Port, global.MPS_CONFIG.Smtp.Username, global.MPS_CONFIG.Smtp.Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		global.MPS_LOG.Error("SendMail failed", zap.Error(err))
		return err
	}

	// 将验证信息存入Redis
	verificationInfo := map[string]interface{}{
		"code":       code,
		"attempts":   0,
		"verified":   false,
		"created_at": time.Now().Unix(),
	}

	if err := global.MPS_REDIS.HSet(context.Background(), global.REDIS_SMTP_PREFIX+in.Email, verificationInfo).Err(); err != nil {
		global.MPS_LOG.Error("redis HSet failed", zap.Error(err))
		return err
	}

	// 设置过期时间
	if err := global.MPS_REDIS.Expire(context.Background(), global.REDIS_SMTP_PREFIX+in.Email, global.SMTP_EXPIRED_TIME).Err(); err != nil {
		global.MPS_LOG.Error("redis Expire failed", zap.Error(err))
		return err
	}

	return nil
}

func VerifyMail(in *request.VerifyMail) (*request.VerifyMailResponse, error) {
	// 获取验证信息
	key := global.REDIS_SMTP_PREFIX + in.Email
	info, err := global.MPS_REDIS.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	if len(info) == 0 {
		return nil, global.ErrorInvalidEmailCode{}
	}

	// 检查尝试次数
	attempts, _ := strconv.Atoi(info["attempts"])
	if attempts >= 5 {
		return nil, errors.New("验证尝试次数过多")
	}

	// 检查验证码是否匹配
	if info["code"] != in.Code {
		// 增加尝试次数
		global.MPS_REDIS.HIncrBy(context.Background(), key, "attempts", 1)
		return nil, global.ErrorInvalidEmailCode{}
	}

	// 检查是否已验证
	if info["verified"] == "true" {
		return nil, errors.New("验证码已被使用")
	}

	// 标记为已验证
	global.MPS_REDIS.HSet(context.Background(), key, "verified", true)

	// 生成验证token
	expiresAt := time.Now().Add(time.Hour).Unix()

	// 创建一个临时用户对象用于生成token
	tempUser := tables.User{
		Email: in.Email,
		UUID:  utils.GenID(),
	}

	token, err := utils.GenToken(tempUser)
	if err != nil {
		return nil, err
	}

	// 存储token
	global.MPS_REDIS.Set(context.Background(),
		global.REDIS_SMTP_PREFIX+"token:"+in.Email,
		token,
		time.Hour)

	return &request.VerifyMailResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// // GetUserTree 获取用户树
//
//	func GetUserTree() (users []tables.User, err error) {
//		userTree, err := mysql.GetUserTreeMap() // 获取user的父子对应关系（此时还是乱序的）
//		users = userTree[0]
//		for i := 0; i < len(users); i++ {
//			err = getUserChildrenList(&users[i], userTree) // 从根节点开始遍历整理user树（从根节点开始有序）
//		}
//		return users, err
//	}

// GetAllUser 获取所有用户
func GetAllUser() (users []response.GetAllUser, err error) {
	return mysql.GetAllUser() // 获取user的父子对应关系（此时还是乱序的）

}

// // getChildrenList 生成一颗关系树
//
//	func getUserChildrenList(user *tables.User, treeMap map[uint][]tables.User) (err error) {
//		user.Children = treeMap[user.ID]
//		for i := 0; i < len(user.Children); i++ {
//			err = getUserChildrenList(&user.Children[i], treeMap)
//		}
//		return err
//	}
//
// GetSelfInfo 获取本用户信息
func GetSelfInfo(uuid int64) (userInfo tables.User, err error) {
	return mysql.GetSelfInfo(uuid)
}

// ChangePassword 修改密码
func ChangePassword(in *request.ChangePassword, uuid int64) error {
	return mysql.ChangePassword(in, uuid)
}

// ChangeHeaderImg 修改头像
func ChangeHeaderImg(c *gin.Context, uuid int64, in *request.ChangeHeaderImg) (url string, err error) {
	// 2. 生成文件名和保存路径
	filename := filepath.Base(in.Data.Filename)
	finalName := fmt.Sprintf("%d_%s", uuid, filename)
	saveFile := filepath.Join("./public/images/avatar/", finalName)
	// 保存文件
	if err = c.SaveUploadedFile(in.Data, saveFile); err != nil {
		global.MPS_LOG.Error("SaveUploadedFile failed", zap.Error(err))
		return
	}
	// 3. 保存路径到数据库
	url = fmt.Sprintf("/public/images/avatar/%s", finalName)
	if err = mysql.ChangeHeaderImg(uuid, url); err != nil {
		return
	}
	return url, nil
}

//
//// ResetPassword 重置密码
//func ResetPassword(uuid int64) error {
//	return mysql.ResetPassword(uuid)
//}
//
//// DeleteUser 删除用户
//func DeleteUser(id uint) error {
//	return mysql.DeleteUser(id)
//}

// SetUserAuthorities 设置用户权限组
func SetUserAuthorities(id uint, authorityIds []uint) error {
	return mysql.SetUserAuthorities(id, authorityIds)
}

// SetUserInfo 设置用户信息
func SetUserInfo(in *request.SetUserInfo) (err error) {
	return mysql.SetUserInfo(&tables.User{
		MPS_MODEL: global.MPS_MODEL{
			ID: in.ID,
		},
		AuthorityId:       in.AuthorityId,
		Sex:               in.Sex,
		Username:          in.Username,
		FirstName:         in.FirstName,
		LastName:          in.LastName,
		HeaderImg:         in.HeaderImg,
		Email:             in.Email,
		Department:        in.Department,
		Phone:             in.Phone,
		Address:           in.Address,
		Education:         in.Education,
		Title:             in.Title,
		Research:          in.Research,
		BlockChainAddress: in.BlockChainAddress,
		Affiliation:       in.Affiliation,
		AffiliationType:   in.AffiliationType,
	})
}

//// SetUserAuthority 切换角色
//func SetUserAuthority(user utils.BaseClaims, authorityId uint) (token string, err error) {
//	if err = mysql.SetUserAuthority(user.ID, authorityId); err != nil {
//		return
//	}
//	// 签发token
//	if token, err = utils.GenToken(tables.User{
//		MPS_MODEL: global.MPS_MODEL{
//			ID: user.ID,
//		},
//		UUID:        user.UUID,
//		Username:    user.Username,
//		AuthorityId: authorityId,
//		Name:        user.Name,
//	}); err != nil {
//		return
//	}
//	return
//}

// func RegisterSendMPS(address string) {

// }

// RegisterUserOnBlockchain 在区块链上注册用户并更新本地数据库
// 注意：这是一个示例实现，具体细节取决于你的智能合约和错误处理策略
func RegisterUserOnBlockchain(req request.Register) (*tables.User, string, error) {
	// 1. (可选) 先在本地数据库创建用户或部分信息，或在合约调用成功后创建/更新
	// ...

	// 2. 连接到以太坊客户端
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL) // 从配置中获取 RPC URL
	if err != nil {
		return nil, "", fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// 3. 加载用于签名的私钥 (应安全存储和管理)
	privateKeyHex := global.MPS_CONFIG.Blockchain.AdminPrivateKey // 从配置中获取私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, "", fmt.Errorf("failed to load private key: %w", err)
	}

	// 4. 获取合约地址
	contractAddress := common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress) // 从配置中获取合约地址

	// 5. 准备交易选项
	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID) // 从配置中获取 ChainID 并转换为 big.Int
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = uint64(global.MPS_CONFIG.Blockchain.GasLimit) // 从配置中读取并设置 GasLimit

	// auth.Nonce = ... // 可能需要手动管理 Nonce
	// auth.Value = ... // 如果需要发送 ETH
	// auth.GasPrice = ... // 对于旧版交易或特定网络，可能需要设置这个
	// 对于 EIP-1559 交易，go-ethereum 库会尝试自动从节点获取 MaxPriorityFeePerGas 和 MaxFeePerGas

	// 6. 调用智能合约的注册函数
	// 假设合约函数是 RegisterUser(address userAddr)
	// req.BlockChainAddress 应该是用户在区块链上的地址
	// req.Username 可以是传递给合约的某个用户标识
	userBlockchainAddress := common.HexToAddress(req.BlockChainAddress)

	// 此处假设合约的注册函数名为 `RegisterUser`，根据新的合约定义进行调用
	// 使用 MPSTransactor 而不是 MPSCaller 来调用写入方法
	mpsContract, err := contracts.NewMPS(contractAddress, client)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create contract instance: %w", err)
	}
	tx, err := mpsContract.RegisterUser(auth, userBlockchainAddress)
	if err != nil {
		return nil, "", fmt.Errorf("failed to call RegisterUser on smart contract: %w", err)
	}

	txHash := tx.Hash().Hex()
	global.MPS_LOG.Info(fmt.Sprintf("Blockchain registration transaction sent: %s", txHash))
	// 添加对交易对象的详细日志输出
	global.MPS_LOG.Info("Blockchain transaction details",
		zap.String("txHash", txHash),
		zap.Uint64("txNonce", tx.Nonce()),
		zap.Stringer("txTo", tx.To()), // tx.To() 返回 *common.Address，它实现了 Stringer 接口
		zap.Stringer("txValue", tx.Value()),
		zap.Uint64("txGasLimit", tx.Gas()),
		zap.Stringer("txGasPrice", tx.GasPrice()), // 对于EIP-1559交易，可能是GasFeeCap()和GasTipCap()
		zap.Any("txData", tx.Data()),              // 交易数据，通常是合约调用的编码
	)

	// 7. (可选) 等待交易被打包确认
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, "", fmt.Errorf("failed waiting for transaction to be mined: %w", err)
	}
	if receipt.Status == 0 {
		return nil, "", fmt.Errorf("transaction failed on blockchain (receipt status 0)")
	}

	// 8. 如果合约调用成功，并且需要在本地数据库中存储/更新用户信息
	// 例如，确认用户的 BlockChainAddress，或存储交易哈希等
	var user tables.User
	// 示例：根据 req.Username 或其他唯一标识查找用户，然后更新
	// err = global.MPS_DB.Where("username = ?", req.Username).First(&user).Error
	// if err != nil { ... }
	// user.BlockChainAddress = req.BlockChainAddress // 确保已设置
	// user.BlockchainTxHash = txHash // 如果有这样的字段
	// err = global.MPS_DB.Save(&user).Error
	// if err != nil { ... }

	// 这里的返回值根据你的实际需求调整
	return &user, txHash, nil
}
