package main
import (
	"fmt"
    "log"
    "os"
    "net/http"
    "html/template"
    "path"
    "golang-mongodb-htmltemplate/pkg/app"
    "golang-mongodb-htmltemplate/pkg/person"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

const COLLECTION = "people"

// Used for HTML Template (ex. .App.Version = app.Application.Version)
type vars struct {
    People []person.Person
    App app.Application
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

//index view handler
func index(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
    app := app.Application{Name: "golang-mongodb-htmltemplate", Version: "1.0.1"}
    
    //url params
    lName := r.URL.Query().Get("last_name")
    fName := r.URL.Query().Get("first_name")

    //mongodb
    session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
    check(err)
    defer session.Close()
    c := session.DB(os.Getenv("MONGODB_DB")).C(COLLECTION)
	var results []person.Person
    
    //decide if params are used
    switch {
        case lName != "" && fName != "":
            err = c.Find(bson.M{"last_name":bson.RegEx{"^"+lName+"$", "i"}, "first_name":bson.RegEx{"^"+fName+"$", "i"}}).All(&results)
        case lName != "":
            err = c.Find(bson.M{"last_name":bson.RegEx{"^"+lName+"$", "i"}}).All(&results)
        case fName != "":
            err = c.Find(bson.M{"first_name":bson.RegEx{"^"+fName+"$", "i"}}).All(&results)
        default:
            err = c.Find(nil).All(&results)
    }
    check(err)

    //set struct for HTML Template
    data := vars{People: results, App: app}
    
    if url == "" {            
        fp := path.Join("templates", "index.html")
        tmpl, err := template.ParseFiles(fp)
        check(err)
        if err := tmpl.Execute(w, data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

//about view handler
func about(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
    if url == "" {            
        fp := path.Join("templates", "about.html")
        tmpl, err := template.ParseFiles(fp)
        check(err)
        if err := tmpl.Execute(w, ""); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func main() {
	fmt.Println("Running local server @ http://localhost:" + os.Getenv("PORT"))
    http.HandleFunc("/", index)
    http.HandleFunc("/about", about)
    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}