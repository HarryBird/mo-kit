package jaeger

type ICollector interface {
	GetEndpoint() string
	GetUsername() string
	GetPassword() string
}
