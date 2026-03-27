# semconv

This package defines the canonical attribute keys for all telemetry (metrics,
traces, logs) across Lantern services. It re-exports standard OTel semconv and
adds custom keys for Lantern-specific concepts.

## Rules

1. **Avoid hardcoded attribute key strings.** Use `semconv.*` constants to aid correlation.
2. **Prefer standard OTel keys** when one exists for the concept. DNS, database
   services, and internal OTel instrumentation already emit standard keys, so
   using them ensures correlation. Only use custom keys for Lantern-specific
   domains (proxies, routes, tracks, bandit, clients, EIPs).
3. **Align with existing usage.** Before adding a new key, check `custom.go`
   and `standard.go`, and grep across services for how similar resources are
   already described.
4. **Naming matters for incident response.** These keys power dashboards and
   alerts. Consistent naming lets on-call engineers query across services.

## Structure

- `standard.go` — re-exports of standard OTel semconv symbols (pinned to
  `v1.37.0`). Add new re-exports here when a service needs an upstream key.
- `custom.go` — Lantern-specific attribute keys, grouped by domain

## Adding Keys

### Standard (OTel)
Add a `const` re-export to `standard.go` under the appropriate group. Use type
inference (no explicit `attribute.Key`). Check
https://pkg.go.dev/go.opentelemetry.io/otel/semconv for available keys.

### Custom (Lantern)
Add a typed `attribute.Key` constant to `custom.go` under the appropriate
domain group. Use the `domain.field_name` naming convention, e.g.
`proxy.protocol`, `route.region`. If no group fits, create a new one with a
comment header.

## Expected Resources

Every service must set these on its telemetry resource. Dashboards and alerts
depend on them for filtering and grouping during incident response.

- `semconv.ServiceNameKey` — identifies the service
- `semconv.ServiceVersionKey` — deployment error rate monitoring
- `semconv.DeploymentEnvironmentNameKey` — environment (prod/staging/dev)
