package paginator

type Requester interface {
	GetResponse() Response
}
