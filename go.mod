module github.com/ipfs-shipyard/gateway-prime

go 1.18

require (
	contrib.go.opencensus.io/exporter/prometheus v0.4.1
	github.com/gabriel-vasile/mimetype v1.4.0
	github.com/ipfs/go-cid v0.2.0
	github.com/ipfs/go-fetcher v1.6.1
	github.com/ipfs/go-log v1.0.5
	github.com/ipfs/go-path v0.3.0
	github.com/ipfs/go-unixfsnode v1.4.1-0.20220502093700-f664db4b2168
	github.com/ipld/go-car/v2 v2.1.1
	github.com/ipld/go-ipld-prime v0.16.0
	github.com/libp2p/go-libp2p v0.19.0
	github.com/libp2p/go-libp2p-core v0.15.1
	github.com/miekg/dns v1.1.48
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multicodec v0.4.1
	github.com/multiformats/go-multihash v0.1.0
	github.com/prometheus/client_golang v1.12.1
	go.opencensus.io v0.23.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/Stebalien/go-bitfield v0.0.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/ipfs/bbloom v0.0.4 // indirect
	github.com/ipfs/go-bitfield v1.0.0 // indirect
	github.com/ipfs/go-block-format v0.0.3 // indirect
	github.com/ipfs/go-blockservice v0.2.1 // indirect
	github.com/ipfs/go-datastore v0.5.1 // indirect
	github.com/ipfs/go-ipfs-blockstore v1.1.2 // indirect
	github.com/ipfs/go-ipfs-chunker v0.0.1 // indirect
	github.com/ipfs/go-ipfs-ds-help v1.1.0 // indirect
	github.com/ipfs/go-ipfs-exchange-interface v0.1.0 // indirect
	github.com/ipfs/go-ipfs-util v0.0.2 // indirect
	github.com/ipfs/go-ipld-cbor v0.0.5 // indirect
	github.com/ipfs/go-ipld-format v0.2.0 // indirect
	github.com/ipfs/go-ipld-legacy v0.1.0 // indirect
	github.com/ipfs/go-log/v2 v2.5.1 // indirect
	github.com/ipfs/go-merkledag v0.5.1 // indirect
	github.com/ipfs/go-metrics-interface v0.0.1 // indirect
	github.com/ipfs/go-verifcid v0.0.1 // indirect
	github.com/ipld/go-codec-dagpb v1.3.0 // indirect
	github.com/jbenet/goprocess v0.1.4 // indirect
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/libp2p/go-buffer-pool v0.0.2 // indirect
	github.com/libp2p/go-eventbus v0.2.1 // indirect
	github.com/libp2p/go-msgio v0.2.0 // indirect
	github.com/libp2p/go-openssl v0.0.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.0.4 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multistream v0.3.0 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/polydawn/refmt v0.0.0-20201211092308-30ac6d18308e // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.33.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/prometheus/statsd_exporter v0.21.0 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/whyrusleeping/cbor-gen v0.0.0-20200123233031-1cdf64d27158 // indirect
	github.com/whyrusleeping/chunker v0.0.0-20181014151217-fe64bd25879f // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/exp v0.0.0-20210615023648-acb5c1269671 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220418201149-a630d4f3e7a2 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	lukechampine.com/blake3 v1.1.7 // indirect
)
