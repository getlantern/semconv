// package semconv defines custom semantic conventions not covered by otel
package semconv

// An semconvCustomKey defines a custom semantic convention for consistent telemetry naming.
type semconvCustomKey string

const (
	ProxyCIDR          semconvCustomKey = "proxy.cidr"
	ProxyGateway       semconvCustomKey = "proxy.gateway"
	HostInterfaceCount semconvCustomKey = "host.interface.count"
)
