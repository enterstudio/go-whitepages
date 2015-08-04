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
	Type     string `json:"type"`
	ValidFor `json:"valid_for"`
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
