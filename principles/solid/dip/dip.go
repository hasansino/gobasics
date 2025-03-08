// Package dip shows how Dependency Inversion Principle can be applied in Go.
//
// https://en.wikipedia.org/wiki/Dependency_inversion_principle
package dip

import (
	"io"
	"net/http"
)

type Worker struct {
	// requester is direct dependency from http package
	requester *http.Client
}

// DoWork makes request using http.Client's Get method.
func (w *Worker) DoWork() error {
	resp, err := w.requester.Get("http://localhost")
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		// ... do something
	}
	return nil
}

// ----------------------------------------------------------------

// Requester is inverted dependency, it can be http.Client or anything else
type Requester interface {
	Get(url string) (HTTPResponse, error)
}

// HTTPResponse is abstraction of any type of response (like http.Response)
type HTTPResponse interface {
	StatusCode() int
	Body() []byte
}

type Worker2 struct {
	requester Requester
}

// DoWork still proceeds to make Get request, but now we can replace requester with any
// object which satisfies Requester interface.
func (w *Worker2) DoWork() error {
	resp, err := w.requester.Get("http://localhost")
	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusOK {
		// ... do something
	}
	return nil
}

// ----------------------------------------------------------------

// STDHTTPRequester is example how to use http.Client as Requester.
type STDHTTPRequester struct {
	client *http.Client
}

// Get makes request using http.Client's Get method.
func (r *STDHTTPRequester) Get(url string) (HTTPResponse, error) {
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	invertedResp := &STDHTTPRequesterResponse{
		statusCode: resp.StatusCode,
		body:       body,
	}

	return invertedResp, err
}

// STDHTTPRequesterResponse is example how to use http.Response as HTTPResponse.
type STDHTTPRequesterResponse struct {
	statusCode int
	body       []byte
}

func (r *STDHTTPRequesterResponse) StatusCode() int {
	return r.statusCode
}

func (r *STDHTTPRequesterResponse) Body() []byte {
	return r.body
}
