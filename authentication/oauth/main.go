package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/twitter"
)

// // Hash keys should be at least 32 bytes long
// var hashKey = []byte("very-secret")

// // Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// // Shorter keys may weaken the encryption used.
// var blockKey = []byte("a-lot-secret")
// var s = securecookie.New(hashKey, blockKey)

var indexTemplate = `
<p><a href="/auth/gplus">Log in with Google Plus</a></p>
<p><a href="/auth/facebook">Log in with Facebook</a></p>
<p><a href="/auth/twitter">Log in with Twitter</a></p>`

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}" width="128" height="128"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>`

type Configuration struct {
	GPlusKey       string
	GPlusSecret    string
	FacebookKey    string
	FacebookSecret string
	TwitterKey     string
	TwitterSecret  string
}

func callbackAuthHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.New("userinfo").Parse(userTemplate)
	t.Execute(res, user)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.New("index").Parse(indexTemplate)
	t.Execute(res, nil)
}

var config Configuration

func init() {
	gothic.Store = sessions.NewCookieStore([]byte("very-secret"))
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	goth.UseProviders(gplus.New(
		config.GPlusKey, config.GPlusSecret, "http://localhost:8080/auth/gplus/callback"),
		facebook.New(config.FacebookKey, config.FacebookSecret, "http://localhost:8080/auth/facebook/callback"),
		twitter.New(config.TwitterKey, config.TwitterSecret, "http://localhost:8080/auth/twitter/callback"),
	)

	r := pat.New()
	r.Get("/auth/{provider}/callback", callbackAuthHandler)
	r.Get("/auth/{provider}", gothic.BeginAuthHandler)
	r.Get("/", indexHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
