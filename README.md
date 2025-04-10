# Flipcash Protobuf API

[![Release](https://img.shields.io/github/v/release/code-payments/flipcash-protobuf-api.svg)](https://github.com/code-payments/flipcash-protobuf-api/releases/latest)
[![GitHub License](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/code-payments/flipcash-protobuf-api/blob/main/LICENSE.md)

The APIs and models for communication between Flipcash clients and server.

## Services

- [Account](https://github.com/code-payments/flipcash-protobuf-api/blob/main/proto/account/v1/account_service.proto)
- [IAP](https://github.com/code-payments/flipcash-protobuf-api/blob/main/proto/iap/v1/iap_service.proto)
- [Push](https://github.com/code-payments/flipcash-protobuf-api/blob/main/proto/push/v1/push_service.proto)

## Code Generation

Generated code can be found under the `generated/` directory. The following languages are directly supported:
- Go
- TypeScript/JavaScript (via [Protobuf-ES](https://github.com/bufbuild/protobuf-es))

To generate all code, run:

```bash
make
```

Or to generate a specific language, run:

```bash
make {go, protobuf-es}
```
