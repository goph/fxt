package fxprometheus

import "go.uber.org/fx"

type ModuleConfig struct {
	// Use the global default registerer and gatherer instances.
	UseGlobal bool

	// MakeGlobal makes the registerer and gatherer instances the global default instances.
	//
	// This option can only be used when UseGlobal is false.
	//
	// Note: this modifies global state which is generally not a good idea.
	// Make sure you know what you do when using this option.
	MakeGlobal bool

	// MergeGlobalGatherer causes the global gatherer instance to be merged with a newly created one.
	// This is useful when some third-party libraries use the global gatherer instead of offering an extension point.
	//
	// This option can only be used when UseGlobal is false.
	MergeGlobalGatherer bool

	// RegisterDefaultCollectors registers the default Process and Go collectors in a newly created registerer instance.
	//
	// This option can only be used when UseGlobal is false.
	//
	// Note: Do not use this option with MergeGlobalGatherer because the new and the global gatherer might collide.
	RegisterDefaultCollectors bool
}

func Module(config ModuleConfig) fx.Option {
	options := []fx.Option{
		fx.Invoke(RegisterCollectors),
	}

	if config.UseGlobal {
		options = append(options, fx.Provide(Global))
	} else {
		if config.MergeGlobalGatherer {
			options = append(options, fx.Provide(NewWithGlobal))
		} else {
			options = append(options, fx.Provide(New))
		}

		if config.RegisterDefaultCollectors {
			options = append(options, fx.Invoke(RegisterDefaultCollectors))
		}

		if config.MakeGlobal {
			options = append(options, fx.Invoke(MakeGlobal))
		}
	}

	return fx.Options(options...)
}
