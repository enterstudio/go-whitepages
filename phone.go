package whitepages

type PhoneResponse struct {
	Results []Result `json:"results"`
	// Messages []Message `json:"messages"`
}

type ID struct {
	Key        string `json:"key"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	UUID       string `json:"uuid"`
	Durability string `json:"durability"`
}

type Result struct {
	LineType            string               `json:"line_type"`
	BelongsTos          []BelongsTo          `json:"belongs_to"`
	AssociatedLocations []AssociatedLocation `json:"associated_locations"`
	IsConnected         bool                 `json:"is_connected"`
	IsValid             bool                 `json:"is_valid"`
	PhoneNumber         string               `json:"phone_number"`
	CountryCallingCode  string               `json:"country_calling_code"`
	Extension           string               `json:"extension"`
	Carrier             string               `json:"carrier"`
	DoNotCall           bool                 `json:"do_not_call"`
	IsPrepaid           bool                 `json:"is_prepaid"`
}

type BelongsTo struct {
	ID struct {
		Key        string `json:"key"`
		Url        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	} `json:"id"`
	Type         string       `json:"type"`
	Names        []Name       `json:"names"`
	Name         string       `json:"name"`
	AgeRange     AgeRange     `json:"age_range"`
	Gender       string       `json:"gender"`
	Locations    []Location   `json:"locations"`
	Phones       []Phone      `json:"phones"`
	BestName     string       `json:"best_name"`
	BestLocation BestLocation `json:"best_location"`
	ValidFor     ValidFor     `json:"valid_for"`
	IsHistorical bool         `json:"is_historical"`
}

type Name struct {
	Salutation string   `json:"salutation"`
	FirstName  string   `json:"first_name"`
	MiddleName string   `json:"middle_name"`
	LastName   string   `json:"last_name"`
	Suffix     string   `json:"suffix"`
	ValidFor   ValidFor `json:"valid_for"`
}

type AgeRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Location struct {
	ID struct {
		Key        string `json:"key"`
		URL        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	} `json:"id"`
	Type                    string   `json:"type"`
	ValidFor                ValidFor `json:"valid_for"`
	LatLong                 LatLong  `json:"lat_long"`
	LegalEntitiesAt         string   `json:"legal_entities_at"`
	City                    string   `json:"city"`
	PostalCode              string   `json:"postal_code"`
	Zip4                    string   `json:"zip4"`
	StateCode               string   `json:"state_code"`
	CountryCode             string   `json:"country_code"`
	Address                 string   `json:"address"`
	House                   string   `json:"house"`
	StreetName              string   `json:"street_name"`
	StreetType              string   `json:"street_type"`
	PreDir                  string   `json:"pre_dir"`
	PostDir                 string   `json:"post_dir"`
	AptNumber               string   `json:"apt_number"`
	AptType                 string   `json:"apt_type"`
	BoxNumber               string   `json:"box_number"`
	IsReceivingMail         bool     `json:"is_receiving_mail"`
	NotReceivingMailReason  string   `json:"not_receiving_mail_reason"`
	Usage                   string   `json:"usage"`
	DeliveryPoint           string   `json:"delivery_point"`
	BoxType                 string   `json:"box_type"`
	AddressType             string   `json:"address_type"`
	IsDeliverable           bool     `json:"is_deliverable"`
	StandardAddressLine1    string   `json:"standard_address_line1"`
	StandardAddressLine2    string   `json:"standard_address_line2"`
	StandardAddressLocation string   `json:"standard_address_location"`
	IsHistorical            bool     `json:"is_historical"`
	ContactType             string   `json:"contact_type"`
	ContactCreationDate     int64    `json:"contact_creation_date"`
}

type ValidFor struct {
	Start Date `json:"start"`
	Stop  Date `json:"stop"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Accuracy  string  `json:"accuracy"`
}

