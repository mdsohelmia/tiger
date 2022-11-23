package controllers

import (
	"github.com/MdSohelMia/tiger/api/controllers/library"
	"github.com/MdSohelMia/tiger/api/controllers/video"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(video.NewVideoController),
	fx.Provide(library.NewLibraryController),
	fx.Provide(NewIndexController),
)
