package tvdb

import (
	"fmt"
	"github.com/dghubble/sling"
)

// EpisodeRecordData Searches results
type EpisodeRecordData struct {
	Data   Episode   `json:"data,omitempty"`
	Errors JSONError `json:"jsonError,omitempty"`
}

// Episode a single episode
type Episode struct {
	ID                 int32    `json:"id,omitempty"`
	AiredSeason        int32    `json:"airedSeason,omitempty"`
	AiredEpisodeNumber int32    `json:"airedEpisodeNumber,omitempty"`
	EpisodeName        string   `json:"episodeName,omitempty"`
	FirstAired         string   `json:"firstAired,omitempty"`
	GuestStars         []string `json:"guestStars,omitempty"`
	Director           string   `json:"director,omitempty"`
	Directors          []string `json:"directors,omitempty"`
	Writers            []string `json:"writers,omitempty"`
	Overview           string   `json:"overview,omitempty"`
	ProductionCode     string   `json:"productionCode,omitempty"`
	ShowURL            string   `json:"showUrl,omitempty"`
	LastUpdated        int32    `json:"lastUpdated,omitempty"`
	DvdDiscid          string   `json:"dvdDiscid,omitempty"`
	DvdSeason          int32    `json:"dvdSeason,omitempty"`
	DvdEpisodeNumber   float32  `json:"dvdEpisodeNumber,omitempty"`
	DvdChapter         float32  `json:"dvdChapter,omitempty"`
	AbsoluteNumber     int32    `json:"absoluteNumber,omitempty"`
	Filename           string   `json:"filename,omitempty"`
	SeriesID           string   `json:"seriesId,omitempty"`
	LastUpdatedBy      string   `json:"lastUpdatedBy,omitempty"`
	AirsAfterSeason    int32    `json:"airsAfterSeason,omitempty"`
	AirsBeforeSeason   int32    `json:"airsBeforeSeason,omitempty"`
	AirsBeforeEpisode  int32    `json:"airsBeforeEpisode,omitempty"`
	ThumbAuthor        int32    `json:"thumbAuthor,omitempty"`
	ThumbAdded         string   `json:"thumbAdded,omitempty"`
	ThumbWidth         string   `json:"thumbWidth,omitempty"`
	ThumbHeight        string   `json:"thumbHeight,omitempty"`
	ImdbID             string   `json:"imdbId,omitempty"`
	SiteRating         float32  `json:"siteRating,omitempty"`
	SiteRatingCount    int32    `json:"siteRatingCount,omitempty"`
}

// EpisodesService the episode service
type EpisodesService struct {
	sling *sling.Sling
}

func (s *EpisodesService) getClient() *sling.Sling {
	return s.sling.New().Path("/episodes")
}

// newSeriesService initialize a new SeriesService
func newEpisodesService(sling *sling.Sling) *EpisodesService {
	return &EpisodesService{
		sling: sling,
	}
}

// Get a single episode
func (s *EpisodesService) Get(id int32) (*Episode, error) {
	client := s.getClient()
	episode := &Episode{}
	jsonError := &JSONError{}

	path := fmt.Sprintf("/%d", id)
	_, err := client.Path(path).Receive(episode, jsonError)
	return episode, relevantError(err, *jsonError)
}
