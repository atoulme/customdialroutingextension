#!/bin/bash

cat <<'EOF' | ./otelcol-dev/otelcol-test \
  --config "yaml:$(cat)" --feature-gates=+pkg.configmiddleware.CustomTCPRouting
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
EOF