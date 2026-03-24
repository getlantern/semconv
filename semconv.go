// Package semconv defines custom semantic conventions not covered by otel
package semconv

import (
	"go.opentelemetry.io/otel/attribute"

	// Re-export all upstream OTel semantic conventions.
	. "go.opentelemetry.io/otel/semconv/v1.37.0"
)

// Anchor the dot-import so the compiler doesn't complain.
var _ = SchemaURL

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
	TrackNameKey            attribute.Key = "track.name"
	TrackIDKey              attribute.Key = "track.id"
	TrackDisabledKey        attribute.Key = "track.disabled"
	TrackTargetRegionsKey   attribute.Key = "track.target_regions"
	TrackTargetPlatformsKey attribute.Key = "track.target_platforms"
	TrackTargetTierKey      attribute.Key = "track.target_tier"
	TrackClientVersionKey   attribute.Key = "track.client_version"
	TrackClientFloorKey     attribute.Key = "track.client_floor"
	TrackClientCeilKey      attribute.Key = "track.client_ceil"
)

// Lantern cloud assignments
const (
	AssignmentCachedKey attribute.Key = "assignment.cached"
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
	ClientVersionKey            attribute.Key = "client.version"
	ClientPreferredRegionKey    attribute.Key = "client.preferred_region"
	ClientIsProKey              attribute.Key = "client.is_pro"
)

// Proxy resource attributes
const (
	ProxyNameKey             attribute.Key = "proxy.name"
	ProxyProtocolKey         attribute.Key = "proxy.protocol"
	ProxyTrackKey            attribute.Key = "proxy.track"
	ProxyProviderKey         attribute.Key = "proxy.provider"
	ProxyFrontendProviderKey attribute.Key = "proxy.frontend_provider"
)

// Proxy connection attributes (sing-box routing)
const (
	ProxyInboundKey     attribute.Key = "proxy.inbound"
	ProxyInboundTypeKey attribute.Key = "proxy.inbound_type"
	ProxyOutboundKey    attribute.Key = "proxy.outbound"
)
