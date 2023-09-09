package pagify

type Paginator struct {

	// configuration fields

	preloadAllPages   bool
	commonPageRequest func(Response, Response) error
	commonHasNext     func(Response) bool
}

func P() *Paginator {
	return &Paginator{}
}

// Coniguration for the paginator

// Defines the common function used to construct and create subsequent pages beyond the first,
// based on details from the previous request.
// Takes a function, whose paramaters are the previous page's response to the request, and is nil if no previous
// page exists.
// Returns the current page's response. A nil body is valid and if
// returned without an error, it is assumed the page does not exist.
func (P *Paginator) SetCommonPageRequestFunc(f func(Response, Response) error) *Paginator {
	P.commonPageRequest = f
	return P
}

func (P *Paginator) SetCommonHasNextFunc(f func(Response) bool) *Paginator {
	P.commonHasNext = f
	return P
}

// Returns the first page of the paginated resource
func (P *Paginator) GetFirstPage(r func(Response, Response) error) (ProcessedPage, error) {
	return (&page{
		requestFunc:     r,
		nextRequestFunc: P.commonPageRequest,
		hasNextFunc:     P.commonHasNext,
	}).InitPage()
}

// Iterator for the paginated resource.
// Takes a function representing thhe frist page request, and a function that will be called on each page.
// If the provided function returns an error, the iteration will stop and the error will be returned.
func (P *Paginator) IteratePages(r func(Response, Response) error, f func(Response) error) error {
	page, err := P.GetFirstPage(r)
	if err != nil {
		return err
	}
	for page != nil {
		err = f(page.GetResponse())
		if err != nil {
			return err
		}
		page, err = page.GetNextPage()
		if err != nil {
			return err
		}
	}
	return nil
}
