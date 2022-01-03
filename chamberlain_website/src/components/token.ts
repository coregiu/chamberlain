import {ajax} from "../api/ajax";

const X_AUTH_TOKEN = "x-auth-token";
const EXPIRES = "expires";
const ROLE = "role";

export const token = {
    methods: {
        async checkChamberlainToken() {
            let token = this.getToken()
            if (token == "") {
                console.log("X_AUTH_TOKEN is empty!")
                return false;
            }
            let headers = {
                "Content-Type" : "application/json;charset=UTF-8",
                "X-AUTH-TOKEN" : token
            }
            let res = await ajax.get("/api/users/token", headers)
            return !((typeof res == "string") && (res.toLocaleString().indexOf("err:") === 0));
        },
        getToken() {
            let cookieValue = document.cookie;
            if (cookieValue == "") {
                console.log("cookieValue is empty!")
                return "";
            }
            return this.getCookieValue(X_AUTH_TOKEN, cookieValue);
        },
        getRole() {
            let cookieValue = document.cookie;
            if (cookieValue == "") {
                console.log("cookieValue is empty!")
                return "";
            }
            return this.getCookieValue(ROLE, cookieValue)
        },
        setChamberlainToken(token, role) {
            let date = new Date();
            date.setTime(date.getTime() + (24 * 60 * 60 * 1000));
            document.cookie = X_AUTH_TOKEN + "=" + token + ";" + EXPIRES + "=" + date.toUTCString()
            document.cookie = ROLE + "=" + role + ";" + EXPIRES + "=" + date.toUTCString()
        },
        clearChamberlainToken() {
            let date = new Date();
            document.cookie = X_AUTH_TOKEN + "=;" + EXPIRES + "=" + date.toUTCString();
            document.cookie = ROLE + "=;" + EXPIRES + "=" + date.toUTCString();
        },
        getCookieValue(cname, cookieValue) {
            let name = cname + "=";
            let item = cookieValue.split(";");
            for(let iLoop = 0; iLoop < item.length; iLoop++) {
                let aItem = item[iLoop].trim();
                if (aItem.indexOf(name) == 0){
                    return aItem.substring(name.length, aItem.length);
                }
            }
            return "";
        }
    },
    name: "token"
}