// package semconv defines custom semantic conventions not covered by otel
package semconv

import "go.opentelemetry.io/otel/attribute"

// Lantern cloud Proxies
const (
	ProxyCIDRKey          attribute.Key = "proxy.cidr"
	ProxyGatewayKey       attribute.Key = "proxy.gateway"
	HostInterfaceCountKey attribute.Key = "host.interface.count"
)

// Lantern cloud routes
const (
	RouteIDKey               attribute.Key = "route.id"
	RouteMultipleIDsKey      attribute.Key = "route.multiple_ids"
	RouteTrackNameKey        attribute.Key = "route.track_name"
	RouteRegionKey           attribute.Key = "route.region"
	RoutePHostTypeKey        attribute.Key = "route.phost_type"
	RoutePHostNameKey        attribute.Key = "route.phost_name"
	RouteIsV6Key             attribute.Key = "route.is_v6"
	RouteEIPKey              attribute.Key = "route.eip"
	RouteEIPProviderKey      attribute.Key = "route.eip_provider"
	RouteEIPLocationKey      attribute.Key = "route.eip_location"
	RouteStaticAddressKey    attribute.Key = "route.static_address"
	RouteStaticFrontendIDKey attribute.Key = "route.static_frontend_id"
	RouteCreatedAtKey        attribute.Key = "route.created_at"
	RouteReleaseForceKey     attribute.Key = "route.release.force"
	RouteDeprecatedKey       attribute.Key = "route.deprecated"
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
	EIPIsV6Key               attribute.Key = "eip.is_v6"
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

// Client Info
const (
	ClientDeviceIDKey           SemconvCustomKey = "client.device_id"
	ClientRegionKey             SemconvCustomKey = "client.region"
	ClientPlatformKey           SemconvCustomKey = "client.platform"
	ClientTierKey               SemconvCustomKey = "client.tier"
	ClientAsnKey                SemconvCustomKey = "client.asn"
	ClientTargetBackendKey      SemconvCustomKey = "client.target_backend"
	ClientSupportedProtocolsKey SemconvCustomKey = "client.supported_protocols"
	ClientIsDevKey              SemconvCustomKey = "client.is_dev"
)
