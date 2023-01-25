package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeaderMiddleWare(t *testing.T) {
	called := false
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		{
			name: "",
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					called = true
					w.Header().Add("Content-Type", "application/json; charset=utf-8")
				}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := headerMiddleware(tt.args.next)

			if _, ok := got.(http.Handler); !ok {
				t.Errorf("headerMiddleWare() = %v, want %v", got, tt.want)
			}
			rr := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/", nil)
			got.ServeHTTP(rr, r)
			if called != true {
				t.Errorf("Failed Called = %v", called)
			}

		})
	}
}
