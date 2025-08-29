package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Generic wrapper response used by other services: { "data": ... }
type respWrapper struct {
	Data json.RawMessage `json:"data"`
}

// Minimal user DTO (optional)
type UserRes struct {
	ID       string `json:"id"`
	FullName string `json:"full_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

type UserService interface {
	// call something like POST /account/batch or GET /account/patients ...
	GetUsers(ctx context.Context, params url.Values, auth string) ([]map[string]interface{}, error)
}

type AppointmentService interface {
	GetAppointments(ctx context.Context, params url.Values, auth string) ([]map[string]interface{}, error)
}

type PrescriptionService interface {
	GetPrescriptions(ctx context.Context, params url.Values, auth string) ([]map[string]interface{}, error)
}

type httpClient struct {
	base string
	c    *http.Client
}

func NewHttpUserClient(base string, timeout time.Duration) UserService {
	return &httpClient{base: strings.TrimRight(base, "/"), c: &http.Client{Timeout: timeout}}
}
func NewHttpAppointmentClient(base string, timeout time.Duration) AppointmentService {
	return &httpClient{base: strings.TrimRight(base, "/"), c: &http.Client{Timeout: timeout}}
}
func NewHttpPrescriptionClient(base string, timeout time.Duration) PrescriptionService {
	return &httpClient{base: strings.TrimRight(base, "/"), c: &http.Client{Timeout: timeout}}
}

// helper: do GET with query params; fallback to POST batch if caller uses different endpoints
func (h *httpClient) doGet(ctx context.Context, path string, params url.Values, auth string) ([]map[string]interface{}, error) {
	u := h.base + path
	if params != nil {
		u = u + "?" + params.Encode()
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if auth != "" {
		req.Header.Set("Authorization", auth)
	}
	resp, err := h.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("status %d body:%s", resp.StatusCode, string(body))
	}
	var w respWrapper
	if err := json.Unmarshal(body, &w); err != nil {
		return nil, err
	}
	// decode into []map
	var arr []map[string]interface{}
	if err := json.Unmarshal(w.Data, &arr); err == nil {
		return arr, nil
	}
	// try decode as object with rows
	var obj map[string]interface{}
	if err := json.Unmarshal(w.Data, &obj); err == nil {
		// try rows or data inside
		if v, ok := obj["rows"]; ok {
			b, _ := json.Marshal(v)
			if err := json.Unmarshal(b, &arr); err == nil {
				return arr, nil
			}
		}
	}
	return nil, fmt.Errorf("unexpected response shape")
}

func (h *httpClient) GetUsers(ctx context.Context, params url.Values, auth string) ([]map[string]interface{}, error) {
	// Only try user-service endpoints
	if arr, err := h.doGet(ctx, "/user-service/account/patients", params, auth); err == nil {
		return arr, nil
	}
	if arr, err := h.doGet(ctx, "/user-service/account/filter", params, auth); err == nil {
		return arr, nil
	}
	return nil, fmt.Errorf("user-service endpoints not reachable")
}

func (h *httpClient) GetAppointments(ctx context.Context, params url.Values, auth string) ([]map[string]interface{}, error) {
	// Only try appointment-service endpoints
	if arr, err := h.doGet(ctx, "/appointment-service/appointment/filter", params, auth); err == nil {
		return arr, nil
	}
	return nil, fmt.Errorf("appointment-service endpoints not reachable")
}

func (h *httpClient) GetPrescriptions(ctx context.Context, params url.Values, auth string) ([]map[string]interface{}, error) {
	// Only try prescription-service endpoints
	if arr, err := h.doGet(ctx, "/prescription-service/prescription/filter", params, auth); err == nil {
		return arr, nil
	}
	return nil, fmt.Errorf("prescription-service endpoints not reachable")
}
