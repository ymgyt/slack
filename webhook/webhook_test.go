package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testEndpoint = "https://slack.com/api/api.test"
)

func TestNew(t *testing.T) {
	cases := []struct {
		cfg  Config
		want error
	}{
		{
			cfg:  Config{URL: "", Channel: "testChannel"},
			want: ErrEmptyURL,
		},
		{
			cfg:  Config{URL: testEndpoint, Channel: ""},
			want: ErrEmptyChannel,
		},
		{
			cfg:  Config{URL: testEndpoint, Channel: "testChannel"},
			want: nil,
		},
	}

	for _, tc := range cases {
		c, got := New(tc.cfg)
		assert.Equal(t, got, tc.want)
		if got == nil {
			assert.NotNil(t, c)
		}
	}
}
