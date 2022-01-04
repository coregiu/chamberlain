import {ajax} from "./ajax";
import {token} from "../components/token";

export default class InputService{
    getInputList(year, month, limit, offset) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        return ajax.get("/api/inputs?year=" + year + "&month=" + month + "&limit=" + limit + "&offset=" + offset, headers)
    }

    addInput(inputInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let inputs = [inputInfo]
        let data = JSON.stringify(inputs);
        return ajax.post("/api/inputs", headers, data)
    }

    addBatchInputs(inputInfoList) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(inputInfoList);
        return ajax.post("/api/inputs", headers, data)
    }

    deleteInput(inputInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(inputInfo);
        return ajax.delete("/api/inputs", headers, data)
    }

    updateInput(inputInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(inputInfo);
        return ajax.put("/api/inputs", headers, data)
    }
}