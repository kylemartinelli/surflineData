package types


type Main struct {
	Associated  Associated  `json:"associated"`
	Units       Units       `json:"units"`
	Spot        Spot        `json:"spot"`
	Forecast    Forecast    `json:"forecast"`
	Report      interface{} `json:"report"`
	Live        Live        `json:"live"`
	Permissions Permissions `json:"permissions"`
}

type Associated struct {
	LocalPhotosURL  interface{} `json:"localPhotosUrl"`
	ChartsURL       string      `json:"chartsUrl"`
	BeachesURL      string      `json:"beachesUrl"`
	SubregionURL    string      `json:"subregionUrl"`
	Href            string      `json:"href"`
	Units           Units       `json:"units"`
	Advertising     Advertising `json:"advertising"`
	Analytics       Analytics   `json:"analytics"`
	WeatherIconPath string      `json:"weatherIconPath"`
	WindStation     interface{} `json:"windStation"`
	UTCOffset       int64       `json:"utcOffset"`
	AbbrTimezone    string      `json:"abbrTimezone"`
	Timezone        string      `json:"timezone"`
}

type Advertising struct {
	SpotID      string `json:"spotId"`
	SubregionID string `json:"subregionId"`
	Sst         string `json:"sst"`
}

type Analytics struct {
	SpotID      string `json:"spotId"`
	SubregionID string `json:"subregionId"`
}

type Units struct {
	Temperature string `json:"temperature"`
	TideHeight  string `json:"tideHeight"`
	SwellHeight string `json:"swellHeight"`
	WaveHeight  string `json:"waveHeight"`
	WindSpeed   string `json:"windSpeed"`
	Pressure    string `json:"pressure"`
}

type Forecast struct {
	Note       interface{} `json:"note"`
	Conditions Conditions  `json:"conditions"`
	Wind       Wind        `json:"wind"`
	WaveHeight WaveHeight  `json:"waveHeight"`
	Tide       Tide        `json:"tide"`
	WaterTemp  WaterTemp   `json:"waterTemp"`
	Wetsuit    Wetsuit     `json:"wetsuit"`
	Weather    Weather     `json:"weather"`
	Swells     []Swell     `json:"swells"`
}

type Conditions struct {
	Value             string `json:"value"`
	SortableCondition int64  `json:"sortableCondition"`
	Human             bool   `json:"human"`
	Expired           bool   `json:"expired"`
}

type Swell struct {
	Height       float64 `json:"height"`
	Period       int64   `json:"period"`
	Direction    float64 `json:"direction"`
	DirectionMin float64 `json:"directionMin"`
	Event        *int64  `json:"event"`
}

type Tide struct {
	Previous Current `json:"previous"`
	Current  Current `json:"current"`
	Next     Current `json:"next"`
}

type Current struct {
	Type      string  `json:"type"`
	Height    float64 `json:"height"`
	Timestamp int64   `json:"timestamp"`
	UTCOffset int64   `json:"utcOffset"`
}

type WaterTemp struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

type WaveHeight struct {
	Min                   int64       `json:"min"`
	Max                   int64       `json:"max"`
	Plus                  bool        `json:"plus"`
	Occasional            interface{} `json:"occasional"`
	HumanRelation         string      `json:"humanRelation"`
	Human                 bool        `json:"human"`
	LastCameraObservation interface{} `json:"lastCameraObservation"`
}

type Weather struct {
	Temperature int64  `json:"temperature"`
	Condition   string `json:"condition"`
}

type Wetsuit struct {
	Thickness string `json:"thickness"`
	Type      string `json:"type"`
}

type Wind struct {
	Speed         int64   `json:"speed"`
	Direction     float64 `json:"direction"`
	DirectionType string  `json:"directionType"`
	Gust          int64   `json:"gust"`
}

type Live struct {
	Wind interface{} `json:"wind"`
}

type Permissions struct {
	Data       []interface{} `json:"data"`
	Violations []Violation   `json:"violations"`
}

type Violation struct {
	Code       int64      `json:"code"`
	Message    string     `json:"message"`
	Permission Permission `json:"permission"`
}

type Permission struct {
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

type Spot struct {
	ID                   string        `json:"_id"`
	Name                 string        `json:"name"`
	Breadcrumb           []Breadcrumb  `json:"breadcrumb"`
	Lat                  float64       `json:"lat"`
	Lon                  float64       `json:"lon"`
	Cameras              []interface{} `json:"cameras"`
	Subregion            Subregion     `json:"subregion"`
	AbilityLevels        []interface{} `json:"abilityLevels"`
	BoardTypes           []interface{} `json:"boardTypes"`
	TravelDetails        TravelDetails `json:"travelDetails"`
	HasLiveWind          bool          `json:"hasLiveWind"`
	LineupEnabled        bool          `json:"lineupEnabled"`
	InsightsCameraID     interface{}   `json:"insightsCameraId"`
	ConsistencyEnabled   bool          `json:"consistencyEnabled"`
	SpectraDataAvailable bool          `json:"spectraDataAvailable"`
}

type Breadcrumb struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type Subregion struct {
	ID             string `json:"_id"`
	ForecastStatus string `json:"forecastStatus"`
}

type TravelDetails struct {
}
