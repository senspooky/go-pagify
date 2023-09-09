package pagify

// Paginator represents a controller for paginating a response or set of responses.
//
// It's type parameter T is the type of data you want to pass between page requests as a refererence to the next page
// of the paginated resource.
type Paginator[T any] interface {
	// SetSubsequentRequestFunc sets the function that will be called to get subsequent pages of the paginated resource
	// beyond the first page.
	SetCommonSubsequentRequestFunc(func(T) (T, error)) Paginator[T]
	// SetCommonHasNextFunc sets the function that will be called to determine if there is a next page of the paginated
	// resource. This function takes in the current page's request return, and should return true if there is a next
	// page, and false if there is not.
	//
	// If this is not set, pages will assume that a zero-valued or nil-valued return from the most recent request
	// indicates that there is no next page.
	SetCommonHasNextFunc(func(T) bool) Paginator[T]
}

type pageController[T any] struct {
	subsequentRequestFunc func(T) (T, error)
	commonHasNextFunc     func(T) bool
}

func (p *pageController[T]) SetCommonHasNextFunc(f func(T) bool) Paginator[T] {
	p.commonHasNextFunc = f
	return p
}

func (p *pageController[T]) SetCommonSubsequentRequestFunc(f func(T) (T, error)) Paginator[T] {
	p.subsequentRequestFunc = f
	return p
}

func (p *pageController[T]) getCommonHasNextFunc() func(T) bool {
	if p.commonHasNextFunc == nil {
		return func(t T) bool {
			return false
		}
	}
	return p.commonHasNextFunc
}

// Returns the first page of the paginated resource
func (p *pageController[T]) GetFirstPage() (Page[T], error) {
	return &page[T]{
		run: true,
	}, nil
}
