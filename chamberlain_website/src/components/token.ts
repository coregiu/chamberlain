const X_AUTH_TOKEN = "x-auth-token";
const EXPIRES = "expires";

export const token = {
    methods: {
        checkChamberlainToken() {
            if (this.getToken() == "") {
                console.log("X_AUTH_TOKEN is empty!")
                return false;
            }
            return true;
        },
        getToken() {
            let cookieValue = document.cookie;
            if (cookieValue == "") {
                console.log("cookieValue is empty!")
                return "";
            }
            return this.getCookieValue(X_AUTH_TOKEN, cookieValue);
        },
        setChamberlainToken(token) {
            let date = new Date();
            date.setTime(date.getTime() + (24 * 60 * 60 * 1000));
            document.cookie = X_AUTH_TOKEN + "=" + token + ";" + EXPIRES + "=" + date.toUTCString();
        },
        clearChamberlainToken() {
            let date = new Date();
            document.cookie = X_AUTH_TOKEN + "=;" + EXPIRES + "=" + date.toUTCString();
        },
        showUnloginMenu() {
            document.getElementById("unLoginTop").style.display = "block";
            document.getElementById("loginTop").style.display = "none";
        },
        showLoginMenu() {
            document.getElementById("unLoginTop").style.display = "none";
            document.getElementById("loginTop").style.display = "block";
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