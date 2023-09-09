package pagify

type Page[T any] interface {
	// SetRequestFunc takes in the function that will be called to get the next page of the paginated resource, and
	// returns the page itself for chaining functions.
	//
	// The passed in function takes in the request return from the previous page, and returns the request return for the
	// current page, and an error if one occurs. It is called when a function on the page that requires the request
	// return is called.
	//
	// Any errors returned by the passed in function will be returned by any function on the page that returns an error.
	SetRequestFunc(func(T) (T, error)) Page[T]
	// SetNextPageRequestFunc takes in the function that will be called to get the next page of the paginated resource,
	// and returns the page itself for chaining functions.
	//
	// This function is not used by the page it is called on, but is used by the next page in the paginated resource,
	// which will have both its request function and next page request function set to the same function, ensuring all
	// subsequent pages in the paginated resource use the same function to get data.
	//
	// Otherwise, this passed function is identical to SetRequestFunc.
	SetNextPageRequestFunc(func(T) (T, error)) Page[T]
	// GetNextPage returns the next page in the paginated resource, and an error if one occurs or has occured in any
	// previous function called on this page.
	//
	// If the page's request function has not been called yet, this function will call the request function to get the
	// data required to build the next page.
	//
	// If there is no next page to request for as indicated by the function set by SetHasNextFunc, this function will
	// return a nil page and no error.
	GetNextPage() (Page[T], error)
	// GetPrevPage returns the previous page in the paginated resource. If this is the first page in the chain, this
	// function will return nil.
	GetPrevPage() Page[T]
	// GetRequestReturn returns the return from the paage's request function. If the request function has not been
	// called yet, this function will call the request function to get it's return, and return that.
	//
	// If an error occured in the request function, this function will return nil.
	GetRequestReturn() T
}

type page[T any] struct {
	run bool
	err error

	reqReturn T

	prev Page[T]
	next Page[T]

	requestFunc         func(T) (T, error)
	nextPageRequestFunc func(T) (T, error)
}

func (p *page[T]) SetRequestFunc(f func(T) (T, error)) Page[T] {
	p.requestFunc = f
	return p
}

func (p *page[T]) SetNextPageRequestFunc(f func(T) (T, error)) Page[T] {
	p.nextPageRequestFunc = f
	return p
}

func (p *page[T]) GetNextPage() (Page[T], error) {
	if p.next != nil {
		return p.next, nil
	}
	return nil, nil
}

func (p *page[T]) GetPrevPage() Page[T] {
	return p.prev
}

func (p *page[T]) GetRequestReturn() T {
	return p.reqReturn
}
