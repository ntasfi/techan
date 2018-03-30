package talib4g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rsTestMockSeries = mockTimeSeries(
	"44.34",
	"44.09",
	"44.15",
	"43.61",
	"44.33",
	"44.83",
	"45.10",
	"45.42",
	"45.84",
	"46.08",
	"45.89",
	"46.03",
	"45.61",
	"46.28",
	"46.28",
)

func TestRelativeStrengthIndexIndicator(t *testing.T) {
	closeIndicator := NewClosePriceIndicator(rsTestMockSeries)
	indicator := NewRelativeStrengthIndexIndicator(closeIndicator, 14)

	assert.EqualValues(t, "70.46", indicator.Calculate(14).FormattedString(2))
}

func TestRelativeStrengthIndicator(t *testing.T) {
	closeIndicator := NewClosePriceIndicator(rsTestMockSeries)
	indicator := NewRelativeStrengthIndicator(closeIndicator, 14)

	assert.EqualValues(t, "2.39", indicator.Calculate(14).FormattedString(2))
}