module chamberlain_mgmt

go 1.14

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/google/uuid v1.2.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/ugorji/go v1.2.4 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
)

replace github.com/russross/blackfriday/v2 v2.1.0 => gopkg.in/russross/blackfriday.v2 v2.1.0
