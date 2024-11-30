package quiz

import (
	_ "embed"
	"encoding/json"
)

//go:embed estonian.json
var estonian []byte

type db struct {
	words []word
}

type word struct {
	Estonian Estonian `json:"et"`
	Russian  []string `json:"ru"`
	English  []string `json:"en"`
}

type Estonian struct {
	FirstForm  string `json:"1st"`
	SecondForm string `json:"2nd"`
	ThirdForm  string `json:"3rd"`
}

func load() (*db, error) {
	var words []word
	err := json.Unmarshal(estonian, &words)
	if err != nil {
		return nil, err
	}

	return &db{words}, nil
}
