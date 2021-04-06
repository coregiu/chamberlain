[run mysql]
docker run -d --name database -p 3306:3306 -v /giu/chamberlain/database:/var/lib/mysql -v /giu/chamberlain/database:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 mysql:latest

[build chamberlain]
go build
upx -o chamberlain chamberlain_mgmt
docker build -t chamberlain:1.0 .
docker run -d --name chamberlain -p 8080:8080 --link database:database  -v /var:/var chamberlain:1.0

[podman]
pd pod create -n giu -p 8080:8080
pd run -d --name database --pod giu -v /root/workspace/database:/var/lib/mysql -v /root/workspace/database:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=199527 docker.io/library/mysql:latest
pd run -d --name chamberlain -v /var:/var --pod giu localhost/chamberlain:1.0