package services

import "github.com/jammy-dodgers/gophenslide/openslide"

type RegionDeepZoom struct {
	TileSize   int
	Overlap    int
	LimitBound int
	Format     string
}

func New(tileSize int, overlap int, limitbound int, format string) *RegionDeepZoom {
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
