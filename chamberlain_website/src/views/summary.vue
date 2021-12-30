<template>
  <iframe id="iframeContent" src = "public/books/summary/_build/html/index.html"></iframe>
</template>

<script>
export default {
  name: "philosophy",
  data() {

  }
}
</script>

<style scoped>
#iframeContent {
  border: 0;
  width: 90%;
  height: 87vh;
  margin-left: 1%;
  overflow: scroll;
}
</style>