import axios from 'axios'

const service = axios.create({
    timeout: 18000
})

export default class TreeNode {
    getTreeNodes() {
        return service.get('../../books/books.json').then(res => res.data.root);
    }
}