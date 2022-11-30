package controllers

import (
	"go.uber.org/fx"
	"gotipath.com/gostream/api/controllers/collection"
	"gotipath.com/gostream/api/controllers/library"
	"gotipath.com/gostream/api/controllers/video"
)

var Module = fx.Options(
	fx.Provide(video.NewVideoController),
	fx.Provide(library.NewLibraryController),
	fx.Provide(NewIndexController),
	fx.Provide(collection.NewCollectionController),
)
