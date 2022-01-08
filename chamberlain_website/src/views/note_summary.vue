<template>
  <div style="float:left; width:10%;">
    <span><p><b><router-link @click="changePanel('todo')" to="#" class="note-link">🖋 待办事务</router-link></b></p></span>
    <span><p><b><router-link @click="changePanel('summary')" to="#" class="note-link">📙 日常记事</router-link></b></p></span>
  </div>
  <div style="float:right; width:90%">
    <DataTable ref="notebookTable" :value="notebookList" :paginator="true" class="p-datatable-customers" :rows="10"
               dataKey="NoteId" :rowHover="true"
               :loading="loading"
               paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
               :rowsPerPageOptions="[10,25,50]"
               currentPageReportTemplate="Showing {first} to {last} of {totalRecords} entries">

      <template #header>
        <div>
          <div style="float:left">我的待办列表</div>
          <div style="float:right">
            <Button type="button" class="p-button-secondary" @click="addNotebookDialog">添加待办</Button>&nbsp;
            <span class="p-input-icon-left">
              <i class="pi pi-search"/>
              <Dropdown v-model="currentQueryObject" :options="queryOptions" optionLabel="name"
                        :placeholder="currentQueryObject.code" @change="changeQuery"/>
            </span>&nbsp;
            <Button label="导出" icon="pi pi-upload" class="p-button-secondary" @click="exportCSV($event)"/>&nbsp;
          </div>
        </div>
      </template>
      <template #empty>
        无数据
      </template>
      <template #loading>
        正在加载数据，请稍等...
      </template>

      <Column header="任务时间" field="NoteTime" :sortable="true" sortField="NoteTime">
        <template #body="noteData">
          <span class="image-text">{{ this.formatDate(noteData.data.NoteTime) }}</span>
        </template>
      </Column>
      <Column field="Content" header="内容" :sortable="true" sortField="Content">
        <template #body="notebookInfo">
          <span :title="notebookInfo.data.Content">{{notebookInfo.data.Content.length > 40 ? notebookInfo.data.Content.substring(0, 40) + " ..." : notebookInfo.data.Content}}</span>
        </template>
      </Column>
      <Column field="Level" header="级别" :sortable="true" sortField="Level">
        <template #body="noteData">
          <span style="background-color: #5d0c28; color: white; font-size: 21px"
                v-if="noteData.data.Level === 'H'">高</span>
          <span style="background-color: #8a6a19; color: white; font-size: 21px"
                v-if="noteData.data.Level === 'M'">中</span>
          <span style="background-color: #0b7ad1; color: white; font-size: 21px"
                v-if="noteData.data.Level === 'L'">低</span>
        </template>
      </Column>
      <Column field="Status" header="状态" :sortable="true" sortField="Status">
        <template #body="noteData">
          <span style="background-color: #8f5902; color: white; font-size: 19px" v-if="noteData.data.Status === 'NEW'">未启动</span>
          <span style="background-color: #026da7; color: white; font-size: 19px"
                v-if="noteData.data.Status === 'DOING'">进行中</span>
          <span style="background-color: #4caf50; color: white; font-size: 19px"
                v-if="noteData.data.Status === 'CLOSED'">已完成</span>
        </template>
      </Column>
      <Column field="FinishTime" header="计划完成时间" :sortable="true" sortField="FinishTime">
        <template #body="noteData">
          <span style="background-color: #5d0c28; color: white; font-size: 17px"
                v-if="!this.compareTodayTime(noteData.data.FinishTime, noteData.data.Status)">
                {{ this.formatDate(noteData.data.FinishTime) }}
          </span>
          <span class="image-text" v-if="this.compareTodayTime(noteData.data.FinishTime, noteData.data.Status)">
            {{ this.formatDate(noteData.data.FinishTime) }}
          </span>
        </template>
      </Column>
      <Column field="Owner" header="责任人" :sortable="true" sortField="Owner"/>
      <Column field="RealFinishTime" header="实际完成时间" :sortable="true" sortField="RealFinishTime">
        <template #body="noteData">
          <span class="image-text">{{ formatRealFinishDate(noteData.data.Status, noteData.data.RealFinishTime) }}</span>
        </template>
      </Column>
      <Column header="操作" headerStyle="width: 8rem; text-align: center"
              bodyStyle="text-align: center; overflow: visible">
        <template #body="notebookInfo">
          <Button type="button" icon="pi pi-pencil" class="p-button-secondary" title="修改"
                  @click="openUpdateNotebookDialog(notebookInfo)"></Button>&nbsp;&nbsp;
          <Button type="button" icon="pi pi-trash" class="p-button-danger" title="删除"
                  @click="openDeleteNotebookDialog(notebookInfo)"></Button>
        </template>
      </Column>
    </DataTable>
  </div>

  <Dialog v-model:visible="isNewNotebookDialogOpen" :style="{width: '350px'}" header="待办信息" :modal="true"
          class="p-fluid">
    <div class="p-field">
      <label for="Description">待办内容:</label>
      <Textarea id="Description" v-model.trim="notebookInfo.Content" rows="10" cols="60"
                :class="{'p-invalid': submitted && !notebookInfo.Content}"/>
      <small class="p-invalid" v-if="submitted && !notebookInfo.Content">**待办内容必须填写**</small>
    </div>
    <br>
    <div class="p-field">
      <label class="p-mb-3">优先级:</label>
      <div class="p-formgrid p-grid">
        <div class="p-field-radiobutton p-col-6">
          <RadioButton id="category1" name="category" value="H" v-model="notebookInfo.Level"
                       :class="{'p-invalid': submitted && !notebookInfo.Level}"/>
          <label for="category1">高</label>&nbsp;&nbsp;&nbsp;&nbsp;
          <RadioButton id="category2" name="category" value="M" v-model="notebookInfo.Level"
                       :class="{'p-invalid': submitted && !notebookInfo.Level}"/>
          <label for="category2">中</label>&nbsp;&nbsp;&nbsp;&nbsp;
          <RadioButton id="category3" name="category" value="L" v-model="notebookInfo.Level"
                       :class="{'p-invalid': submitted && !notebookInfo.Level}"/>
          <label for="category3">低</label>
        </div>
        <small class="p-invalid" v-if="submitted && !notebookInfo.Level">**优先级必须填写**</small>
      </div>
    </div>
    <br>
    <div class="p-field">
      <label for="NoteId">计划完成时间: </label>
      <Calendar id="NoteId" v-model="notebookInfo.FinishTime" dateFormat="yymmdd"
                :class="{'p-invalid': submitted && !notebookInfo.FinishTime}"/>
      <small class="p-invalid" v-if="submitted && !notebookInfo.FinishTime">**计划完成时间必须填写**</small>
    </div>
    <br>

    <div class="p-field">
      <label for="Owner">责任人:</label>
      <InputText id="Owner" v-model.trim="notebookInfo.Owner"/>
    </div>
    <br>

    <div class="p-field" v-if="!isAddOperation">
      <label class="p-mb-3">当前状态:</label>
      <div class="p-formgrid p-grid">
        <div class="p-field-radiobutton p-col-6">
          <RadioButton id="todoStatus1" name="todoStatus" value="NEW" v-model="notebookInfo.Status"
                       :class="{'p-invalid': submitted && !notebookInfo.Status}"/>
          <label for="todoStatus1">未启动</label>&nbsp;&nbsp;&nbsp;&nbsp;
          <RadioButton id="todoStatus2" name="todoStatus" value="DOING" v-model="notebookInfo.Status"
                       :class="{'p-invalid': submitted && !notebookInfo.Status}"/>
          <label for="todoStatus2">进行中</label>&nbsp;&nbsp;&nbsp;&nbsp;
          <RadioButton id="todoStatus3" name="todoStatus" value="CLOSED" v-model="notebookInfo.Status"
                       :class="{'p-invalid': submitted && !notebookInfo.Status}"/>
          <label for="todoStatus3">已完成</label>
        </div>
        <small class="p-invalid" v-if="submitted && !notebookInfo.Level">**优先级必须填写**</small>
      </div>
    </div>
    <br>
    <template #footer>
      <Button label="取消" icon="pi pi-times" class="p-button-text" @click="hideDialog"/>
      <Button label="保存" icon="pi pi-check" class="p-button-text" @click="saveNotebook"/>
    </template>
  </Dialog>

  <Dialog v-model:visible="isDeleteNotebookDialogOpen" :style="{width: '350px'}" header="确认" :modal="true">
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem"/>
      <span v-if="notebookInfo">你确认要删除 <b>{{ notebookInfo.Content }}</b>?</span>
    </div>
    <template #footer>
      <Button label="否" icon="pi pi-times" class="p-button-text" @click="isDeleteNotebookDialogOpen = false"/>
      <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteNotebook"/>
    </template>
  </Dialog>

  <Dialog v-model:visible="tipDisplay" header="事务跟踪提示">{{ tipMessage }}</Dialog>
