# Custom Dial Routing Extension

This extension allows an advanced behavior of the OpenTelemetry Collector where an operator can now configure a HTTP client or gRPC client to dial out via a specific network interface.

This is meant as much of an example of use of the Dialer extension allowed in the collector configuration as a practical use case.

Please offer feedback in issues. This repository is likely short-lived and eyeing a contribution to the OpenTelemetry Collector Contrib repository.

## Usage

Requires v0.165.0 and above.

The collector must run with the feature gate `pkg.configmiddleware.CustomTCPRouting`:

```
./otelcol-contrib --feature-gates=+pkg.configmiddleware.CustomTCPRouting
```

```
extensions:
  customdialrouting:
    network_interface: en0

exporters:
  otlphttp:
    endpoint: http://myhost:4318
    dialer_middleware:
      id: customdialrouting
  debug:
    verbosity: detailed

receivers:
  http_check:
    targets:
        - endpoint: http://example.com
          dialer_middleware:
            id: customdialrouting


service:
  extensions: [customdialrouting]
  pipelines:
    metrics:
      receivers: [http_check]
      exporters: [otlphttp, debug]
```

