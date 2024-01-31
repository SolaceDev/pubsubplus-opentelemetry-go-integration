# Solace: solace.dev/go/messaging-trace/opentelemetry
[![Test](https://github.com/SolaceProducts/pubsubplus-opentelemetry-go-integration/actions/workflows/test.yml/badge.svg)](https://github.com/SolaceProducts/pubsubplus-opentelemetry-go-integration/actions/workflows/test.yml)
[![PkgGoDev](https://pkg.go.dev/badge/solace.dev/go/messaging-trace/opentelemetry.svg)](https://pkg.go.dev/solace.dev/go/messaging-trace/opentelemetry)
[![Go Report Card](https://goreportcard.com/badge/solace.dev/go/messaging-trace/opentelemetry)](https://goreportcard.com/report/solace.dev/go/messaging-trace/opentelemetry)

The Solace PubSub+ OpenTelemetry Integration API for Go is used to handle injection and extraction of tracing information for carrier objects (Solace message) for a Solace PubSub+ Event Broker.

## Getting Started

To get started using the Solace PubSub+ OpenTelemetry Integration API for Go, simply include it as a required module in your Go project by running `go get solace.dev/go/messaging-trace/opentelemetry`. The Solace PubSub+ OpenTelemetry Integration API for Go requires Go version 1.20+ (based on the minimum compatible version for opentelemtry-Go [https://github.com/open-telemetry/opentelemetry-go/blob/main/README.md](https://github.com/open-telemetry/opentelemetry-go/blob/main/README.md#compatibility)).

### Usage

- Sample applications https://github.com/SolaceSamples/solace-samples-go.
- Documentation https://docs.solace.com/Solace-PubSub-Messaging-APIs/Go-API/go-home.htm
- Developer reference https://docs.solace.com/API-Developer-Online-Ref-Documentation/go/

### OS Support

The Solace PubSub+ OpenTelemetry Integration API for Go has a hard depedency on the Solace PubSub+ API which is a wrapper around the high performance Solace C API via Cgo and has support for the following operating systems:
- Linux (x86/x86_64) variants with Linux 2.6 or later (compatible with glibc (desktop/server) and musl-c (Alpine Linux))
- Linux (arm64) variants compatible with glibc (desktop/server)
- Windows WSL 2.0
- macOS 10.15 and later (x86_64 versions)
- macOS 11.0 and later (arm64 versions)

## Contributing

Ensure that you read [CONTRIBUTING](CONTRIBUTING.md) for details on the process for submitting pull requests to us.

## Authors

See the list of [contributors](https://github.com/SolaceProducts/pubsubplus-opentelemetry-go-integration/graphs/contributors) who participated in this project.

## License

This project is licensed under the Apache License, Version 2.0. - See the [LICENSE](LICENSE.txt) file for details.

This project packages and links against the Solace PubSub+ API. See [the licenses for the packgae](https://github.com/SolaceProducts/pubsubplus-go-client/internal/ccsmp/lib/licenses.txt) for details.

## Code of Conduct

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)
Note that this project is released with a Contributor Code of Conduct. By participating in this project, you agree to abide by its terms.

## Support

### Ask Solace Community

Have a question? Ask the [Solace Community](https://dev.solace.com/community/)!

### Ask Solace Support

https://solace.com/support

## Resources

- The [Solace Developer Portal](https://dev.solace.com)
- Understanding [Solace technology](https://solace.com/products/tech/)

Copyright 2024 Solace Corporation. All rights reserved.
