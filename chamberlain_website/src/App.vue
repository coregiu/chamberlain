<template>
  <div class="view">
    <a href="#"><img src="./assets/coregiu-logo.png" id="logo"/></a>
    <div id="unLoginTop" class="top">
      <div class="group">
        <img src="./assets/svgs/blog.svg" class="subMenu"/>博客<img src="./assets/svgs/caret-down.svg"/>
        <ul>
          <li>
            <router-link to="/philosophy" class="subMenu">人生哲学</router-link>
          </li>
          <li>
            <router-link to="/summary" class="subMenu">技术总结</router-link>
          </li>
        </ul>
      </div>
      <div>
        <router-link to="/login" class="menu"><img src="./assets/svgs/login.svg"/>登录</router-link>
      </div>
    </div>

    <div id="loginTop" class="top">
      <div class="group">
        <img src="./assets/svgs/blog.svg" class="subMenu"/>博客<img src="./assets/svgs/caret-down.svg"/>
        <ul>
          <li>
            <router-link to="/philosophy" class="subMenu">人生哲学</router-link>
          </li>
          <li>
            <router-link to="/summary" class="subMenu">技术总结</router-link>
          </li>
        </ul>
      </div>
      <div>
        <router-link to="/notebook" class="menu"><img src="./assets/svgs/notebook.svg"/>记事本</router-link>
      </div>
      <div class="group">
        <img src="./assets/svgs/input.svg" class="subMenu"/>收入管理<img src="./assets/svgs/caret-down.svg"/>
        <ul>
          <li>
            <router-link to="/input_details" class="subMenu">收入明细</router-link>
          </li>
          <li>
            <router-link to="/input_analysis" class="subMenu">收入分析</router-link>
          </li>
        </ul>
      </div>
      <div>
        <img src="./assets/svgs/admin.svg"/>系统管理<img src="./assets/svgs/caret-down.svg"/>
        <ul>
          <li>
            <router-link to="/user_mgmt" class="subMenu">用户管理</router-link>
          </li>
          <li>
            <router-link to="/sys_log" class="subMenu">操作日志</router-link>
          </li>
          <li>
            <router-link to="/sys_back" class="subMenu">备份恢复</router-link>
          </li>
        </ul>
      </div>
      <div>
        <img src="./assets/svgs/accout.svg"/>个人信息<img src="./assets/svgs/caret-down.svg"/>
        <ul>
          <li>
            <router-link to="/user_info" class="subMenu">个人信息</router-link>
          </li>
          <li>
            <router-link to="/update_password" class="subMenu">修改密码</router-link>
          </li>
          <li>
            <router-link to="/logout" class="subMenu">🚀 退出系统</router-link>
          </li>
        </ul>
      </div>
    </div>

    <div id="content">
      <router-view/>
    </div>
    <div id="foot">©2021 COREGIU | <a href="https://github.com/coregiu" target="_blank">https://github.com/coregiu</a> |
      火在水上未济，君子以慎辨物居方。
    </div>
  </div>
</template>