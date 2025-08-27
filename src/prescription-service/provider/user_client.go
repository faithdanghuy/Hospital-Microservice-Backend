package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type UserRes struct {
	ID       string `json:"id"`
	FullName string `json:"full_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

type UserService interface {
	GetUsersByIDs(ctx context.Context, ids []string, authHeader string) (map[string]UserRes, error)
}

type HttpUserService struct {
	baseURL string
	client  *http.Client
}

func NewHttpUserService(baseURL string, timeout time.Duration) UserService {
	return &HttpUserService{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (s *HttpUserService) GetUsersByIDs(ctx context.Context, ids []string, authHeader string) (map[string]UserRes, error) {
	out := make(map[string]UserRes)
	if len(ids) == 0 {
		fmt.Println("⚠️ No IDs provided to fetch users")
		return out, nil
	}

	// -----------------------------
	// Attempt batch endpoint first
	// -----------------------------
	body := map[string][]string{"ids": ids}
	b, _ := json.Marshal(body)
	batchURL := fmt.Sprintf("%s/user-service/account/batch", s.baseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, batchURL, bytes.NewReader(b))
	if err != nil {
		fmt.Println("❌ Failed to create batch request:", err)
	} else {
		req.Header.Set("Content-Type", "application/json")
		if authHeader != "" {
			req.Header.Set("Authorization", authHeader)
		}
		resp, err := s.client.Do(req)
		if err != nil {
			fmt.Println("❌ Batch API request error:", err)
		} else {
			defer resp.Body.Close()
			raw, _ := io.ReadAll(resp.Body)

			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				var wrapper struct {
					Data json.RawMessage `json:"data"`
				}
				if err := json.Unmarshal(raw, &wrapper); err != nil {
					fmt.Println("❌ Failed to unmarshal wrapper:", err)
				} else {
					var users []UserRes
					if err := json.Unmarshal(wrapper.Data, &users); err == nil {
						for _, u := range users {
							out[u.ID] = u
						}
						return out, nil
					}

					var mmap map[string]UserRes
					if err := json.Unmarshal(wrapper.Data, &mmap); err == nil {
						return mmap, nil
					}

				}
			} else {
				fmt.Printf("❌ Batch API returned status %d\n", resp.StatusCode)
			}
		}
	}

	// -----------------------------
	// Fallback to GET detail/{id}
	// -----------------------------
	fmt.Println("🔄 Falling back to per-ID API calls")
	type result struct {
		u   UserRes
		err error
	}

	wg := sync.WaitGroup{}
	ch := make(chan result, len(ids))

	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			url := fmt.Sprintf("%s/user-service/account/detail/%s", s.baseURL, id)
			fmt.Println("➡️ Calling Detail API for ID:", id, "URL:", url)

			rq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				fmt.Println("❌ Failed to create detail request for", id, ":", err)
				ch <- result{err: err}
				return
			}
			if authHeader != "" {
				rq.Header.Set("Authorization", authHeader)
			}
			fmt.Println("📎 Detail API Headers:", rq.Header)

			resp, err := s.client.Do(rq)
			if err != nil {
				fmt.Println("❌ Request error for", id, ":", err)
				ch <- result{err: err}
				return
			}
			defer resp.Body.Close()

			raw, _ := io.ReadAll(resp.Body)

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				ch <- result{err: fmt.Errorf("status %d for id %s", resp.StatusCode, id)}
				return
			}

			var wrapper struct {
				Data json.RawMessage `json:"data"`
			}
			if err := json.Unmarshal(raw, &wrapper); err != nil {
				fmt.Println("❌ Failed to unmarshal wrapper for", id, ":", err)
				ch <- result{err: err}
				return
			}

			var u UserRes
			if err := json.Unmarshal(wrapper.Data, &u); err != nil {
				fmt.Println("❌ Failed to unmarshal UserRes for", id, ":", err)
				ch <- result{err: err}
				return
			}
			fmt.Println("✅ Parsed detail user for", id, ":", u)
			ch <- result{u: u}
		}(id)
	}

	wg.Wait()
	close(ch)

	var firstErr error
	for r := range ch {
		if r.err != nil {
			fmt.Println("⚠️ Detail API error:", r.err)
			if firstErr == nil {
				firstErr = r.err
			}
			continue
		}
		out[r.u.ID] = r.u
	}
	if len(out) == 0 && firstErr != nil {
		return nil, firstErr
	}

	return out, nil
}
