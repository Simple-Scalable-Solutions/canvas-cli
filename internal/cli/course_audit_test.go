package cli

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunCourseAudit(t *testing.T) {
	assignments := []map[string]interface{}{
		// Has due date and rubric — clean
		{"id": 1, "name": "Quiz 1", "due_at": "2026-06-01T23:59:00Z", "has_rubric": true},
		// No due date, no rubric — 2 issues
		{"id": 2, "name": "Essay", "due_at": nil, "has_rubric": false},
	}
	modules := []map[string]interface{}{
		{
			"id": 10, "name": "Week 1",
			"items": []map[string]interface{}{
				{"id": 100, "title": "Lecture Notes", "published": true},
				{"id": 101, "title": "Quiz", "published": false}, // unpublished
			},
		},
	}
	pages := []map[string]interface{}{
		{"page_id": 1, "title": "Welcome", "body": "Hello students"},    // clean
		{"page_id": 2, "title": "Placeholder", "body": ""},              // empty
		{"page_id": 3, "title": "Nil body", "body": nil},                // nil body
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "assignments"):
			_ = json.NewEncoder(w).Encode(assignments)
		case strings.Contains(r.URL.Path, "modules"):
			// Canvas returns modules without items inline by default;
			// we ask for include[]=items
			_ = json.NewEncoder(w).Encode(modules)
		case strings.Contains(r.URL.Path, "pages"):
			_ = json.NewEncoder(w).Encode(pages)
		default:
			_ = json.NewEncoder(w).Encode([]interface{}{})
		}
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	var buf bytes.Buffer
	err := runCourseAudit(cc, "42", &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	// 5 findings: Essay (no_due_date + no_rubric = 2), Quiz unpublished, Placeholder empty, Nil body empty
	if !strings.Contains(out, "no_due_date") {
		t.Error("expected no_due_date finding")
	}
	if !strings.Contains(out, "no_rubric") {
		t.Error("expected no_rubric finding")
	}
	if !strings.Contains(out, "unpublished") {
		t.Error("expected unpublished finding")
	}
	if !strings.Contains(out, "empty_page") {
		t.Error("expected empty_page finding")
	}
	// Clean items should not appear
	if strings.Contains(out, "Quiz 1") {
		t.Error("Quiz 1 has due date and rubric, should not appear in findings")
	}
	if strings.Contains(out, "Welcome") {
		t.Error("Welcome page has content, should not appear in findings")
	}
	// Count: 5 findings (Essay×2, Quiz unpublished, Placeholder, Nil body)
	if !strings.Contains(out, "Findings (5)") {
		t.Errorf("expected 'Findings (5)' in output, got:\n%s", out)
	}
}

func TestRunCourseAudit_Clean(t *testing.T) {
	assignments := []map[string]interface{}{
		{"id": 1, "name": "Quiz 1", "due_at": "2026-06-01T23:59:00Z", "has_rubric": true},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "assignments") {
			_ = json.NewEncoder(w).Encode(assignments)
		} else {
			_ = json.NewEncoder(w).Encode([]interface{}{})
		}
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	var buf bytes.Buffer
	err := runCourseAudit(cc, "42", &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "No issues") {
		t.Error("expected 'No issues' for clean course")
	}
}
