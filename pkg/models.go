package pkg

type CompanyInfo struct {
	INN  string `json:"inn"`
	Name string `json:"raw_name"`
	FIO  string `json:"ceo_name"`
	URL  string `json:"url"`
}

type FullCompanyInfo struct {
	defaultInfo *CompanyInfo
	LPP         string
}

const (
	address = "https://www.rusprofile.ru"
	query   = "ajax.php?query=%s&action=search"
)
