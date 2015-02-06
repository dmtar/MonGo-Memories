package middlewares

import (
	"net/http"

	"encoding/json"

	"github.com/dmtar/pit/lib"
	"github.com/zenazn/goji/web"
)

func JSON(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Content-Type") == "application/json" {
			var params lib.Params
			c.Env["ParamsError"] = json.NewDecoder(r.Body).Decode(&params)
			c.Env["Params"] = params
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
