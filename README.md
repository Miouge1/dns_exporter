# dns_exporter

A Prometheus exporter for measuring DNS queries

## Usage
```
$ ./dns_exporter google.com github.com
$ curl localhost:9100/metrics
dns_a_replies{domain="google.com"} 2
dns_a_replies{domain="github.com"} 2
```
