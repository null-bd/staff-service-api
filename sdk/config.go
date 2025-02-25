package sdk

import "time"

type Environment string

const (
	Production  Environment = "production"
	Staging     Environment = "staging"
	Development Environment = "development"

	defaultTimeout    = 30 * time.Second
	defaultMaxRetries = 3
	defaultRetryDelay = 100 * time.Millisecond
	defaultMaxDelay   = 2 * time.Second
)

// Config holds the SDK configuration
type Config struct {
	Environment Environment
	BaseURL     string
	Token       string
	Timeout     time.Duration
	RetryConfig *RetryConfig
	UserAgent   string
	Debug       bool
}

// RetryConfig holds the retry configuration
type RetryConfig struct {
	MaxRetries           int
	RetryDelay           time.Duration
	MaxDelay             time.Duration
	RetryableStatusCodes []int
}

// DefaultConfig returns the default SDK configuration
func DefaultConfig() *Config {
	return &Config{
		Environment: Production,
		Timeout:     defaultTimeout,
		RetryConfig: &RetryConfig{
			MaxRetries:           defaultMaxRetries,
			RetryDelay:           defaultRetryDelay,
			MaxDelay:             defaultMaxDelay,
			RetryableStatusCodes: []int{408, 429, 500, 502, 503, 504},
		},
		UserAgent: "domain-sdk/1.0",
	}
}

// WithEnvironment sets the environment
func (c *Config) WithEnvironment(env Environment) *Config {
	c.Environment = env
	return c
}

// WithBaseURL sets the base URL
func (c *Config) WithBaseURL(url string) *Config {
	c.BaseURL = url
	return c
}

// WithToken sets the authentication token
func (c *Config) WithToken(token string) *Config {
	c.Token = token
	return c
}

// WithTimeout sets the request timeout
func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

// WithRetryConfig sets the retry configuration
func (c *Config) WithRetryConfig(rc *RetryConfig) *Config {
	c.RetryConfig = rc
	return c
}

// WithDebug enables or disables debug logging
func (c *Config) WithDebug(debug bool) *Config {
	c.Debug = debug
	return c
}

// WithUserAgent sets the User-Agent header
func (c *Config) WithUserAgent(userAgent string) *Config {
	c.UserAgent = userAgent
	return c
}
