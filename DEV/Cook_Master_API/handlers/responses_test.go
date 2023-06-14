package handlers

// import (
// 	"gastroguru/user"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strconv"
// 	"testing"

// 	"gopkg.in/mgo.v2/bson"
// )

// type response struct {
// 	header http.Header
// 	code   int
// 	body   []byte
// }

// const (
// 	dbPath = "users.db"
// )

// type mockWriter response

// func newMockWriter() *mockWriter {
// 	return &mockWriter{
// 		header: http.Header{},
// 		body:   []byte{},
// 	}
// }

// func (mw *mockWriter) Write(b []byte) (int, error) {
// 	mw.body = make([]byte, len(b))
// 	for k, v := range b {
// 		mw.body[k] = v
// 	}
// 	return len(b), nil
// }

// func (mw *mockWriter) WriteHeader(statusCode int) {
// 	mw.code = statusCode
// }

// func (mw *mockWriter) Header() http.Header {
// 	return mw.header
// }

// func TestMain(m *testing.M) {
// 	m.Run()
// 	os.Remove(dbPath)

// }

// func prepDB(n int) error {
// 	os.Remove(dbPath)
// 	for i := 0; i < n; i++ {
// 		u := &user.User{
// 			FirstName:     "JJ",
// 			Name:          "Abrams" + strconv.Itoa(i),
// 			Email:         "abrams@gmail.com",
// 			Role:          "Student",
// 			Address:       "1 rue de la Paix",
// 			Phone:         "0123456789",
// 			Subscribed_ID: bson.NewObjectId().Time().Day(),
// 			Loyality_ID:   bson.NewObjectId().Time().Day(),
// 		}
// 		err := u.Save()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func makeRequest() (*http.Request, error) {
// 	u, err := url.Parse("/users")

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &http.Request{
// 		URL:    u,
// 		Header: http.Header{},
// 		Method: http.MethodGet,
// 	}, nil
// }

// func getAll(b *testing.B, r *http.Request) {
// 	prepDB(100)
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		mw := newMockWriter()
// 		b.StartTimer()
// 		UsersRouter(mw, r)
// 	}
// }

// func BenchmarkGetAllNonCached(b *testing.B) {
// 	r, err := makeRequest()
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	r.Header.Add("Cache-Control", "no-cache")
// 	getAll(b, r)
// }

// func BenchmarkGetAllCached(b *testing.B) {
// 	r, err := makeRequest()
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	getAll(b, r)
// }
