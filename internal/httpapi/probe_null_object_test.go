package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProbeRejectsNullInsteadOfObject(t *testing.T) {
	a := New()
	req := httptest.NewRequest(http.MethodPost, "/permissions", strings.NewReader("null"))
	rec := httptest.NewRecorder()
	a.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}
