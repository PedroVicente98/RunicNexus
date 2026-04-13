// Package grpcproxy provides a built-in RunicNexus module that bridges
// the Go engine to a C++ gRPC back-end service. It implements the
// runic.Module interface and proxies incoming Envelope packets over a
// gRPC channel.
package grpcproxy

import "github.com/PedroVicente98/RunicNexus/pkg/runic"

// Proxy is the gRPC bridge module.
type Proxy struct {
	// Target is the gRPC server address to proxy requests to,
	// e.g. "localhost:50051".
	Target string
}

// Name returns the module identifier.
func (p *Proxy) Name() string { return "grpcproxy" }

// Init satisfies the runic.Module interface. It establishes the
// underlying gRPC client connection.
func (p *Proxy) Init(_ *runic.Engine) error {
	// TODO: dial p.Target and store the client connection.
	return nil
}
