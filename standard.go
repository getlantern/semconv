package semconv

import (
	otel "go.opentelemetry.io/otel/semconv/v1.37.0"

	"go.opentelemetry.io/otel/attribute"
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

// Service convenience constructors
var (
	ServiceName      = otel.ServiceName
	ServiceVersion   = otel.ServiceVersion
	ServiceNamespace = otel.ServiceNamespace
)

// Deployment

// DeploymentEnvironmentNameKey is "prod", "staging", or "dev".
const DeploymentEnvironmentNameKey = otel.DeploymentEnvironmentNameKey

var DeploymentEnvironmentName = otel.DeploymentEnvironmentName

// Host
const (
	// HostNameKey is the OS hostname.
	HostNameKey = otel.HostNameKey
	// HostIDKey is a unique host identifier.
	HostIDKey = otel.HostIDKey
	// HostArchKey is the CPU architecture, e.g. "amd64", "arm64".
	HostArchKey = otel.HostArchKey
)

// Network
const (
	// NetworkPeerAddressKey is the remote IP, e.g. "10.0.0.5".
	NetworkPeerAddressKey attribute.Key = otel.NetworkPeerAddressKey
	// NetworkPeerPortKey is the remote port, e.g. 443.
	NetworkPeerPortKey attribute.Key = otel.NetworkPeerPortKey
)

// HTTP
const (
	// HTTPRequestMethodKey is "GET", "POST", etc.
	HTTPRequestMethodKey attribute.Key = otel.HTTPRequestMethodKey
	// HTTPResponseStatusCodeKey is the integer status, e.g. 200, 502.
	HTTPResponseStatusCodeKey attribute.Key = otel.HTTPResponseStatusCodeKey
	// HTTPRouteKey is the matched route pattern, e.g. "/users/:id".
	HTTPRouteKey attribute.Key = otel.HTTPRouteKey
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
const ErrorTypeKey attribute.Key = otel.ErrorTypeKey
