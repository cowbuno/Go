package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

type templateData struct {
	StringMap   map[string]string
	StringInt   map[string]int
	StringFloat map[string]float32
	Data        map[string]interface{}
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("../ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("../ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

func (app *application) newTemplateData(r *http.Request) templateData {
	StringMap := make(map[string]string)
	StringInt := make(map[string]int)
	StringFloat := make(map[string]float32)
	Data := make(map[string]interface{})
	StringInt["CurrentYear"] = time.Now().Year()
	return templateData{
		StringInt:   StringInt,
		StringMap:   StringMap,
		StringFloat: StringFloat,
		Data:        Data,
	}
}
