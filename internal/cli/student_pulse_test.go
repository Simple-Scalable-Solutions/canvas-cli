package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunStudentPulse(t *testing.T) {
	enrollments := []map[string]interface{}{
		{"user_id": 1, "user": map[string]interface{}{"name": "Alice"}, "grades": map[string]interface{}{"current_score": 95.0}},
		{"user_id": 2, "user": map[string]interface{}{"name": "Bob"}, "grades": map[string]interface{}{"current_score": 82.0}},
		{"user_id": 3, "user": map[string]interface{}{"name": "Carol"}, "grades": map[string]interface{}{"current_score": 55.0}},
		{"user_id": 4, "user": map[string]interface{}{"name": "Dave"}, "grades": map[string]interface{}{"current_score": nil}},
	}
	// Only Alice has submitted recently; Bob submitted long ago; Carol and Dave never submitted
	submissions := []map[string]interface{}{
		{"user_id": 1, "submitted_at": "2026-05-10T10:00:00Z", "workflow_state": "submitted"},
		{"user_id": 2, "submitted_at": "2026-04-01T10:00:00Z", "workflow_state": "submitted"},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "enrollments") {
			json.NewEncoder(w).Encode(enrollments)
		} else {
			json.NewEncoder(w).Encode(submissions)
		}
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	var buf bytes.Buffer
	err := runStudentPulse(cc, "42", 30, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	// Grade distribution
	if !strings.Contains(out, "A") {
		t.Error("expected A bucket in output")
	}
	if !strings.Contains(out, "F") {
		t.Error("expected F bucket in output")
	}
	// At-risk: Bob (submitted > 30 days ago), Carol and Dave (never submitted)
	if !strings.Contains(out, "Bob") {
		t.Error("expected Bob in at-risk")
	}
	if !strings.Contains(out, "Carol") {
		t.Error("expected Carol in at-risk")
	}
	if !strings.Contains(out, "Dave") {
		t.Error("expected Dave in at-risk")
	}
	// Alice submitted recently, should NOT be at-risk
	atRiskSection := out[strings.Index(out, "At-Risk"):]
	if strings.Contains(atRiskSection, "Alice") {
		t.Error("Alice submitted recently, should not be at-risk")
	}
}

func TestRunStudentPulse_Empty(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[]`)
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	var buf bytes.Buffer
	err := runStudentPulse(cc, "42", 30, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "No students") {
		t.Error("expected 'No students' message for empty course")
	}
}
