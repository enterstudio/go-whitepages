package whitepages

// Request Fields
type Request struct {
	APIKey string `json:"-"`
	*RequestAddress

	// Name is the complete name of a person or business
	Name string `json:"name,omitempty"`

	FirstName    string `json:"firstname,omitempty"`
	LastName     string `json:"lastname,omitempty"`
	EmailAddress string `json:"email_address",omitempty` // Lead Verify supports all standard email addresses
	IPAddress    string `json:"ip_address",omitempty`    // IPv4 IP Addresses.
	Phone        string `json:"phone",omitempty`         // Contains a raw unparsed or a formatted phone number
}

// RequestAddress is the request address - the struct inside of the response looks different than the one it shows online as the actual request object.
type RequestAddress struct {
	PostalCode  string `json:"address.postal_code",omitempty`
	StateCode   string `json:"address.state",omitempty`
	StreetLine1 string `json:"address.street_line_1",omitempty`
	StreetLine2 string `json:"address.street_line_2",omitempty`
	City        string `json:"address_city",omitempty`
}

// LeadVerifyResponse is the full response that the Lead Verify API returns
// If any of the checks are left out of the request (query), those check structures will be nil.
type LeadVerifyResponse struct {
	AddressChecks      *AddressChecks      `json:"address_checks"`
	EmailAddressChecks *EmailAddressChecks `json:"email_address_checks"`
	IPAddressChecks    *IPAddressChecks    `json:"ip_address_checks"`
	NameChecks         *NameChecks         `json:"name_checks"`
	PhoneChecks        *PhoneChecks        `json:"phone_checks"`
	Errors             []*error            `json:"errors"`
	Request            Request             `json:"request"`
}

// AddressChecks indicates whether the address is real and active and verifies if the resident matches
// the input name provided on the lead. If lead verify + append is used, response also provides
// certain address metadata attributes and resident’s demographic attributes that can further help
// prioritize the leads.
type AddressChecks struct {
	// A 1-4 score on whether this is a valid address and resident matches to the input name provided for the lead.
	// Score of 1 indicates a valid address with the resident name matching to the input name provided with the lead.
	// 4 indicates very high confidence that the lead cannot be reached via this address.
	AddressContactScore int `json:"address_contact_score"`

	// Verification result whether the resident name for the address matches the input name. Possible values:
	// Possible values are: Match, No Match, null
	AddressToName *string `json:"address_to_name"`

	// Indicates if the address is currently receiving mail. Possible values are true, false, or null.
	IsActive *bool `json:"is_active"`

	// A boolean value indicating if the address is a valid existing address.
	// Possible values are true, false, and null.
	IsValid      *bool  `json:"is_valid"`
	ResidentName string `json:"resident_name"`

	// Resident’s age in a 5 year range, e.g. 30-34.
	ResidentAgeRange string `json:"resident_age_range,omitempty"`

	// Resident’s gender, either “Male” or “Female”.
	ResidentGender string `json:"resident_gender,omitempty"`

	/*
		Indicates delivery point for the address.
		Possible values:
			Commercial mail drop
			Multi unit
			Single unit
			PO Box
			PO Box Throwback
			Unknown address type
	*/
	Type string `json:"type,omitempty"`

	// Indicates if the address is a business address. Possible values are true, false, or null.
	IsCommercial *bool `json:"is_commercial,omitempty"`

	Error    error    `json:"error"`
	Warnings []string `json:"warnings"`
}

// EmailAddressChecks indicates whether the email is valid or malformed, active or inactive, and verifies if the email registered name matches the input name provided.
// No data is appended by lead verify append instead of lead verify.
type EmailAddressChecks struct {
	/*
		diagnostics: [
			"Syntax OK, domain exists, and mailbox does not reject mail"
		]

		Diagnostic message for the is_valid flag. This is the reason why we call the email valid or invalid.
		Valid messages:
			Domain does not support validation (accepts all mailboxes)
			Syntax OK and domain valid according to the domain database
			Syntax OK, domain exists, and mailbox does not reject mail
		Invalid messages:
			The mailbox is invalid or the username does not exist at the domain
			Address does not have an @ sign
			Addresses with that domain are not allowed
			Addresses with that username are not allowed
			Domain cannot receive email
			Domain does not exist
			Domain does not have a valid IP address
			Invalid domain syntax
			Invalid top-level-domain (TLD) in address
			Invalid username syntax
			Invalid username syntax for that domain
			Mail is not accepted for this domain
			Mailbox is full and can not receive email at this time
	*/
	Diagnostics []string `json:"diagnostics"`

	// EmailContactScore is a 1-4 score on whether this is a valid email address and matches to the input name provided for the lead.
	// Score of 1 indicates a valid email with the registered name matching to the input name provided with the lead.
	// 4 indicates very high confidence that the lead cannot be reached via this email.
	EmailContactScore int `json:"email_contact_score"`

	// Verification result whether the registered name for the email matches the input name.
	// Possible values: Match, No match, No name found, null
	EmailToName *string `json:"email_to_name"`

	// Indicates whether the email is potentially auto generated. Possible values are True, False, or null.
	IsAutoGenerated *bool `json:"is_autogenerated"`

	// Indicates whether the email domain is disposable. Possible values are True, False, or null.
	IsDisposable *bool `json:"is_disposable"`

	// IsValid is a warning message that returns “Email Address Invalid” or null
	IsValid *bool `json:"is_valid"`

	// Returns the name that we have on record for the supplied email
	RegisteredName string `json:"registered_name"`
	Error          error  `json:"error"`

	// A warning message that returns “Email Address Invalid” or null
	Warnings []string `json:"warnings"`
}

