package tvdb

import (
	"github.com/dghubble/sling"
)

// EpisodesService the episode service
type EpisodesService struct {
	sling *sling.Sling
}

// newSeriesService initialize a new SeriesService
func newEpisodesService(sling *sling.Sling) *EpisodesService {
	return &EpisodesService{
		sling: sling,
	}
}
