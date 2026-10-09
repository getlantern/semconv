// Package semconv defines semantic conventions for Lantern telemetry.
// Standard OTel conventions are re-exported in standard.go; this file
// contains Lantern-specific keys.
package semconv

import "go.opentelemetry.io/otel/attribute"

// Lantern cloud Proxies
const (
	ProxyCIDRKey           attribute.Key = "proxy.cidr"
	ProxyGatewayKey        attribute.Key = "proxy.gateway"
	HostInterfaceCountKey  attribute.Key = "host.interface.count"
	ProxyFrontendSyncCount attribute.Key = "proxy.frontend_sync_count"
	ProxyFrontendName      attribute.Key = "proxy.frontend_name"
	ProxyHostSyncCount     attribute.Key = "proxy.host_sync_count"
	ProxyHostNameKey       attribute.Key = "proxy.host_name"
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

// Bandit proxy assignment — span names
const (
	BanditSpanSelect            = "bandit.select"
	BanditSpanCallback          = "bandit.callback"
	BanditSpanReaper            = "bandit.reaper"
	BanditSpanReaperExpired     = "bandit.reaper.expired"
	BanditSpanArmCallbackReaper = "bandit.arm_callback_reaper"
)

// Bandit proxy assignment — metric names
const (
	BanditMetricSelections           = "bandit.selections"
	BanditMetricCallbacks            = "bandit.callbacks"
	BanditMetricRepeatCallbacks      = "bandit.repeat_callbacks"
	BanditMetricCallbackLatency      = "bandit.callback_latency_ms"
	BanditMetricReward               = "bandit.reward"
	BanditMetricProbesExpired        = "bandit.probes_expired"
	BanditMetricExpiredProbeAge      = "bandit.expired_probe_age_ms"
	BanditMetricVPSRoutes            = "bandit.vps_routes"
	BanditMetricVPSProvision         = "bandit.vps_provision_ms"
	BanditMetricRoutesBlocked        = "bandit.routes_blocked"
	BanditMetricRoutesBlockedPending = "bandit.routes_blocked_pending"
	BanditMetricRoutesDeprecated     = "bandit.routes_deprecated"
	BanditMetricArmCallbackAbsences  = "bandit.arm_callback_absences"
	// BanditMetricProbeStalls counts callbacks a proxy reported with
	// verdict=stalled: the probe request arrived but the client did not
	// acknowledge the full response body.
	BanditMetricProbeStalls = "bandit.probe_stalls"
)

// Bandit proxy assignment — span attribute keys
const (
	BanditArmIDKey      attribute.Key = "bandit.arm_id"
	BanditASNKey        attribute.Key = "bandit.asn"
	BanditCountryKey    attribute.Key = "bandit.country"
	BanditTrackIDKey    attribute.Key = "bandit.track_id"
	BanditTrackNameKey  attribute.Key = "bandit.track_name"
	BanditRegionNameKey attribute.Key = "bandit.region_name"
	BanditProviderKey   attribute.Key = "bandit.provider"
	BanditRouteIDKey    attribute.Key = "bandit.route_id"
	// BanditCallbackLatencyKey is the proxy roundtrip time (in ms) AFTER
	// subtracting the client-reported queue delay. Paired with
	// BanditCallbackLatencyTotalMsKey and BanditClientQueueDelayMsKey to
	// decompose stall location (client queue vs proxy RTT).
	BanditCallbackLatencyKey attribute.Key = "bandit.callback_latency"
	// BanditCallbackLatencyTotalMsKey is the end-to-end probe time
	// (probe-created → callback-received on the server) BEFORE subtracting
	// the client-reported queue delay.
	BanditCallbackLatencyTotalMsKey attribute.Key = "bandit.callback_latency_total_ms"
	// BanditClientQueueDelayMsKey is the client-reported time (ms) that the
	// URL-test spent in the client's worker-pool queue before the outbound
	// HTTP request fired. Clamped server-side at 80% of the observed latency
	// to prevent reward manipulation via inflated queue reports.
	BanditClientQueueDelayMsKey attribute.Key = "bandit.client_queue_delay_ms"
	// BanditProbeVerdictKey is the proxy's judgement of whether the client
	// received the probe response: delivered, stalled or unknown. Set only on
	// callbacks a proxy answered itself.
	BanditProbeVerdictKey attribute.Key = "bandit.probe_verdict"
	// BanditProbeDrainMsKey is the time (ms) the proxy spent waiting for the
	// client to acknowledge the probe response. It is subtracted from the
	// callback latency so proxy-answered probes rank like relayed ones.
	BanditProbeDrainMsKey attribute.Key = "bandit.probe_drain_ms"
	// BanditProbeAckedBytesKey is the number of bytes the client's TCP stack
	// had acknowledged on the probe connection when the proxy reached its
	// verdict.
	BanditProbeAckedBytesKey attribute.Key = "bandit.probe_acked_bytes"
	// BanditProbeRetransmitsKey is the TCP retransmission count on the probe
	// connection as seen by the proxy.
	BanditProbeRetransmitsKey attribute.Key = "bandit.probe_retransmits"
	// BanditProbeRTTMsKey is the proxy's smoothed TCP RTT (ms) to the client
	// on the probe connection.
	BanditProbeRTTMsKey      attribute.Key = "bandit.probe_rtt_ms"
	BanditProbeAgeSecondsKey attribute.Key = "bandit.probe_age_seconds"
	BanditFirstCallbackKey   attribute.Key = "bandit.first_callback"
	// BanditTokenPrefixKey is a short prefix of the callback probe token used
	// in logs/spans so operators can correlate entries without leaking the
	// full token (which acts as the callback capability).
	BanditTokenPrefixKey         attribute.Key = "bandit.token_prefix"
	BanditNumCandidateRegionsKey attribute.Key = "bandit.num_candidate_regions"
	BanditNumCandidateArmsKey    attribute.Key = "bandit.num_candidate_arms"
	BanditBlockedArmsKey         attribute.Key = "bandit.blocked_arms"
	BanditNumSelectedKey         attribute.Key = "bandit.num_selected"
	BanditSelectedArmsKey        attribute.Key = "bandit.selected_arms"
	// BanditSelectedTracksKey lists the track name of each selected arm,
	// index-aligned with BanditSelectedArmsKey.
	BanditSelectedTracksKey             attribute.Key = "bandit.selected_tracks"
	BanditBlockedRouteCountKey          attribute.Key = "bandit.blocked_route_count"
	BanditDeprecatedCountKey            attribute.Key = "bandit.deprecated_count"
	BanditReaperExpiredProbesKey        attribute.Key = "bandit.reaper.expired_probes"
	BanditArmCallbackReaperNegativesKey attribute.Key = "bandit.arm_callback_reaper.negatives"
	BanditActiveKey                     attribute.Key = "bandit.active"
	BanditCacheHitKey                   attribute.Key = "bandit.cache_hit"
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
	ClientISPKey                attribute.Key = "client.isp"
	ClientAppKey                attribute.Key = "client.app"
	ClientArchitectureKey       attribute.Key = "client.arch"
)

// Proxy resource attributes
const (
	ProxyNameKey             attribute.Key = "proxy.name"
	ProxyProtocolKey         attribute.Key = "proxy.protocol"
	ProxyTrackKey            attribute.Key = "proxy.track"
	ProxyProviderKey         attribute.Key = "proxy.provider"
	ProxyFrontendProviderKey attribute.Key = "proxy.frontend_provider"
	// ProxyBinaryKey names the program serving the proxy (lantern-box,
	// http-proxy-lantern). service.name says where it is deployed, and VPS
	// deployments report vps-proxy for both binaries.
	ProxyBinaryKey attribute.Key = "proxy.binary"
)

// Proxy connection attributes
const (
	ProxyInboundKey     attribute.Key = "proxy.inbound"      // sing-box
	ProxyInboundTypeKey attribute.Key = "proxy.inbound_type" // sing-box
	ProxyOutboundKey    attribute.Key = "proxy.outbound"     // sing-box
)

// Proxy metric names
const (
	ProxyMetricIO          = "proxy.io"          // generic: bytes through a proxy, by direction
	ProxyMetricActivations = "proxy.activations" // unbounded: times proxy mode was turned on and proxied traffic
)

// Geneva server-side packet manipulation — metric names
//
// Deliberately absent: any metric or attribute carrying a strategy DNA or
// candidate identifier. Candidate identity is attributed through server-issued
// opaque tokens in the brain's own tables, never through telemetry labels, so a
// per-candidate label would both leak that identity into SigNoz and make the
// series cardinality unbounded in the GA's population size.
const (
	GenevaMetricPacketsIn      = "geneva.engine.packets_in"
	GenevaMetricPacketsOut     = "geneva.engine.packets_out"
	GenevaMetricBytesIn        = "geneva.engine.bytes_in"
	GenevaMetricBytesOut       = "geneva.engine.bytes_out"
	GenevaMetricOutcomes       = "geneva.engine.outcomes"
	GenevaMetricErrors         = "geneva.engine.errors"
	GenevaMetricPacketOverhead = "geneva.engine.packet_overhead"
	GenevaMetricByteOverhead   = "geneva.engine.byte_overhead"
	GenevaMetricVerdicts       = "geneva.runtime.verdicts"
	GenevaMetricReinjections   = "geneva.runtime.reinjections"
	GenevaMetricStrategySwaps  = "geneva.strategy_swaps"
	GenevaMetricUptime         = "geneva.uptime"

	// GenevaMetricInboundTCP counts inbound TCP packets on the steered port by
	// GenevaTCPEventKey. It is the box-side censor-reachability signal: clients
	// behind a censor cannot report a connection that never completed, so the
	// syn-versus-data ratio per market is the only evidence of a test box IP
	// being burned.
	GenevaMetricInboundTCP = "geneva.censor.inbound_tcp"
)

// Geneva server-side packet manipulation — attribute keys
const (
	// GenevaModeKey is "prod" or "eval".
	GenevaModeKey attribute.Key = "geneva.mode"
	// GenevaOutcomeKey classifies what the strategy did to a packet:
	// "unchanged", "dropped", "tampered", or "expanded".
	GenevaOutcomeKey attribute.Key = "geneva.outcome"
	// GenevaVerdictKey is the NFQUEUE verdict issued: "accepted", "dropped", or
	// "modified".
	GenevaVerdictKey attribute.Key = "geneva.verdict"
	// GenevaReinjectionKey is the raw-socket reinjection result: "ok" or "failed".
	GenevaReinjectionKey attribute.Key = "geneva.reinjection"
	// GenevaTCPEventKey classifies an observed TCP packet by its flags and
	// payload: "syn", "rst", "fin", "ack_only", or "data".
	GenevaTCPEventKey attribute.Key = "geneva.tcp_event"
)

// Grants
const (
	GrantCodeKey      attribute.Key = "grant.code"      // e.g. "otfsurge", "drlvision", "neabuild"
	GrantIndicatorKey attribute.Key = "grant.indicator" // number only (e.g. "2.3", "1.1")
)

// Process crashes (a Go runtime crash dump reported on the next start). The
// crash site uses the standard Code* and Exception* keys; these cover what
// OTel has no key for.
const (
	// CrashTypeKey classifies the crash, e.g. "runtime_panic".
	CrashTypeKey attribute.Key = "crash.type"
	// CrashSignatureKey identifies a crash independent of the values involved,
	// so the same bug groups as one: the panic line with numbers normalised,
	// plus the crashing function, e.g. "panic: runtime error: slice bounds out
	// of range [:N] with capacity N @ github.com/sagernet/sing-vmess/vless.(*VisionConn).filterTLS".
	CrashSignatureKey attribute.Key = "crash.signature"
)
