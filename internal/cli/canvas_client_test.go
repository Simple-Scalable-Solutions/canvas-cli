package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	page := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page++
		w.Header().Set("Content-Type", "application/json")
		if page == 1 {
			w.Header().Set("Link", fmt.Sprintf(`<%s/next>; rel="next"`, "http://"+r.Host))
			fmt.Fprint(w, `[{"id":1}]`)
		} else {
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

func TestNewCanvasClientFromConfig_NoToken(t *testing.T) {
	t.Setenv("CANVAS_LMS_TOKEN", "")
	_, err := newCanvasClientFromConfig("")
	if err == nil {
		t.Fatal("expected error when no token set")
	}
}
