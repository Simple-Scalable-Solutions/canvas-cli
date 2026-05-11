package cli

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunActivityReport(t *testing.T) {
	enrollments := []map[string]interface{}{
		{"user_id": 1001, "user": map[string]interface{}{"name": "Alice"}},
		{"user_id": 1002, "user": map[string]interface{}{"name": "Bob"}},
	}
	summaries := []map[string]interface{}{
		{"id": 1001, "page_views": 42, "participations": 15, "last_activity_at": "2026-05-10T10:00:00Z"},
		{"id": 1002, "page_views": 0, "participations": 0, "last_activity_at": nil},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "analytics") {
			json.NewEncoder(w).Encode(summaries)
		} else {
			json.NewEncoder(w).Encode(enrollments)
		}
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	var buf bytes.Buffer
	err := runActivityReport(cc, "42", &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Alice") {
		t.Error("expected Alice in report")
	}
	if !strings.Contains(out, "Bob") {
		t.Error("expected Bob in report")
	}
	// Alice has activity
	if !strings.Contains(out, "42") {
		t.Error("expected 42 page views for Alice")
	}
	// Bob is at-risk
	if !strings.Contains(out, "At-Risk") {
		t.Error("expected At-Risk section")
	}
	if strings.Count(out, "Bob") < 2 {
		t.Error("Bob should appear in table AND in at-risk section")
	}
	// Alice should NOT be in at-risk (she has page views)
	atRiskIdx := strings.Index(out, "At-Risk")
	if atRiskIdx >= 0 && strings.Contains(out[atRiskIdx:], "Alice") {
		t.Error("Alice has page views, should not be in at-risk")
	}
}

func TestRunActivityReport_NoAnalytics(t *testing.T) {
	enrollments := []map[string]interface{}{
		{"user_id": 1001, "user": map[string]interface{}{"name": "Alice"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "analytics") {
			json.NewEncoder(w).Encode([]interface{}{}) // empty analytics
		} else {
			json.NewEncoder(w).Encode(enrollments)
		}
	}))
	defer srv.Close()

	cc := newCanvasClient("testtoken", srv.URL)
	var buf bytes.Buffer
	err := runActivityReport(cc, "42", &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Alice with no analytics data should appear in at-risk
	if !strings.Contains(buf.String(), "Alice") {
		t.Error("expected Alice in output even with no analytics")
	}
}
