package types

type Main struct {
	Spot        Spot        `json:"spot"`
}

type Spot struct {
	ID                   string        `json:"_id"`
	Name                 string        `json:"name"`
	Breadcrumb           []Breadcrumb  `json:"breadcrumb"`
	Lat                  float64       `json:"lat"`
	Lon                  float64       `json:"lon"`
	TravelDetails        TravelDetails `json:"travelDetails"`
}

type Breadcrumb struct {
	Name string `json:"name"`
	Href string `json:"href"`
}



type TravelDetails struct {
	AbilityLevels    AbilityLevels `json:"abilityLevels"`
	Best             Best          `json:"best"`
	Bottom           Bottom        `json:"bottom"`
	CrowdFactor      CrowdFactor   `json:"crowdFactor"`
	LocalVibe        CrowdFactor   `json:"localVibe"`
	ShoulderBurn     CrowdFactor   `json:"shoulderBurn"`
	SpotRating       CrowdFactor   `json:"spotRating"`
	WaterQuality     CrowdFactor   `json:"waterQuality"`
	BreakType        []string      `json:"breakType"`
	Description      string        `json:"description"`
	Hazards          string        `json:"hazards"`
	BoardTypes       []string      `json:"boardTypes"`
}

type AbilityLevels struct {
	Description string   `json:"description"`
	Levels      []string `json:"levels"`
	Summary     string   `json:"summary"`
}

type Best struct {
	Season         Bottom `json:"season"`
	Tide           Bottom `json:"tide"`
	Size           Bottom `json:"size"`
	WindDirection  Bottom `json:"windDirection"`
	SwellDirection Bottom `json:"swellDirection"`
}

type Bottom struct {
	Description string   `json:"description"`
	Value       []string `json:"value"`
}

type CrowdFactor struct {
	Description string `json:"description"`
	Rating      int64  `json:"rating"`
	Summary     string `json:"summary"`
}
