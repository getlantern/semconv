package semconv

import "go.opentelemetry.io/otel/attribute"

// Deprecated bandit attribute keys — use the canonical keys instead.
// These aliases exist so downstream code compiles during migration.
// Remove after all consumers have been updated.
const (
	// Deprecated: Use ClientAsnKey instead.
	BanditASNKey attribute.Key = ClientAsnKey
	// Deprecated: Use GeoCountryISOCodeKey instead.
	BanditCountryKey attribute.Key = GeoCountryISOCodeKey
	// Deprecated: Use TrackIDKey instead.
	BanditTrackIDKey attribute.Key = TrackIDKey
	// Deprecated: Use ProxyTrackKey instead.
	BanditTrackNameKey attribute.Key = ProxyTrackKey
	// Deprecated: Use RouteRegionKey instead.
	BanditRegionNameKey attribute.Key = RouteRegionKey
	// Deprecated: Use ProxyProviderKey instead.
	BanditProviderKey attribute.Key = ProxyProviderKey
	// Deprecated: Use RouteIDKey instead.
	BanditRouteIDKey attribute.Key = RouteIDKey
)
