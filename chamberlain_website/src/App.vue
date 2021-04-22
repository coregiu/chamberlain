<template>
  <div class="view">
    <div id="top">
      <a href="#"><img src="./assets/regiu-logo.png" id="logo"/></a>
      <div>
        <router-link to="/blog" class="menu"><img src="./assets/svgs/blog.svg"/>博客</router-link>
      </div>
      <div>
        <router-link to="/notebook" class="menu"><img src="./assets/svgs/notebook.svg"/>记事本</router-link>
      </div>
      <div class="group">
        <img src="./assets/svgs/input.svg" class="subMenu">收入管理
        <ul>
          <li>
            <router-link to="/inputs_details" class="subMenu">收入明细</router-link>
          </li>
          <li>
            <router-link to="/inputs_analysis" class="subMenu">收入分析</router-link>
          </li>
        </ul>
      </div>
      <div>
        <img src="./assets/svgs/admin.svg"/>系统管理
        <ul>
          <li>
            <router-link to="/urmgmt" class="subMenu">用户管理</router-link>
          </li>
          <li>
            <router-link to="/syslog" class="subMenu">操作日志</router-link>
          </li>
          <li>
            <router-link to="/sysback" class="subMenu">备份恢复</router-link>
          </li>
        </ul>
      </div>
      <div>
        <router-link to="/login" class="menu"><img src="./assets/svgs/login.svg"/>登录</router-link>
      </div>
      <div>
        <img src="./assets/svgs/accout.svg"/>个人信息
        <ul>
          <li>
            <router-link to="/userinfo" class="subMenu">个人信息</router-link>
          </li>
          <li>
            <router-link to="/updatepassword" class="subMenu">修改密码</router-link>
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
    <div id="foot">©2021 REGIU | <a href="https://gitee.com/regiu" target="_blank">https://gitee.com/regiu</a> |
      火在水上未济，君子以慎辨物居方。
    </div>
  </div>
</template>