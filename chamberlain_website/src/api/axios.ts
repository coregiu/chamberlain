import axios from 'axios'

const service = axios.create({
    timeout: 18000
})

const request = (url, method, headers, data) => {
    const options = Object.assign({}, {url, method, headers, data})
    options.headers = options.headers || {}
    return new Promise((resolve, reject) => {
        service.request(options).then(
            res => {
                const data = res.data;
                const status = res.status;
                console.log(res)
                if (status >= 200) {
                    resolve(data);
                }
            }
        ).catch(res => {
            console.error("Failed to send request");
            resolve("err:" + res);
        })
    })
}

export const ajax = {
    get(url, headers) {
        return request(url, 'get', headers, null,);
    },
    post(url, headers, data) {
        if (!headers) {
            headers = {
                'Content-Type': 'application/json;charset=UTF-8'
            }
        }
        return request(url, 'post', headers, data)
    },
    put(url, headers, data) {
        if (!headers) {
            headers = {
                'Content-Type': 'application/json;charset=UTF-8'
            }
        }
        return request(url, 'put', headers, data)
    },
    delete(url, headers, data) {
        if (!headers) {
            headers = {
                'Content-Type': 'application/json;charset=UTF-8'
            }
        }
        return request(url, 'delete', headers, data)
    }
}