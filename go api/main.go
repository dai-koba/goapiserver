package main

import (

    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

   _ "github.com/mattn/go-sqlite3"

)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")

    dbInit()

    //Index
    router.GET("/", func(ctx *gin.Context) {
        users := dbGetAll()
        ctx.HTML(200, "index.html", gin.H{
            "users": users,
        })
    })

    //Create
    router.POST("/new", func(ctx *gin.Context) {
        name := ctx.PostForm("name")
        age := ctx.PostForm("age")
        email := ctx.PostForm("email")
        dbInsert(name, age,email)
        ctx.Redirect(302, "/")
    })

    //Detail
    router.GET("/detail/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        user := dbGetOne(id)
        ctx.HTML(200, "detail.html", gin.H{"user": user})
    })

    //Update
    router.POST("/update/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        name := ctx.PostForm("name")
        age := ctx.PostForm("age")
        email := ctx.PostForm("email")
        dbUpdate(id, name, age, email)
        ctx.Redirect(302, "/")
    })

    //削除確認
    router.GET("/delete_check/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        user := dbGetOne(id)
        ctx.HTML(200, "delete.html", gin.H{"user": user})
    })

    //Delete
    router.POST("/delete/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        dbDelete(id)
        ctx.Redirect(302, "/")

    })

    router.Run()
}



type User struct {
    gorm.Model
    Name   string
    Age  string
    Email string
}

//DB初期化
func dbInit() {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInit）")
    }
    db.AutoMigrate(&User{})
    defer db.Close()
}

//DB追加
func dbInsert(name string, age string,email string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbInsert)")
    }
    db.Create(&User{Name: name, Age: age,Email: email})
    defer db.Close()
}

//DB更新
func dbUpdate(id int, name string, age string,email string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbUpdate)")
    }
    var user User
    db.First(&user, id)
    user.Name = name
    user.Age = age
    user.Email=email
    db.Save(&user)
    db.Close()
}

//DB削除
func dbDelete(id int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！（dbDelete)")
    }
    var user User
    db.First(&user, id)
    db.Delete(&user)
    db.Close()
}

//DB全取得
func dbGetAll() []User {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！(dbGetAll())")
    }
    var users []User
    db.Order("created_at desc").Find(&users)
    db.Close()
    return users
}

//DB一つ取得
func dbGetOne(id int) User {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("データベース開けず！(dbGetOne())")
    }
    var user User
    db.First(&user, id)
    db.Close()
    return user
}