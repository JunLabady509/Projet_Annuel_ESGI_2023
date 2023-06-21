package cache

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

type mockWriter response

func newMockWriter() *mockWriter {
	return &mockWriter{
		header: http.Header{},
		body:   []byte{},
	}
}

func (mw *mockWriter) Write(b []byte) (int, error) {
	mw.body = make([]byte, len(b))
	for k, v := range b {
		mw.body[k] = v
	}
	return len(b), nil
}

func (mw *mockWriter) WriteHeader(statusCode int) {
	mw.code = statusCode
}

func (mw *mockWriter) Header() http.Header {
	return mw.header
}

func TestWriter(t *testing.T) {
	mw := newMockWriter()

	res := "/test/url?with=params"
	u, err := url.Parse(res)

	if err != nil {
		t.Fatal("Invalid URL")
	}

	req := &http.Request{
		URL: u,
	}

	t.Log("Testing NewWriter")
	w := NewWriter(mw, req)
	if w.resource != res {
		t.Errorf("Expected %s, got %s", res, w.resource)
	}
	if w.writer != mw {
		t.Fatal("Writer not set")
	}

	t.Log("Testing Header")
	h := w.Header()
	h.Add("test", "value")
	h2 := w.response.header
	if h2.Get("test") != "value" {
		t.Error("Header not set")
	}

	t.Log("Testing WriteHeader")
	c := 201
	w.WriteHeader(c)
	if w.response.code != c {
		t.Errorf("Expected %d, got %d", c, w.response.code)
	}
	if mw.code != c {
		t.Errorf("Status code not written")
	}
	h2 = mw.header
	if h2.Get("test") != "value" {
		t.Error("Header not written")
	}

	t.Log("Testing Write")
	b := []byte{1, 2, 3, 4, 5}
	l, err := w.Write(b)
	if err != nil {
		t.Errorf("Error writing %s", err)
	}
	if l != len(b) {
		t.Errorf("Expected %d, got %d", len(b), l)
	}

	if &w.response.body == &b {
		t.Error("Body assigned, not copied")
	}

	if !reflect.DeepEqual(w.response.body, b) {
		t.Error("Body not passed through")
		t.Error(mw.body)
		t.Error(b)
	}

}
