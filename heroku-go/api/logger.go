package api

import (
	"fmt"
	"net/http"
	"time"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			recorder := &StatusRecorder{
				ResponseWriter: w,
				Status:         200,
			}

			defer func() {
				fmt.Printf(
					"%s %+v %s %s %dms\n",
					r.Proto,
					recorder.Status,
					r.Method,
					r.URL.Path,
					time.Since(t1)/time.Millisecond,
					// middleware.GetReqID(r.Context()))
				)
			}()

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
