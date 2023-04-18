package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)


// get general company info
func GetMainInfo(inn string) (*CompanyInfo, error) {
	url := fmt.Sprintf(
		address + query,
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

	var data struct {
		Items []CompanyInfo `json:"ul"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	var comp *CompanyInfo
	for _, el := range data.Items {
		el.INN = strings.Trim(el.INN, "!~")
		if el.INN == inn {
			comp = &el
			break
		}
	}

	if comp == nil {
		return nil, errors.New("no companies with provided INN")
	}


	fmt.Println(comp)

	return comp, nil
}


