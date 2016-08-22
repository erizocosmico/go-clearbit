package clearbit

type Company struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	LegalName     string   `json:"legalName"`
	Domain        string   `json:"domain"`
	DomainAliases []string `json:"domainAliases"`
	URL           string   `json:"url"`
	Site          struct {
		URL             string   `json:"url"`
		Title           string   `json:"title"`
		H1              string   `json:"h1"`
		MetaDescription string   `json:"metaDescription"`
		MetaAuthor      string   `json:"metaAuthor"`
		PhoneNumbers    []string `json:"phoneNumbers"`
		EmailAddresses  []string `json:"emailAddresses"`
	} `json:"site"`
	Category struct {
		Sector        string `json:"sector"`
		IndustryGroup string `json:"industryGroup"`
		Industry      string `json:"industry"`
		SubIndustry   string `json:"subIndustry"`
	} `json:"category"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	FoundedYear int64    `json:"foundedYear"`
	Location    string   `json:"location"`
	TimeZone    string   `json:"timeZone"`
	UTCOffset   int64    `json:"utcOffset"`
	Logo        string   `json:"logo"`
	Geo         struct {
		StreetNumber string  `json:"streetNumber"`
		SubPremise   string  `json:"subPremise"`
		City         string  `json:"city"`
		PostalCode   string  `json:"postalCode"`
		State        string  `json:"state"`
		StateCode    string  `json:"stateCode"`
		Country      string  `json:"country"`
		CountryCode  string  `json:"countryCode"`
		Lat          float64 `json:"lat"`
		Lng          float64 `json:"lng"`
	} `json:"geo"`
	Metrics struct {
		Raised          int64 `json:"raised"`
		AlexaUSRank     int64 `json:"alexaUsRank"`
		AlexaGlobalRank int64 `json:"alexaglobalRank"`
		Employees       int64 `json:"emloyees"`
		MarketCap       int64 `json:"marketCap"`
		AnnualRevenue   int64 `json:"annualRevenue"`
	} `json:"metrics"`
	Facebook struct {
		Handle string `json:"handle"`
	} `json:"facebook"`
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
	AngelList struct {
		Handle    string `json:"handle"`
		Bio       string `json:"bio"`
		Blog      string `json:"blog"`
		Followers string `json:"followers"`
		Site      string `json:"site"`
		Avatar    string `json:"avatar"`
	} `json:"angellist"`
	LinkedIn struct {
		Handle string `json:"handle"`
	} `json:"linkedin"`
	CrunchBase struct {
		Handle string `json:"handle"`
	} `json:"crunchbase"`
	EmailProvider bool        `json:"emailProvider"`
	Type          CompanyType `json:"type"`
	Ticker        string      `json:"ticker"`
	Phone         string      `json:"phone"`
	Tech          []string    `json:"tech"`
}

type CompanyType string

const (
	Education  CompanyType = "education"
	Government CompanyType = "government"
	NonProfit  CompanyType = "nonprofit"
	Private    CompanyType = "private"
	Public     CompanyType = "public"
	Personal   CompanyType = "personal"
)
