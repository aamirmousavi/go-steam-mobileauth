package mobileauth

import (
	"fmt"

	"github.com/aamirmousavi/htmlreader"
)

func confirmationDescription(html string, confIds [][]string) ([]string, error) {
	node, err := htmlreader.New(html)
	if err != nil {
		return nil, err
	}
	confDescs := make([]string, 0)
	for _, id := range confIds {
		content, ok := node.GetElementById("conf" + id[1])
		if !ok {
			return nil, fmt.Errorf("id is not currect")
		}
		raw, err := htmlreader.ConvReader(content).ToString()
		if err != nil {
			return nil, err
		}
		confDescs = append(confDescs, raw)
	}
	return confDescs, nil
}
