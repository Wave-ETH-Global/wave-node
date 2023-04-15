package repositories

import "go.uber.org/fx"

var FxModule = fx.Module("repositories",
	fx.Provide(
		NewProfileRepository,
	),
)
