package cli

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunBulkEnroll(t *testing.T) {
	enrolled := []string{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "bad form", http.StatusBadRequest)
			return
		}
		userID := r.FormValue("enrollment[user_id]")
		role := r.FormValue("enrollment[type]")
		if userID == "9999" {
			http.Error(w, `{"message":"user not found"}`, http.StatusNotFound)
			return
		}
		enrolled = append(enrolled, userID)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id":1,"user_id":%s,"type":"%s"}`, userID, role)
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	input := strings.NewReader("1001\n1002\n9999\n")
	var buf bytes.Buffer
	err := runBulkEnroll(cc, "42", "StudentEnrollment", false, input, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Enrolled user 1001") {
		t.Error("expected enrolled message for 1001")
	}
	if !strings.Contains(out, "Enrolled user 1002") {
		t.Error("expected enrolled message for 1002")
	}
	if !strings.Contains(out, "Failed user 9999") {
		t.Error("expected failed message for 9999")
	}
	if !strings.Contains(out, "2 enrolled, 1 failed") {
		t.Errorf("expected summary line, got:\n%s", out)
	}
}

func TestRunBulkEnroll_DryRun(t *testing.T) {
	apiCalled := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiCalled = true
		http.Error(w, "should not be called", http.StatusInternalServerError)
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	input := strings.NewReader("1001\n1002\n")
	var buf bytes.Buffer
	err := runBulkEnroll(cc, "42", "StudentEnrollment", true, input, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if apiCalled {
		t.Error("API should not be called in dry-run mode")
	}
	out := buf.String()
	if !strings.Contains(out, "1001") || !strings.Contains(out, "1002") {
		t.Errorf("expected user IDs in dry-run output, got:\n%s", out)
	}
}
