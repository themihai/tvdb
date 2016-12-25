package tvdb

import (
	"github.com/aliem/tvdb/swagger"
	"github.com/dghubble/sling"
)

// SearchService provides an interface to the search engine
type SearchService struct {
	sling *sling.Sling
}

// NewSearchService returns a new SearchService
func NewSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("/search/series"),
	}
}

// SearchParams all optional search parameters
type SearchParams struct {
	Name     string `url:"name,omitempty"`
	ImdbID   string `url:"imdbId,omitempty"`
	Zap2itID string `url:"zap2itId,omitempty"`
}

// Search search by SearchParams
func (s *SearchService) Search(params *SearchParams) (*swagger.SeriesSearchData, error) {
	series := new(swagger.SeriesData)
	errors := new(swagger.SeriesData)
	return s.sling.New().Get("/").QueryStruct(params).Receive(series, errors)
}

// ByName Search series by name
func (s *SearchService) ByName(name string) (*swagger.SeriesSearchData, error) {
	params := &SearchParams{Name: name}
	return s.Search(params)
}

// ByImdbID Search by IMDB id
func (s *SearchService) ByImdbID(id string) (*swagger.SeriesSearchData, error) {
	params := &SearchParams{ImdbID: id}
	return s.Search(params)
}

// ByZap2itID Search by Zap2it id
func (s *SearchService) ByZap2itID(id string) (*swagger.SeriesSearchData, error) {
	params := &SearchParams{Zap2itID: id}
	return s.Search(params)
}
