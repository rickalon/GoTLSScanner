package data

type URI interface {
}

type MockURL struct {
	Data []string `json:"urls"`
}

type URL struct {
	Data []UrlObj
}

type UrlObj struct {
	Url     string   `json:"url"`
	Result  string   `json:"resutl"`
	To      string   `json:"to"`
	From    string   `json:"From"`
	Country []string `json:"country"`
	ExpDate string   `json:"expDate"`
	EmiDate string   `json:"emiDate"`
	Alg     string   `json:"alg"`
	DNS     []string `json:"dns"`
	IsCA    bool     `json:"isCA"`
}

func NewURL(data ...[]string) *MockURL {
	dft := &MockURL{}
	if len(data) != 0 {
		dft.Data = data[0]
	}
	return dft
}
