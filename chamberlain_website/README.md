npm install --save @popperjs/core
npm install --save fs-filesystem
npm install path

Run test:
~~~shell
npm run dev
~~~

Run build:
~~~shell
npm run build
docker build regiu:xx .
~~~ 

Run on server:
```angular2html
podman run -d --name chamberlain -v /var:/var --pod giu swr.cn-east-3.myhuaweicloud.com/coregiu/chamberlain:1.11
podman run -d --name regiu --pod giu swr.cn-east-3.myhuaweicloud.com/coregiu/regiu:1.11
```

TODO:
- 3 记事本／个人待办
- 4 日志
- 5 定制主页
- 6 备份恢复
- 7 手机版