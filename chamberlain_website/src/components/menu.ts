import {token} from "./token";

const UN_LOGIN_MENU = [
    {
        "name": "🖍博客 ▽",
        "image": "../../src/assets/svgs/blog.svg",
        "url": "-",
        "subMenu": [
            {
                "name": "🕸 人生哲学",
                "url": "/philosophy"
            },
            {
                "name": "🖥 技术总结",
                "url": "/summary"
            }
        ]
    }, {
        "name": "🔐登录",
        "image": "../../src/assets/svgs/login.svg",
        "url": "/login"
    }]

const USER_LOGIN_MENU = [
    {
        "name": "🖍博客 ▽",
        "image": "../../src/assets/svgs/blog.svg",
        "url": "-",
        "subMenu": [
            {
                "name": "🕸 人生哲学",
                "url": "/philosophy"
            },
            {
                "name": "🖥 技术总结",
                "url": "/summary"
            }
        ]
    }, {
        "name": "📙记事本",
        "image": "../../src/assets/svgs/notebook.svg",
        "url": "/notebook"
    }, {
        "name": "👨个人信息 ▽",
        "image": "../../src/assets/svgs/accout.svg",
        "url": "-",
        "subMenu": [
            {
                "name": "👨 个人信息",
                "url": "/user_info"
            },
            {
                "name": "🔑 修改密码",
                "url": "/update_password"
            },
            {
                "name": "🚀 退出系统",
                "url": "/logout"
            }
        ]
    }]

const ADMIN_LOGIN_MENU = [
    {
        "name": "🖍博客 ▽",
        "image": "../../src/assets/svgs/blog.svg",
        "url": "-",
        "subMenu": [
            {
                "name": "🕸 人生哲学",
                "url": "/philosophy"
            },
            {
                "name": "🖥 技术总结",
                "url": "/summary"
            }
        ]
    }, {
        "name": "📙记事本",
        "image": "../../src/assets/svgs/notebook.svg",
        "url": "/notebook"
    }, {
        "name": "📊️收入管理 ▽",
        "image": "../../src/assets/svgs/input.svg",
        "url": "-",
        "subMenu": [
            {
                "name": "📝 收入明细",
                "url": "/input_details"
            },
            {
                "name": "📊️ 收入分析",
                "url": "/input_analysis"
            }
        ]
    }, {
        "name": "🛠系统管理 ▽",
        "image": "../../src/assets/svgs/admin.svg",
        "url": "-",
        "subMenu": [
            {
                "name": "🛠 用户管理",
                "url": "/user_mgmt"
            },
            {
                "name": "💾 操作日志",
                "url": "/sys_log"
            },
            {
                "name": "🖱 系统检查",
                "url": "/sys_back"
            }
        ]
    },{
        "name": "👨个人信息 ▽",
        "image": "../../src/assets/svgs/accout.svg",
        "url": "-",
        "subMenu": [
            {
                "name": "👨 个人信息",
                "url": "/user_info"
            },
            {
                "name": "🔑 修改密码",
                "url": "/update_password"
            },
            {
                "name": "🚀 退出系统",
                "url": "/logout"
            }
        ]
    }]

export const menu = {
    methods: {
        getMenus() {
            let role = token.methods.getRole()
            if (role === "user") {
                return USER_LOGIN_MENU
            } else if (role === "admin") {
                return ADMIN_LOGIN_MENU
            } else {
                return UN_LOGIN_MENU
            }
        },
        setMenuTop() {
            let menuList = this.getMenus()
            let topElement = document.getElementById("menuTop")
            while (topElement.lastChild) {
                topElement.removeChild(topElement.lastChild);
            }

            menuList.forEach(function (menu) {
                let divElement = document.createElement("div")
                if (menu.url !== "-") {
                    let routerElement = document.createElement("a")
                    routerElement.setAttribute("href", "#" + menu.url)
                    routerElement.setAttribute("class", "menu")
                    routerElement.innerText = menu.name
                    divElement.appendChild(routerElement)
                } else {
                    divElement.innerText = menu.name
                    divElement.setAttribute("class", "group")

                    let ulElement = document.createElement("ul")
                    menu.subMenu.forEach(function (subMenu) {
                        let liElement = document.createElement("li")
                        let hrefElement = document.createElement("a")
                        hrefElement.setAttribute("class", "subMenu")
                        hrefElement.setAttribute("href", "#" + subMenu.url)
                        hrefElement.innerText = subMenu.name
                        liElement.appendChild(hrefElement)
                        ulElement.appendChild(liElement)
                    })
                    divElement.appendChild(ulElement)
                }

                topElement.appendChild(divElement)
            })
        }
    }
}