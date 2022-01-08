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
    <Editor v-model="currentSummaryValue" editorStyle="height: 800px"/>
  </div>

  <Dialog v-model:visible="isDeleteDialogOpen" :style="{width: '350px'}" header="确认" :modal="true" v-if="currentSummaryNode">
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem"/>
      <span>你确认要删除吗{{currentSummaryNode.data}}?</span>
    </div>
    <template #footer>
      <Button label="否" icon="pi pi-times" class="p-button-text" @click="isDeleteDialogOpen = false"/>
      <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteNoteSummary"/>
    </template>
  </Dialog>

  <Dialog v-model:visible="tipDisplay" header="事务跟踪提示">{{ tipMessage }}</Dialog>
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
      isDeleteDialogOpen: false,
      tipDisplay: false,
      tipMessage: "",
      items: [
        {
          label: '新建',
          icon: 'pi pi-fw pi-plus',
          command:() => this.opNewBookDialog(),
        },
        {
          label: '删除',
          icon: 'pi pi-fw pi-trash',
          command:() => this.openDeleteDialog(),
        },
        {
          separator: true
        },
        {
          label: '移动',
          icon:'pi pi-fw pi-external-link',
          command:() => this.openMoveBookDialog(),
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

    opNewBookDialog() {

    },

    openDeleteDialog() {
      if (this.currentSummaryNode === null) {
        return
      }
      this.isDeleteDialogOpen = true
    },

    deleteNoteSummary() {
      this.nodeService.deleteNoteSummary(this.currentSummaryNode).then(res => {
        if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
          this.tipDisplay = true;
          this.tipMessage = "删除失败！";
        } else {
          location.reload()
        }
      })
    },

    openMoveBookDialog() {
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