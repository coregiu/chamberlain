export default class SelfFormat {
    formatDate(time) {
        let date = new Date(time)
        let year = date.getFullYear()
        let month = date.getMonth() + 1 //月份是从0开始的
        let day = date.getDate()
        return year + '-' + month + '-' + day
    }
}