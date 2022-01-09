<template>
  <div style="float:left; width:20%;">
    <div style="margin-bottom: 0;">
      <Button type="button" class="p-button-text" label="▽展开" @click="expandAll"/>
      <Button type="button" class="p-button-text" label="▷折叠" @click="collapseAll"/>
      <Button type="button" class="p-button-text" label="✎新建" @click="opNewBookDialog"/>
      <Button type="button" class="p-button-text" label="✕删除" @click="openDeleteDialog"/>
    </div>
    <Tree :value="nodes" :expandedKeys="expandedKeys" @contextmenu="onMenuSelect"
          selectionMode="single" v-model:selectionKeys="selectedKeys"
          @node-select="onNodeSelect" @node-unselect="onNodeUnSelect"/>
    <ContextMenu ref="menu" :model="items"/>
  </div>
  <div style="float:right; width:80%;">
    <Editor v-model="currentSummaryNode.Content" editorStyle="height: 800px" @focusout="saveNoteContent"
            @focusin="this.editChangeCheckValue = this.currentSummaryNode.Content"/>
  </div>

  <Dialog v-model:visible="isDeleteDialogOpen" :style="{width: '350px'}" header="确认" :modal="true"
          v-if="currentSummaryNode">
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem"/>
      <span>你确认要删除 <b>{{ currentSummaryNode.data }}</b> 吗?</span>
    </div>
    <template #footer>
      <Button label="否" icon="pi pi-times" class="p-button-text" @click="isDeleteDialogOpen = false"/>
      <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteNoteSummary"/>
    </template>
  </Dialog>

  <Dialog v-model:visible="isNewDialogOpen" :style="{width: '300px'}" header="请输入" :modal="true" class="p-fluid">
    <div class="p-field">
      <label for="Username">请输入文件名:</label>
      <InputText id="Username" v-model.trim="newFileName" required="true" autofocus @keyup.enter.native="addNewBook()"
                 :class="{'p-invalid': submitted && !newFileName}"/>
      <small class="p-invalid" v-if="submitted && !newFileName">**名称必须填写**</small>
    </div>
    <template #footer>
      <Button label="取消" icon="pi pi-times" class="p-button-text" @click="isNewDialogOpen = false"/>
      <Button label="确认" icon="pi pi-check" class="p-button-text" @click="addNewBook"/>
    </template>
  </Dialog>

  <Dialog v-model:visible="tipDisplay" header="提示">{{ tipMessage }}</Dialog>
</template>

<script>
import Uuid from '../util/uuid.ts';
import NoteSummaryService from '../api/note_summary.ts';

export default {
  name: "note_summary",
  data() {
    return {
      nodes: null,
      expandedKeys: {},
      editChangeCheckValue: "",
      selectedKeys: null,
      currentSummaryNode: {"Content": "文本编辑框失去焦点后会自动保存文档..."},
      isDeleteDialogOpen: false,
      isNewDialogOpen: false,
      tipDisplay: false,
      submitted: false,
      tipMessage: "",
      newFileName: "",
      items: [
        {
          label: '新建',
          icon: 'pi pi-fw pi-plus',
          command: () => this.opNewBookDialog(),
        },
        {
          label: '删除',
          icon: 'pi pi-fw pi-trash',
          command: () => this.openDeleteDialog(),
        },
        {
          separator: true
        },
        {
          label: '移动',
          icon: 'pi pi-fw pi-external-link',
          command: () => this.openMoveBookDialog(),
        }
      ]
    }
  },
  nodeService: null,
  uuidService: null,
  created() {
    this.nodeService = new NoteSummaryService();
    this.uuid = new Uuid()
  },
  mounted() {
    this.nodeService.getTreeNodes().then(data => {
      this.nodes = data
      this.expandAll()
    })
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
      this.isNewDialogOpen = true
    },

    addNewBook() {
      this.submitted = true
      if (!this.newFileName) {
        return
      }
      let parentBookId = this.currentSummaryNode === null ? "0" : this.currentSummaryNode.key
      let noteSummary = {
        "BookId": this.uuid.getUuid(),
        "BookName": this.newFileName,
        "ParentBookId": parentBookId,
        "BookTime": new Date()
      }
      this.nodeService.addNoteSummary(noteSummary).then(res => {
        if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
          this.tipDisplay = true;
          this.tipMessage = "添加失败！";
        } else {
          location.reload()
        }
      })
    },

    openDeleteDialog() {
      if (this.currentSummaryNode === null) {
        this.tipDisplay = true;
        this.tipMessage = "请先选择要删除的节点！";
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
    },

    onNodeSelect(node) {
      this.currentSummaryNode = node
      if (!this.currentSummaryNode.Content) {
        this.nodeService.getNoteSummaryContent(this.currentSummaryNode.key).then(res => this.currentSummaryNode.Content = res.Content)
      }
    },

    onNodeUnSelect() {
    },

    saveNoteContent(event) {
      if (!this.currentSummaryNode || this.editChangeCheckValue === this.currentSummaryNode.Content) {
        return
      }
      this.nodeService.updateNoteSummary({"BookId": this.currentSummaryNode.key, "Content": this.currentSummaryNode.Content})
    }
  }
}
</script>