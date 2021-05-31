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
        "X-AUTH-TOKEN" : token.methods.getToken()
    }
    return ajax.post("/api/users/logout", headers, data)
}

export const refreshBlogs = function () {
    let data = {}
    let headers = {
        "Content-Type" : "application/json;charset=UTF-8",
        "X-AUTH-TOKEN" : token.methods.getToken()
    }
    return ajax.post("/api/blogs", headers, data)
}

export const getUserByToken = function () {
    let headers = {
        "Content-Type" : "application/json;charset=UTF-8",
        "X-AUTH-TOKEN" : token.methods.getToken()
    }
    return ajax.get("/api/users/token", headers)
}

export const resetPassword = function (username, password, newPassword) {
    let data = {"Username":username, "Password": password, "NewPassword": newPassword}
    let headers = {
        "Content-Type" : "application/json;charset=UTF-8",
        "X-AUTH-TOKEN" : token.methods.getToken()
    }
    return ajax.put("/api/users/password", headers, data)
}