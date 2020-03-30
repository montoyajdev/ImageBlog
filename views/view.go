package views

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)

	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
	}
}

// render is used to render the view with predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{})  {
	w.Header().Set("Content-Type", "text/html")
	switch data.(type) {
	case Data:
		// Do nothing
	default:
		data = Data{
			Yield: data,
		}
	}
	var buf *bytes.Buffer
	if err := v.Template.ExecuteTemplate(&buf, v.Layout, data); err != nil {
		http.Error(w, "Something went wrong.  If the problem persists, please email support @lenslocked.com", http.StatusInternalServerError)
		return
	}
	io.Copy(w,&buf)
}

// layout files returns a slice of strings representing
// the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	return files
}

// add template path takes in a slice of strings
// representing file paths for templates and it prepends
// the TemplateDir directory to each string in the slice

// Ex. the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

//addTemlateExt takes in a slice of strings
// representing file paths for templates and it appends
// the TemplateExt extenstion to each string in the slice

// Ex. the input {"home"} would result in the output
// {"home.gohtml"} if TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
