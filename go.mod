module solace.dev/go/trace/propagation

go 1.17

require solace.dev/go/messaging v1.4.0

replace solace.dev/go/trace/propagation => ./pubsubplus-opentelemetry-go-integration
replace solace.dev/go/messaging v1.4.0 => ../pubsubplus-go-client
