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

var csv_rows []CSVRow