// IPAddressChecks is a structure that gives info about the validity of the IP address in the query
// No IP address given in query will return a null response.
type IPAddressChecks struct {
	DistanceFromAddress float32 `json:"distance_from_address"`

	// GeoLocation is the location of the IP address. Includes postal_code, city_name, country_name, continent_code, and country_code when available.
	// 	geolocation: {
	// 		postal_code: "29205",
	// 		city_name: "Columbia",
	// 		country_name: "United States",
	// 		continent_code: "NA",
	// 		country_code: "US"
	// },
	GeoLocation *GeoLocation `json:"geolocation"`

	// IsProxy indicates whether the IP address is a known proxy. Possible values are true, false, or null.
	IsProxy *bool `json:"is_proxy"`
	IsValid *bool `json:"is_valid"`
	Error   error `json:"error"`

	// Warnings are a warning message that returns “Invalid Input” or null.
	Warnings []string `json:"warnings"`
}

// NameChecks is a structure that gives info about the validity of the name in the query
// No name given will return a null response.
type NameChecks struct {
	// CelebrityName indicates if the input name matches any known celebrity names. Possible values are True and False.
	CelebrityName *bool `json:"celebrity_name"`

	// FakeName indicates if the input name seems to be fake. Possible values are True and False.
	FakeName *bool `json:"fake_name"`

	// A warning message that returns “Missing Input” or null
	Warnings []string `json:"warnings"`
	Error    error    `json:"error"`
}

// Address is a generic whitepages address struct
type Address struct {
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	State       string `json:"state"`
	StateCode   string `json:"state_code"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// PhoneChecks indicate whether the phone matches the input name provided on the lead, and whether
// the phone number is valid or not, and is currently in service at time of inquiry.
// If lead verify + append is used, the response also provides certain phone metadata attributes and
// subscriber’s demographic attributes that help prioritize the lead and manage TCPA compliance.
type PhoneChecks struct {
	// IsConnected indicates whether the phone is connected or not in service.
	// Possible values are True, False, or null.
	IsConnected *bool `json:"is_connected"`

	// A boolean value indicating if the phone is a valid phone number.
	// Possible values are true, false or null.
	IsValid *bool `json:"is_valid"`

	// PhoneContactScore is a 1-4 score on whether this is a valid phone number and matches to the input name provided for the lead.
	// Score of 1 indicates a valid, connected number with the subscriber name matching to the input name provided with the lead.
	// 4 indicates very high confidence that the lead cannot be reached via this phone number
	PhoneContactScore int `json:"phone_contact_score"`

	// PhoneToName is a verification result whether the subscriber name for the phone matches the input name.
	// Possible values: Match, No Match, null
	PhoneToName *string `json:"phone_to_name"`

	// SubscriberName is the name of the input phone
	SubscriberName string `json:"subscriber_name"`
	Error          error  `json:"error"`

	// A warning message that returns “Missing Input” or null
	Warnings []string `json:"warnings"`

	// SubscriberAgeRange is the subscriber’s age in a 5 year range, e.g. 30-34.
	SubscriberAgeRange string `json:"subscriber_age_range,omitempty"`

	// Full address of the subscriber. Includes House Number, Street, City, State, Postal and Country.
	SubscriberAddress *Address `json:"subscriber_address,omitempty"`

	// SubscriberGender is the subscribers gender
	SubscriberGender string `json:"subscriber_gender"`

	// Country code for the input phone
	CountryCode string `json:"country_code,omitempty"`

	/*
		Line type for the input phone.
		Possible values:
			Mobile
			Landline
			Fixed VOIP
			Non-fixed VOIP
			Premium
			Tollfree
			Voicemail
			Other
			Unknown
	*/
	LineType string `json:"line_type,omitempty"`

	// The company that provides voice and/or data services for this phone number. Carriers are returned at the MVNO level.
	Carrier string `json:"carrier,omitempty"`

	// Indicates if the phone is associated with a prepaid account. Possible values are true, false, or null.
	IsPrepaid *bool `json:"is_prepaid,omitempty"`

	// Indicates if the phone is on the National Do Not Call registry. Possible values are true, false, or null.
	IsDoNotCallRegistered *bool `json:"is_do_not_call_registered"`

	// Indicates if the phone is registered to a business. Possible values are true, false, or null.
	IsCommercial *bool `json:"is_commercial"`
}

// GeoLocation is just simple address info like city, state, country and continent code
type GeoLocation struct {
	CityName      string `json:"city_name"`
	ContinentCode string `json:"continent_code"`
	CountryCode   string `json:"country_code"`
	CountryName   string `json:"country_name"`
	PostalCode    string `json:"postal_code"`
}
