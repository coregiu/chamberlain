<template>

</template>

<script>
export default {
  name: "logout",
  created() {
    this.$token.methods.clearChamberlainToken();
    this.$token.methods.showUnloginMenu();
    this.$router.push("/");
  }
}
</script>