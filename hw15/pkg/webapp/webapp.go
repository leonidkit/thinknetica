package webapp

import (
	"encoding/json"
	"fmt"
	"gosearch/pkg/engine"
	"net/http"
	"text/template"
)

type WebApp struct {
	engine engine.Service
}

func New(engine engine.Service) *WebApp {
	wa := &WebApp{
		engine: engine,
	}
	return wa
}

func (wa *WebApp) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		resp := map[string]string{
			"error": "method not allowed",
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.New("index").Parse(`<ul>
		{{ range $index, $content := . }}
		<li><b>{{ $index }}</b> - {{ range $val := $content }}{{$val}}, {{ end }}</li>
		{{ end }}
		</ul>`)

	if err != nil {
		resp := map[string]string{
			"error": fmt.Sprintf("error occured during the parsing of a template: %s", err.Error()),
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	err = tpl.Execute(w, wa.engine.Index.Recieve())
	if err != nil {
		resp := map[string]string{
			"error": fmt.Sprintf("error occured during the executing of a template: %s", err.Error()),
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}
}

func (wa *WebApp) Docs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		resp := map[string]string{
			"error": "method not allowed",
		}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.New("docs").Parse(`<ul>
		{{ range . }}
		<li><b>{{ .URL }}</b> - {{ .Title }}</li>
		{{ end }}
		</ul>`)

	if err != nil {
		resp := map[string]string{
			"error": fmt.Sprintf("error occured during the parsing of a template: %s", err.Error()),
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	err = tpl.Execute(w, wa.engine.Data)
	if err != nil {
		resp := map[string]string{
			"error": fmt.Sprintf("error occured during the executing of a template: %s", err.Error()),
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}
}

func (wa *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/docs":
		wa.Docs(w, r)
	case "/index":
		wa.Index(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		resp := map[string]string{
			"error": "unknown endpoint",
		}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}
}
