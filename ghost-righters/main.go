package main

import (
	// "fmt"
	"encoding/json"
	"html/template"
	// "io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
	// "github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	// "google.golang.org/appengine"
	// "google.golang.org/appengine/datastore"
	// "google.golang.org/appengine/log"
)

type User struct {
	Email string
	Alias string
}
type Post struct {
	Email    string
	Alias    string
	Text     string
	Category string
	Time     time.Time
}

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Index)
	r.GET("/home/", Home)
	r.POST("/home/post", MakePost)
	r.POST("/home/getposts", GetPosts)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // maybe not needed b/c of schmidt router
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "index.html", nil)
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(req)
	u := user.Current(c)
	email := u.String()
	//	url, _ := user.LogoutURL(c, "/")
	if u == nil {
		url, err := user.LoginURL(c, req.URL.String())
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("Location", url)
		res.WriteHeader(http.StatusFound)
		return
	}
	key := datastore.NewKey(c, "Users", email, 0, nil)
	var us User
	err := datastore.Get(c, key, &us)
	log.Infof(c, "ERR: %v", err)
	if err != nil {
		user := User{
			Email: email,
			Alias: Getalias(),
		}
		key := datastore.NewKey(c, "Users", user.Email, 0, nil)
		key, err = datastore.Put(c, key, &user)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		tpl.ExecuteTemplate(res, "home.html", &user)
	} else {
		user := User{
			Email: us.Email,
			Alias: us.Alias,
		}
		tpl.ExecuteTemplate(res, "home.html", &user)
	}
	return
}

func MakePost(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(req)
	u := user.Current(c)
	email := u.String()
	post := Post{
		Email:    email,
		Alias:    req.FormValue("Alias"),
		Text:     req.FormValue("Post"),
		Category: req.FormValue("Category"),
		Time:     time.Now(),
	}

	//	key := datastore.NeweKey(c, "Posts", email, 0, nil)
	k := datastore.NewIncompleteKey(c, "Posts", nil)
	_, err := datastore.Put(c, k, &post)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

func GetPosts(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(req)
	var posts []Post
	q := datastore.NewQuery("Posts").Order("-Time").Limit(25)

	_, err := q.GetAll(c, &posts)
	js, err := json.Marshal(posts)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}
