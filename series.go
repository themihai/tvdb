package tvdb

import (
	"github.com/dghubble/sling"
)

// SeriesService tv series service
type SeriesService struct {
	sling *sling.Sling
}

// NewSeriesService initialize a new SeriesService
func newSeriesService(sling *sling.Sling) *SeriesService {
	return &SeriesService{
		sling: sling,
	}
}
