package controls

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Noush-012/Login-Page-Server/db"
	"github.com/Noush-012/Login-Page-Server/helper"
	"github.com/google/uuid"
)

// To render register Page
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	if _, ok := helper.SessionAndCookie(r); ok {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-control", "no-cache, no-store, must-revalidate")
	regTempl := helper.CreateTemplate(db.Register, db.RegisterPath)
	regTempl.Execute(w, db.RegError)
}

// Register submission
func RegisterSubmit(w http.ResponseWriter, r *http.Request) {
	if _, ok := helper.SessionAndCookie(r); ok {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//get values from form
	fName := r.PostFormValue("name")   //form name
	fEmail := r.PostFormValue("email") //form mail
	fPass1 := r.PostFormValue("fpass") //form first pass
	fPass2 := r.PostFormValue("spass") // form second pass

	// validate the user value
	if fName == "" {
		db.RegError.ErrorName = true
	}
	if fEmail == "" {
		db.RegError.ErrorEmail = true
	}
	if fPass1 == "" || fPass2 == "" || fPass1 != fPass2 {
		db.RegError.ErrorPass = true
	}

	//check if the user already exists
	if _, ok := db.DataBase[fEmail]; ok {

		db.LoginMessage.Color = "text-success"
		db.LoginMessage.Message = "You are already a User"

		http.Redirect(w, r, "/", http.StatusSeeOther) //render login page to show the message
		return
	}

	//check any true error in regError
	//then render same page with showing this errors
	if db.RegError.ErrorEmail || db.RegError.ErrorName || db.RegError.ErrorPass {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	//if all validation is completed add value to localDB and show the login page
	db.DataBase[fEmail] = db.UserDetails{
		Name:  fName,
		Email: fEmail,
		Pass:  fPass1,
	}
	//set login text class and message
	db.LoginMessage.Color = "text-success"
	db.LoginMessage.Message = "Successfully Registered. Please Login!"

	clearRegError()
	//redirect to login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func clearRegError() {
	db.RegError.ErrorEmail = false
	db.RegError.ErrorName = false
	db.RegError.ErrorPass = false
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginPage")

	if _, ok := helper.SessionAndCookie(r); ok {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//session is not avaliable then render login page
	clearRegError() //to clear register form errors

	tmpl := helper.CreateTemplate(db.Login, db.LoginPath)
	tmpl.Execute(w, db.LoginMessage)

	//clear all login error messages
	db.LoginMessage.Color = ""
	db.LoginMessage.Message = ""
}

func LoginSubmit(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login submit start")

	if _, ok := helper.SessionAndCookie(r); ok {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	// w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	userEmail := r.FormValue("email")
	userPass := r.FormValue("pass")

	//check user entered email
	if userEmail == "" {
		//setting
		db.LoginMessage.Color = "text-danger"
		db.LoginMessage.Message = "Enter Email Properly"
		//after setting message call login handler to render login page
		LoginPage(w, r)
		return
	}
	// check this email contains in our mapDB
	singleUser, ok := db.DataBase[userEmail]

	//check user is exist or not then check password
	if !ok {
		db.LoginMessage.Color = "text-danger"
		db.LoginMessage.Message = "You are not a registered User! you can register"

		LoginPage(w, r)
		return
	} else if userPass != singleUser.Pass { // user exist password not match
		db.LoginMessage.Color = "text-danger"
		db.LoginMessage.Message = "Incorrect Password"

		LoginPage(w, r)
		return
	}

	//create session
	sessionToken := uuid.NewString() //create a new random session token

	//sessionToken := "token"
	sessionTime := time.Now().Add(3 * time.Minute) //expire time current time plus two minute

	newSession := db.Session{
		Username: singleUser.Name,
		Expiry:   sessionTime,
	}

	//add this sessoin to session database
	db.Sessions[sessionToken] = newSession

	//set cookie
	newCookie := &http.Cookie{
		Name:    "session",
		Value:   sessionToken,
		Expires: sessionTime,
	}
	http.SetCookie(w, newCookie)

	HomePage(w, r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	session, ok := helper.SessionAndCookie(r)

	if !ok { //if session no availabe

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	homeTmpl := helper.CreateTemplate(db.Home, db.HomePath)
	homeTmpl.Execute(w, session)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout page ")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	_, ok := helper.SessionAndCookie(r)

	if ok {
		delete(db.Sessions, db.CookieID)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// this function check the user is entering invalid url from login page or home according to that function redirect to that page
func ErrorHandleFunc(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if _, ok := helper.SessionAndCookie(r); ok {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
