# jaeger-storage-badger

This repository contains a proof of concept to use Badger as a datastore for [Jaeger](https://github.com/jaegertracing/jaeger). The implementation is a copy-paste of this [PR](https://github.com/burmanm/jaeger/pull/760) from [burmanm](https://github.com/burmanm).

It's archived now because the PoC was using [Go plugin](https://golang.org/pkg/plugin/) and a [gRPC approach](https://github.com/jaegertracing/jaeger/tree/v1.19.2/plugin/storage/grpc) was preferred. Details can still be found [here](https://github.com/jaegertracing/jaeger/pull/1050)
