# Go Gin REST API with JWT

Product REST API with CRUD & Login & Register

- Authentication
- Authorization multi-level user
- Authorization access product by ID

Flow:

user register -> user login -> login success -> generate token

admin -> all product authorization -> CRUD product

user -> product authorization by ID -> create & read product

## Dependencies

- `go get github.com/asaskevich/govalidator`
- `go get github.com/golang-jwt/jwt/v5`
- `go get github.com/gin-gonic/gin`
- `go get golang.org/x/crypto`
- `go get gorm.io/driver/postgres`
- `go get gorm.io/gorm`
- `go get github.com/joho/godotenv`

## Setup DB (Postgres)

- Login psql: `psql -U postgres`
- Show databases: `\list` or `\l`
- Create database: `CREATE DATABASE db_go_jwt_api;`
- Select database: `\c db_go_jwt_api`
- Show tables: `dt`

## Create random string

`openssl rand -base64 32`
