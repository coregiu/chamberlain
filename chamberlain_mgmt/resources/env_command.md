[run mysql]
docker run -d --name database -p 3306:3306 -v /giu/chamberlain/database:/var/lib/mysql -v /giu/chamberlain/database:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 mysql:latest

[build chamberlain]
go build
sudo GOOS=linux CGO_ENABLED=0 /usr/local/go/bin/go build -o apps -ldflags "-s -w" -i chamberlain_mgmt
upx -o app apps
docker build -t chamberlain:1.0 .

# delete temporary images.
docker rmi $(docker images -q -f dangling=true)   
docker image prune

docker run -d --name chamberlain -p 8080:8080 --link database:database  -v /var:/var chamberlain:1.0
docker run -it --rm --name chamberlaintest -p 8080:8080 --link database:database  -v /var:/var chamberlain:1.1

[podman]
docker pod create -n giu -p 8080:8080 -p 80:80
docker run -d --name database --pod giu -v /root/workspace/database:/var/lib/mysql -v /root/workspace/database:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=199527 docker.io/library/mysql:latest
docker run -d --name chamberlain -v /var:/var -v /giu/chamberlain/books:/giu/chamberlain/books -v /usr/bin/git:/usr/bin/git -v /usr/lib/git-core:/usr/lib/git-core -v /usr/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu -v ~/.gitconfig:/root/.gitconfig -v ~/.git-credentials:/root/.git-credentials -v /usr/share/git-core/templates:/usr/share/git-core/templates --pod giu localhost/chamberlain:1.7
docker run -d --name regiu --pod giu -v /giu/chamberlain/books:/usr/share/nginx/html/books localhost/regiu:1.5

docker run -d --name chamberlain -e MYSQL_ROOT_PASSWORD=199527 -v /var:/var -p 8080:8080 --link database:database --link redis:redis swr.cn-east-3.myhuaweicloud.com/coregiu/chamberlain:1.0.1
docker run -d -p 80:80 --link chamberlain:chamberlain --name regiu swr.cn-east-3.myhuaweicloud.com/coregiu/regiu:1.0.1
