docker run -d --name database -p 3306:3306 -v /giu/chamberlain/database:/var/lib/mysql -v /giu/chamberlain/database:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 mysql:latest
