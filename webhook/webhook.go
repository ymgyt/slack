package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	contentTypeJSON  = "application/json"
	defaultUsername  = "slack-webhook-client"
	defaultIconEmoji = "+1"
)

var (
	ErrEmptyURL     = errors.New("URL must not empty")
	ErrEmptyChannel = errors.New("Channel must not empty")
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
	// flag whether dump response
	Dump bool
}

func (cfg *Config) SetDefualts() {
	if cfg.Username == "" {
		cfg.Username = defaultUsername
	}
	if cfg.IconEmoji == "" {
		cfg.IconEmoji = defaultIconEmoji
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
func New(c Config) (*Client, error) {
	c.SetDefualts()
	if err := checkConfig(&c); err != nil {
		return nil, err
	}
	return &Client{
		cfg: c,
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

// Field which is used in Attachment is displayed in message table.
type Field struct {
	// Required Field Title.
	Title string `json:"title"`
	// Text value of the field.
	Value string `json:"value"`
	// Optional flag indicating whether the `value` is short enough to be displayed
	// side-by-side with other values.
	Short bool `json:"short"`
}

// Attachment contains message formatting info.
type Attachment struct {
	// Required text summary of the attachment that is shown by clients that understand attachments
	// but choose not to show them.
	Fallback string `json:"fallback"`
	// Optional text that should appear within the attachment.
	Text string `json:"text"`
	// Optional text that should appear above the formatted data.
	Pretext string `json:"pretext"`
	// Can either be one of 'good', 'warning', 'danger', or any hex color code.
	Color  string   `json:"color"`
	Fields []*Field `json:"fields"`
}

// Payload represent POST request body.
// it is intended to encode json.
type Payload struct {
	Text        string        `json:"text"`
	Channel     string        `json:"channel"`
	Username    string        `json:"username,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

func (c *Client) setDefault(cfg Config, p *Payload) {
	if p.Channel == "" {
		p.Channel = cfg.Channel
	}
	if p.Username == "" {
		p.Username = cfg.Username
	}
	if p.IconEmoji == "" {
		p.IconEmoji = cfg.IconEmoji
	}
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

func (c *Client) handleResponse(res *http.Response) error {
	if c.cfg.Dump {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return err
		}
		fmt.Println(string(dump))
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("failed to send messages: %s", res.Status)
	}
	return nil
}

func (c *Client) SendPayload(p *Payload) (err error) {
	c.setDefault(c.cfg, p)
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
	return c.handleResponse(res)
}

func (c *Client) Send(msg string) error {
	p := &Payload{Text: msg}
	return c.SendPayload(p)
}
