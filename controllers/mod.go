package controllers

import "go.uber.org/fx"

var FxModule = fx.Module("controllers",
	fx.Provide(
		NewProfileController,
	),
)
