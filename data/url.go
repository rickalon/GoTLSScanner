package data

type URI interface {
}

type URL struct {
	Data []string `json:"urls"`
}

func NewURL(data ...[]string) *URL {
	dft := &URL{}
	if len(data) != 0 {
		dft.Data = data[0]
	}
	return dft
}
