import {ajax} from "./ajax";
import {token} from "../components/token";

export default class NoteSummaryService {
    getTreeNodes() {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        return ajax.get("/api/summarybooks", headers).then(res => this.buildNoteTree(res))
    }

    buildNoteTree(noteList) {
        let treeNodes = []
        noteList.forEach(function (aNoteSummary) {
            if (aNoteSummary.ParentBookId === '0') {
                treeNodes.push({
                    "key": aNoteSummary.BookId,
                    "label": aNoteSummary.BookName,
                    "data": aNoteSummary.BookName,
                    "icon": "pi pi-fw pi-calendar",
                    "children": []
                })
            } else {
                let appendNode = undefined
                treeNodes.forEach(function getAppendNode(aNode) {
                    if (aNode.key === aNoteSummary.ParentBookId) {
                        appendNode = aNode
                        return aNode
                    }

                    aNode.children.forEach(aChildren => {
                        let childAppendNode = getAppendNode(aChildren)
                        if (childAppendNode !== undefined) {
                            appendNode = childAppendNode
                            return childAppendNode
                        }
                    })
                    return undefined
                })

                if (appendNode === undefined) {
                    return
                }

                appendNode.children.push({
                    "key": aNoteSummary.BookId,
                    "label": aNoteSummary.BookName,
                    "data": aNoteSummary.BookName,
                    "icon": "pi pi-fw pi-calendar-plus",
                    "children": []
                })
            }
        })
        return treeNodes
    }

    getNoteSummaryContent(bookId) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        return ajax.get("/api/summarybooks/content?book_id=" + bookId, headers)
    }

    addNoteSummary(notebookInfo) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let notebooks = [notebookInfo]
        let data = JSON.stringify(notebooks);
        return ajax.post("/api/notebooks", headers, data)
    }

    addBatchNoteSummaries(notebookInfoList) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let data = JSON.stringify(notebookInfoList);
        return ajax.post("/api/notebooks", headers, data)
    }

    deleteNoteSummary(notebookInfo) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let data = JSON.stringify(notebookInfo);
        return ajax.delete("/api/notebooks", headers, data)
    }

    updateNoteSummary(notebookInfo) {
        let headers = {
            "Content-Type": "application/json;charset=UTF-8",
            "X-AUTH-TOKEN": token.methods.getToken()
        }
        let data = JSON.stringify(notebookInfo);
        return ajax.put("/api/notebooks", headers, data)
    }
}