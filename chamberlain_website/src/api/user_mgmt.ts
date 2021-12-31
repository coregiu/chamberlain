import {ajax} from "./ajax";
import {token} from "../components/token";

export default class UserService{
    getUserList(username, limit, offset) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        if (username != "") {
            return ajax.get("/api/users?username=" + username, headers)
        } else {
            return ajax.get("/api/users?limit=" + limit + "&offset=" + offset, headers)
        }
    }

    addUser(userInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(userInfo);
        return ajax.post("/api/users", headers, data)
    }

    deleteUser(userInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(userInfo);
        return ajax.delete("/api/users", headers, data)
    }

    updateUser(userInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(userInfo);
        return ajax.put("/api/users", headers, data)
    }
}