//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/game"
	"mrcAPI/pkg/game/internal/application"
	"mrcAPI/pkg/game/internal/infrastructure"
	"mrcAPI/pkg/system"
)

func provideGameMysqlRepository() (infrastructure.GameMysqlRepository, error) {
	panic(wire.Build(
		system.ProvideMySQL,
		infrastructure.NewGameMysqlRepository,
	))
}

func ProvideGame() (game.IGame, error) {
	panic(wire.Build(
		provideGameMysqlRepository,
		application.NewGameService,
		wire.Bind(new(game.IGame), new(application.Game)),
	))
}
