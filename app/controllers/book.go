package controllers

import (
    "github.com/revel/revel"
    "golang-login-validation/app/models"
    "golang-login-validation/app/routes"
    "fmt"
    "strings"
)

type Book struct{
    App
}

func (c Book) cekUser() revel.Result{
    if user := c.connected(); user == nil{
        return c.Redirect(routes.App.Index())
    }
    return nil
}

func (c Book) Index() revel.Result{
    var books []models.Book
    book := c.Orm.Find(&books)
    if book.Error != nil {
        panic(book.Error)
    }
    name := strings.ToTitle(c.connected().Name)
    return c.Render(books, name)
}

func (c Book) New() revel.Result{
    return c.Render()
}

func (c Book) Save(book models.Book) revel.Result{
    c.Validation.Required(book.Title).Message("Tidak boleh kosong")
    c.Validation.Required(book.Description).Message("Tidak boleh kosong")
    book.Validate(c.Validation)
    if c.Validation.HasErrors(){
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(routes.Book.New())
    }
    book.Title = strings.ToUpper(book.Title)
    book.Description = strings.ToUpper(book.Description)
    err := c.Orm.Save(&book)
    if err.Error != nil{
        panic(err.Error)
    }
    return c.Redirect(routes.Book.Index())
}

func (c Book) Edit(id int) revel.Result{
    book := c.getBook(id)
    fmt.Println(book)
    return c.Render(book)
}

func (c Book) Update(id int, title string, description string) revel.Result{
    book := c.getBook(id)
    book.Title = strings.ToUpper(title)
    book.Description = strings.ToUpper(description)
    err := c.Orm.Save(&book)
    if err.Error != nil{
        panic(err.Error)
    }
    return c.Redirect(routes.Book.Index())
}

func (c Book) Show(id int) revel.Result  {
    book := c.getBook(id)
    return c.Render(book)
}

func (c Book) Delete(id int) revel.Result{
    var buku models.Book
    err := c.Orm.Where("ID=?", id).Delete(&buku)
    if err.Error != nil{
        panic(err.Error)
    }
    return c.Redirect(routes.Book.Index())

}


func (c Book) getBook(id int) *models.Book{
    var books models.Book
    book := c.Orm.Find(&books, id)
    if book.Error != nil{
        panic(book.Error)
    }
    return &books
}
