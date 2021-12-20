# Currency Service

[![GoDoc](https://godoc.org/github.com/qiangxue/go-rest-api?status.png)](https://bitbucket.org/lgaetecl/microservices-test/src/master/)
[![Build Status](https://github.com/qiangxue/go-rest-api/workflows/build/badge.svg)](https://github.com/jkarlos000/technical-challenge/actions?query=workflow%3Abuild)
[![Code Coverage](https://codecov.io/gh/qiangxue/go-rest-api/branch/master/graph/badge.svg)](https://app.codecov.io/gh/jkarlos000/technical-challenge)

The currency service is a gRPC service which provides up to date exchange rates and currency conversion capabilities, the is obtain from [CurrencyLayer](https://currencylayer.com/)


This project is part of technical challenge of CleverIT

## Building protos
To build the gRPC client and server interfaces, first install protoc:

### Linux
```shell
sudo apt install protobuf-compiler
```

### Mac
```shell
brew install protoc
```

Then install the Go gRPC plugin:

```shell
go get google.golang.org/grpc
```

Then run the build command:

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/v1/currency.proto
```

