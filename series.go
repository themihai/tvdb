package tvdb

import (
	"fmt"
	"github.com/dghubble/sling"
)

// Series type
type Series struct {
	ID              string   `json:"id"`
	SeriesName      string   `json:"seriesName"`
	Aliases         []string `json:"aliases,omitempty"`
	Banner          string   `json:"banner,omitempty"`
	SeriesID        int32    `json:"seriesId,omitempty"`
	Status          string   `json:"status,omitempty"`
	FirstAired      string   `json:"firstAired,omitempty"`
	Network         string   `json:"network,omitempty"`
	NetworkID       string   `json:"networkId,omitempty"`
	Runtime         string   `json:"runtime,omitempty"`
	Genre           []string `json:"genre,omitempty"`
	Overview        string   `json:"overview,omitempty"`
	LastUpdated     int32    `json:"lastUpdated,omitempty"`
	AirsDayOfWeek   string   `json:"airsDayOfWeek,omitempty"`
	AirsTime        string   `json:"airsTime,omitempty"`
	Rating          string   `json:"rating,omitempty"`
	ImdbID          string   `json:"imdbId,omitempty"`
	Zap2itID        string   `json:"zap2itId,omitempty"`
	Added           string   `json:"added,omitempty"`
	SiteRating      float32  `json:"siteRating,omitempty"`
	SiteRatingCount int32    `json:"siteRatingCount,omitempty"`
}

// SeriesService TV Series service
type SeriesService struct {
	sling *sling.Sling
}

func (s *SeriesService) getClient() *sling.Sling {
	return s.sling.New().Path("/series")
}

// NewSeriesService initialize a new SeriesService
func newSeriesService(sling *sling.Sling) *SeriesService {
	return &SeriesService{
		sling: sling,
	}
}

// Get one TV Serie by ID
func (s *SeriesService) Get(id string) (*Series, error) {
	client := s.getClient()
	series := new(Series)
	jsonError := new(JSONError)
	path := fmt.Sprintf("/%s", id)

	_, err := client.Path(path).Receive(series, jsonError)
	return series, relevantError(err, *jsonError)
}
