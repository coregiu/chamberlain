<template>
  <div id="login_div">
    用户 <InputText type="text" v-model="username" required="true" autofocus @keyup.enter.native="doLogin()"/><br><br>
    密码 <InputText type="password" v-model="password" required="true" autofocus @keyup.enter.native="doLogin()"/><br><br>
    <Button label="submit" @click="doLogin()" v-bind:disabled="disableBn">登录</Button><br>
    <Dialog v-model:visible="display" header="登录失败">{{message}}</Dialog>
  </div>
</template>

<script>
import {login} from '../api/system_api'
export default {
  name: "login",
  data() {
    return {
      username: "",
      password: "",
      disableBn: false,
      display: false,
      message: ""
    }
  },
  methods: {
    async doLogin() {
      this.disableBn = true;
      let res = await login(this.username, this.password);
      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this.display = true;
        this.message = "感谢抬爱，请正确输入用户名和密码！";
      } else {
        this.$token.methods.setChamberlainToken(res.TokenId, res.User.Role);
        this.$menu.methods.setMenuTop()
        this.$router.push("/");
      }
      this.disableBn = false;
    }
  }
}
</script>

<style scoped>
#login_div {
  position: relative;
  top: 25%;
  text-align: center;
}
</style>