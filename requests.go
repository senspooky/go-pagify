package pagify

import (
	"errors"
	"net/http"

	"github.com/senspooky/go-pagify/internal/utils"
)

// Represents a newly created response object
type Response interface {
	GetRequest() *http.Request           // Get the request object that was used to make the request
	GetResponse() *http.Response         // Get the response object that was returned by the request
	GetMetadata() interface{}            // Get set metadata of the response
	CreateResponse() (Response, error)   // Create a new response object
	SetRequest(*http.Request) Response   // Get the request object that was used to make the request
	SetResponse(*http.Response) Response // Get the response object that was returned by the request
	SetMetadata(interface{}) error       // Get set metadata of the response
}

type response struct {
	request  *http.Request
	response *http.Response
	metadata interface{}
}

func (r response) GetRequest() *http.Request {
	return r.request
}

func (r *response) SetRequest(request *http.Request) Response {
	r.request = request
	return r
}

func (r response) GetResponse() *http.Response {
	return r.response
}

func (r *response) SetResponse(response *http.Response) Response {
	r.response = response
	return r
}

func (r response) GetMetadata() interface{} {
	return r.metadata
}

func (r *response) SetMetadata(m interface{}) error {
	if r.metadata == nil || utils.CompareTyes(r.metadata, m) {
		r.metadata = m
		return nil
	}
	return errors.New("metadata type mismatch")
}

func (r response) CreateResponse() (Response, error) {
	if r.metadata == nil {
		return &response{}, nil
	}
	v, err := utils.GetNewPointerToInterface(r.metadata)
	if err != nil {
		return nil, err
	}
	return &response{
		metadata: v, // Allows for type "enforcement" when passing objects between requests
	}, nil
}
