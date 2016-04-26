package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

var noteStore = make(map[string]Note)
var templates map[string]*template.Template
var id int = 0

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/add.html", "templates/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/edit.html", "templates/base.html"))
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base", noteStore)
}
func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}
func deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		http.Error(w, "Could not find the resource to delete", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", 302)
}
func updateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpdate Note
	if note, ok := noteStore[k]; ok {
		r.ParseForm()
		noteToUpdate.Title = r.PostFormValue("title")
		noteToUpdate.Description = r.PostFormValue("description")
		noteToUpdate.CreatedOn = note.CreatedOn
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
	} else {
		http.Error(w, "Could not find resource", http.StatusBadRequest)
	}
	// renderTemplate(w, "index", "base", noteStore)
	http.Redirect(w, r, "/", 302)
}

type EditNote struct {
	Note
	Id string
}

func editNote(w http.ResponseWriter, r *http.Request) {
	var viewModel EditNote
	vars := mux.Vars(r)
	k := vars["id"]
	if note, ok := noteStore[k]; ok {
		viewModel = EditNote{note, k}
	} else {
		http.Error(w, "Could not find the resource to edit", http.StatusBadRequest)
	}
	renderTemplate(w, "edit", "base", viewModel)
}
func saveNote(w http.ResponseWriter, r *http.Request) {
	// renderTemplate(w, "index", "base", noteStore)
	r.ParseForm()
	title := r.PostFormValue("title")
	desc := r.PostFormValue("description")
	note := Note{title, desc, time.Now()}
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	http.Redirect(w, r, "/", 302)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
	r.HandleFunc("/notes/update/{id}", updateNote)
	r.HandleFunc("/notes/delete/{id}", deleteNote)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listen...")
	server.ListenAndServe()
}
