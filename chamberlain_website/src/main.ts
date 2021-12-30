import {createApp} from 'vue'
import router from './router'
import App from './App.vue'
import './index.scss'
import {token} from './components/token'

import PrimeVue from 'primevue/config'
import 'primeicons/primeicons.css'
import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.min.css'
// @ts-ignore
import ToastService from 'primevue/toastservice'
import {ajax} from './api/ajax'
import Checkbox from "primevue/checkbox";
import Chart from 'primevue/chart';
import InputText from 'primevue/inputtext';
import Timeline from "primevue/timeline";
import MultiSelect from "primevue/multiselect";
import OverlayPanel from "primevue/overlaypanel";
import Button from "primevue/button";
import Chips from "primevue/chips";
import Dialog from "primevue/dialog";
import Toast from "primevue/toast";
import InputSwitch from "primevue/inputswitch";
import ListBox from "primevue/listbox";
import FieldSet from "primevue/fieldset";
import Card from "primevue/card";
import Panel from "primevue/panel";
import DataView from "primevue/dataview";
import Avatar from "primevue/avatar";
import AutoComplete from "primevue/autocomplete";
import Tree from "primevue/tree";
import FileUpload from "primevue/fileupload";
import Toolbar from "primevue/toolbar";
import DataTable from "primevue/datatable";

const app = createApp(App);
app.use(router);
app.use(ToastService);
app.use(PrimeVue);

app.component("Checkbox", Checkbox);
app.component("Chart", Chart);
app.component("MultiSelect", MultiSelect);
app.component("InputText", InputText);
app.component("OverlayPanel", OverlayPanel);
app.component("Button", Button);
app.component("Timeline", Timeline);
app.component("Card", Card);
app.component("FieldSet", FieldSet);
app.component("ListBox", ListBox);
app.component("InputSwitch", InputSwitch);
app.component("Toast", Toast);
app.component("Dialog", Dialog);
app.component("Chips", Chips);
app.component("Panel", Panel);
app.component("DataView", DataView);
app.component("Avatar", Avatar);
app.component("AutoComplete", AutoComplete);
app.component("Tree", Tree);
app.component("FileUpload", FileUpload);
app.component("Toolbar", Toolbar);
app.component("DataTable", DataTable);

app.config.globalProperties.$axios = ajax;
app.config.globalProperties.$token = token;
app.config.globalProperties.$router = router;
app.mount('#app');

if (!token.methods.checkChamberlainToken()) {
    token.methods.showUnloginMenu();
} else {
    token.methods.showLoginMenu();
}