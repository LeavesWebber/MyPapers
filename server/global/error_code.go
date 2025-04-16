package global

import (
	"errors"
)

var (
	ErrorUserExist         = errors.New("用户名已存在")
	ErrorUserNotExist      = errors.New("用户名不存在")
	ErrorUserEmailExist    = errors.New("邮箱已存在")
	ErrorUserEmailNotExist = errors.New("邮箱不存在")
	ErrUserHasChild        = errors.New("此用户存在子用户不可删除")
	ErrorAuthorityNoExist  = errors.New("角色不存在")
	ErrorAuthorityUsing    = errors.New("此角色有用户正在使用禁止删除")
	ErrorAuthorityHasChild = errors.New("此角色存在子角色不允许删除")
	ErrorInvalidPassword   = errors.New("密码错误")
	ErrRoleExistence       = errors.New("存在相同角色id")
	ErrUserNoAuthority     = errors.New("该用户无此角色")

	ErrMenuExistence     = errors.New("存在重复name，请修改name")
	ErrMenuHasChild      = errors.New("此菜单存在子菜单不可删除")
	ErrPaperReviewed     = errors.New("该论文已审核")
	ErrPaperNotAccepted  = errors.New("该论文未被接受")
	ErrMinRechargeAmount = errors.New("充值金额过小")
	ErrMaxRechargeAmount = errors.New("充值金额过大")

	ErrIssueHasPaper = errors.New("Issue下存在论文不可删除")
)

type ErrorInvalidEmailReSend struct{}

func (e ErrorInvalidEmailReSend) Error() string {
	return "邮件重发间隔过短"
}

type ErrorInvalidEmailCode struct{}

func (e ErrorInvalidEmailCode) Error() string {
	return "邮箱验证码错误"
}
