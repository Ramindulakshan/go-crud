package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func newTestHandler() *Handler {
	svc := NewService(stubRepo{})
	return NewHandler(svc)
}

func TestRegisterRoutesUsesPatchForUpdate(t *testing.T) {
	h := newTestHandler()
	r := chi.NewRouter()
	h.RegisterRoutes(r)

	req := httptest.NewRequest(http.MethodPatch, "/users/not-a-uuid", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for patch invalid id, got %d", res.Code)
	}

	reqPut := httptest.NewRequest(http.MethodPut, "/users/not-a-uuid", nil)
	resPut := httptest.NewRecorder()
	r.ServeHTTP(resPut, reqPut)
	if resPut.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405 for put route, got %d", resPut.Code)
	}
}

func TestPatchRejectsEmptyPayload(t *testing.T) {
	h := newTestHandler()
	r := chi.NewRouter()
	h.RegisterRoutes(r)

	body, _ := json.Marshal(map[string]any{})
	req := httptest.NewRequest(http.MethodPatch, "/users/550e8400-e29b-41d4-a716-446655440000", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for empty patch payload, got %d", res.Code)
	}
}
