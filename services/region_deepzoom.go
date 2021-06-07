package services

import (
	"github.com/jammy-dodgers/gophenslide/openslide"
)

type RegionDeepZoom struct {
	TileSize   int64
	Overlap    int
	LimitBound int
	Format     string
}

func New(tileSize int64, overlap int, limitbound int, format string) *RegionDeepZoom {
	return &RegionDeepZoom{tileSize, overlap, limitbound, format}
}

func (dz *RegionDeepZoom) LevelCount(path string) int32 {
	openslide, err := openslide.Open(path)
	if err != nil {
		panic("Fail to open slide image")
	}
	levelCount := openslide.LevelCount()
	return levelCount
}

func (dz *RegionDeepZoom) LevelTileDimension(path string, level int32) (int64, int64) {
	openslide, err := openslide.Open(path)
	if err != nil {
		panic("Fail to open slide image")
	}
	w, h := openslide.LevelDimensions(level)
	tileW := w / dz.TileSize
	tileH := h / dz.TileSize
	return tileH, tileW
}
