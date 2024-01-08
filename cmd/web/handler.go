package main

import (
	"SDP/internal/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// files := []string{
	// 	"../ui/html/base.html",
	// 	"../ui/html/pages/home.html",
	// 	"../ui/html/partials/nav.html",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serveError(w, r, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serveError(w, r, err)
	// }

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serveError(w, r, err)
	}

	for _, s := range snippets {
		fmt.Fprintf(w, "%+v\n", s)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serveError(w, r, err)
		}
		return
	}

	files := []string{
		"../ui/html/base.html",
		"../ui/html/partials/nav.html",
		"../ui/html/pages/view.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, r, err)
		return
	}

	data := make(map[string]interface{})
	data["Snippet"] = snippet

	err = ts.ExecuteTemplate(w, "base", templateData{
		Data: data,
	})
	if err != nil {
		app.serveError(w, r, err)
	}
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serveError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
