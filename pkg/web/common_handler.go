package web

import (
	"encoding/json"
	"net/http"
)

type WebHandlerFunc func(r *http.Request) (int, any, error)

func ToHandlerFunc(f WebHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, data, err := f(r)
		if err != nil {
			// TODO
			// exception handler的なやつをやりたい
			// エラーレスポンス型ならそのままそれを返す
			http.Error(w, err.Error(), status)
			return
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(data)
	}
}
