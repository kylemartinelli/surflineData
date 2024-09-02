package types

type Main struct {
	Associated  Associated  `json:"associated"`
	Units       Units       `json:"units"`
	Spot        Spot        `json:"spot"`
	Forecast    Forecast    `json:"forecast"`
	Report      Report      `json:"report"`
	Live        Live        `json:"live"`
	Permissions Permissions `json:"permissions"`
}

type Associated struct {
	LocalPhotosURL  string      `json:"localPhotosUrl"`
	ChartsURL       string      `json:"chartsUrl"`
	BeachesURL      string      `json:"beachesUrl"`
	SubregionURL    string      `json:"subregionUrl"`
	Href            string      `json:"href"`
	Units           Units       `json:"units"`
	Advertising     Advertising `json:"advertising"`
	Analytics       Analytics   `json:"analytics"`
	WeatherIconPath string      `json:"weatherIconPath"`
	WindStation     WindStation `json:"windStation"`
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

type WindStation struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Provider string   `json:"provider"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
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
	SortableCondition float64  `json:"sortableCondition"`
	Human             bool   `json:"human"`
	Expired           bool   `json:"expired"`
}

type Swell struct {
	Height       float64 `json:"height"`
	Period       int64   `json:"period"`
	Direction    float64 `json:"direction"`
	DirectionMin float64 `json:"directionMin"`
	Event        int64   `json:"event"`
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
	Speed         float64 `json:"speed"`
	Direction     float64   `json:"direction"`
	DirectionType string  `json:"directionType"`
	Gust          float64 `json:"gust"`
}

type Live struct {
	Wind Wind `json:"wind"`
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

type Report struct {
	Timestamp  int64      `json:"timestamp"`
	Forecaster Forecaster `json:"forecaster"`
	Body       string     `json:"body"`
	Expired    bool       `json:"expired"`
}

type Forecaster struct {
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
	Title   string `json:"title"`
}

type Spot struct {
	ID                   string        `json:"_id"`
	Name                 string        `json:"name"`
	Breadcrumb           []Breadcrumb  `json:"breadcrumb"`
	Lat                  float64       `json:"lat"`
	Lon                  float64       `json:"lon"`
	Cameras              []Camera      `json:"cameras"`
	Subregion            Subregion     `json:"subregion"`
	AbilityLevels        []string      `json:"abilityLevels"`
	BoardTypes           []string      `json:"boardTypes"`
	MetaDescription      string        `json:"metaDescription"`
	TitleTag             string        `json:"titleTag"`
	TravelDetails        TravelDetails `json:"travelDetails"`
	HasLiveWind          bool          `json:"hasLiveWind"`
	LegacyID             int64         `json:"legacyId"`
	LegacyRegionID       int64         `json:"legacyRegionId"`
	LineupEnabled        bool          `json:"lineupEnabled"`
	InsightsCameraID     interface{}   `json:"insightsCameraId"`
	ConsistencyEnabled   bool          `json:"consistencyEnabled"`
	SpectraDataAvailable bool          `json:"spectraDataAvailable"`
}

type Breadcrumb struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type Camera struct {
	ID                              string      `json:"_id"`
	Control                         string      `json:"control"`
	Title                           string      `json:"title"`
	StreamURL                       string      `json:"streamUrl"`
	StillURL                        string      `json:"stillUrl"`
	PixelatedStillURL               string      `json:"pixelatedStillUrl"`
	RewindBaseURL                   string      `json:"rewindBaseUrl"`
	IsPremium                       bool        `json:"isPremium"`
	IsPrerecorded                   bool        `json:"isPrerecorded"`
	LastPrerecordedClipStartTimeUTC string      `json:"lastPrerecordedClipStartTimeUTC"`
	LastPrerecordedClipEndTimeUTC   string      `json:"lastPrerecordedClipEndTimeUTC"`
	Alias                           string      `json:"alias"`
	SupportsHighlights              bool        `json:"supportsHighlights"`
	SupportsInsights                bool        `json:"supportsInsights"`
	SupportsSmartRewinds            bool        `json:"supportsSmartRewinds"`
	SupportsCrowds                  bool        `json:"supportsCrowds"`
	Status                          Status      `json:"status"`
	Highlights                      interface{} `json:"highlights"`
	Nighttime                       bool        `json:"nighttime"`
	RewindClip                      string      `json:"rewindClip"`
}

type Status struct {
	IsDown     bool   `json:"isDown"`
	Message    string `json:"message"`
	SubMessage string `json:"subMessage"`
	AltMessage string `json:"altMessage"`
}

type Subregion struct {
	ID             string `json:"_id"`
	ForecastStatus string `json:"forecastStatus"`
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
	TravelID         interface{}   `json:"travelId"`
	Access           string        `json:"access"`
	BreakType        []string      `json:"breakType"`
	Description      string        `json:"description"`
	Hazards          string        `json:"hazards"`
	RelatedArticleID string        `json:"relatedArticleId"`
	Status           string        `json:"status"`
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
