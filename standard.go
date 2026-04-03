package semconv

import (
	otel "go.opentelemetry.io/otel/semconv/v1.34.0"
)

// Re-exported standard OTel semantic conventions.
// Only symbols actively used across Lantern services belong here.
const SchemaURL = otel.SchemaURL

// Service resource attributes
const (
	// ServiceNameKey is the logical service name, e.g. "router", "bandit".
	ServiceNameKey = otel.ServiceNameKey
	// ServiceVersionKey is the deployed version, e.g. "v2.4.1".
	ServiceVersionKey = otel.ServiceVersionKey
	// ServiceNamespaceKey groups related services, e.g. "lantern".
	ServiceNamespaceKey = otel.ServiceNamespaceKey
)

// Cloud
const (
	// CloudProviderKey is the cloud provider, e.g. "aws", "gcp", "azure".
	CloudProviderKey = otel.CloudProviderKey
	// CloudAvailabilityZoneKey is the zone, e.g. "us-east-1a".
	CloudAvailabilityZoneKey = otel.CloudAvailabilityZoneKey
)

// Deployment

// DeploymentEnvironmentNameKey is "prod", "staging", or "dev".
const DeploymentEnvironmentNameKey = otel.DeploymentEnvironmentNameKey

// Host
const (
	// HostNameKey is the OS hostname.
	HostNameKey = otel.HostNameKey
	// HostIDKey is a unique host identifier.
	HostIDKey = otel.HostIDKey
	// HostArchKey is the CPU architecture, e.g. "amd64", "arm64".
	HostArchKey = otel.HostArchKey
)

// OS
const (
	// OSNameKey is the operating system name, e.g. "linux", "darwin".
	OSNameKey = otel.OSNameKey
	// OSVersionKey is the OS version string, e.g. "14.0", "22.04".
	OSVersionKey = otel.OSVersionKey
)

// Network
const (
	// NetworkPeerAddressKey is the remote IP, e.g. "10.0.0.5".
	NetworkPeerAddressKey = otel.NetworkPeerAddressKey
	// NetworkPeerPortKey is the remote port, e.g. 443.
	NetworkPeerPortKey = otel.NetworkPeerPortKey
	// NetworkIODirectionKey is "transmit" or "receive".
	NetworkIODirectionKey = otel.NetworkIODirectionKey
	// NetworkProtocolNameKey is the protocol, e.g. "tcp", "udp".
	NetworkProtocolNameKey = otel.NetworkProtocolNameKey
)

// HTTP
const (
	// HTTPRequestMethodKey is "GET", "POST", etc.
	HTTPRequestMethodKey = otel.HTTPRequestMethodKey
	// HTTPResponseStatusCodeKey is the integer status, e.g. 200, 502.
	HTTPResponseStatusCodeKey = otel.HTTPResponseStatusCodeKey
	// HTTPRouteKey is the matched route pattern, e.g. "/users/:id".
	HTTPRouteKey = otel.HTTPRouteKey
)

// Geo
const (
	// GeoContinentCodeKey is a two-letter continent, e.g. "NA", "EU".
	GeoContinentCodeKey = otel.GeoContinentCodeKey
	// GeoCountryISOCodeKey is ISO 3166-1 alpha-2, e.g. "US", "IR".
	GeoCountryISOCodeKey = otel.GeoCountryISOCodeKey
	// GeoRegionISOCodeKey is ISO 3166-2, e.g. "US-CA", "DE-BY".
	GeoRegionISOCodeKey = otel.GeoRegionISOCodeKey
	// GeoLocalityNameKey is the city name, e.g. "San Francisco".
	GeoLocalityNameKey = otel.GeoLocalityNameKey
	// GeoLocationLatKey is the latitude, e.g. 37.7749.
	GeoLocationLatKey = otel.GeoLocationLatKey
	// GeoLocationLonKey is the longitude, e.g. -122.4194.
	GeoLocationLonKey = otel.GeoLocationLonKey
)

// Error

// ErrorTypeKey classifies the error, e.g. "timeout", "cancel".
const ErrorTypeKey = otel.ErrorTypeKey