</template>

<script>
import NotebookService from '../api/notebook.ts';

export default {
  name: "note_summary",
  data() {
    return {
      currentQueryObject: {"name": "全部", "code": "all"},
      queryOptions: [{"name": "全部", "code": "all"},
        {"name": "今日待办", "code": "today"},
        {"name": "未完成", "code": "unfin"},
        {"name": "今日未完", "code": "tounfin"}],
      notebookList: null,
      loading: true,
      notebookInfo: null,
      isNewNotebookDialogOpen: false,
      submitted: false,
      tipDisplay: false,
      tipMessage: "",
      isDeleteNotebookDialogOpen: false,
      isAddOperation: true
    }
  },
  notebookService: null,
  created() {
    this.notebookService = new NotebookService();
  },
  mounted() {
    this.notebookService.getNotebookList("", "", 10000, 0).then(data => this.notebookList = data);
    this.loading = false;
  },
  methods: {
    addNotebookDialog() {
      let today = new Date()
      today.setHours(0)
      today.setMinutes(0)
      today.setSeconds(0)
      this.notebookInfo = {"Level": "H", "FinishTime": today}
      this.submitted = false;
      this.isNewNotebookDialogOpen = true;
      this.isAddOperation = true;
    },

    hideDialog() {
      this.isNewNotebookDialogOpen = false;
      this.submitted = false;
    },

    async saveNotebook() {
      this.submitted = true;
      if (this.isAddOperation) {
        this.notebookInfo.NoteId = this.getUuid()
        this.notebookInfo.NoteTime = new Date()
        this.notebookInfo.Status = "NEW"
        this.notebookInfo.RealFinishTime = new Date(this.notebookInfo.FinishTime.getFullYear() + 1, this.notebookInfo.FinishTime.getMonth() + 1, this.notebookInfo.FinishTime.getDate(), 0, 0, 0)
      }

      if (!this.notebookInfo.Content
          || !this.notebookInfo.FinishTime
          || !this.notebookInfo.Level) {
        return
      }

      if (this.notebookInfo.Status === "CLOSED") {
        this.notebookInfo.RealFinishTime = new Date()
      }

      let res = ""
      if (this.isAddOperation) {
        res = await this.notebookService.addNotebook(this.notebookInfo);
      } else {
        res = await this.notebookService.updateNotebook(this.notebookInfo)
      }

      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this.tipDisplay = true;
        this.tipMessage = "操作失败，请检查输入信息！";
      } else {
        location.reload()
      }
    },

    openDeleteNotebookDialog(notebookInfo) {
      this.notebookInfo = notebookInfo.data;
      this.isDeleteNotebookDialogOpen = true;
    },
    async deleteNotebook() {
      this.isDeleteNotebookDialogOpen = false;
      let res = await this.notebookService.deleteNotebook(this.notebookInfo)
      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this    .tipDisplay = true;
        this.tipMessage = "删除失败！";
      } else {
        this.notebookList = this.notebookList.filter(val => val.NoteId !== this.notebookInfo.NoteId);
        this.notebookInfo = {};
      }
    },

    openUpdateNotebookDialog(notebookInfo) {
      this.isAddOperation = false
      this.notebookInfo = {...notebookInfo.data};
      this.submitted = false;
      this.isNewNotebookDialogOpen = true;
    },

    exportCSV() {
      this.$refs.notebookTable.exportCSV();
    },

    getUuid() {
      function S4() {
        return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
      }

      return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
    },

    formatDate(time) {
      let date = new Date(time)
      let year = date.getFullYear()
      let month = date.getMonth() + 1 //月份是从0开始的
      let day = date.getDate()
      return year + '-' + month + '-' + day
    },

    formatRealFinishDate(status, realFinishTIme) {
      if (status === 'CLOSED') {
        return this.formatDate(realFinishTIme)
      } else {
        return ""
      }
    },

    compareTodayTime(time, status) {
      if (status === 'CLOSED') {
        return true
      }
      let compareDate = new Date(time)
      return Date.parse(compareDate) - Date.parse(new Date()) >= 0
    },

    changeQuery() {
      if (this.currentQueryObject.code === 'today') {
        let today = new Date()
        let todayTime = today.getFullYear() * 10000 + (today.getMonth() + 1) * 100 + today.getDate()
        this.notebookService.getNotebookList(todayTime, "", 10000, 0).then(data => this.notebookList = data);
      } else if (this.currentQueryObject.code === 'unfin') {
        this.notebookService.getNotebookList("", "NEW,DOING", 10000, 0).then(data => this.notebookList = data);
      } else if (this.currentQueryObject.code === 'tounfin') {
        let today = new Date()
        let todayTime = today.getFullYear() * 10000 + (today.getMonth() + 1) * 100 + today.getDate()
        this.notebookService.getNotebookList(todayTime, "NEW,DOING", 10000, 0).then(data => this.notebookList = data);
      } else {
        this.notebookService.getNotebookList("", "", 10000, 0).then(data => this.notebookList = data);
      }
    },

    changePanel(panel){
      alert("------------------------" + panel)
    }
  }
}
</script>

<style>
.note-link{
  color: #3a0d14;
  text-decoration: none;
}
.note-link:hover{
   color: #026da7;
 }
</style>