type Phone struct {
	ID struct {
		Key        string `json:"key"`
		URL        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	} `json:"id"`
	LineType            string       `json:"line_type"`
	BelongsTo           string       `json:"belongs_to"`
	AssociatedLocations string       `json:"associated_locations"`
	IsValid             bool         `json:"is_valid"`
	PhoneNumber         string       `json:"phone_number"`
	CountryCallingCode  string       `json:"country_calling_code"`
	Extension           string       `json:"extension"`
	Carrier             string       `json:"carrier"`
	DoNotCall           bool         `json:"do_not_call"`
	Reputation          Reputation   `json:"reputation"`
	IsPrepaid           bool         `json:"is_prepaid"`
	IsConnected         bool         `json:"is_connected"`
	BestLocation        BestLocation `json:"best_location"`
	ValidFor            ValidFor     `json:"valid_for"`
	ContactType         string       `json:"contact_type"`
	ContactCreationDate int64        `json:"contact_creation_date"`
}

type Reputation struct {
	SpamScore int `json:"spam_score"`
	SpamIndex int `json:"spam_index"`
	Level     int `json:"level"`
	Details   []struct {
		Score    int    `json:"score"`
		Type     string `json:"type"`
		Category string `json:"category"`
	} `json:"details"`
}

type BestLocation struct {
	ID struct {
		Key        string `json:"key"`
		URL        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	} `json:"id"`
	Type                    string   `json:"type"`
	ValidFor                ValidFor `json:"valid_for"`
	LegalEntitiesAt         string   `json:"legal_entities_at"`
	City                    string   `json:"city"`
	PostalCode              string   `json:"postal_code"`
	Zip4                    string   `json:"zip4"`
	StateCode               string   `json:"state_code"`
	CountryCode             string   `json:"country_code"`
	Address                 string   `json:"address"`
	House                   string   `json:"house"`
	StreetName              string   `json:"street_name"`
	StreetType              string   `json:"street_type"`
	PreDir                  string   `json:"pre_dir"`
	PostDir                 string   `json:"post_dir"`
	AptNumber               string   `json:"apt_number"`
	AptType                 string   `json:"apt_type"`
	BoxNumber               string   `json:"box_number"`
	IsReceivingMail         bool     `json:"is_receiving_mail"`
	NotReceivingMailReason  string   `json:"not_receiving_mail_reason"`
	Usage                   string   `json:"usage"`
	DeliveryPoint           string   `json:"delivery_point"`
	BoxType                 string   `json:"box_type"`
	AddressType             string   `json:"address_type"`
	LatLong                 LatLong  `json:"lat_long"`
	IsDeliverable           bool     `json:"is_deliverable"`
	StandardAddressLine1    string   `json:"standard_address_line1"`
	StandardAddressLine2    string   `json:"standard_address_line2"`
	StandardAddressLocation string   `json:"standard_address_location"`
}

type AssociatedLocation struct {
	ID struct {
		Key        string `json:"key"`
		URL        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	} `json:"id"`
	Type                    string   `json:"type"`
	ValidFor                ValidFor `json:"valid_for"`
	LegalEntitiesAt         string   `json:"legal_entities_at"`
	City                    string   `json:"city"`
	PostalCode              string   `json:"postal_code"`
	Zip4                    string   `json:"zip4"`
	StateCode               string   `json:"state_code"`
	CountryCode             string   `json:"country_code"`
	Address                 string   `json:"address"`
	House                   string   `json:"house"`
	StreetName              string   `json:"street_name"`
	StreetType              string   `json:"street_type"`
	PreDir                  string   `json:"pre_dir"`
	PostDir                 string   `json:"post_dir"`
	AptNumber               string   `json:"apt_number"`
	AptType                 string   `json:"apt_type"`
	BoxNumber               string   `json:"box_number"`
	IsReceivingMail         bool     `json:"is_receiving_mail"`
	NotReceivingMailReason  string   `json:"not_receiving_mail_reason"`
	Usage                   string   `json:"usage"`
	DeliveryPoint           string   `json:"delivery_point"`
	BoxType                 string   `json:"box_type"`
	AddressType             string   `json:"address_type"`
	LatLong                 LatLong  `json:"lat_long"`
	IsDeliverable           bool     `json:"is_deliverable"`
	StandardAddressLine1    string   `json:"standard_address_line1"`
	StandardAddressLine2    string   `json:"standard_address_line2"`
	StandardAddressLocation string   `json:"standard_address_location"`
	IsHistorical            bool     `json:"is_historical"`
	ContactType             string   `json:"contact_type"`
	ContactCreationDate     int64    `json:"contact_creation_date"`
}
