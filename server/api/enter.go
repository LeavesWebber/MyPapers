package api

// ApiGroup 一个结构体对应一类操作的方法
type ApiGroup struct {
	UserApi          // 用户相关方法
	AuthorityApi     // 角色相关方法
	AuthorityMenuApi // 菜单相关方法
	CommitteeApi     // 委员相关方法
	ConferenceApi    // 会议相关方法
	JournalApi       // 期刊相关方法
	PaperApi         // 论文相关方法
	ReviewApi        // 审核相关方法
}

var ApiGroupApp = new(ApiGroup)
