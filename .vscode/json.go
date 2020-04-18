package main

// TimelineJSON represents user timeline
type TimelineJSON struct {
	Locations []Location
}

// ActivityWithConfidenceJSON from Google API
type ActivityWithConfidenceJSON struct {
	Type       string
	Confidence int
}

// ActivityJSON from Google API
type ActivityJSON struct {
	TimestampMs                int64
	ActivityWithConfidenceJSON []ActivityWithConfidenceJSON
}

// Location from Google API
type Location struct {
	Timestamp   int64
	LatitudeE7  int64
	LongitudeE7 int64
	Accuracy    int
	Activity    []ActivityJSON
}

type MapMarkerJSON struct {
	Name string
	Data string
	Lat  float64
	Lng  float64
}

type InfectedPlaceJSON struct {
	MapMarker MapMarkerJSON
	Intensity int
}

type InfectedPlacesJSON struct {
	Places []InfectedPlaceJSON
}