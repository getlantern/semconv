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
	ServiceNameKey      = otel.ServiceNameKey
	ServiceVersionKey   = otel.ServiceVersionKey
	ServiceNamespaceKey = otel.ServiceNamespaceKey
)

// Service convenience constructors
var (
	ServiceName      = otel.ServiceName
	ServiceVersion   = otel.ServiceVersion
	ServiceNamespace = otel.ServiceNamespace
)

// Deployment
const DeploymentEnvironmentNameKey = otel.DeploymentEnvironmentNameKey

var DeploymentEnvironmentName = otel.DeploymentEnvironmentName

// Host
const (
	HostNameKey = otel.HostNameKey
	HostIDKey   = otel.HostIDKey
	HostArchKey = otel.HostArchKey
)

// Network
const (
	NetworkPeerAddressKey attribute.Key = otel.NetworkPeerAddressKey
	NetworkPeerPortKey    attribute.Key = otel.NetworkPeerPortKey
)

// HTTP
const (
	HTTPRequestMethodKey      attribute.Key = otel.HTTPRequestMethodKey
	HTTPResponseStatusCodeKey attribute.Key = otel.HTTPResponseStatusCodeKey
	HTTPRouteKey              attribute.Key = otel.HTTPRouteKey
)

// Error
const ErrorTypeKey attribute.Key = otel.ErrorTypeKey
