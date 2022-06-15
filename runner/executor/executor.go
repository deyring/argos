package executor

import (
	"context"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"

	"github.com/deyring/argos/models"
)

type Executor interface {
	Run() (*models.EndpointCheckResult, error)
}

type executor struct {
	request *http.Request
	timeout int
}

func New(endpointCheck *models.EndpointCheck) (Executor, error) {
	request, err := createRequest(endpointCheck)
	if err != nil {
		return nil, err
	}

	timeout := 5
	if endpointCheck.Timeout > 0 {
		timeout = endpointCheck.Timeout
	}

	return &executor{
		request: request,
		timeout: timeout,
	}, nil
}

func createRequest(endpointCheck *models.EndpointCheck) (*http.Request, error) {
	var body io.Reader
	switch endpointCheck.Method {
	case "GET", "HEAD", "OPTIONS", "TRACE", "DELETE":
		body = nil
	default:
		if len(endpointCheck.Body) > 0 {
			body = strings.NewReader(endpointCheck.Body)
		}
	}

	request, err := http.NewRequest(string(endpointCheck.Method), endpointCheck.URL, body)
	if err != nil {
		return nil, err
	}

	for key, value := range endpointCheck.Headers {
		request.Header.Set(key, value)
	}

	return request, nil
}

func (e *executor) Run() (*models.EndpointCheckResult, error) {
	client := createClient(e.timeout)

	var start, connect, dns, tlsHandshake time.Time

	var dnsDone, tlsHandshakeDone, connectDone, firstByteReceived, totalDone time.Duration

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			dnsDone = time.Since(dns)
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			tlsHandshakeDone = time.Since(tlsHandshake)
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			connectDone = time.Since(connect)
		},

		GotFirstResponseByte: func() {
			firstByteReceived = time.Since(start)
		},
	}

	e.request = e.request.WithContext(httptrace.WithClientTrace(e.request.Context(), trace))
	start = time.Now()

	response, err := client.Transport.RoundTrip(e.request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	totalDone = time.Since(start)

	result := &models.EndpointCheckResult{
		StatusCode:           response.StatusCode,
		Body:                 body,
		Headers:              response.Header,
		DNSDuration:          dnsDone,
		TLSHandshakeDuration: tlsHandshakeDone,
		ConnectDuration:      connectDone,
		FirstByteDuration:    firstByteReceived,
		TotalDuration:        totalDone,
	}

	return result, nil
}

func createClient(timeout int) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: defaultTransportDialContext(&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}),
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          1,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeout) * time.Second,
	}
}

func defaultTransportDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	return dialer.DialContext
}
