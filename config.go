package otelpubsub

import (
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type config struct {
	tracerProvider trace.TracerProvider
	propagators    propagation.TextMapPropagator
}

// Option applies a configuration value.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

// Apply will apply the option to the config, c.
func (o optionFunc) apply(c *config) {
	o(c)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.tracerProvider = provider
		}
	})
}

// WithPropagators sets the propagators used in the middleware.
// Defaults to global.Propagators().
func WithPropagators(propagators propagation.TextMapPropagator) Option {
	return optionFunc(func(c *config) {
		if propagators != nil {
			c.propagators = propagators
		}
	})
}
