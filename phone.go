package whitepages

import (
	"encoding/json"
	"time"
)

type V2PhoneResponse struct {
	Messages []interface{} `json:"messages"`
	Results  []struct {
		AssociatedLocations []struct {
			Address             string      `json:"address"`
			AddressType         interface{} `json:"address_type"`
			AptNumber           interface{} `json:"apt_number"`
			AptType             interface{} `json:"apt_type"`
			BoxNumber           interface{} `json:"box_number"`
			BoxType             interface{} `json:"box_type"`
			City                string      `json:"city"`
			ContactCreationDate interface{} `json:"contact_creation_date"`
			ContactType         interface{} `json:"contact_type"`
			CountryCode         string      `json:"country_code"`
			DeliveryPoint       interface{} `json:"delivery_point"`
			House               interface{} `json:"house"`
			ID                  struct {
				Durability string `json:"durability"`
				Key        string `json:"key"`
				Type       string `json:"type"`
				URL        string `json:"url"`
				UUID       string `json:"uuid"`
			} `json:"id"`
			IsDeliverable   interface{} `json:"is_deliverable"`
			IsHistorical    bool        `json:"is_historical"`
			IsReceivingMail interface{} `json:"is_receiving_mail"`
			LatLong         struct {
				Accuracy  string  `json:"accuracy"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"lat_long"`
			LegalEntitiesAt         interface{} `json:"legal_entities_at"`
			NotReceivingMailReason  interface{} `json:"not_receiving_mail_reason"`
			PostDir                 interface{} `json:"post_dir"`
			PostalCode              string      `json:"postal_code"`
			PreDir                  interface{} `json:"pre_dir"`
			StandardAddressLine1    string      `json:"standard_address_line1"`
			StandardAddressLine2    string      `json:"standard_address_line2"`
			StandardAddressLocation string      `json:"standard_address_location"`
			StateCode               string      `json:"state_code"`
			StreetName              interface{} `json:"street_name"`
			StreetType              interface{} `json:"street_type"`
			Type                    string      `json:"type"`
			Usage                   interface{} `json:"usage"`
			ValidFor                interface{} `json:"valid_for"`
			Zip4                    interface{} `json:"zip4"`
		} `json:"associated_locations"`
		BelongsTo []struct {
			AgeRange struct {
				End   int `json:"end"`
				Start int `json:"start"`
			} `json:"age_range"`
			BestLocation struct {
				Address       string      `json:"address"`
				AddressType   string      `json:"address_type"`
				AptNumber     interface{} `json:"apt_number"`
				AptType       interface{} `json:"apt_type"`
				BoxNumber     interface{} `json:"box_number"`
				BoxType       interface{} `json:"box_type"`
				City          string      `json:"city"`
				CountryCode   string      `json:"country_code"`
				DeliveryPoint string      `json:"delivery_point"`
				House         string      `json:"house"`
				ID            struct {
					Durability string `json:"durability"`
					Key        string `json:"key"`
					Type       string `json:"type"`
					URL        string `json:"url"`
					UUID       string `json:"uuid"`
				} `json:"id"`
				IsDeliverable   bool `json:"is_deliverable"`
				IsReceivingMail bool `json:"is_receiving_mail"`
				LatLong         struct {
					Accuracy  string  `json:"accuracy"`
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"lat_long"`
				LegalEntitiesAt         interface{} `json:"legal_entities_at"`
				NotReceivingMailReason  interface{} `json:"not_receiving_mail_reason"`
				PostDir                 interface{} `json:"post_dir"`
				PostalCode              string      `json:"postal_code"`
				PreDir                  string      `json:"pre_dir"`
				StandardAddressLine1    string      `json:"standard_address_line1"`
				StandardAddressLine2    string      `json:"standard_address_line2"`
				StandardAddressLocation string      `json:"standard_address_location"`
				StateCode               string      `json:"state_code"`
				StreetName              string      `json:"street_name"`
				StreetType              string      `json:"street_type"`
				Type                    string      `json:"type"`
				Usage                   string      `json:"usage"`
				ValidFor                interface{} `json:"valid_for"`
				Zip4                    string      `json:"zip4"`
			} `json:"best_location"`
			BestName string `json:"best_name"`
			Gender   string `json:"gender"`
			ID       struct {
				Durability string `json:"durability"`
				Key        string `json:"key"`
				Type       string `json:"type"`
				URL        string `json:"url"`
				UUID       string `json:"uuid"`
			} `json:"id"`
			IsHistorical bool `json:"is_historical"`
			Locations    []struct {
				Address             string      `json:"address"`
				AddressType         string      `json:"address_type"`
				AptNumber           interface{} `json:"apt_number"`
				AptType             interface{} `json:"apt_type"`
				BoxNumber           interface{} `json:"box_number"`
				BoxType             interface{} `json:"box_type"`
				City                string      `json:"city"`
				ContactCreationDate int         `json:"contact_creation_date"`
				ContactType         string      `json:"contact_type"`
				CountryCode         string      `json:"country_code"`
				DeliveryPoint       string      `json:"delivery_point"`
				House               string      `json:"house"`
				ID                  struct {
					Durability string `json:"durability"`
					Key        string `json:"key"`
					Type       string `json:"type"`
					URL        string `json:"url"`
					UUID       string `json:"uuid"`
				} `json:"id"`
				IsDeliverable   bool `json:"is_deliverable"`
				IsHistorical    bool `json:"is_historical"`
				IsReceivingMail bool `json:"is_receiving_mail"`
				LatLong         struct {
					Accuracy  string  `json:"accuracy"`
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"lat_long"`
				LegalEntitiesAt         interface{} `json:"legal_entities_at"`
				NotReceivingMailReason  interface{} `json:"not_receiving_mail_reason"`
				PostDir                 interface{} `json:"post_dir"`
				PostalCode              string      `json:"postal_code"`
				PreDir                  string      `json:"pre_dir"`
				StandardAddressLine1    string      `json:"standard_address_line1"`
				StandardAddressLine2    string      `json:"standard_address_line2"`
				StandardAddressLocation string      `json:"standard_address_location"`
				StateCode               string      `json:"state_code"`
				StreetName              string      `json:"street_name"`
				StreetType              string      `json:"street_type"`
				Type                    string      `json:"type"`
				Usage                   string      `json:"usage"`
				ValidFor                struct {
					Start struct {
						Day   int `json:"day"`
						Month int `json:"month"`
						Year  int `json:"year"`
					} `json:"start"`
					Stop interface{} `json:"stop"`
				} `json:"valid_for"`
				Zip4 string `json:"zip4"`
			} `json:"locations"`
			Names []struct {
				FirstName  string      `json:"first_name"`
				LastName   string      `json:"last_name"`
				MiddleName string      `json:"middle_name"`
				Salutation interface{} `json:"salutation"`
				Suffix     interface{} `json:"suffix"`
				ValidFor   interface{} `json:"valid_for"`
			} `json:"names"`
			Phones []struct {
				AssociatedLocations interface{} `json:"associated_locations"`
				BelongsTo           interface{} `json:"belongs_to"`
				BestLocation        interface{} `json:"best_location"`
				Carrier             string      `json:"carrier"`
				ContactCreationDate int         `json:"contact_creation_date"`
				ContactType         string      `json:"contact_type"`
				CountryCallingCode  string      `json:"country_calling_code"`
				DoNotCall           bool        `json:"do_not_call"`
				Extension           interface{} `json:"extension"`
				ID                  struct {
					Durability string `json:"durability"`
					Key        string `json:"key"`
					Type       string `json:"type"`
					URL        string `json:"url"`
					UUID       string `json:"uuid"`
				} `json:"id"`
				IsPrepaid   bool   `json:"is_prepaid"`
				IsValid     bool   `json:"is_valid"`
				LineType    string `json:"line_type"`
				PhoneNumber string `json:"phone_number"`
				Reputation  struct {
					SpamIndex int `json:"spam_index"`
					SpamScore int `json:"spam_score"`
				} `json:"reputation"`
				ValidFor interface{} `json:"valid_for"`
			} `json:"phones"`
			Type     string      `json:"type"`
			ValidFor interface{} `json:"valid_for"`
		} `json:"belongs_to"`
		BestLocation struct {
			Address       string      `json:"address"`
			AddressType   interface{} `json:"address_type"`
			AptNumber     interface{} `json:"apt_number"`
			AptType       interface{} `json:"apt_type"`
			BoxNumber     interface{} `json:"box_number"`
			BoxType       interface{} `json:"box_type"`
			City          string      `json:"city"`
			CountryCode   string      `json:"country_code"`
			DeliveryPoint interface{} `json:"delivery_point"`
			House         interface{} `json:"house"`
			ID            struct {
				Durability string `json:"durability"`
				Key        string `json:"key"`
				Type       string `json:"type"`
				URL        string `json:"url"`
				UUID       string `json:"uuid"`
			} `json:"id"`
			IsDeliverable   interface{} `json:"is_deliverable"`
			IsReceivingMail interface{} `json:"is_receiving_mail"`
			LatLong         struct {
				Accuracy  string  `json:"accuracy"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"lat_long"`
			LegalEntitiesAt         interface{} `json:"legal_entities_at"`
			NotReceivingMailReason  interface{} `json:"not_receiving_mail_reason"`
			PostDir                 interface{} `json:"post_dir"`
			PostalCode              string      `json:"postal_code"`
			PreDir                  interface{} `json:"pre_dir"`
			StandardAddressLine1    string      `json:"standard_address_line1"`
			StandardAddressLine2    string      `json:"standard_address_line2"`
			StandardAddressLocation string      `json:"standard_address_location"`
			StateCode               string      `json:"state_code"`
			StreetName              interface{} `json:"street_name"`
			StreetType              interface{} `json:"street_type"`
			Type                    string      `json:"type"`
			Usage                   interface{} `json:"usage"`
			ValidFor                interface{} `json:"valid_for"`
			Zip4                    interface{} `json:"zip4"`
		} `json:"best_location"`
		Carrier            string      `json:"carrier"`
		CountryCallingCode string      `json:"country_calling_code"`
		DoNotCall          bool        `json:"do_not_call"`
		IsConnected        bool        `json:"is_connected"`
		Extension          interface{} `json:"extension"`
		ID                 struct {
			Durability string `json:"durability"`
			Key        string `json:"key"`
			Type       string `json:"type"`
			URL        string `json:"url"`
			UUID       string `json:"uuid"`
		} `json:"id"`
		IsPrepaid   bool   `json:"is_prepaid"`
		IsValid     bool   `json:"is_valid"`
		LineType    string `json:"line_type"`
		PhoneNumber string `json:"phone_number"`
		Reputation  struct {
			SpamIndex int `json:"spam_index"`
			SpamScore int `json:"spam_score"`
		} `json:"reputation"`
	} `json:"results"`
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
