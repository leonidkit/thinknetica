package webapp

import (
	"gosearch/pkg/engine"
	"net/http"
	"text/template"
)

type WebApp struct {
	engine engine.Service
	host   string
	port   string
}

func New(engine engine.Service) *WebApp {
	wa := &WebApp{
		engine: engine,
	}

	return wa
}

func (wa *WebApp) handlers() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/index", wa.Index)
	r.HandleFunc("/docs", wa.Docs)

	return r
}

func (wa *WebApp) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.New("index").Parse(`<ul>
		{{ range $index, $content := . }}
		<li><b>{{ $index }}</b> - {{ range $val := $content }}{{$val}}, {{ end }}</li>
		{{ end }}
		</ul>`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured during the parsing of a template: " + err.Error()))
		return
	}

	err = tpl.Execute(w, wa.engine.Index.Index())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured during the executing of a template: " + err.Error()))
		return
	}
}

func (wa *WebApp) Docs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.New("docs").Parse(`<ul>
		{{ range . }}
		<li><b>{{ .URL }}</b> - {{ .Title }}</li>
		{{ end }}
		</ul>`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured during the parsing of a template: " + err.Error()))
		return
	}

	err = tpl.Execute(w, wa.engine.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured during the executing of a template: " + err.Error()))
		return
	}
}
