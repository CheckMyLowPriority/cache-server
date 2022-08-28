package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func main() {
	handler := new(Handler)
	if err := http.ListenAndServe(":8000", handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}

var templates = template.Must(template.ParseFiles("upload.html"))

// Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		return
	}
}

var AllFiles []string

func uploadFile(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//GET displays the upload form.
	case "GET":
		display(w, "upload", nil)

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//get the multipart reader for the request.
		reader, err := r.MultipartReader()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//copy each part to destination.
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}

			//if part.FileName() is empty, skip this iteration.
			if part.FileName() == "" {
				continue
			}

			dir, err := os.CreateTemp("temp-images", "upload-*.png")
			if err != nil {
				return
			}

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if _, err := io.Copy(dir, part); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			AllFiles = append(AllFiles, part.FileName())
		}
		//display success message.
		display(w, "upload", "Upload successful.")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getdocs(w http.ResponseWriter, r *http.Request) {
	for i := range AllFiles {
		fmt.Fprintf(w, "%s\n", AllFiles[i])
		//w.Write([]byte(AllFiles[i]))
	}
}
func CreateUser(login, password string) (users User) {
	users = User{Id: rand.Intn(10), Login: login, Pass: password}
	return
}

/*
func GetUsers(u User) (users []User){

		//users = []User{}
	}
*/
func Register(w http.ResponseWriter, r *http.Request) {
	ts, _ := template.ParseFiles("user.html")
	ts.Execute(w, nil)
	user := CreateUser(r.FormValue("login"), r.FormValue("password"))

	//db.Query("insert into users values ($1,$2)", user.Login, user.Pass)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("nothing")
	}

}
func Auth(w http.ResponseWriter, r *http.Request) {

}
func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/getdocs", getdocs)
	//http.HandleFunc("/user", do)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
