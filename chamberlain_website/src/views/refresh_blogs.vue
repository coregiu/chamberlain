<template>
{{message}}
</template>

<script>
import {refreshBlogs} from "../api/api";

export default {
  name: "refresh_blogs",
  data() {
    return {
      message: "正在生成博客，请不要重复提交。博客刷新后会自动跳转到博客页面，请耐心等待..."
    }
  },
  async created() {
    await refreshBlogs();
    this.$router.push("/blog");
  }
}
</script>