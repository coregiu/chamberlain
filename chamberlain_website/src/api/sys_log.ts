import {ajax} from "./ajax";
import {token} from "../components/token";

export default class SyslogService {
    getSyslogList(username, operation, limit, offset) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        return ajax.get("/api/syslogs?username=" + username + "&operation=" + operation + "&limit=" + limit + "&offset=" + offset, headers)
    }

    deleteSyslog(syslogs) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let data = JSON.stringify(syslogs);
        return ajax.delete("/api/syslogs", headers, data)
    }
}