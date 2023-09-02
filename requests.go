package paginator

import (
	"net/http"
)

type Response interface {
	GetRequest() *http.Request     // Get the request object that was used to make the request
	GetResponse() *http.Response   // Get the response object that was returned by the request
	GetBodyInterface() interface{} // Get the body of the response as an interface
	GetBodyAs(interface{}) error   // Get the body of the response as the type of the passed interface
}

// creates a new object implementing the Response and PrevPresonse interface
// accepts a http.Request, http.Response and zero or one interfaces representing the request's body
// will not write a body if a number of interfaces other than 1 are passed
func NewResponse(req *http.Request, resp *http.Response, b ...interface{}) Response {
	r := &response{
		request:  req,
		response: resp,
	}
	if len(b) == 1 {
		r.body = b[0]
	}
	return r
}

type response struct {
	request  *http.Request
	response *http.Response
	body     interface{}
}

func (r *response) GetRequest() *http.Request {
	return r.request
}

func (r *response) SetRequest(request *http.Request) {
	r.request = request
}

func (r *response) GetResponse() *http.Response {
	return r.response
}

func (r *response) SetResponse(response *http.Response) {
	r.response = response
}

func (r *response) GetBodyInterface() interface{} {
	return r.body
}

func (r *response) SetBody(body interface{}) {
	r.body = body
}

func (r *response) GetBodyAs(v interface{}) error {
	return nil
}
