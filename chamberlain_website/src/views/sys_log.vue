<template>
  <DataTable ref="syslogTable" :value="syslogList" :paginator="true" class="p-datatable-customers" :rows="10"
             dataKey="LogId" :rowHover="true"
             :loading="loading"
             paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
             :rowsPerPageOptions="[10,25,50]"
             currentPageReportTemplate="Showing {first} to {last} of {totalRecords} entries"
             v-model:selection="selectedLogIds" selectionMode="multiple">

    <template #header>
      <div>
        <div style="float:left">系统操作日志列表</div>
        <div style="float:right">
          <Button type="button" class="p-button-secondary" @click="openDeleteSyslogDialog">删除</Button>&nbsp;&nbsp;&nbsp;&nbsp;
          <span class="p-input-icon-left">
              <i class="pi pi-search"/>
              <InputText v-model="username" placeholder="操作人" @keyup.enter.native="queryByCondition"/>
          </span>&nbsp;
          <span class="p-input-icon-left">
              <i class="pi pi-search"/>
              <InputText v-model="operation" placeholder="操作功能" @keyup.enter.native="queryByCondition"/>
          </span>
        </div>
      </div>
    </template>
    <template #empty>
      无数据
    </template>
    <template #loading>
      正在加载数据，请稍等...
    </template>
    <Column selectionMode="multiple" headerStyle="width: 3em"></Column>
    <Column field="selectIds"><Checkbox name="city" v-model="selectedLogIds" /></Column>
    <Column field="LogId" header="日志ID" :sortable="true" sortField="LogId"/>
    <Column field="Username" header="操作人" :sortable="true" sortField="Username"/>
    <Column field="Operation" header="操作功能" :sortable="true" sortField="Operation"/>
    <Column field="OpTime" header="操作时间" :sortable="true" sortField="OpTime"/>
    <Column field="OpResult" header="操作结果" :sortable="true" sortField="OpResult"/>
    <Column field="Description" header="描述" :sortable="true" sortField="Description"/>
  </DataTable>

  <Dialog v-model:visible="isDeleteInputDialogOpen" :style="{width: '350px'}" header="确认" :modal="true">
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem"/>
      <span v-if="selectedLogIds">你确认要删除这些日志吗?</span>
    </div>
    <template #footer>
      <Button label="否" icon="pi pi-times" class="p-button-text" @click="isDeleteInputDialogOpen = false"/>
      <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteSyslog"/>
    </template>
  </Dialog>
  <Dialog v-model:visible="tipDisplay" header="系统日志管理提示">{{ tipMessage }}</Dialog>
</template>

<script>
import SyslogService from '../api/sys_log.ts';

export default {
  name: "syslog",
  data() {
    return {
      username: "",
      operation: "",
      selectedLogIds: null,
      syslogList: null,
      loading: true,
      tipDisplay: false,
      tipMessage: "",
      isDeleteInputDialogOpen: false,
    }
  },
  syslogService: null,
  created() {
    this.syslogService = new SyslogService();
  },
  mounted() {
    this.syslogService.getSyslogList("", "", 1000, 0).then(data => this.syslogList = data);
    this.loading = false;
  },
  methods: {
    openDeleteSyslogDialog() {
      this.isDeleteInputDialogOpen = true;
    },

    queryByCondition() {
      this.syslogService.getSyslogList(this.username, this.operation, 1000, 0).then(data => this.syslogList = data);
    },

    async deleteSyslog() {
      this.isDeleteInputDialogOpen = false
      console.log(this.selectedLogIds)
      let res = await this.syslogService.deleteSyslog(this.selectedLogIds)
      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this.tipDisplay = true;
        this.tipMessage = "删除失败！";
      } else {
        location.reload()
      }
    }
  }
}
</script>