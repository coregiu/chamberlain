podman run -d --name regiu --pod giu \
    -v /giu/chamberlain/books:/usr/share/nginx/html/books \
    swr.cn-east-3.myhuaweicloud.com/coregiu/regiu:1.8

podman run -d --name chamberlain \n
    -v /var:/var -v /giu/chamberlain/books:/giu/chamberlain/books \
    -v /usr/bin/git:/usr/bin/git \
    -v /usr/lib/git-core:/usr/lib/git-core \
    -v /usr/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu \
    -v ~/.gitconfig:/root/.gitconfig \
    -v ~/.git-credentials:/root/.git-credentials \
    -v /usr/share/git-core/templates:/usr/share/git-core/templates \
    --pod giu swr.cn-east-3.myhuaweicloud.com/coregiu/chamberlain:1.10