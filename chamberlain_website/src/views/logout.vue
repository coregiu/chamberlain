<template>

</template>

<script>
import {logout} from "../api/api";

export default {
  name: "logout",
  created() {
    logout();
    this.$token.methods.clearChamberlainToken();
    this.$token.methods.showUnloginMenu();
    this.$router.push("/");
  }
}
</script>