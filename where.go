package goq

import (
	"bytes"
	"strings"
)

type where struct {
	logicalOperator string
	fieldComparison string
	value           interface{}
}

// String is a string representation of the where struct
// and option with logical operator (AND|OR) or not
func (w *where) String(withLogicalOperator bool) string {
	var where bytes.Buffer

	if withLogicalOperator {
		where.WriteString(strings.ToUpper(w.logicalOperator) + " ")
	}

	where.WriteString(w.fieldComparison)

	return where.String()
}
