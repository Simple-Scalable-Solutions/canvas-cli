package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"canvas-pp-cli/internal/config"
)

var linkNextRe = regexp.MustCompile(`<([^>]+)>;\s*rel="next"`)

type canvasClient struct {
	token   string
	baseURL string
	http    *http.Client
}

func newCanvasClient(token, baseURL string) *canvasClient {
	return &canvasClient{
		token:   token,
		baseURL: strings.TrimRight(baseURL, "/"),
		http:    &http.Client{},
	}
}

func (c *canvasClient) getAll(path string, params url.Values) ([]json.RawMessage, error) {
	next := c.baseURL + path
	if len(params) > 0 {
		next += "?" + params.Encode()
	}
	var all []json.RawMessage
	for next != "" {
		req, err := http.NewRequest(http.MethodGet, next, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+c.token)
		resp, err := c.http.Do(req)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("canvas API %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
		}
		var page []json.RawMessage
		if err := json.Unmarshal(body, &page); err != nil {
			return nil, fmt.Errorf("decoding response: %w", err)
		}
		all = append(all, page...)
		next = ""
		if link := resp.Header.Get("Link"); link != "" {
			if m := linkNextRe.FindStringSubmatch(link); m != nil {
				next = m[1]
			}
		}
	}
	return all, nil
}

func (c *canvasClient) post(path string, body url.Values) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = strings.NewReader(body.Encode())
	}
	req, err := http.NewRequest(http.MethodPost, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("canvas API %d: %s", resp.StatusCode, strings.TrimSpace(string(b)))
	}
	return b, nil
}

func newCanvasClientFromConfig(configPath string) (*canvasClient, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, err
	}
	if cfg.CanvasLmsToken == "" {
		return nil, fmt.Errorf("no Canvas token configured: set CANVAS_LMS_TOKEN or run 'canvas-pp-cli auth set-token'")
	}
	return newCanvasClient(cfg.CanvasLmsToken, cfg.BaseURL), nil
}
