<template>
  <div style="float:left; width:20%;">
    <div style="margin-bottom: 1em;">
      <Button type="button" icon="pi pi-plus" class="p-button-text" label="全部展开" @click="expandAll"/>
      <Button type="button" icon="pi pi-minus" class="p-button-text" label="全部关闭" @click="collapseAll"/>
    </div>
    <Tree :value="nodes" :expandedKeys="expandedKeys" @contextmenu="onMenuSelect"
          selectionMode="single" v-model:selectionKeys="selectedKeys"
          @node-select="onNodeSelect" @node-unselect="onNodeUnSelect"/>
    <ContextMenu ref="menu" :model="items"/>
  </div>
  <div style="float:right; width:80%;">
    <Editor v-model="currentSummaryValue" editorStyle="height: 800px">
      <template slot="toolbar">
        <span class="ql-formats">
          <button class="ql-bold"></button>
          <button class="ql-italic"></button>
          <button class="ql-underline"></button>
        </span>
      </template>
    </Editor>
  </div>
</template>

<script>
import NoteSummaryService from '../api/note_summary.ts';

export default {
  name: "note_summary",
  data() {
    return {
      nodes: null,
      expandedKeys: {},
      currentSummaryValue: "",
      selectedKeys: null,
      currentSummaryNode: null,
      items: [
        {
          label: '新建',
          icon: 'pi pi-fw pi-plus',
          command:() => this.newBook(),
        },
        {
          label: '删除',
          icon: 'pi pi-fw pi-trash',
          command:() => this.deleteBook(),
        },
        {
          separator: true
        },
        {
          label: '移动',
          icon:'pi pi-fw pi-external-link',
          command:() => this.moveBook(),
        }
      ]
    }
  },
  nodeService: null,
  created() {
    this.nodeService = new NoteSummaryService();
  },
  mounted() {
    this.nodeService.getTreeNodes().then(data => this.nodes = data);
  },
  methods: {
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
    },

    onMenuSelect(event) {
      this.$refs.menu.show(event);
    },

    newBook() {
      
    },

    deleteBook() {

    },

    moveBook() {
      console.log("-----------unsupported now-------------")
    },

    onNodeSelect(node) {
      this.currentSummaryNode = node
      this.nodeService.getNoteSummaryContent(this.currentSummaryNode.key).then(res => this.currentSummaryValue = res.Content)
    },

    onNodeUnSelect() {
      this.currentSummaryNode = null
    }
  }
}
</script>