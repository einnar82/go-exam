// handlers/form.go
package handlers

import (
	"html/template"
	"net/http"
)

func FormHandler(authToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>File Upload</title>
        </head>
        <body>
            <form action="/upload" method="post" enctype="multipart/form-data">
                <input type="hidden" name="auth" value="{{.AuthToken}}">
                <input type="file" name="data">
                <input type="submit" value="Upload">
            </form>
        </body>
        </html>
    `
		t, _ := template.New("form").Parse(tmpl)
		t.Execute(w, map[string]interface{}{"AuthToken": authToken})
	}
}
