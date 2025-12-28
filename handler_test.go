// test workflow2
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImageHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ImageHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba status 200, pero se obtuvo %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		t.Errorf("La respuesta debe incluir Content-Type")
	}
}
