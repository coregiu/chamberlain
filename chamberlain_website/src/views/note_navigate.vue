<template>
  <div style="float:left; width:10%;">
    <span><p><b><router-link :to="{name:'NoteBook'}" class="note-link">🖋 待办事务</router-link></b></p></span>
    <span><p><b><router-link :to="{name:'NoteSummary'}" class="note-link">📙 工作便笺</router-link></b></p></span>
  </div>
  <div style="float:right; width:90%">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "note-navigate",
}
</script>

<style>
.note-link {
  color: #3a0d14;
  text-decoration: none;
}

.note-link:hover {
  color: #026da7;
}
</style>