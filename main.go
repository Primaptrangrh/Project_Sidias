package main

import (
	"fmt"
	"html/template"
	"net/http"

	authcontroller "github.com/jeypc/go-auth/controllers"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("views/*.html"))
}
func RunIndex(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}
func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/dashboard", authcontroller.Dashboard)
	http.HandleFunc("/permission", authcontroller.Permission)
	http.HandleFunc("/calendar", authcontroller.Calendar)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/adminlogbook", authcontroller.LogbookAdmin)
	http.HandleFunc("/admin_home", authcontroller.AdminHome)
	http.HandleFunc("/active_user", authcontroller.ActiveEmployee)
	http.HandleFunc("/FindAllUser", authcontroller.ActiveEmployee)
	http.HandleFunc("/adduser", authcontroller.AddUser)
	http.HandleFunc("/edituser", authcontroller.EditUser)
	http.HandleFunc("/deleteuser", authcontroller.DeleteUser)
	http.HandleFunc("/adminpermission", authcontroller.AdminPermission)
	http.HandleFunc("/FindAllPermission", authcontroller.AdminPermission)
	http.HandleFunc("/addpermission", authcontroller.AddPermission)

	fmt.Println("Server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
