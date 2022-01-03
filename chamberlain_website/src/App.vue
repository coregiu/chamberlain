<template>
  <div class="view">
    <a href="#"><img src="./assets/coregiu-logo.png" id="logo"/></a>
    <div id="menuTop" class="top"/>

    <div id="content">
      <router-view/>
    </div>
    <div id="foot">©2021 COREGIU | <a href="https://github.com/coregiu" target="_blank">https://github.com/coregiu</a> |
      火在水上未济，君子以慎辨物居方。
    </div>
  </div>
</template>

<script>
  export default {
    name: "ChamberlainApp",
    date(){

    },
    created() {
      this.showMenu()
    },
    methods: {
      async showMenu() {
        let tokenValid = await this.$token.methods.checkChamberlainToken()
        if (!tokenValid) {
          this.$token.methods.clearChamberlainToken()
        }
        this.$menu.methods.setMenuTop()
      }
    }
  }
</script>