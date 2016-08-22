package clearbit

type Person struct {
	ID   string `json:"id"`
	Name struct {
		Given  string `json:"givenName"`
		Family string `json:"familyName"`
		Full   string `json:"fullName"`
	} `json:"name"`
	Gender    Gender `json:"gender"`
	Location  string `json:"location"`
	TimeZone  string `json:"timeZone"`
	UTCOffset int64  `json:"utcOffset"`
	Geo       struct {
		City        string  `json:"city"`
		State       string  `json:"state"`
		StateCode   string  `json:"stateCode"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
	} `json:"geo"`
	Bio        string `json:"bio"`
	Site       string `json:"site"`
	Avatar     string `json:"avatar"`
	Employment struct {
		Domain    string `json:"domain"`
		Name      string `json:"name"`
		Title     string `json:"title"`
		Role      string `json:"role"`
		Seniority string `json:"seniority"`
	} `json:"employment`
	Facebook struct {
		Handle string `json:"handle"`
	} `json:"facebook"`
	GitHub struct {
		Handle    string `json:"handle"`
		Avatar    string `json:"avatar"`
		Company   string `json:"company"`
		Blog      string `json:"blog"`
		Following int64  `json:"following"`
		Followers int64  `json:"followers"`
	} `json:"github"`
	Twitter struct {
		Handle    string `json:"handle"`
		ID        string `json:"id"`
		Bio       string `json:"bio"`
		Followers int64  `json:"followers"`
		Following int64  `json:"following"`
		Location  string `json:"location"`
		Site      string `json:"site"`
		Avatar    string `json:"avatar"`
	} `json:"twitter"`
	LinkedIn struct {
		Handle string `json:"handle"`
	} `json:"linkedin"`
	GooglePlus struct {
		Handle string `json:"handle"`
	} `json:"googleplus"`
	AngelList struct {
		Handle    string `json:"handle"`
		Bio       string `json:"bio"`
		Blog      string `json:"blog"`
		Followers string `json:"followers"`
		Site      string `json:"site"`
		Avatar    string `json:"avatar"`
	} `json:"angellist"`
	AboutMe struct {
		Handle string `json:"handle"`
		Bio    string `json:"bio"`
		Avatar string `json:"avatar"`
	} `json:"aboutme"`
	Gravatar struct {
		Handle  string           `json:"handle"`
		URLs    []GravatarURL    `json:"urls"`
		Avatar  string           `json:"avatar"`
		Avatars []GravatarAvatar `json:"avatars"`
	} `json:"gravatar"`
	Fuzzy bool `json:"fuzzy"`
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = ""
)

type GravatarURL struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type GravatarAvatar struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}
