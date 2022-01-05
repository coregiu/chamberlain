import {ajax} from "./ajax";
import {token} from "../components/token";

export default class NotebookService{
    getNotebookList(finishTime, status, limit, offset) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        return ajax.get("/api/notebooks?finish_time=" + finishTime + "&status=" + status + "&limit=" + limit + "&offset=" + offset, headers)
    }

    addNotebook(notebookInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let notebooks = [notebookInfo]
        let data = JSON.stringify(notebooks);
        return ajax.post("/api/notebooks", headers, data)
    }

    addBatchNotebooks(notebookInfoList) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(notebookInfoList);
        return ajax.post("/api/notebooks", headers, data)
    }

    deleteNotebook(notebookInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(notebookInfo);
        return ajax.delete("/api/notebooks", headers, data)
    }

    updateNotebook(notebookInfo) {
        let headers = {
            "Content-Type" : "application/json;charset=UTF-8",
            "X-AUTH-TOKEN" : token.methods.getToken()
        }
        let data = JSON.stringify(notebookInfo);
        return ajax.put("/api/notebooks", headers, data)
    }
}