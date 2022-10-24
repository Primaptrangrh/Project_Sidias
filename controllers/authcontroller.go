package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
	"github.com/jeypc/go-auth/libraries"
	"github.com/jeypc/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()
var permissionModel = models.NewPermissionModel()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, data)
		}

	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	fmt.Print(session)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/admin_dashboard.html")
			temp.Execute(w, data)
			
		}

	}
}

func AdminHome(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	fmt.Print(session)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/admin_home.html")
			temp.Execute(w, data)
			
		}

	}
}

func Permission(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/permission.html")
			temp.Execute(w, data)
			fmt.Print(data)
		}

	}
}

func Calendar(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/calendar.html")
			temp.Execute(w, data)
		}

	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// proses login
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		errorMessages := validation.Struct(UserInput)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {
			var user entities.User
			userModel.Where(&user, "username", UserInput.Username)

			var message error
			if user.Username == "" {
				message = errors.New("username atau password salah!")
			} else {
				// pengecekan password
				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
				if errPassword != nil {
					fmt.Print(errPassword.Error())	
				}
			}

			if message != nil {

				data := map[string]interface{}{
					"error": message,
				}

				temp, _ := template.ParseFiles("views/login.html")
				temp.Execute(w, data)
			} else {
				// set session
				session, _ := config.Store.Get(r, config.SESSION_ID)

				session.Values["loggedIn"] = true
				session.Values["email"] = user.Email
				session.Values["username"] = user.Username
				session.Values["nama_lengkap"] = user.NamaLengkap
				session.Values["role"] = user.Role
				session.Save(r, w)
				// http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				// fmt.Print(session)
				if session.Values["role"] != "User" {
					http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
					
				} else if session.Values["role"] == "User"{
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			} 
						
			}
			
		}

	}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}


func LogbookAdmin(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, _ := template.ParseFiles("views/admin_logbook.html")
			temp.Execute(w, data)
		}

	}
}

func ActiveEmployee(w http.ResponseWriter, r *http.Request) {

	user, _ := userModel.FindAll()

	data := map[string]interface{}{
		"user": user,
	}

	temp, err := template.ParseFiles("views/active_user.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}


func AddUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/add_user.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// melakukan proses registrasi

		// mengambil inputan form
		r.ParseForm()

		user := entities.User{
			NamaLengkap: r.Form.Get("nama_lengkap"),
			Email:       r.Form.Get("email"),
			Username:    r.Form.Get("username"),
			Password:    r.Form.Get("password"),
			Cpassword:   r.Form.Get("cpassword"),
			Role:   r.Form.Get("role"),
		}

		errorMessages := validation.Struct(user)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}

			temp, _ := template.ParseFiles("views/add_user.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			// insert ke database
			userModel.Create(user)

			data := map[string]interface{}{
				"pesan": "Registrasi berhasil",
			}
			temp, _ := template.ParseFiles("views/add_user.html")
			temp.Execute(w, data)
		}
	}

}

func EditUser(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var user entities.User
		userModel.Find(id, &user)

		data := map[string]interface{}{
			"user": user,
		}

		temp, err := template.ParseFiles("views/edit_user.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var user entities.User
		user.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		user.NamaLengkap = request.Form.Get("nama_lengkap")
		user.Email = request.Form.Get("email")
		user.Username = request.Form.Get("username")
		user.Password = request.Form.Get("password")
		user.Cpassword = request.Form.Get("cpassword")
		user.Role = request.Form.Get("role")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(user)

		if vErrors != nil {
			data["user"] = user
			data["validation"] = vErrors
		} else {
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)
			data["pesan"] = "Data user berhasil diperbarui"
			userModel.Update(user)
		}

		temp, _ := template.ParseFiles("views/edit_user.html")
		temp.Execute(response, data)
		fmt.Print(data)
	}

}

func DeleteUser(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	userModel.Delete(id)

	http.Redirect(response, request, "/active_user", http.StatusSeeOther)
}

func AdminPermission(w http.ResponseWriter, r *http.Request) {

	permission, _ := permissionModel.FindAllPermission()

	data := map[string]interface{}{
		"permission": permission,
	}

	temp, err := template.ParseFiles("views/adminpermission.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func AddPermission(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/addpermission.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// melakukan proses registrasi

		// mengambil inputan form
		r.ParseForm()

		permission := entities.Permission{
	
			NamaLengkap: r.Form.Get("nama_lengkap"),
			Email:       r.Form.Get("email"),
			Departemen:    r.Form.Get("departement"),
			Position:    r.Form.Get("position"),
			Reason:   r.Form.Get("reason"),
		}

		errorMessages := validation.Struct(permission)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"permission":       permission,
			}

			temp, _ := template.ParseFiles("views/addpermission.html")
			temp.Execute(w, data)
		} else {

			// insert ke database
			permissionModel.CreatePermission(permission)

			data := map[string]interface{}{
				"pesan": "Data permission berhasil ditambahkan",
			}
			temp, _ := template.ParseFiles("views/addpermission.html")
			temp.Execute(w, data)
			fmt.Print(data)
		}
	}
}