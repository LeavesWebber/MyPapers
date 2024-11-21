package utils

import (
	"server/global"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// Casbin 初始化
func Casbin() *casbin.SyncedEnforcer {
	// https://www.jianshu.com/p/9506406e745f
	// https://www.cnblogs.com/tomtellyou/p/13489145.html
	// https://www.cnblogs.com/Mail-maomao/p/11951482.html（√） 有较详细的说明参考着理解
	once.Do(func() { // 用的是最简单的模型，没有涉及到g
		a, _ := gormadapter.NewAdapterByDB(global.MPS_DB) // 通过现有的gorm实例创建gorm适配器 // 将数据库连接同步给插件， 插件用来操作数据库
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text) // 从包含模型文本的字符串创建模型 // 从CSV文件adapter加载策略规则
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(m, a) // 通过文件或数据库创建同步执行器 // 在多线程环境下使用Enforcer对象的接口，必须使用casbin.NewSyncedEnforcer创建Enforcer
	})
	_ = syncedEnforcer.LoadPolicy() // 加载数据库中的策略 // 把数据库记录的东西读出来
	return syncedEnforcer
}
