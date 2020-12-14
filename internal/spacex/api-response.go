package spacex

type LaunchPads []LaunchPadElement

type LaunchPadElement struct {
	Fairings           *Fairings     `json:"fairings"`
	Links              Links         `json:"links"`
	StaticFireDateUTC  *string       `json:"static_fire_date_utc"`
	StaticFireDateUnix *int64        `json:"static_fire_date_unix"`
	Tbd                bool          `json:"tbd"`
	Net                bool          `json:"net"`
	Window             *int64        `json:"window"`
	Rocket             Rocket        `json:"rocket"`
	Success            *bool         `json:"success"`
	Details            *string       `json:"details"`
	Crew               []string      `json:"crew"`
	Ships              []string      `json:"ships"`
	Capsules           []string      `json:"capsules"`
	Payloads           []string      `json:"payloads"`
	Launchpad          Launchpad     `json:"launchpad"`
	AutoUpdate         bool          `json:"auto_update"`
	Failures           []Failure     `json:"failures"`
	FlightNumber       int64         `json:"flight_number"`
	Name               string        `json:"name"`
	DateUTC            string        `json:"date_utc"`
	DateUnix           int64         `json:"date_unix"`
	DateLocal          string        `json:"date_local"`
	DatePrecision      DatePrecision `json:"date_precision"`
	Upcoming           bool          `json:"upcoming"`
	Cores              []Core        `json:"cores"`
	ID                 string        `json:"id"`
}

type Core struct {
	Core           *string      `json:"core"`
	Flight         *int64       `json:"flight"`
	Gridfins       *bool        `json:"gridfins"`
	Legs           *bool        `json:"legs"`
	Reused         *bool        `json:"reused"`
	LandingAttempt *bool        `json:"landing_attempt"`
	LandingSuccess *bool        `json:"landing_success"`
	LandingType    *LandingType `json:"landing_type"`
	Landpad        *Landpad     `json:"landpad"`
}

type Failure struct {
	Time     int64  `json:"time"`
	Altitude *int64 `json:"altitude"`
	Reason   string `json:"reason"`
}

type Fairings struct {
	Reused          *bool  `json:"reused"`
	RecoveryAttempt *bool  `json:"recovery_attempt"`
	Recovered       *bool  `json:"recovered"`
	Ships           []Ship `json:"ships"`
}

type Links struct {
	Patch     Patch   `json:"patch"`
	Reddit    Reddit  `json:"reddit"`
	Flickr    Flickr  `json:"flickr"`
	Presskit  *string `json:"presskit"`
	Webcast   *string `json:"webcast"`
	YoutubeID *string `json:"youtube_id"`
	Article   *string `json:"article"`
	Wikipedia *string `json:"wikipedia"`
}

type Flickr struct {
	Small    []interface{} `json:"small"`
	Original []string      `json:"original"`
}

type Patch struct {
	Small *string `json:"small"`
	Large *string `json:"large"`
}

type Reddit struct {
	Campaign *string `json:"campaign"`
	Launch   *string `json:"launch"`
	Media    *string `json:"media"`
	Recovery *string `json:"recovery"`
}

type LandingType string

type Landpad string

type DatePrecision string

const (
	Day   DatePrecision = "day"
	Hour  DatePrecision = "hour"
	Month DatePrecision = "month"
)

type Ship string

type Launchpad string

type Rocket string
