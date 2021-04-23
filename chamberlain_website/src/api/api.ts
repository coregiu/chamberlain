import {ajax} from "./axios";

const headers = {"Content-Type" : "application/json;charset=UTF-8"};

export const login = function (username, password) {
    let data = {"Username":username, "Password": password}
    return ajax.post("/api/users/login", headers, data)
}