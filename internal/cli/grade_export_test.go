package cli

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunGradeExport(t *testing.T) {
	enrollments := []map[string]interface{}{
		{"user_id": 1001, "user": map[string]interface{}{"name": "Alice"}, "grades": map[string]interface{}{"final_score": 90.0}},
		{"user_id": 1002, "user": map[string]interface{}{"name": "Bob"}, "grades": map[string]interface{}{"final_score": 75.0}},
	}
	assignments := []map[string]interface{}{
		{"id": 201, "name": "Midterm", "points_possible": 100.0},
		{"id": 202, "name": "Final", "points_possible": 100.0},
	}
	submissions := []map[string]interface{}{
		{"user_id": 1001, "assignment_id": 201, "score": 85.0, "submitted_at": "2026-04-15T10:00:00Z"},
		{"user_id": 1002, "assignment_id": 201, "score": 72.0, "submitted_at": "2026-04-15T11:00:00Z"},
		// Alice and Bob have no submission for Final (id: 202)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "enrollments"):
			json.NewEncoder(w).Encode(enrollments)
		case strings.Contains(r.URL.Path, "assignments"):
			json.NewEncoder(w).Encode(assignments)
		default:
			json.NewEncoder(w).Encode(submissions)
		}
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	var buf bytes.Buffer
	err := runGradeExport(cc, "42", &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	lines := strings.Split(strings.TrimSpace(out), "\n")

	// Header row
	if lines[0] != "student_name,student_id,assignment_name,assignment_id,score,possible,submitted_at" {
		t.Errorf("unexpected header: %s", lines[0])
	}
	// 4 data rows: 2 students × 2 assignments
	if len(lines) != 5 {
		t.Errorf("expected 5 lines (header + 4 data), got %d:\n%s", len(lines), out)
	}
	// Verify Alice's Midterm row
	if !strings.Contains(out, "Alice,1001,Midterm,201,85,100,2026-04-15T10:00:00Z") {
		t.Errorf("missing Alice Midterm row in:\n%s", out)
	}
	// Verify Bob's Midterm row
	if !strings.Contains(out, "Bob,1002,Midterm,201,72,100,2026-04-15T11:00:00Z") {
		t.Errorf("missing Bob Midterm row in:\n%s", out)
	}
	// Alice Final: no submission — score and submitted_at empty
	if !strings.Contains(out, "Alice,1001,Final,202,,100,") {
		t.Errorf("missing Alice Final (no submission) row in:\n%s", out)
	}
}
