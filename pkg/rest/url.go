package rest

import (
	"fmt"
	"strings"
)

type FilterOptions struct {
	Field    string
	Operator string
	Values   []string
}

func (fo *FilterOptions) ToString() string {
	return fmt.Sprintf("%s%s", fo.Operator, strings.Join(fo.Values, ","))
}
