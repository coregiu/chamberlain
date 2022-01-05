<template>
  <DataTable :value="userInfoList" :paginator="true" class="p-datatable-customers" :rows="10"
             dataKey="Username" :rowHover="true" :filters="filters"
             filterDisplay="row"
             :loading="loading"
             paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
             :rowsPerPageOptions="[10,25,50]"
             currentPageReportTemplate="Showing {first} to {last} of {totalRecords} entries">

    <template #header>
      <div>
        <div style="float:left">用户列表</div>
        <div style="float:right">
          <Button type="button" class="p-button-secondary" @click="addUserDialog">添加用户</Button>&nbsp;&nbsp;&nbsp;&nbsp;
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

    <Column field="Username" header="用户名" :sortable="true" sortField="Username" filterField="Username"
            filterMatchMode="contains">
      <template #filter>
        <InputText type="text" v-model="filters['Username']" class="p-column-filter" placeholder="按用户名检索"/>
      </template>
    </Column>
    <Column field="Role" header="角色" :sortable="true" sortField="Role" filterField="Role" filterMatchMode="contains">
      <template #filter>
        <InputText type="text" v-model="filters['Role']" class="p-column-filter" placeholder="按角色检索"/>
      </template>
    </Column>
    <Column header="操作" headerStyle="width: 8rem; text-align: center" bodyStyle="text-align: center; overflow: visible">
      <template #body="userInfo">
        <Button type="button" icon="pi pi-pencil" class="p-button-secondary" title="修改"
                @click="openUpdateUserDialog(userInfo)"></Button>&nbsp;&nbsp;
        <Button type="button" icon="pi pi-trash" class="p-button-danger" title="删除"
                @click="openDeleteUserDialog(userInfo)"></Button>
      </template>
    </Column>
  </DataTable>

  <Dialog v-model:visible="isNewUserDialogOpen" :style="{width: '350px'}" header="用户信息" :modal="true" class="p-fluid">
    <div class="p-field">
      <label for="Username">用户名: <span class="p-invalid" v-if="!isAddOperation">{{userInfo.Username}}</span></label>
      <InputText id="Username" v-model.trim="userInfo.Username" required="true" autofocus
                 :class="{'p-invalid': submitted && !userInfo.Username}" v-if="isAddOperation"/>
      <small class="p-invalid" v-if="submitted && !userInfo.Username">**名称必须填写**</small>
    </div>
    <br>

    <div class="p-field">
      <label class="p-mb-3">角色:</label>
      <div class="p-formgrid p-grid">
        <div class="p-field-radiobutton p-col-6">
          <RadioButton id="category1" name="category" value="admin" v-model="userInfo.Role" :class="{'p-invalid': submitted && !userInfo.Role}"/>
          <label for="category1">管理员</label>&nbsp;&nbsp;&nbsp;&nbsp;
          <RadioButton id="category2" name="category" value="user" v-model="userInfo.Role" :class="{'p-invalid': submitted && !userInfo.Role}"/>
          <label for="category2">普通用户</label>
        </div>
      </div>
      <small class="p-invalid" v-if="submitted && !userInfo.Role">**角色必须填写**</small>
    </div>
    <br>

    <div class="p-field">
      <label for="Password">密码:</label>
      <InputText type="password" id="Password" v-model.trim="userInfo.Password" required="true" autofocus
                 :class="{'p-invalid': submitted && !userInfo.Password}"/>
      <small class="p-invalid" v-if="submitted && !userInfo.Password">**密码必须填写**</small>
    </div>
    <br>

    <div class="p-field">
      <label for="CPassword">确认密码:</label>
      <InputText type="password" id="CPassword" v-model.trim="cPassword" required="true" autofocus
                 :class="{'p-invalid': submitted && userInfo.Password !== cPassword}"/>
      <small class="p-invalid" v-if="submitted && userInfo.Password !== cPassword">**两次密码输入不一致**</small>
    </div>
    <template #footer>
      <Button label="取消" icon="pi pi-times" class="p-button-text" @click="hideDialog"/>
      <Button label="保存" icon="pi pi-check" class="p-button-text" @click="saveUser"/>
    </template>
  </Dialog>

  <Dialog v-model:visible="isDeleteUserDialogOpen" :style="{width: '350px'}" header="确认" :modal="true">
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem"/>
      <span v-if="userInfo">你确认要删除 <b>{{ userInfo.Username }}</b>?</span>
    </div>
    <template #footer>
      <Button label="否" icon="pi pi-times" class="p-button-text" @click="isDeleteUserDialogOpen = false"/>
      <Button label="是" icon="pi pi-check" class="p-button-text" @click="deleteUser"/>
    </template>
  </Dialog>
  <Dialog v-model:visible="tipDisplay" header="用户管理提示">{{ tipMessage }}</Dialog>
</template>

<script>
import UserService from '../api/user_mgmt.ts';

export default {
  data() {
    return {
      userInfoList: null,
      loading: true,
      filters: {},
      userInfo: {},
      cPassword: "",
      isNewUserDialogOpen: false,
      submitted: false,
      tipDisplay: false,
      tipMessage: "",
      isDeleteUserDialogOpen: false,
      isAddOperation: true
    }
  },
  userService: null,
  created() {
    this.userService = new UserService();
  },
  mounted() {
    this.userService.getUserList("", 100, 0).then(data => this.userInfoList = data);
    this.loading = false;
  },
  methods: {
    addUserDialog() {
      this.userInfo = {"Role": "user"};
      this.submitted = false;
      this.isNewUserDialogOpen = true;
      this.isAddOperation = true;
    },
    hideDialog() {
      this.isNewUserDialogOpen = false;
      this.submitted = false;
    },
    async saveUser() {
      this.submitted = true;
      if (this.userInfo.Username && this.userInfo.Role && this.cPassword !== "" && this.userInfo.Password === this.cPassword) {
        let res = ""
        if (this.isAddOperation) {
          res = await this.userService.addUser(this.userInfo);
        } else {
          res = await this.userService.updateUser(this.userInfo)
        }

        if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
          this.tipDisplay = true;
          this.tipMessage = "操作失败，请检查输入信息！";
        } else {
          this.isNewUserDialogOpen = false;
          if (!this.isAddOperation) {
            this.userInfoList = this.userInfoList.filter(val => val.Username !== this.userInfo.Username);
          }
          this.userInfoList.push(this.userInfo)
          this.userInfoList.sort(function(a, b){return a.Username.localeCompare(b.Username)})
          this.userInfo = {};
          this.cPassword = "";
        }
      }
    },

    openDeleteUserDialog(userInfo) {
      this.userInfo = userInfo.data;
      this.isDeleteUserDialogOpen = true;
    },
    async deleteUser() {
      this.isDeleteUserDialogOpen = false;
      let res = await this.userService.deleteUser(this.userInfo)
      if ((typeof res == "string") && (res.indexOf("err:") === 0)) {
        this.tipDisplay = true;
        this.tipMessage = "删除失败！";
      } else {
        this.userInfoList = this.userInfoList.filter(val => val.Username !== this.userInfo.Username);
        this.userInfo = {};
      }
    },

    openUpdateUserDialog(userInfo) {
      this.isAddOperation = false
      this.userInfo = {...userInfo.data};
      this.submitted = false;
      this.isNewUserDialogOpen = true;
    }
  }
}
</script>