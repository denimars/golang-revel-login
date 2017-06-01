package controllers


import "github.com/revel/revel"

func init(){
    revel.InterceptMethod((*ModelController).Begin, revel.BEFORE)
    revel.InterceptMethod((*ModelController).CreateBook, revel.BEFORE)
    revel.InterceptMethod((*ModelController).CreateUser, revel.BEFORE)
    revel.InterceptMethod(App.AddUser, revel.BEFORE)
    revel.InterceptMethod(Book.cekUser, revel.BEFORE)
}