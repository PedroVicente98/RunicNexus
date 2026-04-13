package runic

// Module is the extension interface that all RunicNexus modules must
// implement. Built-in modules (e.g. grpcproxy) and user-defined modules
// both satisfy this interface.
type Module interface {
	// Init is called once by the Engine during startup. The engine
	// reference allows modules to register additional handlers or
	// interact with other modules.
	Init(e *Engine) error

	// Name returns a unique human-readable identifier for the module,
	// used for logging and dependency resolution.
	Name() string
}
