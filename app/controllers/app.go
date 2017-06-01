package controllers

import (
	"github.com/revel/revel"
	"golang-login-validation/app/models"
	"golang-login-validation/app/routes"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"strings"
)

type App struct {
	ModelController
}

func (c App) AddUser() revel.Result{
	if user := c.connected(); user != nil{
		c.ViewArgs["user"] = user
	}
	return nil
}

func (c App) connected() *models.User{
	c.cekSessionUser(c.Session["user"])
	if username, ok := c.Session["user"]; ok{
		return c.getUser(username)
	}

	if c.ViewArgs["user"] != nil{
		return c.ViewArgs["user"].(*models.User)
	}
	
	return nil
}

func (c App) cekSessionUser(username string){
	var user models.User
	err := c.Orm.Where("username=?", username).Find(&user)
	if err.RecordNotFound(){
		c.Logout()
	}
}

func (c App) getUser(username string) *models.User{
	var user models.User
	err := c.Orm.Where("username=?", username).Find(&user)
	if err.RecordNotFound(){
		fmt.Println("Not Found")
	}
	return &user
}

func (c App) Index() revel.Result {
	if c.connected() != nil{
		return c.Redirect(routes.Book.Index())
	}
	return c.Render()
}


func (c App) Register() revel.Result{
	return c.Render()
}

func (c App) Save(user models.User,password string, cekpassword string) revel.Result{
	c.Validation.Required(password).Message("Required pwd broooh")
	c.Validation.MaxSize(password, 255)
	c.Validation.MinSize(password, 7)
	c.Validation.Required(cekpassword).Message("Required Brooh")
	c.Validation.Required(cekpassword==password).Message("Not same")
	user.Validate(c.Validation)
	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Register())
	}
	user.Password, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := c.Orm.Save(&user)
	if err.Error != nil{
		c.Flash.Error("user atas nama "+user.Username+" sudah digunakan")
		return c.Redirect(routes.App.Register())
	}
	c.Session["user"] = user.Username
	c.Flash.Success("welcome, "+strings.Title(user.Username))
	return c.Redirect(routes.Book.Index())

}

func (c App) Login(username string, password string, remember bool) revel.Result{
	user := c.getUser(username)
	if user != nil{
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err == nil{
			c.Session["user"] = username
			if remember{
				c.Session.SetDefaultExpiration()
			}else{
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("welcome "+ strings.Title(username))
			return c.Redirect(routes.Book.Index())
		}
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login Failed")
	return c.Redirect(routes.App.Index())
}

func (c App) Logout() revel.Result{
	for k:= range c.Session{
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}