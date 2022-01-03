<template>

</template>

<script>
import {logout} from "../api/system_api";

export default {
  name: "logout",
  created() {
    logout();
    this.$token.methods.clearChamberlainToken();
    this.$menu.methods.setMenuTop()
    this.$router.push("/");
  }
}
</script>