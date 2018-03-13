package prometheus

import "go.uber.org/fx"

type ModuleConfig struct {
	// Use the global default registerer and gatherer instances.
	UseGlobal bool

	// MakeGlobal makes the registerer and gatherer instances the global default instances.
	// This option can only be used when UseGlobal is false.
	//
	// Note: this modifies global state which is generally not a good idea.
	// Make sure you know what you do when using this option.
	MakeGlobal bool
}

func Module(config ModuleConfig) fx.Option {
	options := []fx.Option{
		fx.Invoke(RegisterCollectors),
	}

	if config.UseGlobal {
		options = append(options, fx.Provide(Global))
	} else {
		options = append(options, fx.Provide(New))

		if config.MakeGlobal {
			options = append(options, fx.Invoke(MakeGlobal))
		}
	}

	return fx.Options(options...)
}
