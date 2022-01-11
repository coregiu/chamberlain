<template>
  <div style="float:left; width:20%;">
    <div style="margin-bottom: 0;">
      <Button type="button" class="p-button-text" label="▽展开" @click="expandAll"/>
      <Button type="button" class="p-button-text" label="▷折叠" @click="collapseAll"/>
      <Button type="button" class="p-button-text" label="+新建" @click="opNewBookDialog"/>
      <Button type="button" class="p-button-text" label="✕删除" @click="openDeleteDialog"/>
    </div>
    <Tree :value="nodes" :expandedKeys="expandedKeys" @contextmenu="onMenuSelect"
          selectionMode="single" v-model:selectionKeys="selectedKeys"
          @node-select="onNodeSelect" @node-unselect="onNodeUnSelect"/>
    <ContextMenu ref="menu" :model="items"/>
  </div>
  <div style="float:right; width:80%;">
    <Editor v-model="currentSummaryNode.content" editorStyle="height: 800px" @focusout="saveNoteContent"
            @focusin="this.editChangeCheckValue = this.currentSummaryNode.content"/>
  </div>

  <Dialog v-model:visible="isChangeDialogOpen" :style="{width: '350px'}" header="选择变更节点" :modal="true"
          v-if="currentSummaryNode">
    <div class="confirmation-content">
      <Tree :value="[{'key':'0', 'label':'根节点','data':'根节点', 'icon': 'pi pi-fw pi-inbox', 'children':nodes}]" :expandedKeys="expandedKeys"
            selectionMode="single" v-model:selectionKeys="selectedKeys"
            @node-select="onChangeNodeSelect"/>
    </div>
    <template #footer>
      <Button label="否" icon="pi pi-times" class="p-button-text" @click="isChangeDialogOpen = false"/>
      <Button label="是" icon="pi pi-check" class="p-button-text" @click="changeSummaryNode"/>
    </template>
  </Dialog>

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
      <label for="FileName">
        <span v-if="isNewOperation && currentSummaryNode.key!=='0'">添加 <b>{{currentSummaryNode.label}}</b> 的子文件，请输入文件名:</span>
        <span v-if="isNewOperation && currentSummaryNode.key==='0'">添加根文件，请输入文件名:</span>
        <span v-if="!isNewOperation">修改 <b>{{currentSummaryNode.label}}</b> 的文件名，请输入新文件名:</span>
      </label>
      <InputText id="FileName" v-model.trim="newFileName" required="true" autofocus @keyup.enter.native="addNewBook()"
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
      currentSummaryNode: {"key": "0", "content": "文本编辑框失去焦点后会自动保存文档..."},
      isDeleteDialogOpen: false,
      isChangeDialogOpen: false,
      changeToNode: {},
      isNewDialogOpen: false,
      isNewOperation: true,
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
          label: '重命名',
          icon: 'pi pi-fw pi-pencil',
          command: () => this.openUpdateDialog(),
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
      this.isNewOperation = true
    },
    addNewBook() {
      this.submitted = true
      if (!this.newFileName) {
        return
      }
      if (this.isNewOperation) {
        let noteSummary = {
          "BookId": this.uuid.getUuid(),
          "BookName": this.newFileName,
          "ParentBookId": this.currentSummaryNode.key,
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
      } else {
        if (this.newFileName === this.currentSummaryNode.label) {
          return
        }
        this.nodeService.updateNoteSummary({"BookId": this.currentSummaryNode.key, "BookName": this.newFileName})
        location.reload()
      }
    },
    openDeleteDialog() {
      if (this.currentSummaryNode.key === "0") {
        this.tipDisplay = true;
        this.tipMessage = "请先选择要删除的节点！";
        return
      }
      this.isDeleteDialogOpen = true
    },
    openUpdateDialog() {
      if (this.currentSummaryNode.key === "0") {
        this.tipDisplay = true;
        this.tipMessage = "请先选择要重命名的节点！";
        return
      }
      this.isNewDialogOpen = true
      this.isNewOperation = false
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
      if (this.currentSummaryNode.key === "0") {
        this.tipDisplay = true;
        this.tipMessage = "请先选择要移动的节点！";
        return
      }
      this.isChangeDialogOpen = true
    },
    onNodeSelect(node) {
      this.currentSummaryNode = node
      if (!this.currentSummaryNode.content) {
        this.nodeService.getNoteSummaryContent(this.currentSummaryNode.key).then(res => this.currentSummaryNode.content = res.Content)
      }
    },
    onNodeUnSelect() {
    },
    saveNoteContent(event) {
      if (this.currentSummaryNode.key === "0" || this.editChangeCheckValue === this.currentSummaryNode.content) {
        return
      }
      this.nodeService.updateNoteSummary({"BookId": this.currentSummaryNode.key, "Content": this.currentSummaryNode.content})
    },
    onChangeNodeSelect(node) {
      this.changeToNode = node
    },
    changeSummaryNode() {
      this.isChangeDialogOpen = false
      if (this.changeToNode.key === this.currentSummaryNode.key) {
        return
      }
      this.nodeService.updateNoteSummary({"BookId": this.currentSummaryNode.key, "ParentBookId": this.changeToNode.key})
      location.reload()
    }
  }
}
</script>