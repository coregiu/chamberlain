<template>
  <DataTable :value="inputInfoList" :paginator="true" class="p-datatable-customers" :rows="10"
             dataKey="InputTime" :rowHover="true" :filters="filters"
             filterDisplay="row"
             :loading="loading"
             paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
             :rowsPerPageOptions="[10,25,50]"
             currentPageReportTemplate="Showing {first} to {last} of {totalRecords} entries">

    <template #header>
      <div>
        <div style="float:left">收入列表</div>
        <div style="float:right">
          <Button type="button" class="p-button-secondary" @click="addInputDialog">添加信息</Button>&nbsp;&nbsp;&nbsp;&nbsp;
          <span class="p-input-icon-left">
              <i class="pi pi-search"/>
              <InputText v-model="filters['global']" placeholder="全局检索"/>
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

    <Column field="InputTime" header="时间" :sortable="true" sortField="InputTime"/>
    <Column field="Year" header="年度" :sortable="true" sortField="Year" filterField="Year" filterMatchMode="contains">
      <template #filter>
        <InputText type="text" v-model="filters['Year']" class="p-column-filter" placeholder="按年度检索" size="4"/>
      </template>
    </Column>
    <Column field="Month" header="月份" :sortable="true" sortField="Month" filterField="Month" filterMatchMode="contains">
      <template #filter>
        <InputText type="text" v-model="filters['Month']" class="p-column-filter" placeholder="按月份检索" size="2"/>
      </template>
    </Column>
    <Column field="Type" header="收入类型" :sortable="true" sortField="Type" filterField="Type" filterMatchMode="contains">
      <template #filter>
        <InputText type="text" v-model="filters['Type']" class="p-column-filter" placeholder="按收入类型检索" size="6"/>
      </template>
    </Column>
    <Column field="AllInput" header="税前收入" :sortable="true" sortField="AllInput"/>
    <Column field="Actual" header="税后收入" :sortable="true" sortField="Actual"/>
    <Column field="Tax" header="缴税额" :sortable="true" sortField="Tax"/>
    <Column field="Base" header="基本工资" :sortable="true" sortField="Base"/>
    <Column field="Description" header="备注" :sortable="true" sortField="Description"/>
    <Column header="操作" headerStyle="width: 8rem; text-align: center" bodyStyle="text-align: center; overflow: visible">
      <template #body="inputInfo">
        <Button type="button" icon="pi pi-pencil" class="p-button-secondary" title="修改"
                @click="openUpdateInputDialog(inputInfo)"></Button>&nbsp;&nbsp;
        <Button type="button" icon="pi pi-trash" class="p-button-danger" title="删除"
                @click="openDeleteInputDialog(inputInfo)"></Button>
      </template>
    </Column>
  </DataTable>

  <Dialog v-model:visible="isNewInputDialogOpen" :style="{width: '350px'}" header="收入信息" :modal="true" class="p-fluid">
    <div class="p-field">
      <label for="InputTime">收入时间: <span class="p-invalid"
                                         v-if="!isAddOperation">{{ inputInfo.InputTime }}</span></label>
      <Calendar id="InputTime" v-model="inputTime" dateFormat="yymmdd"
                :class="{'p-invalid': submitted && !inputTime}" v-if="isAddOperation"/>
      <small class="p-invalid" v-if="submitted && !inputTime">**收入时间必须填写**</small>
    </div>
    <br>

    <div class="p-field">
      <label class="p-mb-3">收入类型:</label>
      <div class="p-formgrid p-grid">
        <div class="p-field-radiobutton p-col-6">
          <RadioButton id="category1" name="category" value="salary" v-model="inputInfo.Type"/>
          <label for="category1">工资</label>&nbsp;&nbsp;&nbsp;&nbsp;
          <RadioButton id="category2" name="category" value="award" v-model="inputInfo.Type"/>
          <label for="category2">资金</label>&nbsp;&nbsp;&nbsp;&nbsp;
          <RadioButton id="category3" name="category" value="dividends" v-model="inputInfo.Type"/>
          <label for="category2">分红</label>
        </div>
      </div>
      <small class="p-invalid" v-if="submitted && !inputInfo.Type">**收入类型必须填写**</small>
    </div>
    <br>

    <div class="p-field">
      <label for="AllInput">税前收入:</label>
      <InputText id="AllInput" v-model.trim="inputInfo.AllInput" required="true" autofocus
                 :class="{'p-invalid': submitted &&  (!inputInfo.AllInput || isNaN(inputInfo.AllInput))}"/>
      <small class="p-invalid"
             v-if="submitted &&  (!inputInfo.AllInput || isNaN(inputInfo.AllInput))">**税前收入必须填写且是数字**</small>
    </div>
    <br>

    <div class="p-field">
      <label for="Actual">税后收入:</label>
      <InputText id="Actual" v-model.trim="inputInfo.Actual" required="true" autofocus
                 :class="{'p-invalid': submitted &&  (!inputInfo.Actual || isNaN(inputInfo.Actual))}"/>
      <small class="p-invalid"
             v-if="submitted && (!inputInfo.Actual || isNaN(inputInfo.Actual))">**税后收入必须填写且是数字**</small>
    </div>
    <br>

    <div class="p-field">
      <label for="Base">基本工资:</label>
      <InputText id="Base" v-model.trim="inputInfo.Base" required="true" autofocus
                 :class="{'p-invalid': submitted && (!inputInfo.Base || isNaN(inputInfo.Base))}"/>
      <small class="p-invalid" v-if="submitted && (!inputInfo.Base || isNaN(inputInfo.Base))">**基本工资必须填写且是数字**</small>
    </div>
    <br>

    <div class="p-field">
      <label for="Description">备注:</label>
      <Textarea id="Description" v-model.trim="inputInfo.Description" rows="3" cols="20"/>
    </div>
    <br>
    <template #footer>
      <Button label="取消" icon="pi pi-times" class="p-button-text" @click="hideDialog"/>
      <Button label="保存" icon="pi pi-check" class="p-button-text" @click="saveInput"/>
    </template>
  </Dialog>

  <Dialog v-model:visible="isDeleteInputDialogOpen" :style="{width: '350px'}" header="确认" :modal="true">
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem"/>
      <span v-if="inputInfo">你确认要删除 <b>{{ inputInfo.InputTime }}</b>?</span>
    </div>
    <template #footer>
      <Button label="否" icon="pi pi-times" class="p-button-text" @click="isDeleteInputDialogOpen = false"/>
      <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteInput"/>
    </template>
  </Dialog>
  <Dialog v-model:visible="tipDisplay" header="收入管理提示">{{ tipMessage }}</Dialog>
