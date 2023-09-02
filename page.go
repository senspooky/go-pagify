package pagify

import "errors"

// Represents the state of a page before the request has been run
type NewPage interface {
	SetRequestFunc(func(Response) (Response, error))
	SetNextPageRequestFunc(func(Response) (Response, error))
	SetHasNextFunc(func(Response) bool)
	InitPage() (ProcessedPage, error)
}
type ProcessedPage interface {
	GetResponse() Response
	GetNextPage() (ProcessedPage, error)
	GetPrevPage() ProcessedPage
}

type page struct {
	// data fields

	resp Response // nil if no request has been run

	// links

	next    ProcessedPage
	prev    ProcessedPage
	hasNext bool

	// configuration fields

	requestFunc     func(Response) (Response, error)
	nextRequestFunc func(Response) (Response, error)
	hasNextFunc     func(Response) bool
}

func (P *page) getPrevResponse() Response {
	if P.prev == nil {
		return nil
	}
	return P.prev.GetResponse()
}

func (P *page) GetResponse() Response {
	return P.resp
}

func (P *page) GetNextPage() (ProcessedPage, error) {
	if !P.hasNext {
		return nil, nil
	}
	if P.next != nil { // if the next page has already been created, return it
		return P.next, nil
	}
	next, err := (&page{ // Copy over all the configuration fields
		prev:            P,
		requestFunc:     P.nextRequestFunc,
		nextRequestFunc: P.nextRequestFunc,
		hasNextFunc:     P.hasNextFunc,
	}).InitPage()
	if err != nil {
		return nil, err
	}
	P.next = next
	return next, nil
}

func (P *page) GetPrevPage() ProcessedPage {
	return P.prev
}

// Setters
// These are available on a page that has been created from scratch
// and not from a paginator object
func (P *page) SetRequestFunc(f func(Response) (Response, error)) {
	P.requestFunc = f
}

func (P *page) SetNextPageRequestFunc(f func(Response) (Response, error)) {
	P.nextRequestFunc = f
}

func (P *page) SetHasNextFunc(f func(Response) bool) {
	P.hasNextFunc = f
}

// Runs the request on the page, fills out dependant data, and returns itself
func (P *page) InitPage() (ProcessedPage, error) {
	if P.requestFunc == nil {
		return nil, errors.New("request function not set")
	}
	resp, err := P.requestFunc(P.getPrevResponse())
	if err != nil {
		return nil, err
	}
	if P.hasNextFunc == nil {
		return nil, errors.New("has next function not set")
	}
	P.hasNext = P.hasNextFunc(resp)
	P.resp = resp
	return P, nil
}
