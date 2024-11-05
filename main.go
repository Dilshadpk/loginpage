package main

import (
	"net/http"
	"text/template"

	"github.com/icza/session"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("templates/*.html"))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-control","no-store,no-cache,must-revalidate,max-age=0")
	w.Header().Set("pragma","no-cache")
	w.Header().Set("Expires","0")

	temp.ExecuteTemplate(w, "index.html", nil)

}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "dilshad" && password == "dilu1234" {
		secsion := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{
				"username": username,
			},
		})
		session.Add(secsion, w)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)

	} else {
		data := map[string]interface{}{
			"error": "invalid username and password",
		}
		temp.ExecuteTemplate(w, "index.html", data)

	}

}
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-control","no-store,no-cache,must-revalidate,max-age=0")
	w.Header().Set("pragma","no-cache")
	w.Header().Set("Expires","0")
	sess := session.Get(r)
	if sess == nil {
	
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username := sess.CAttr("username")
	data := map[string]interface{}{
		"username": username,
	}
	temp.ExecuteTemplate(w, "welcome.html", data)


}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":1111", nil)

}