<template>

</template>

<script>
import {logout} from "../api/api";
export default {
  name: "logout",
  created() {
    this.$token.methods.clearChamberlainToken();
    this.$token.methods.showUnloginMenu();
    logout();
    this.$router.push("/");
  }
}
</script>