</template>

<script>
import InputService from '../api/input_mgmt.ts';

export default {
  name: "input_details",
  data() {
    return {
      inputInfoList: null,
      loading: true,
      filters: {},
      inputInfo: {},
      inputTime: new Date(),
      isNewInputDialogOpen: false,
      submitted: false,
      tipDisplay: false,
      tipMessage: "",
      isDeleteInputDialogOpen: false,
      isAddOperation: true
    }
  },
  inputService: null,
  created() {
    this.inputService = new InputService();
  },
  mounted() {
    this.inputService.getInputList(0, 0, 10000, 0).then(data => this.inputInfoList = data);
    this.loading = false;
  },
  methods: {
    addInputDialog() {
      this.inputInfo = {"Type": "salary"};
      this.inputTime = new Date();
      this.submitted = false;
      this.isNewInputDialogOpen = true;
      this.isAddOperation = true;
    },

    hideDialog() {
      this.isNewInputDialogOpen = false;
      this.submitted = false;
    },
    async saveInput() {
      this.submitted = true;
      if (this.isAddOperation) {
        this.inputInfo.Year = this.inputTime.getFullYear()
        this.inputInfo.Month = this.inputTime.getMonth() + 1
        let day = this.inputTime.getDate()

        this.inputInfo.InputTime = this.inputTime.getFullYear() * 10000 + (this.inputTime.getMonth() + 1) * 100 + day
      }

      if (!this.inputInfo.Type
          || !this.inputInfo.AllInput || isNaN(this.inputInfo.AllInput)
          || !this.inputInfo.Actual || isNaN(this.inputInfo.Actual)
          || !this.inputInfo.Base || isNaN(this.inputInfo.Base)) {
        return
      }

      try {
        let allInput = Number(this.inputInfo.AllInput)
        this.inputInfo.AllInput = allInput

        let actual = Number(this.inputInfo.Actual)
        this.inputInfo.Actual = actual

        this.inputInfo.Base = Number(this.inputInfo.Base)

        this.inputInfo.Tax = allInput - actual
      } catch (error) {
        console.log("failed to convert number")
        return
      }

      let res = ""
      if (this.isAddOperation) {
        res = await this.inputService.addInput(this.inputInfo);
      } else {
        res = await this.inputService.updateInput(this.inputInfo)
      }

      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this.tipDisplay = true;
        this.tipMessage = "操作失败，请检查输入信息！";
      } else {
        this.isNewInputDialogOpen = false;
        this.submitted = false;
        if (!this.isAddOperation) {
          this.inputInfoList = this.inputInfoList.filter(val => val.InputTime !== this.inputInfo.InputTime);
        }
        this.inputInfoList.push(this.inputInfo)
        this.inputInfoList.sort(function(a,b){return b.InputTime - a.InputTime})
        this.inputInfo = {};
      }
    },

    openDeleteInputDialog(inputInfo) {
      this.inputInfo = inputInfo.data;
      this.isDeleteInputDialogOpen = true;
    },
    async deleteInput() {
      this.isDeleteInputDialogOpen = false;
      let res = await this.inputService.deleteInput(this.inputInfo)
      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this.tipDisplay = true;
        this.tipMessage = "删除失败！";
      } else {
        this.inputInfoList = this.inputInfoList.filter(val => val.InputTime !== this.inputInfo.InputTime);
        this.inputInfo = {};
      }
    },

    openUpdateInputDialog(inputInfo) {
      this.isAddOperation = false
      this.inputInfo = {...inputInfo.data};
      this.submitted = false;
      this.isNewInputDialogOpen = true;
    }
  }
}
</script>

<style scoped>

</style>