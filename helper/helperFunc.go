package helper

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/Noush-012/Login-Page-Server/db"
)

// these funtions are used to help other function
func CreateTemplate(name string, path string) *template.Template {

	tmpl, err := template.New(name).ParseFiles(path)

	if CheckError(err, "template "+name) { //to check error
		//if found error parse error page
		errTmpl, _ := template.New(db.ErrorPage).ParseFiles(db.ErrorPagePath)
		return errTmpl
	}

	return tmpl
}

func CheckError(err error, name string) bool {
	if err != nil {
		fmt.Println("Error found at ", name)
		return true
	}

	return false
}

// Func to check session active for user
func SessionAndCookie(r *http.Request) (db.Session, bool) {

	if cookieVal, ok1 := GetCookieVal(r); ok1 { //get coookie val if cookie not get return false

		if session, ok2 := db.Sessions[cookieVal]; ok2 { //get session using cookie value other return flase and empty session

			if !session.IsExpired() { //check session is expired is expired then delete otherwise return that session
				return session, true
			}

			//delete sessoin if session is expired

			delete(db.Sessions, cookieVal)
		}
	}

	return db.Session{}, false //return nill session
}

// get cookie if need
func GetCookieVal(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("session") // creating cookie name in client side by name "session"

	if CheckError(err, " getting Cokkie") {
		fmt.Println("session and cokkie func error to get cookie")
		return "", false
	}
	db.CookieID = cookie.Value
	return cookie.Value, true
}
