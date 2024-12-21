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
            try {
                // 检查缓存中的菜单数据
                const menuStr = localStorage.getItem('menu');
                if (!menuStr) {
                    console.warn('No menu data in localStorage');
                    return;
                }

                // 解析菜单数据
                let menu;
                try {
                    menu = JSON.parse(menuStr);
                } catch (error) {
                    console.error('Failed to parse menu data:', error);
                    return;
                }

                // 验证菜单数据
                if (!Array.isArray(menu)) {
                    console.error('Menu data is not an array:', menu);
                    return;
                }

                // 更新状态
                state.menu = menu;
                const menuArray = [];

                // 处理菜单项
                for (const item of menu) {
                    if (item.children) {
                        // 处理子菜单
                        item.children = item.children.map(child => {
                            if (!child.url) {
                                console.warn('Child menu item missing url:', child);
                                return child;
                            }
                            
                            try {
                                child.component = () => import(/* webpackChunkName: "center" */ `../views/center/${child.url}`);
                            } catch (error) {
                                console.error(`Failed to load component for ${child.url}:`, error);
                            }
                            return child;
                        });
                        menuArray.push(...item.children);
                    } else if (item.url) {
                        // 处理主菜单
                        try {
                            item.component = () => import(/* webpackChunkName: "center" */ `../views/center/${item.url}`);
                            menuArray.push(item);
                        } catch (error) {
                            console.error(`Failed to load component for ${item.url}:`, error);
                        }
                    }
                }

                // 添加路由
                menuArray.forEach(item => {
                    if (item.path && item.component) {
                        try {
                            router.addRoute('Center', item);
                        } catch (error) {
                            console.error(`Failed to add route for ${item.path}:`, error);
                        }
                    } else {
                        console.warn('Invalid route item:', item);
                    }
                });

            } catch (error) {
                console.error('Error in addMenu:', error);
            }
        }
    },
}