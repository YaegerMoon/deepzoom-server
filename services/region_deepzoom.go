package services

import (
	"math"

	"github.com/jammy-dodgers/gophenslide/openslide"
)

type Area struct {
	Top    int64
	Left   int64
	Right  int64
	Bottom int64
}

type RegionDeepZoom struct {
	slide      openslide.Slide
	TileSize   int64
	Overlap    int32
	LimitBound int
	Format     string
	Bound      Area
}

func New(path string, tileSize int64, overlap int32, limitbound int, format string, bound Area) *RegionDeepZoom {
	openslide, err := openslide.Open(path)
	if err != nil {
		panic("Fail to open slide image")
	}
	return &RegionDeepZoom{openslide, tileSize, overlap, limitbound, format, bound}
}

func (dz *RegionDeepZoom) LevelCount() int32 {

	levelCount := dz.slide.LevelCount()
	return levelCount
}

func (dz *RegionDeepZoom) TileDimension(level int32) (int64, int64) {
	w, h := dz.slide.LevelDimensions(level)
	tileW := w / dz.TileSize
	tileH := h / dz.TileSize
	return tileH, tileW
}

func (dz *RegionDeepZoom) TileCount(level int32) int64 {
	tileW, tileH := dz.TileDimension(level)
	return tileW * tileH
}

func (dz *RegionDeepZoom) level0ZoomDownsample(level int32) float64 {
	levelCount := dz.LevelCount()
	downsample := math.Pow(2, float64(levelCount-level-1))
	return downsample
}

func (dz *RegionDeepZoom) level0LevelDownsample(level int32) float64 {
	return dz.slide.LevelDownsample(level)
}

func (dz *RegionDeepZoom) levelZoomDownsmaple(level int32) float64 {
	levelCount := dz.LevelCount()
	ds := int32(dz.level0LevelDownsample(levelCount))
	return dz.level0ZoomDownsample(level) / dz.level0LevelDownsample(ds)
}

func (dz *RegionDeepZoom) deepZoomLevel(level int32) int32 {
	downsample := dz.level0ZoomDownsample(level)
	deepzoomLevel := dz.slide.BestLevelForDownsample(downsample)
	return deepzoomLevel
}

// func (dz *RegionDeepZoom) Tile(level int32, col int64, row int64) {
// 	tile, error :=n dz.slide.ReadRegion()
// }

func (dz *RegionDeepZoom) tileInfo(level int32, col int32, row int32) {
	tileWithLimit, tileHeightLimit := dz.TileDimension(level)

	if level < 0 || level >= dz.LevelCount() {
		panic("Invalid Level")
	}
	if col < 0 || int64(col) >= tileWithLimit {
		panic("Invalid column")
	}
	if row < 0 || int64(row) >= tileHeightLimit {
		panic("Invalid row")
	}
	slideLevel := dz.deepZoomLevel(level)
	overlapTop := dz.Overlap
	overlapBottom := dz.Overlap
	overlapLeft := dz.Overlap
	overlapRight := dz.Overlap
	if col == 0 {
		overlapLeft = 0
	}
	if row == 0 {
		overlapTop = 0
	}
	if tileWithLimit == int64(col+1) {
		overlapRight = 0
	}
	if tileHeightLimit == int64(row+1) {
		overlapBottom = 0
	}
	width := int32(math.Min(float64(dz.TileSize), float64(tileWithLimit-dz.TileSize*int64(col)))) + overlapLeft + overlapRight
	height := int32(math.Min(float64(dz.TileSize), float64(tileHeightLimit-dz.TileSize*int64(row)))) + overlapTop + overlapBottom

}

func (dz *RegionDeepZoom) levelFromLevel0(slideLevel int32, level0 int32) float64 {
	return float64(level0) * dz.level0LevelDownsample(slideLevel)
}

func (dz *RegionDeepZoom) level0FromLevel(slideLevel int32, level int32) float64 {
	return dz.level0LevelDownsample(slideLevel) * float64(level)
}

func (dz *RegionDeepZoom) levelFromZoom(dzLevel int32, zoom float64) float64 {
	return dz.levelZoomDownsmaple(dzLevel) * zoom
}

func (dz *RegionDeepZoom) zoomFromTile(tile int32) int64 {
	return dz.TileSize * int64(tile)
}
