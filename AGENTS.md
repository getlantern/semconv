# semconv

This package defines the canonical attribute keys for all telemetry (metrics,
traces, logs) across Lantern services. It re-exports standard OTel semconv and
adds custom keys for Lantern-specific concepts.

## Rules

1. **Avoid hardcoded attribute key strings.** Use `semconv.*` constants to aid correlation.
2. **Prefer standard OTel keys** when one exists for the concept. Only use
   custom keys for Lantern-specific domains (proxies, routes, tracks, bandit,
   clients, EIPs).
3. **Align with existing usage.** Before adding a new key, check `semconv.go`
   and grep across services for how similar resources are already described.
4. **Naming matters for incident response.** These keys power dashboards and
   alerts. Consistent naming lets on-call engineers query across services.

## Structure

- `semconv.go` — all key definitions, grouped by domain
- Standard OTel keys are available via dot-import of
  `go.opentelemetry.io/otel/semconv/v1.37.0`

## Expected Resources

Every service must set these on its telemetry resource. Dashboards and alerts
depend on them for filtering and grouping during incident response.

- `semconv.ServiceNameKey` — identifies the service
- `semconv.ServiceVersionKey` — deployment error rate monitoring
- `semconv.DeploymentEnvironmentNameKey` — environment (prod/staging/dev)
