package controllers

import (
    "golang-login-validation/app/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    r "github.com/revel/revel"
    "fmt"
)

var DB *gorm.DB

func init(){
    var err error
    DB, err = gorm.Open("mysql", "root:root@/book?charset=utf8&parseTime=True&loc=Local")
    if err !=nil{
        panic(err.Error)
    }else{
        fmt.Println("Berhasil terkoneksi")
    }
}

type ModelController struct{
    *r.Controller
    Orm *gorm.DB
}

func (c *ModelController) Begin() r.Result{
    c.Orm = DB
    return nil
}

func (c *ModelController) CreateBook() r.Result{
    c.Orm.CreateTable(models.Book{})
    return nil
}

func (c *ModelController) CreateUser() r.Result{
    c.Orm.CreateTable(models.User{})
    return nil
}

