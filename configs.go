package main

//Change the names of the struct property with the names of each column of your csv file found on the first row
type CSVRow struct {
	PolicyID           string `json:"policyID"`
	Statecode          string `json:"statecode"`
	County             string `json:"county"`
	Eq_site_limit      string `json:"eq_site_limit"`
	Hu_site_limit      string `json:"hu_site_limit"`
	Fl_site_limit      string `json:"fl_site_limit"`
	Fr_site_limit      string `json:"fr_site_limit"`
	Tiv_2011           string `json:"tiv_2011"`
	Tiv_2012           string `json:"tiv_2012"`
	Eq_site_deductible string `json:"eq_site_deductible"`
	Hu_site_deductible string `json:"hu_site_deductible"`
	Fl_site_deductible string `json:"fl_site_deductible"`
	Fr_site_deductible string `json:"fr_site_deductible"`
	Point_latitude     string `json:"point_latitude"`
	Point_longitude    string `json:"point_longitude"`
	Line               string `json:"line"`
	Construction       string `json:"construction"`
	Point_granularity  string `json:"point_granularity"`
}

type CSVRow1 struct {
	TemplateId                                                    string `json:"templateId"`
	VerificationToken                                             string `json:"verificationToken"`
	Email                                                         string `json:"email"`
	Credential__referenceId                                       string `json:"credential__referenceId"`
	Credential__expires                                           string `json:"credential__expires"`
	Credential__claim__type                                       string `json:"credential__claim__type"`
	Credential__claim__givenName                                  string `json:"credential__claim__givenName"`
	Credential__claim__familyName                                 string `json:"credential__claim__familyName"`
	Credential__claim__additionalName                             string `json:"credential__claim__additionalName"`
	Credential__claim__achievement__type                          string `json:"credential__claim__achievement__type"`
	Credential__claim__achievement__dateEffective                 string `json:"credential__claim__achievement__dateEffective"`
	Credential__claim__achievement__certification__type           string `json:"credential__claim__achievement__certification__type"`
	Credential__claim__achievement__certification__credentialType string `json:"credential__claim__achievement__certification__credentialType"`
	Credential__claim__achievement__certification__name           string `json:"credential__claim__achievement__certification__name"`
	Credential__claim__achievement__certification__description    string `json:"credential__claim__achievement__certification__description"`
	Credential__claim__achievement__certification__acronym        string `json:"credential__claim__achievement__certification__acronym"`
	Credential__claim__achievement__certification__criteria       string `json:"credential__claim__achievement__certification__criteria"`
	Credential__claim__achievement__certification__image          string `json:"credential__claim__achievement__certification__image"`
}

var json_objs []string

/*templateId	verificationToken	email	credential__referenceId	credential__expires	credential__claim__type	credential__claim__givenName	credential__claim__familyName	credential__claim__additionalName	credential__claim__achievement__type	credential__claim__achievement__dateEffective	credential__claim__achievement__certification__type	credential__claim__achievement__certification__credentialType	credential__claim__achievement__certification__name	credential__claim__achievement__certification__description	credential__claim__achievement__certification__acronym	credential__claim__achievement__certification__criteria	credential__claim__achievement__certification__image
 */
