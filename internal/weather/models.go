package weather

// WeatherResponse represents the weather API response
// type WeatherResponse struct {
// 	Location struct {
// 		Name      string  `json:"name"`
// 		Region    string  `json:"region"`
// 		Country   string  `json:"country"`
// 		Lat       float64 `json:"lat"`
// 		Lon       float64 `json:"lon"`
// 		LocalTime string  `json:"localtime"`
// 	} `json:"location"`
// 	Current struct {
// 		TempC     float64 `json:"temp_c"`
// 		TempF     float64 `json:"temp_f"`
// 		Condition struct {
// 			Text string `json:"text"`
// 			Icon string `json:"icon"`
// 		} `json:"condition"`
// 		WindKph   float64 `json:"wind_kph"`
// 		WindDir   string  `json:"wind_dir"`
// 		Humidity  int     `json:"humidity"`
// 		Cloud     int     `json:"cloud"`
// 		FeelsLike float64 `json:"feelslike_c"`
// 		UV        float64 `json:"uv"`
// 	} `json:"current"`
// }


type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

// Location represents the location information
type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	LocalTime string  `json:"localtime"`
}

// Current represents current weather conditions
type Current struct {
	TempC     float64   `json:"temp_c"`
	TempF     float64   `json:"temp_f"`
	Condition Condition `json:"condition"`
	WindKph   float64   `json:"wind_kph"`
	WindDir   string    `json:"wind_dir"`
	Humidity  int       `json:"humidity"`
	Cloud     int       `json:"cloud"`
	FeelsLike float64   `json:"feelslike_c"`
	UV        float64   `json:"uv"`
}

// Condition represents weather condition with icon
type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}
