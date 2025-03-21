export default {
    state: {
        isCollapse: false, // 控制菜单的展开还是收起
        menu: [
            // {
            //     path: "/center",
            //     name: "information",
            //     label: "Information",
            //     icon: "user",
            //     url: "Center",
            // },
        ]
    },
    mutations: {
        // 修改菜单展开收起的方法
        collapseMenu(state) {
            state.isCollapse = !state.isCollapse
        },
        // 设置menu的数据
        setMenu(state, val) {
            state.menu = val
            // Cookies.set('menu', JSON.stringify(val))
            localStorage.setItem("menu", JSON.stringify(val));
        },
        // 动态注册路由
        addMenu(state, router) {
            // 判断当前缓存中是否有数据
            // if (!Cookies.get('menu')) return
            if (!localStorage.getItem('menu')) return
            // const menu = JSON.parse(Cookies.get('menu')) // 转成数组
            const menu = JSON.parse(localStorage.getItem('menu')) // 转成数组
            state.menu = menu // 更新state中的数据
            // 组装动态路由的数据
            console.log(menu, 'menu')
            const menuArray = []
            menu.forEach(item => { // 遍历数组
                if (item.children) {
                    item.children = item.children.map(item => {
                        item.component = () => import(`../views/center/${item.url}`)
                        return item
                    })
                    menuArray.push(...item.children)
                } else {
                    item.component = () => import(`../views/center/${item.url}`)
                    menuArray.push(item)
                }
            });
            console.log(menuArray, 'menuArray')
            // 路由的动态添加
            menuArray.forEach(item => {
                router.addRoute('Center', item)
            })
        }
    },
    getters: {
        isLoggedIn: () => {
            return !!localStorage.getItem('token')
        }
    }
}