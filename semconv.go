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
	RouteIDKey                        attribute.Key = "route.id"
	RouteMultipleIDsKey               attribute.Key = "route.multiple_ids"
	RouteTrackNameKey                 attribute.Key = "route.track_name"
	RoutePhostTypeKey                 attribute.Key = "route.phost_type"
	RouteIsV6Key                      attribute.Key = "route.is_v6"
	LanternCloudRouteEIP              attribute.Key = "route.eip"
	LanternCloudRouteEIPProvider      attribute.Key = "route.eip_provider"
	LanternCloudRouteEIPLocation      attribute.Key = "route.eip_location"
	LanternCloudRouteStaticAddress    attribute.Key = "route.static_address"
	LanternCloudRouteStaticFrontendID attribute.Key = "route.static_frontend_id"
	LanternCloudRouteCreatedAt        attribute.Key = "route.created_at"
	RouteReleaseForceKey              attribute.Key = "route.release.force"
	RouteDeprecatedKey                attribute.Key = "route.deprecated"
)

// Lantern cloud EIPs
const (
	EIPIDKey                 attribute.Key = "eip.id"
	EIPProviderIDKey         attribute.Key = "eip.provider_id"
	EIPAddressKey            attribute.Key = "eip.address"
	EIPFrontendIDKey         attribute.Key = "eip.frontend_id"
	EIPFrontendProviderIDKey attribute.Key = "eip.frontend_provider_id"
	EIPSlotProviderIDKey     attribute.Key = "eip.slot_provider_id"
	EIPSlotPrivateAddressKey attribute.Key = "eip.slot_private_address"
	EIPV6Key                 attribute.Key = "eip.is_v6"
)

// Lantern cloud tracks
const (
	TrackNameKey attribute.Key = "track.name"
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
