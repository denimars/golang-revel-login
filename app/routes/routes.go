// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tModelController struct {}
var ModelController tModelController


func (_ tModelController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ModelController.Begin", args).URL
}

func (_ tModelController) CreateBook(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ModelController.CreateBook", args).URL
}

func (_ tModelController) CreateUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ModelController.CreateUser", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddUser", args).URL
}

func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Register", args).URL
}

func (_ tApp) Save(
		user interface{},
		password string,
		cekpassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "cekpassword", cekpassword)
	return revel.MainRouter.Reverse("App.Save", args).URL
}

func (_ tApp) Login(
		username string,
		password string,
		remember bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "remember", remember)
	return revel.MainRouter.Reverse("App.Login", args).URL
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).URL
}


type tBook struct {}
var Book tBook


func (_ tBook) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Book.Index", args).URL
}

func (_ tBook) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Book.New", args).URL
}

func (_ tBook) Save(
		book interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "book", book)
	return revel.MainRouter.Reverse("Book.Save", args).URL
}

func (_ tBook) Edit(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Book.Edit", args).URL
}

func (_ tBook) Update(
		id int,
		title string,
		description string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "description", description)
	return revel.MainRouter.Reverse("Book.Update", args).URL
}

func (_ tBook) Show(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Book.Show", args).URL
}

func (_ tBook) Delete(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Book.Delete", args).URL
}


