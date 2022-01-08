import {ajax} from "./ajax";
import {token} from "../components/token";

export default class NoteSummaryService {
    getTreeNodes() {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        return ajax.get("/api/notesummaryies", headers).then(res => {
            let treeNodes = []
            this.buildNoteTree(res, "0", treeNodes)
            return treeNodes
        })
    }

    buildNoteTree(noteList, parentBookId, treeNodes) {
        for(let aNoteSummary of noteList) {
            if (aNoteSummary.ParentBookId === parentBookId) {
                let classIcon = parentBookId === "0" ? "pi pi-fw pi-inbox" : "pi pi-fw pi-calendar"
                let subNode = {
                    "key": aNoteSummary.BookId,
                    "label": aNoteSummary.BookName,
                    "data": aNoteSummary.BookName,
                    "icon": classIcon,
                    "children": []
                };
                treeNodes.push(subNode)
                this.buildNoteTree(noteList, aNoteSummary.BookId, subNode.children)
            }
        }
    }

    getNoteSummaryContent(bookId) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        return ajax.get("/api/notesummaryies/content?book_id=" + bookId, headers)
    }

    addNoteSummary(noteSummaryInfo) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let noteSummaries = [noteSummaryInfo]
        let data = JSON.stringify(noteSummaries);
        return ajax.post("/api/notesummaryies", headers, data)
    }

    addBatchNoteSummaries(noteSummaryInfoList) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let data = JSON.stringify(noteSummaryInfoList);
        return ajax.post("/api/notesummaryies", headers, data)
    }

    deleteNoteSummary(noteSummaryInfo) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let data = JSON.stringify({"BookId": noteSummaryInfo.key});
        return ajax.delete("/api/notesummaryies", headers, data)
    }

    updateNoteSummary(noteSummaryInfo) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let data = JSON.stringify(noteSummaryInfo);
        return ajax.put("/api/notesummaryies", headers, data)
    }
}