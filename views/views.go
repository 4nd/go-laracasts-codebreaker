package views

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
)

const (
	viewsDir    = "templates"
	layoutsDir  = viewsDir + "/layouts"
	partialsDir = viewsDir + "/partials"
)

var (
	//go:embed templates
	files embed.FS
	Views map[string]*template.Template
)

func LoadViews() error {
	if Views == nil {
		Views = make(map[string]*template.Template)
	}

	viewFiles, err := fs.ReadDir(files, viewsDir)
	if err != nil {
		return err
	}

	for _, view := range viewFiles {
		if view.IsDir() {
			continue
		}

		parsedTemplate, err := template.ParseFS(
			files,
			viewsDir+"/"+view.Name(),
			layoutsDir+"/*.html",
			partialsDir+"/*.html",
		)
		if err != nil {
			return err
		}

		Views[view.Name()] = parsedTemplate
	}

	return nil
}

func RenderTemplate(
	resp http.ResponseWriter,
	viewFile string,
	data map[string]interface{},
) {
	if viteHot() == true {
		data["viteHot"] = true
	} else {
		data["viteHot"] = false
		data["styles"], data["scripts"] = ParseManifest()
	}

	view, ok := Views[viewFile]
	if !ok {
		fmt.Printf("view %s not found", viewFile)
		return
	}

	if err := view.Execute(resp, data); err != nil {
		fmt.Println(err)
	}
}

func viteHot() bool {
	_, err := os.Stat("./vite-hot")

	return err == nil
}
