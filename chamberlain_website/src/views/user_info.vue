<template>
  <div id="user_info">
    <br><br><br><br><br><br>
    <p><b>Username: {{ username }}</b></p><br><br>
    <p><b>Role: {{ role }}</b></p><br><br>
    <p><b>Token: {{ token }}</b></p><br><br>
    <p><b>Token Expire Time: {{ tokenExpire }}</b></p>
  </div>
</template>

<script>
import {getUserByToken, login} from "../api/api";
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
    this.username =userInfo.User.Username
    this.role =userInfo.User.Role
    this.token =userInfo.TokenId
    this.tokenExpire =userInfo.ExpireTime
  }
}
</script>

<style>
#user_info {
  width: 60%;
  height: 87vh;
  background-color: #fff;
  text-align: left;
  overflow: scroll;
  margin: 0 auto;
  padding-left: 200px;
}
</style>