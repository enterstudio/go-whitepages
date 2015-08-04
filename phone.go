package whitepages

import (
	"encoding/json"
	"time"
)

type V2PhoneResponse struct {
	Results
}

type Results []struct {
	ID struct {
		Key        string `json:"key"`
		URL        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	}
	LineType           string `json:"line_type"`
	BelongsTo          `json:"belongs_to"`
	IsConnected        bool   `json:"is_connected"`
	IsValid            bool   `json:"is_valid"`
	PhoneNumber        string `json:"phone_number"`
	CountryCallingCode string `json:"country_calling_code"`
	Exension           string `json:"extension"`
	Carrier            string `json:"carrier"`
	DoNotCall          bool   `json:"do_not_call"`
	IsPrepaid          bool   `json:"is_prepaid"`
}

type BelongsTo []struct {
	ID struct {
		Key        string `json:"key"`
		Url        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	} `json:"id"`
	Type      string `json:"type"`
	Names     `json:"names"`
	AgeRange  `json:"age_range"`
	Gender    string `json:"gender"`
	Locations `json:"locations"`
	Phones    `json:"phones"`
}

type Names []struct {
	Salutation string `json:"salutation"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Suffix     string `json:"suffix"`
	ValidFor   string `json:"valid_for"`
}

type AgeRange []struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Locations []struct {
	ID struct {
		Key        string `json:"key"`
		URL        string `json:"url"`
		Type       string `json:"type"`
		UUID       string `json:"uuid"`
		Durability string `json:"durability"`
	} `json:"id"`
	Type                    string `json:"type"`
	ValidFor                `json:"valid_for"`
	LatLong                 `json:"lat_long"`
	LegalEntitiesAt         string `json:"legal_entities_at"`
	City                    string `json:"city"`
	PostalCode              string `json:"postal_code"`
	Zip4                    string `json:"zip4"`
	StateCode               string `json:"state_code"`
	CountryCode             string `json:"country_code"`
	Address                 string `json:"address"`
	House                   string `json:"house"`
	StreetName              string `json:"street_name"`
	StreetType              string `json:"street_type"`
	PreDir                  string `json:"pre_dir"`
	PostDir                 string `json:"post_dir"`
	AptNumber               string `json:"apt_number"`
	AptType                 string `json:"apt_type"`
	BoxNumber               string `json:"box_number"`
	IsReceivingMail         bool   `json:"is_receiving_mail"`
	NotReceivingMailReason  string `json:"not_receiving_mail_reason"`
	Usage                   string `json:"usage"`
	DeliveryPoint           string `json:"delivery_point"`
	BoxType                 string `json:"box_type"`
	AddressType             string `json:"address_type"`
	IsDeliverable           bool   `json:"is_deliverable"`
	StandardAddressLine1    string `json:"standard_address_line1"`
	StandardAddressLine2    string `json:"standard_address_line2"`
	StandardAddressLocation string `json:"standard_address_location"`
	IsHistorical            bool   `json:"is_historical"`
	ContactType             string `json:"contact_type"`
	ContactCreationDate     int    `json:"contact_creation_date"`
}

type ValidFor struct {
	Start struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"start"`
	Stop struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"stop"`
}

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Accuracy  string  `json:"accuracy"`
}

type Phones []struct {
	LineType string `json:"line_type"`
}

func (c *V2Client) Phone(params map[string]string, timeout time.Duration) (error, V2PhoneResponse) {
	p := V2PhoneResponse{}
	err, response := c.request("phone.json", timeout, params)
	if err != nil {
		return err, p
	}
	if err = json.Unmarshal(response, &p); err != nil {
		return err, p
	}
	return nil, p

}
