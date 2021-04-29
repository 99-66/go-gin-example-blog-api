# golang-gin example blog APIs

Golang Gin + Gorm Example blog API 

- [Gin WebFramework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)
- [Blog DB Modeling](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

## Set Environment
```shell
# if you want DEBUG mode: 
# export GIN_MODE=debug
export GIN_MODE=release
export GIN_PORT=8000

export MYSQL_HOST=mysql_host
export MYSQL_PORT=3306
export MYSQL_NAME=mysql_db_name
export MYSQL_USER=mysql_user_name
export MYSQL_PASSWORD=mysql_password
```
## Migration
If you haven't created the table yourself, uncomment the next line and run it.

Then gorm will automatically create the table
```shell
# services/database/mysql.go 43~54 line

## Before
	// AutoMigrate
	//err = db.AutoMigrate(&blog_m.User{},
	//&blog_m.Post{},
	//&blog_m.PostMeta{},
	//&blog_m.PostComment{},
	//&blog_m.Category{},
	//&blog_m.PostCategory{},
	//&blog_m.Tag{},
	//&blog_m.PostTag{},
	//)
	//if err != nil {
	//	return nil, err
	//}

## After
	//AutoMigrate
	err = db.AutoMigrate(&blog_m.User{},
	&blog_m.Post{},
	&blog_m.PostMeta{},
	&blog_m.PostComment{},
	&blog_m.Category{},
	&blog_m.PostCategory{},
	&blog_m.Tag{},
	&blog_m.PostTag{},
	)
	if err != nil {
		return nil, err
	}
```
##No Build running
```shell
go run main.go
```
##Build
```shell
go build main.go
./main.go
```