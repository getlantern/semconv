# semconv

Standard and custom semantic conventions for creating consistent, comparable `attribute.Key`s across services.

## usage

Import `semconv` and use it exclusively for defining attribute keys in telemetry.
```go
import "github.com/getlantern/semconv"
```

Example attribute map:
```go
attrs := map[attribute.Key]attribute.Value{
	semconv.GeoCountryISOCodeKey: attribute.StringValue("US"),
	semconv.HostNameKey:          attribute.StringValue("phost-abcd"),
	semconv.TrackNameKey:         attribute.StringValue("nidoran"),
	semconv.ProxyProtocolKey:     attribute.StringValue("HTTPS"),
}
```

Example attribute slice:
```go
attrs := []attribute.KeyValue{
	{
		Key:   semconv.GeoCountryISOCodeKey,
		Value: attribute.StringValue("US"),
	},
	{
		Key:   semconv.HostNameKey,
		Value: attribute.StringValue("phost-abcd"),
	},
	{
		Key:   semconv.TrackNameKey,
		Value: attribute.StringValue("nidoran"),
	},
	{
		Key:   semconv.ProxyProtocolKey,
		Value: attribute.StringValue("HTTPS"),
	},
}
```

## reference

### standard
- https://opentelemetry.io/docs/concepts/semantic-conventions/
- https://pkg.go.dev/go.opentelemetry.io/otel/semconv

### custom
- [`semconv.go`](./semconv.go)

