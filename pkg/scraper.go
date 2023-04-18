package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

// GetMainInfo gets general company info
func GetMainInfo(inn string) (*CompanyInfo, error) {
	url := fmt.Sprintf(
		address+"/"+query,
		inn)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var companies struct {
		Items []CompanyInfo `json:"ul"`
	}

	if err := json.Unmarshal(b, &companies); err != nil {
		return nil, err
	}

	var comp *CompanyInfo
	for _, el := range companies.Items {
		el.INN = strings.Trim(el.INN, "!~")
		if el.INN == inn {
			comp = &el
			break
		}
	}

	if comp == nil {
		return nil, errors.New("no companies with provided INN")
	}

	return comp, nil
}

func GetCompanyKPP(company *CompanyInfo) (string, error) {
	url := address + company.URL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(resp.Body)

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	domElements := html.Find("#clip_kpp")

	if domElements.Length() == 0 {
		return "", errors.New("KPP was not found")
	}

	return domElements.Nodes[0].FirstChild.Data, nil
}
