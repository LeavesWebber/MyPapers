package logic

import (
	"errors"
	"server/model/request"
	"server/utils"
	"strconv"
)

// UpdateCasbin 更新casbin权限
func UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := utils.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// ClearCasbin 清除匹配的权限
func ClearCasbin(v int, p ...string) bool {
	e := utils.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
