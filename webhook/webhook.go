package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	contentTypeJSON  = "application/json"
	defaultUsername  = "slack-webhook-client"
	defaultIconEmoji = "+1"
)

var (
	ErrNilConfig    = errors.New("Config must not nil")
	ErrEmptyURL     = errors.New("URL must not empty")
	ErrEmptyChannel = errors.New("Channel must not empty")
	defaultTimeout  = 3 * time.Second
)

type Config struct {
	// WebhookURL
	URL string
	// Channel to which post messages
	Channel string
	// Username by whom message is posted
	Username string
	// Emoji representing user
	IconEmoji string
	// Timeout used by http.Client
	Timeout time.Duration
}

func (cfg *Config) SetDefualts() {
	if cfg.Username == "" {
		cfg.Username = defaultUsername
	}
	if cfg.IconEmoji == "" {
		cfg.IconEmoji = defaultIconEmoji
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = defaultTimeout
	}
}

// Client keeps config and underlying http client.
// store config value not pointer for immutability.
type Client struct {
	cfg    Config
	client *http.Client
}

// New return Client by provided configuration.
// if it is not given, use default one
func New(c *Config) (*Client, error) {
	if c == nil {
		return nil, ErrNilConfig
	}
	c.SetDefualts()
	if err := checkConfig(c); err != nil {
		return nil, err
	}
	return &Client{
		cfg: *c,
		client: &http.Client{
			Timeout: c.Timeout,
		},
	}, nil
}

func checkConfig(c *Config) error {
	return applyChecks(c,
		checkURLIsNotEmpty,
		checkChannelIsNotEmpty,
	)
}

func checkURLIsNotEmpty(c *Config) error {
	if c.URL == "" {
		return ErrEmptyURL
	}
	return nil
}

func checkChannelIsNotEmpty(c *Config) error {
	if c.Channel == "" {
		return ErrEmptyChannel
	}
	return nil
}

func applyChecks(c *Config, fns ...func(*Config) error) error {
	for _, fn := range fns {
		if err := fn(c); err != nil {
			return err
		}
	}
	return nil
}

// Payload represent POST request body.
// it is intended to encode json.
type Payload struct {
	Text      string `json:"text"`
	Channel   string `json:"channel"`
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

func (c *Client) setHeaders(req *http.Request) *http.Request {
	req.Header.Set("Content-Type", contentTypeJSON)
	return req
}

func (c *Client) buildRequest(p *Payload) (*http.Request, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.cfg.URL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return c.setHeaders(req), nil
}

func (c *Client) SendPayload(p *Payload) (err error) {
	req, err := c.buildRequest(p)
	if err != nil {
		return err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		clsErr := res.Body.Close()
		if err == nil {
			err = clsErr
		}
	}()
	return nil
}

func (c *Client) Send(msg string) error {
	p := &Payload{
		Text:      msg,
		Channel:   c.cfg.Channel,
		Username:  c.cfg.Username,
		IconEmoji: c.cfg.IconEmoji,
	}
	return c.SendPayload(p)
}
