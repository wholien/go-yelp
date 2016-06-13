package yelp

// PhoneOptions provide search options for phone search. Phone number is
// required, while country code and category are optional. The latter two
// enhance the quality of the query.
type PhoneOptions struct {
	Phone    string // Parameter that specifies the business phone number to search for. Outside of the US and Canada, include the international dialing code (e.g. +442079460000) or use the 'cc' parameter
	Cc       string // ISO 3166-1 alpha-2 country code. Default country to us when parsing the phone number. United States = US, Canada = CA, United Kingdom = GB (not UK).
	Category string // Category to filter search results with. See the list of supported categories. The category filter can be a list of comma delimited categories. For example, 'bars,french' will filter by Bars and French. The category identifier should be used (for example 'discgolf', not 'Disc Golf').
}

// getParameters will reflect over the values of the given
// struct, and provide a type appropriate set of querystring parameters
func (o *PhoneOptions) getParameters() (params map[string]string, err error) {
	params = make(map[string]string)
	if o.Phone != "" {
		params["phone"] = o.Phone
	} else {
		return params, errUnspecifiedPhone
	}
	if o.Cc != "" {
		params["cc"] = o.Cc
	}
	if o.Category != "" {
		params["category"] = o.Category
	}
	return params, nil
}
