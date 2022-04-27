package task_models

type Type string

const (
	TypeEndpoint    = "end:point"
	TypeNetCats     = "net:cats"
	TypePageSpeeds  = "page:speeds"
	TypePings       = "ping:s"
	TypeTraceRoutes = "trace:routes"
)

const (
	QueueEndpoint    = "endpoint"
	QueueNetCats     = "net_cats"
	QueuePageSpeeds  = "page_speeds"
	QueuePings       = "pings"
	QueueTraceRoutes = "trace_routes"
)
