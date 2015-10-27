package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	//	"io"
	//	"bytes"
	//	"google.golang.org/appengine/memcache"
)

var store = sessions.NewCookieStore([]byte("session-name"))

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string
}

var tpl *template.Template

func init() {
	// tpl, _ = template.ParseFiles("static/home.html", "static/login.html", "static/signup.html", "static/headersFooters.html")
	// tpl = template.New("roottemplate")
	// tpl = template.Must(tpl.ParseGlob("templates/*.html"))
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/fail", Fail)
	r.GET("/form/login", Login)
	r.GET("/form/signup", Signup)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // maybe not needed b/c of schmidt router
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func Fail(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "fail.html", nil)
}
func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
}

func checkUserName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	log.Infof(ctx, "REQUEST BODY: %v", sbs)
	var user User
	key := datastore.NewKey(ctx, "Users", sbs, 0, nil)
	err = datastore.Get(ctx, key, &user)
	// if there is an err, there is NO user
	log.Infof(ctx, "ERR: %v", err)
	if err != nil {
		// there is an err, there is a NO user
		fmt.Fprint(res, "false")
		return
	} else {
		fmt.Fprint(res, "true")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	NewUser := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: req.FormValue("password"),
	}
	q := datastore.NewQuery("Users").Filter("Email =", req.FormValue("email"))

	var count int = 0
	for t := q.Run(ctx); ; {
		var x User
		count++
		_, err1 := t.Next(&x)
		if err1 == datastore.Done {
			log.Infof(ctx, "ERR: %v", err1)
			log.Infof(ctx, "ERR: %v", count)
			break
		}
	}
	if count <= 1 {
		key := datastore.NewKey(ctx, "Users", NewUser.UserName, 0, nil)
		key, err := datastore.Put(ctx, key, &NewUser)

		if err != nil {
			log.Errorf(ctx, "error adding todo: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		http.Redirect(res, req, "/", 302)
	} else {
		http.Redirect(res, req, "/fail", 302)
	}
}
