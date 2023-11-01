## How to Contribute to a Solace Project

#### **Did you find a bug?**

* **Ensure the bug was not already reported** by searching on GitHub under [Issues](https://github.com/SolaceProducts/pubsubplus-opentelemetry-go-integration/issues).

* If you're unable to find an open issue that addresses the problem, [open a new one](https://github.com/SolaceProducts/pubsubplus-opentelemetry-go-integration/issues/new). Be sure to include a **title and clear description**, as much relevant information as possible, and a **code sample** or an **executable test case** demonstrating the expected behavior that is not occurring.

#### **Did you write a patch that fixes a bug?**

* Open a new GitHub pull request with the patch.

* Ensure the PR description clearly describes the problem and solution. Include the relevant issue number if applicable.

#### **Do you intend to add a new feature or change an existing one?**

* Open a GitHub [enhancement request issue](https://github.com/SolaceProducts/pubsubplus-opentelemetry-go-integration/issues/new) and describe the new functionality.

#### **Do you have questions about the source code?**

* Ask any question about the code or how to use Solace PubSub+ OpenTelemetry Integration API in the [Solace community](https://solace.dev/community/).

## Developer Guide

### Overview

The Solace PubSub+ OpenTelemetry Integration API for Go is used to handle injection and extraction of tracing information for carrier objects (Solace message) for a Solace PubSub+ Event Broker. It has a dependency on the [Solace PubSub+ Messaging API for Go](https://docs.solace.com/API/Messaging-APIs/Go-API/go-home.htm).

### Repository Contents

The following are the directories that are part of the repository:
- `./`: contains the API's main entrypoints such as tracing.
- `propagation`: contains the API's interfaces, constants and basic data structures
- `internal`: contains the implementation for the PubSub+ OpenTelemetry Integration API for Go
- `test`: contains integration tests for the PubSub+ OpenTelemetry Integration API for Go

### Prerequisites

There are a handful of prerequisites for developing the Solace PubSub+ OpenTelemetry Integration API for Go:
- Golang version 1.17+
- A golang enabled code editor, preferably with format on save
    - https://github.com/fatih/vim-go
    - https://code.visualstudio.com/docs/languages/go
- (optional) Docker used for testing
- (optional) Godoc Static to generate a static version of the documentation

Note:
- The integration tests are their own go module. This is currently not compatible with the [go language server](https://github.com/golang/tools/blob/master/gopls/README.md) which is the default used in VS Code. In particular, there is only support for a single module per workspace. The solution to this is to open the tests in their own workspace.

### Testing

Any new features added to the API must have corresponding tests added to the [integration test project](./test). See [the test README](./test/README.md) for details on running the tests. Any new feature must not decrease overall test coverage.

#### Github Actions

The integration tests are run on every new commit via Github Actions. These tests are run on the latest Go version on Linux.

### Generating Documentation

Documentation should work in all cases using `godoc` in the module and navigating to the local webserver. 

#### Generating Static Documentation

In addition, static documentation can be generated. To do this, first install godoc-static, and then run `go run docs-templates/generate_documentation.go`.
