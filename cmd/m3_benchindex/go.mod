module github.com/chronosphereiox/high_cardinality_microbenchmark/cmd/m3_benchindex

go 1.13

require (
	github.com/MichaelTJones/pcg v0.0.0-20180122055547-df440c6ed7ed // indirect
	github.com/apache/thrift v0.13.0 // indirect
	github.com/apache/thrift/lib/go/thrift v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/couchbase/vellum v0.0.0-20190829182332-ef2e028c01fd // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/mock v1.4.0 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.12.2 // indirect
	github.com/hydrogen18/stalecucumber v0.0.0-20180226003526-6de214d141dd // indirect
	github.com/jhump/protoreflect v1.6.0 // indirect
	github.com/leanovate/gopter v0.2.7 // indirect
	github.com/m3db/bitset v2.0.0+incompatible // indirect
	github.com/m3db/bloom v3.0.0+incompatible // indirect
	github.com/m3db/m3 v0.14.2
	github.com/m3db/pilosa v0.0.0-20190128031222-ac8920c6e1ab // indirect
	github.com/m3db/prometheus_client_golang v0.8.1 // indirect
	github.com/m3db/prometheus_client_model v0.1.0 // indirect
	github.com/m3db/prometheus_common v0.1.0 // indirect
	github.com/m3db/prometheus_procfs v0.8.1 // indirect
	github.com/m3db/stackadler32 v0.0.0-20180104200216-bfebcd73ef6f // indirect
	github.com/m3db/stackmurmur3 v0.0.0-20171110233611-744c0229c12e // indirect
	github.com/m3db/vellum v0.0.0-20190111185746-e766292d14de // indirect
	github.com/mauricelam/genny v0.0.0-20190320071652-0800202903e5 // indirect
	github.com/prometheus/client_golang v1.4.0 // indirect
	github.com/prometheus/prometheus v1.8.2-0.20190818123050-43acd0e2e93f
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	github.com/uber-go/atomic v1.4.0
	github.com/uber-go/tally v3.3.14+incompatible // indirect
	github.com/uber/jaeger-client-go v2.22.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/uber/tchannel-go v1.16.0 // indirect
	go.etcd.io/etcd v0.5.0-alpha.5.0.20191023171146-3cf2f69b5738
	go.uber.org/config v1.4.0 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	gopkg.in/validator.v2 v2.0.0-20191107172027-c3144fdedc21 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1 // indirect
)

// Use some specific dependencies
replace github.com/apache/thrift/lib/go/thrift => github.com/m3dbx/thrift/lib/go/thrift v0.0.0-20200106002022-da72b4507a76

replace github.com/apache/thrift => github.com/m3dbx/thrift v0.0.0-20200106002022-da72b4507a76

replace google.golang.org/grpc => google.golang.org/grpc v1.23.1
