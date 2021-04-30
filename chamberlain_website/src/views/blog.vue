<template>
  <div id="iframeLeft">
    <div style="text-align: center">
      <Button type="button" icon="pi pi-plus" label="全部展开" @click="expandAll"/> &nbsp;
      <Button type="button" icon="pi pi-minus" label="全部关闭" @click="collapseAll" /><br><br>
    </div>
    我的博客
    <Tree :value='nodes' selectionMode="single" v-model:selectionKeys="selectedKey" @node-select="onNodeSelect" v-model:expandedKeys="expandedKeys"></Tree>
  </div>
  <iframe id="iframeContent" name="iframeContent"></iframe>
</template>

<script>
import TreeNode from '../api/tree'

export default {
  name: "blog",
  data() {
    return {
      selectedKey: null,
      expandedKeys: {},
      nodes: null
    }
  },
  treeNode: null,
  created() {
    this.treeNode = new TreeNode();
  },
  async mounted() {
    await this.treeNode.getTreeNodes().then(data => this.nodes = data);
  },
  methods: {
    onNodeSelect(node) {
      document.getElementById("iframeContent").src = node.link;
    },
    expandAll() {
      for (let node of this.nodes) {
        this.expandNode(node);
      }

      this.expandedKeys = {...this.expandedKeys};
    },
    collapseAll() {
      this.expandedKeys = {};
    },
    expandNode(node) {
      if (node.children && node.children.length) {
        this.expandedKeys[node.key] = true;

        for (let child of node.children) {
          this.expandNode(child);
        }
      }
    }
  }
}
</script>

<style scoped>
#iframeLeft {
  width: 20%;
  height: 87vh;
  float: left;
  border: 0;
  background-color: #fff;
  text-align: left;
  padding-left: 5px;
  padding-top:10px;
}

#iframeContent {
  border: 0;
  width: 79%;
  height: 87vh;
  margin-left: 1%;
  background-color: #fff;
}
button {
  background-color: #ffffff;
  color: #0d3c61;
}
</style>