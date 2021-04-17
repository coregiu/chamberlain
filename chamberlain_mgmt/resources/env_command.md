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
pd pod create -n giu -p 8080:8080
pd run -d --name database --pod giu -v /root/workspace/database:/var/lib/mysql -v /root/workspace/database:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=199527 docker.io/library/mysql:latest
pd run -d --name chamberlain -v /var:/var --pod giu localhost/chamberlain:1.0