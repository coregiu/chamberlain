import {ajax} from "./ajax";
import {token} from "../components/token";

export const login = function (username, password) {
    let headers = {"Content-Type" : "application/json;charset=UTF-8"};
    let data = {"Username":username, "Password": password}
    return ajax.post("/api/users/login", headers, data)
}

export const logout = function () {
    let data = {}
    let headers = {
        "Content-Type" : "application/json;charset=UTF-8",
        "X-AUTH-TOKEN" : '"' + token.methods.getToken() + '"'
    }
    return ajax.post("/api/users/login", headers, data)
}

export const refreshBlogs = function () {
    let data = {}
    let headers = {
        "Content-Type" : "application/json;charset=UTF-8",
        "X-AUTH-TOKEN" : token.methods.getToken()
    }
    return ajax.post("/api/blogs", headers, data)
}