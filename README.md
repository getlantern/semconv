# semconv

Standard and custom semantic conventions for creating consistent, comparable `attribute.Key`s across services.

## usage

Import `semconv` and use it exclusively for defining attribute keys in telemetry.
```go
import "github.com/getlantern/semconv"
```

Standard OTel keys may have convenience functions, but custom keys do not.

```go
attrs := map[attribute.Key]attribute.Value{
	semconv.GeoCountryISOCodeKey: attribute.StringValue("US"),         // standard
	semconv.HostNameKey:          attribute.StringValue("phost-abcd"), // standard
	semconv.TrackNameKey:         attribute.StringValue("nidoran"),    // custom
	semconv.ProxyProtocolKey:     attribute.StringValue("HTTPS"),      // custom
}
```

```go
attrs := []attribute.KeyValue{
	semconv.GeoCountryISOCode("US"),
	semconv.HostName("phost-abcd"),
	attribute.KeyValue{
		Key:   semconv.TrackNameKey,
		Value: attribute.StringValue("nidoran"),
	},
	attribute.KeyValue{
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

