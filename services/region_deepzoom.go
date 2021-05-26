package services

type RegionDeepZoom struct {
	TileSize   int
	Overlap    int
	LimitBound int
	Format     string
}

func New(tileSize int, overlap int, limitbound int, format string) *RegionDeepZoom {
	return &RegionDeepZoom{tileSize, overlap, limitbound, format}
}
