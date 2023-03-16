package db

import "time"

// To store user login data
type UserDetails struct {
	Name  string
	Email string
	Pass  string
}

// Using map key as email and struct ass value
var DataBase = make(map[string]UserDetails)

// each session contains the username of the user and  the time at which it expire.
type Session struct {
	Username string
	Expiry   time.Time
}

// method to determine if the session has expired
func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

// this map stores the usser sessions
var Sessions = make(map[string]Session)

// message struct for login form

type Messages struct {
	Color   string
	Message string
}

// this to pass message to login page
var LoginMessage = Messages{}

// struct to store all erorrs what user made when user registering using boolean value
// if a fields is false means its no error otherwise there is an error
type regFormErrors struct {
	ErrorName  bool
	ErrorEmail bool
	ErrorPass  bool
}

var RegError = regFormErrors{}

var Home = "home.html"
var HomePath = "templates/home.html"
var Login = "login.html"
var LoginPath = "templates/login.html"
var Register = "register.html"
var RegisterPath = "templates/register.html"

var CookieID string

// templates for if any error
var ErrorPage = "errorPage.html"
var ErrorPagePath = "templates/errorPage.html"
