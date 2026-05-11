package gmaps

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGmapJobKeepsGeoCoordinatesForBrowserContext(t *testing.T) {
	job := NewGmapJob("job-1", "en", "turkish restaurant", 1, false, "52.52, 13.405", 14)

	require.Equal(t, "52.52,13.405", job.GeoCoordinates)
	require.Equal(
		t,
		"https://www.google.com/maps/search/turkish+restaurant/@52.52,13.405,14z?hl=en",
		job.GetFullURL(),
	)
}

func TestParseGeoCoordinates(t *testing.T) {
	latitude, longitude, err := parseGeoCoordinates("52.52, 13.405")

	require.NoError(t, err)
	require.Equal(t, 52.52, latitude)
	require.Equal(t, 13.405, longitude)
}

func TestParseGeoCoordinatesRejectsInvalidValues(t *testing.T) {
	_, _, err := parseGeoCoordinates("berlin")

	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid geo coordinates")
}
