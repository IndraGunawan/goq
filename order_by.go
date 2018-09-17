package goq

import (
	"strings"
)

type orderBy struct {
	field string
	sort  string
}

// String is a string representation of the orderBy struct
func (o *orderBy) String() string {
	order := o.field
	if o.sort != "" {
		order += " " + strings.ToUpper(o.sort)
	}

	return order
}
