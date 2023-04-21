package services

import (
	"go.uber.org/fx"
)

var FxModule = fx.Module("services",
	fx.Provide(
		NewLoginService,
	),
)
