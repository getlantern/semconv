// package semconv defines custom semantic conventions not covered by otel
package semconv

import "go.opentelemetry.io/otel/attribute"

// An SemconvCustomKey defines a custom semantic convention for consistent telemetry naming.
type SemconvCustomKey attribute.Key

// Lantern cloud Proxies
const (
	ProxyCIDR          SemconvCustomKey = "proxy.cidr"
	ProxyGateway       SemconvCustomKey = "proxy.gateway"
	HostInterfaceCount SemconvCustomKey = "host.interface.count"
)

// Lantern cloud routes
const (
	LanternCloudRouteID               SemconvCustomKey = "route.id"
	LanternCloudRouteMultipleIDs      SemconvCustomKey = "route.multiple_ids"
	LanternCloudRouteTrackName        SemconvCustomKey = "route.track_name"
	LanternCloudRoutePHostType        SemconvCustomKey = "route.phost_type"
	LanternCloudRouteIsV6             SemconvCustomKey = "route.is_v6"
	LanternCloudRouteEIP              SemconvCustomKey = "route.eip"
	LanternCloudRouteEIPProvider      SemconvCustomKey = "route.eip_provider"
	LanternCloudRouteEIPLocation      SemconvCustomKey = "route.eip_location"
	LanternCloudRouteStaticAddress    SemconvCustomKey = "route.static_address"
	LanternCloudRouteStaticFrontendID SemconvCustomKey = "route.static_frontend_id"
	LanternCloudRouteCreatedAt        SemconvCustomKey = "route.created_at"
	LanternCloudRouteReleaseForce     SemconvCustomKey = "route.release.force"
	LanternCloudRouteDeprecated       SemconvCustomKey = "route.deprecated"
)

// Lantern cloud EIPs
const (
	LanternCloudEIPID                 SemconvCustomKey = "eip.id"
	LanternCloudEIPProviderID         SemconvCustomKey = "eip.provider_id"
	LanternCloudEIPAddress            SemconvCustomKey = "eip.address"
	LanternCloudEIPFrontendID         SemconvCustomKey = "eip.frontend_id"
	LanternCloudEIPFrontendProviderID SemconvCustomKey = "eip.frontend_provider_id"
	LanternCloudEIPSlotProviderID     SemconvCustomKey = "eip.slot_provider_id"
	LanternCloudEIPSlotPrivateAddress SemconvCustomKey = "eip.slot_private_address"
	LanterCloudEIPV6                  SemconvCustomKey = "eip.is_v6"
)

// Lantern cloud tracks
const (
	LanternCloudTrackName SemconvCustomKey = "track.name"
)
