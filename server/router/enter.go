package router

// RouterGroup 路由组汇集，一个结构体对应一组路由接口
type RouterGroup struct {
	BaseRouter       // 基本路由组（无需鉴权）
	UserRouter       // 用户相关路由组
	MenuRouter       // 菜单相关路由组
	AuthorityRouter  // 角色相关路由组
	CommitteeRouter  // 委员会相关路由组
	ConferenceRouter // 委员会相关路由组
	JournalRouter    // 期刊相关路由组
	PaperRouter      // 论文相关路由组
	ReviewRouter     // 审核相关路由组
	MPSRouter
}

// RouterGroupApp 初始化所有路由组结构体
var RouterGroupApp = new(RouterGroup)
