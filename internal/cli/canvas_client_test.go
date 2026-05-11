package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetAll_SinglePage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer testtoken" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	items, err := cc.getAll("/courses", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("want 2 items, got %d", len(items))
	}
}

func TestGetAll_Pagination(t *testing.T) {
	var srv *httptest.Server
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/courses" {
			// first page — link to second
			w.Header().Set("Link", fmt.Sprintf(`<%s/page2>; rel="next"`, srv.URL))
			fmt.Fprint(w, `[{"id":1}]`)
		} else {
			// second page — no next link
			fmt.Fprint(w, `[{"id":2}]`)
		}
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	items, err := cc.getAll("/courses", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("want 2 items across pages, got %d", len(items))
	}
}

func TestGetAll_ErrorStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"message":"not found"}`, http.StatusNotFound)
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	_, err := cc.getAll("/courses", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestPost(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":99,"name":"test"}`)
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	b, err := cc.post("/courses/1/enrollments", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
}

func TestPost_WithBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad form", http.StatusBadRequest)
			return
		}
		if r.FormValue("enrollment[type]") != "StudentEnrollment" {
			http.Error(w, "missing field", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1}`)
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	body := url.Values{"enrollment[type]": {"StudentEnrollment"}}
	_, err := cc.post("/courses/1/enrollments", body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewCanvasClientFromConfig_NoToken(t *testing.T) {
	t.Setenv("CANVAS_LMS_TOKEN", "")
	_, err := newCanvasClientFromConfig("")
	if err == nil {
		t.Fatal("expected error when no token set")
	}
}
