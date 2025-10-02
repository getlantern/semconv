// package semconv defines custom semantic conventions not covered by otel
package semconv

// An lcSemconv defines a custom semantic convention for consistent telemetry naming.
type lcSemconv string

const (
	ProxyCIDR          lcSemconv = "proxy.cidr"
	ProxyGateway       lcSemconv = "proxy.gateway"
	HostInterfaceCount lcSemconv = "host.interface.count"
)
