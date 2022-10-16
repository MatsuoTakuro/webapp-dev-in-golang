package ch12

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

func Sub() {
	// list_12_2()
	// list_12_3()
	// list_12_4()
	// list_12_5()
	list_12_6()
}

func DefaultOkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK (in Response Body)\n"))
	log.Printf("request body: %s\n", r.Body)
}

func list_12_2() {
	// curl -v localhost:8888/ok
	http.Handle("/ok", MyMiddleware(http.HandlerFunc(DefaultOkHandler)))
	http.ListenAndServe(":8888", nil)
}

func MyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		h.ServeHTTP(w, r)
		t := time.Now()
		d := t.Sub(s).Milliseconds()
		log.Printf("end %s(%d ms)\n", t.Format(time.RFC3339), d)
	})
}

func list_12_3() {
	vmw := VersionAdder("1.0.1")
	// curl -v localhost:8888/users
	http.Handle("/users", vmw(http.HandlerFunc(DefaultOkHandler)))
	http.ListenAndServe(":8888", nil)
}

type AppVersion string

func VersionAdder(v AppVersion) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("App-Version", string(v))
			next.ServeHTTP(w, r)
		})
	}
}

func list_12_4() {
	// curl -v localhost:8888/recovery
	http.Handle("/recovery", RecoveryMiddleware(http.HandlerFunc(PanicHandler)))
	http.ListenAndServe(":8888", nil)
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("panic happened")
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				errBody, _ := json.Marshal(map[string]string{
					"error": fmt.Sprintf("%v", err),
				})

				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(errBody)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func list_12_5() {
	// curl -X POST localhost:8888/logreqbody -d '"test":"test"'
	http.Handle("/logreqbody", RequestBodyLogMiddleware(http.HandlerFunc(DefaultOkHandler)))
	http.ListenAndServe(":8888", nil)
}

func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		// err = errors.New("test")
		if err != nil {
			log.Printf("Failed to log request body: %v\n", zap.Error(err))
			http.Error(w, "Failed to get request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}

func list_12_6() {
	l := log.New(os.Stdout, "logger: ", 0)
	loggerMw := NewLogger(l)
	// curl -X POST localhost:8888/logresp
	http.Handle("/logresp", loggerMw(http.HandlerFunc(DefaultOkHandler)))
	http.ListenAndServe(":8888", nil)
}

type rwWrapper struct {
	rw     http.ResponseWriter
	mw     io.Writer
	status int
}

func NewRwWrapper(rw http.ResponseWriter, buf io.Writer) *rwWrapper {
	return &rwWrapper{
		rw:     rw,
		mw:     io.MultiWriter(rw, buf),
		status: 0,
	}
}

func (r *rwWrapper) Write(i []byte) (n int, err error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.mw.Write(i)
}

func (r *rwWrapper) Header() http.Header {
	return r.rw.Header()
}

func (r *rwWrapper) WriteHeader(statusCode int) {
	r.status = statusCode
	r.rw.WriteHeader(r.status)
}

func NewLogger(l *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buf := &bytes.Buffer{}
			rww := NewRwWrapper(w, buf)
			next.ServeHTTP(rww, r)
			l.Printf("response body: %s", buf)
			l.Printf("status: %d", rww.status)
		})
	}
}
