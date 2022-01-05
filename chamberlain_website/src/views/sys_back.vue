<template>
  <div class="system-info-div">
    <span style="font-size:28px;color: blue"><p><b>结论： 系统当前运行状态良好</b></p></span>
    <table class="system-info-table">
      <tr>
        <td align="right"><p><b>剩余磁盘空间:</b></p></td>
        <td align="left"><p><b> {{ diskRemain }}</b></p></td>
      </tr>
      <tr>
        <td align="right"><p><b>平均CPU使用率:</b></p></td>
        <td align="left"><p><b> {{ cpuUsage }}</b></p></td>
      </tr>
      <tr>
        <td align="right"><p><b>平均内存使用:</b></p></td>
        <td align="left"><p><b> {{ memUsed }}</b></p></td>
      </tr>
      <tr>
        <td align="right"><p><b>网络时延:</b></p></td>
        <td align="left"><p><b> {{ netDelay }}</b></p></td>
      </tr>
      <tr>
        <td align="right"><p><b>数据库状态:</b></p></td>
        <td align="left"><p><b> {{ dbStatus }}</b></p></td>
      </tr>
    </table>
  </div>
</template>

<script>
import {getUserByToken} from "../api/system_api";

export default {
  name: "user_info",
  data() {
    return {
      diskRemain: "100G",
      cpuUsage: "58%",
      memUsed: "2048M",
      netDelay: "正常",
      dbStatus: "正常"
    }
  },
  created() {

  }
}
</script>

<style>
.system-info-div{
  position: relative;
  top:10%;
}
.system-info-table {
  text-align: center;
  margin:0 auto;
  font-size:20px;
  color: #42A5F5;
}
</style>