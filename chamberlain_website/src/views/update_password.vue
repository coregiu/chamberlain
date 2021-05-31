<template>
  <div>
    用户 <InputText type="text" v-model="username"/><br><br>
    旧密码 <InputText type="password" v-model="password"/><br><br>
    设置新密码 <InputText type="password" v-model="newPassword"/><br><br>
    确认新密码 <InputText type="password" v-model="newPasswordVery"/><br><br>
    <Button label="submit" @click="doSetPassword()" v-bind:disabled="disableBn">修改密码</Button><br>
    <Dialog v-model:visible="display" header="修改密码失败">{{message}}</Dialog>
  </div>
</template>

<script>
import {getUserByToken} from '../api/api'
import {resetPassword} from '../api/api'
export default {
  name: "reset_password",
  data() {
    return {
      username: "",
      password: "",
      newPassword: "",
      newPasswordVery: "",
      disableBn: false,
      display: false,
      message: ""
    }
  },
  async created() {
    let userInfo = await getUserByToken();
    this.username =userInfo.User.Username
  },
  methods: {
    async doSetPassword() {
      this.disableBn = true;
      if (this.newPassword !== this.newPasswordVery) {
        this.display = true;
        this.message = "两次密码输入不一致，请重新输入密码！";
        this.disableBn = false;
        return
      }
      let res = await resetPassword(this.username, this.password, this.newPassword);
      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this.display = true;
        this.message = "修改失败，请正确输入！";
      } else {
        this.$token.methods.setChamberlainToken(res.TokenId);
        this.$token.methods.showLoginMenu();
        this.$router.push("/");
      }
      this.disableBn = false;
    }
  }
}
</script>