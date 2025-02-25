package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type sdk struct {
	config     *Config
	httpClient *http.Client
}

func NewSDK(config *Config) ISDK {
	if config == nil {
		config = DefaultConfig()
	}

	client := &sdk{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}

	return client
}

func (s *sdk) do(ctx context.Context, method, path string, body, result interface{}) error {
	url := fmt.Sprintf("%s%s", s.config.BaseURL, path)

	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", s.config.UserAgent)
	if s.config.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.config.Token))
	}

	// Handle retries if configured
	var lastError error
	retryCount := 0

	for {
		resp, err := s.httpClient.Do(req)
		if err != nil {
			lastError = fmt.Errorf("failed to execute request: %w", err)
			if !s.shouldRetry(retryCount, resp) {
				return lastError
			}
		} else {
			defer resp.Body.Close()

			// Check if we need to retry based on status code
			if s.shouldRetry(retryCount, resp) {
				lastError = fmt.Errorf("request failed with status: %d", resp.StatusCode)
			} else if resp.StatusCode >= 400 {
				var apiError APIError
				if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
					return fmt.Errorf("failed to decode error response: %w", err)
				}
				return &apiError
			} else {
				if result != nil {
					if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
						return fmt.Errorf("failed to decode response: %w", err)
					}
				}
				return nil
			}
		}

		retryCount++
		if !s.shouldRetry(retryCount, resp) {
			break
		}

		// Calculate retry delay using exponential backoff
		delay := s.calculateRetryDelay(retryCount)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
			continue
		}
	}

	return lastError
}

func (s *sdk) shouldRetry(retryCount int, resp *http.Response) bool {
	if s.config.RetryConfig == nil || retryCount >= s.config.RetryConfig.MaxRetries {
		return false
	}

	if resp == nil {
		return true
	}

	for _, code := range s.config.RetryConfig.RetryableStatusCodes {
		if resp.StatusCode == code {
			return true
		}
	}

	return false
}

func (s *sdk) calculateRetryDelay(retryCount int) time.Duration {
	delay := s.config.RetryConfig.RetryDelay * time.Duration(1<<uint(retryCount))
	if delay > s.config.RetryConfig.MaxDelay {
		delay = s.config.RetryConfig.MaxDelay
	}
	return delay
}
