[![Language](https://img.shields.io/badge/language-go-blue.svg)](README.md)
[![Language](https://img.shields.io/badge/language-vue-blue.svg)](README.md)
[![License](https://img.shields.io/badge/license-Apache2.0-brightgreen.svg)](LICENSE)

# Self website for code example and experies summary.

## Function:
1. Self blogs, synchronize from github repository and convert from markdown to html automaticly.
2. Self memorandum and diary, edit, display and search online.
3. Selt inputs management, analysis types of income and the trend of income.
4. Has completable system management, like user management, auth check, log management, system backup and resoter and so on.
5. At the last, I want to implements a game of my hometown named the cannon bombards the dogfaces.

## Get started:
1. Install golang, node, npm, git.
2. Clone code.
3. Compile the chamberlain_website and make docker image:
```shell
npm run build
docker build -t regiu:1.0 .
```
4. Compile the chamberlain_mtmt and make docker image:
```shell
docker build -t chamberlain:1.0 .
```
5. Set the image version to docker-compose.yml
6. Pull mysql image
7. Run docker compose:
```shell
docker-compose up&
```
8. Docker exec the mysql container, init database and tables. The sql is chamberlain.sql.
9. Then you can open the page in browse and operate it.