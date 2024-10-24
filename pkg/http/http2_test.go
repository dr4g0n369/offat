package http_test

import (
	"testing"

	"github.com/owasp-offat/offat/pkg/http"

	c "github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/valyala/fasthttp"
)

func TestHttp2Client(t *testing.T) {
	// http2 client
	requestsPerSecond := 10
	skipTlsVerification := false
	proxy := ""
	hc := http.NewConfigHttp2(&requestsPerSecond, &skipTlsVerification, &proxy)

	url := "https://example.com"
	hc.Requests = append(hc.Requests, c.NewRequest(url, fasthttp.MethodGet, nil, nil, nil))
	hc.Requests = append(hc.Requests, c.NewRequest(url, fasthttp.MethodGet, nil, nil, nil))

	t.Run("Concurrent Requests Test", func(t *testing.T) {
		hc.Responses = c.MakeConcurrentRequests(hc.Requests, hc)

		for _, connResp := range hc.Responses {
			if connResp.Error != nil {
				t.Fatalf("failed to make concurrent request: %v\n", connResp.Error)
			}
		}
	})

}
