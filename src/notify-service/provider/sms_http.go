package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type HTTPSmsClient struct {
	base   string
	client *http.Client
}

func NewHTTPSmsClient(base string, timeout time.Duration) *HTTPSmsClient {
	if base == "" {
		return &HTTPSmsClient{base: "", client: &http.Client{Timeout: timeout}}
	}
	return &HTTPSmsClient{base: strings.TrimRight(base, "/"), client: &http.Client{Timeout: timeout}}
}

func (s *HTTPSmsClient) SendSMS(ctx context.Context, to []string, body string) error {
	if len(to) == 0 {
		return nil
	}
	if s.base == "" {
		// noop / mock
		fmt.Printf("SMS mock send to %v: %s\n", to, body)
		return nil
	}
	reqBody := strings.NewReader(body)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, s.base+"/send", reqBody)
	req.Header.Set("Content-Type", "text/plain")
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("sms provider status %d", resp.StatusCode)
	}
	return nil
}
