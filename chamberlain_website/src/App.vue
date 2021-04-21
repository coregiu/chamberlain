<template>
  <div class="view">
    <div id="top">
      <a href="#"><img src="./assets/regiu-logo.png" id="logo"/></a>
      <div>
        <router-link to="/blog"><img src="./assets/svgs/blog.svg"/>博客</router-link>
      </div>
      <div>
        <router-link to="/notebook"><img src="./assets/svgs/notebook.svg"/>记事本</router-link>
      </div>
      <div class="group">
        <img src="./assets/svgs/input.svg">收入管理
        <ul>
          <li>
            <router-link to="/inputs_details">收入明细</router-link>
          </li>
          <li>
            <router-link to="/inputs_analysis">收入分析</router-link>
          </li>
        </ul>
      </div>
      <div>
        <img src="./assets/svgs/admin.svg"/>系统管理
        <ul>
          <li>
            <router-link to="/urmgmt">用户管理</router-link>
          </li>
          <li>
            <router-link to="/syslog">操作日志</router-link>
          </li>
          <li>
            <router-link to="/sysback">备份恢复</router-link>
          </li>
        </ul>
      </div>
      <div>
        <router-link to="/login"><img src="./assets/svgs/login.svg"/>登录</router-link>
      </div>
      <div>
        <img src="./assets/svgs/accout.svg"/>个人信息
        <ul>
          <li>
            <router-link to="/userinfo">个人信息</router-link>
          </li>
          <li>
            <router-link to="/updatepassword">修改密码</router-link>
          </li>
          <li>
            <router-link to="/logout">🚀 退出系统</router-link>
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