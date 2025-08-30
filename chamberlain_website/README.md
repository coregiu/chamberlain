```shell
#使用NVM安装node 20版本
使用 NVM 安装（国内镜像）
NVM（Node Version Manager）是一个灵活的 Node.js 版本管理工具，可以通过国内镜像加速下载。

步骤：
1. **安装 NVM（使用国内镜像）**：
curl -o- https://gitee.com/mirrors/nvm/raw/v0.39.3/install.sh | bash
- 这里使用了 Gitee 的镜像源（`gitee.com/mirrors/nvm`）。

2. **重新加载终端配置**：
source ~/.bashrc

3. **安装 Node.js 20.19.0**：
NVM_NODEJS_ORG_MIRROR=https://npmmirror.com/mirrors/node nvm install 20.19.0
- 这里使用了淘宝的 Node.js 镜像源（`npmmirror.com`）。

4. **验证安装**：
node -v # 应输出 v20.19.0

5. **切换 Node.js 版本**（可选）：
nvm use 20.19.0
```
清缓存
```shell
npm cache clean --force
rm -rf apckage-lock.json
rm -rf node_modules
```

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
- 5 定制主页
- 7 手机版