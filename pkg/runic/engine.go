// Package runic is the public-facing core of the RunicNexus framework.
// Users import this package to create and configure a RunicNexus Engine.
package runic

// Config holds the top-level configuration for a RunicNexus Engine.
type Config struct {
	// Addr is the TCP/WebSocket listen address, e.g. ":7000".
	Addr string
}

// Engine is the central runtime of the RunicNexus framework. It wires
// together the network listeners, broker, and all registered modules.
type Engine struct {
	cfg     Config
	modules []Module
}

// NewEngine creates a new Engine with the provided configuration.
func NewEngine(cfg Config) *Engine {
	return &Engine{cfg: cfg}
}

// Register attaches a Module to the engine. Modules are initialised in
// the order they are registered when Engine.Start is called.
func (e *Engine) Register(m Module) {
	e.modules = append(e.modules, m)
}

// Start initializes all registered modules and begins accepting
// connections on the configured address.
func (e *Engine) Start() error {
	for _, m := range e.modules {
		if err := m.Init(e); err != nil {
			return err
		}
	}
	return nil
}
