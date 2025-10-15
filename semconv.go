// package semconv defines custom semantic conventions not covered by otel
package semconv

// An semconvCustomKey defines a custom semantic convention for consistent telemetry naming.
type semconvCustomKey string

// Lantern cloud Proxies
const (
	ProxyCIDR          semconvCustomKey = "proxy.cidr"
	ProxyGateway       semconvCustomKey = "proxy.gateway"
	HostInterfaceCount semconvCustomKey = "host.interface.count"
)

// Lantern cloud routes
const (
	LanternCloudRouteID               semconvCustomKey = "route.id"
	LanternCloudRouteMultipleIDs      semconvCustomKey = "route.multiple_ids"
	LanternCloudRouteTrackName        semconvCustomKey = "route.track_name"
	LanternCloudRoutePHostType        semconvCustomKey = "route.phost_type"
	LanternCloudRouteIsV6             semconvCustomKey = "route.is_v6"
	LanternCloudRouteEIP              semconvCustomKey = "route.eip"
	LanternCloudRouteEIPProvider      semconvCustomKey = "route.eip_provider"
	LanternCloudRouteEIPLocation      semconvCustomKey = "route.eip_location"
	LanternCloudRouteStaticAddress    semconvCustomKey = "route.static_address"
	LanternCloudRouteStaticFrontendID semconvCustomKey = "route.static_frontend_id"
	LanternCloudRouteCreatedAt        semconvCustomKey = "route.created_at"
)

// Lantern cloud EIPs
const (
	LanternCloudEIPID                 semconvCustomKey = "eip.id"
	LanternCloudEIPProviderID         semconvCustomKey = "eip.provider_id"
	LanternCloudEIPAddress            semconvCustomKey = "eip.address"
	LanternCloudEIPFrontendID         semconvCustomKey = "eip.frontend_id"
	LanternCloudEIPFrontendProviderID semconvCustomKey = "eip.frontend_provider_id"
	LanternCloudEIPSlotProviderID     semconvCustomKey = "eip.slot_provider_id"
	LanternCloudEIPSlotPrivateAddress semconvCustomKey = "eip.slot_private_address"
	LanterCloudEIPV6                  semconvCustomKey = "eip.is_v6"
)

// Lantern cloud tracks
const (
	LanternCloudTrackName semconvCustomKey = "track.name"
)
