package bsprite

import "net/http"

// ServeHTTP allows us to satisfy the http.Handler interface
// so Sprite can be passed to http.Handle()
func (sprite Sprite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i, v := range sprite.Headers() {
		w.Header().Set(i, v)
	}
	w.Write(sprite.Body())
}
