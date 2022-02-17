package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/madjiebimaa/isyarah/models"
	"github.com/madjiebimaa/isyarah/requests"
)

type locationService struct {
	locationRepository models.LocationRepository
	contextTimeout     time.Duration
}

func NewLocationService(locationRepository models.LocationRepository, contextTimeout time.Duration) models.LocationService {
	return &locationService{
		locationRepository,
		contextTimeout,
	}
}

func (l *locationService) Create(c context.Context, req *requests.LocationCreate) (models.Location, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()

	key := os.Getenv("OPENCAGEDATA_KEY")
	loc := strings.ReplaceAll(req.Location, " ", "%20")
	url := fmt.Sprintf("https://api.opencagedata.com/geocode/v1/json?key=%s&q=%s&pretty=1&no_annotations=1", key, loc)
	resp, _ := http.Get(url)

	var query models.Query
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &query)

	location := query.GetMostConfidenceLocation()
	location.Description = req.Description
	if err := l.locationRepository.Create(ctx, &location); err != nil {
		return models.Location{}, err
	}

	return location, nil
}
