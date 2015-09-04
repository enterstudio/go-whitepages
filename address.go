package whitepages

type AddressResponse struct {
	Messages []interface{} `json:"messages"`
	Results  []struct {
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
		LegalEntitiesAt []struct {
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
				PreDir                  interface{} `json:"pre_dir"`
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
				AptNumber           string      `json:"apt_number"`
				AptType             string      `json:"apt_type"`
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
				PreDir                  interface{} `json:"pre_dir"`
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
					Stop struct {
						Day   int `json:"day"`
						Month int `json:"month"`
						Year  int `json:"year"`
					} `json:"stop"`
				} `json:"valid_for"`
				Zip4 string `json:"zip4"`
			} `json:"locations"`
			Names []struct {
				FirstName  string      `json:"first_name"`
				LastName   string      `json:"last_name"`
				MiddleName string      `json:"middle_name"`
				Salutation interface{} `json:"salutation"`
				Suffix     string      `json:"suffix"`
				ValidFor   interface{} `json:"valid_for"`
			} `json:"names"`
			Phones   []interface{} `json:"phones"`
			Type     string        `json:"type"`
			ValidFor struct {
				Start struct {
					Day   int `json:"day"`
					Month int `json:"month"`
					Year  int `json:"year"`
				} `json:"start"`
				Stop interface{} `json:"stop"`
			} `json:"valid_for"`
		} `json:"legal_entities_at"`
		NotReceivingMailReason  interface{} `json:"not_receiving_mail_reason"`
		PostDir                 interface{} `json:"post_dir"`
		PostalCode              string      `json:"postal_code"`
		PreDir                  interface{} `json:"pre_dir"`
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
	} `json:"results"`
}
