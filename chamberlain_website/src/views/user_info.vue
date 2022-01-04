<template>
  <div class="user-info-div">
    <table class="user-info-table">
      <tr>
        <td align="right"><p><b>Username:</b></p></td>
        <td align="left"><p><b> {{ username }}</b></p></td>
      </tr>
      <tr>
        <td align="right"><p><b>Role:</b></p></td>
        <td align="left"><p><b> {{ role }}</b></p></td>
      </tr>
      <tr>
        <td align="right"><p><b>Token:</b></p></td>
        <td align="left"><p><b> {{ token }}</b></p></td>
      </tr>
      <tr>
        <td align="right"><p><b>Token Expire Time:</b></p></td>
        <td align="left"><p><b> {{ tokenExpire }}</b></p></td>
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
      username: "",
      role: "",
      token: "",
      tokenExpire: ""
    }
  },
  async created() {
    let userInfo = await getUserByToken();
    this.username = userInfo.User.Username
    this.role = userInfo.User.Role
    this.token = userInfo.TokenId
    this.tokenExpire = userInfo.ExpireTime
  }
}
</script>

<style>
.user-info-div{
  position: relative;
  top:25%;
}
.user-info-table {
  text-align: center;
  margin:0 auto;
  font-size:20px;
}
</style>