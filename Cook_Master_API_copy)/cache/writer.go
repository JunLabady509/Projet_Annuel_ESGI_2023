package cache

import "net/http"

// Le Writer est un Wrapper autour de http.ResponseWriter qui permet de capturer le code de statut, les en-têtes et le corps de la réponse.
type Writer struct {
	writer   http.ResponseWriter
	response response
	resource string
}

// interface implementation check
var (
	_ http.ResponseWriter = (*Writer)(nil)
)

// NewWriter retourne le writer qui sera utilisé pour écrire la réponse HTTP dans le cache.
func NewWriter(w http.ResponseWriter, r *http.Request) *Writer {
	return &Writer{
		writer:   w,
		resource: MakeResource(r),
		response: response{
			header: http.Header{},
		},
	}
}

// Header retourne l'en-tête HTTP qui sera envoyé dans la réponse.
func (w *Writer) Header() http.Header {
	return w.response.header
}

// WriteHeader enregistre le code de statut HTTP qui sera envoyé dans la réponse.
func (w *Writer) WriteHeader(statusCode int) {
	copyHeader(w.response.header, w.writer.Header())
	w.response.code = statusCode
	w.writer.WriteHeader(statusCode)
}

func (w *Writer) Write(b []byte) (int, error) {
	w.response.body = make([]byte, len(b))
	for k, v := range b {
		w.response.body[k] = v
	}
	copyHeader(w.Header(), w.writer.Header())
	set(w.resource, &w.response)
	return w.writer.Write(b)
}
