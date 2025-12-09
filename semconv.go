// package semconv defines custom semantic conventions not covered by otel
package semconv

import "go.opentelemetry.io/otel/attribute"

// Lantern cloud Proxies
const (
	ProxyCIDR          attribute.Key = "proxy.cidr"
	ProxyGateway       attribute.Key = "proxy.gateway"
	HostInterfaceCount attribute.Key = "host.interface.count"
)

// Lantern cloud routes
const (
	LanternCloudRouteID               attribute.Key = "route.id"
	LanternCloudRouteMultipleIDs      attribute.Key = "route.multiple_ids"
	LanternCloudRouteTrackName        attribute.Key = "route.track_name"
	LanternCloudRoutePHostType        attribute.Key = "route.phost_type"
	LanternCloudRouteIsV6             attribute.Key = "route.is_v6"
	LanternCloudRouteEIP              attribute.Key = "route.eip"
	LanternCloudRouteEIPProvider      attribute.Key = "route.eip_provider"
	LanternCloudRouteEIPLocation      attribute.Key = "route.eip_location"
	LanternCloudRouteStaticAddress    attribute.Key = "route.static_address"
	LanternCloudRouteStaticFrontendID attribute.Key = "route.static_frontend_id"
	LanternCloudRouteCreatedAt        attribute.Key = "route.created_at"
	LanternCloudRouteReleaseForce     attribute.Key = "route.release.force"
	LanternCloudRouteDeprecated       attribute.Key = "route.deprecated"
)

// Lantern cloud EIPs
const (
	LanternCloudEIPID                 attribute.Key = "eip.id"
	LanternCloudEIPProviderID         attribute.Key = "eip.provider_id"
	LanternCloudEIPAddress            attribute.Key = "eip.address"
	LanternCloudEIPFrontendID         attribute.Key = "eip.frontend_id"
	LanternCloudEIPFrontendProviderID attribute.Key = "eip.frontend_provider_id"
	LanternCloudEIPSlotProviderID     attribute.Key = "eip.slot_provider_id"
	LanternCloudEIPSlotPrivateAddress attribute.Key = "eip.slot_private_address"
	LanternCloudEIPV6                 attribute.Key = "eip.is_v6"
)

// Lantern cloud tracks
const (
	TrackName attribute.Key = "track.name"
)

// Client Info
const (
	ClientDeviceIDKey           attribute.Key = "client.device_id"
	ClientRegionKey             attribute.Key = "client.region"
	ClientPlatformKey           attribute.Key = "client.platform"
	ClientTierKey               attribute.Key = "client.tier"
	ClientAsnKey                attribute.Key = "client.asn"
	ClientTargetBackendKey      attribute.Key = "client.target_backend"
	ClientSupportedProtocolsKey attribute.Key = "client.supported_protocols"
	ClientIsDevKey              attribute.Key = "client.is_dev"
)
