package pagify

type Requester interface {
	GetResponse() Response
}
