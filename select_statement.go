package goq

import (
	"bytes"
	"strings"
)

// SelectInterface provides all needed method for select statement
type SelectInterface interface {
	Select(fields ...string) SelectInterface
	Distinct(distinct bool) SelectInterface
	From(table string) SelectInterface
	Where(fieldComparison string, value interface{}) SelectInterface
	AndWhere(fieldComparison string, value interface{}) SelectInterface
	OrWhere(fieldComparison string, value interface{}) SelectInterface
	GroupBy(field string) SelectInterface
	OrderBy(field string, sort string) SelectInterface
	GetBindingParameters() []interface{}
	ToSQL() string
}

// selectParts to store all part of select statement
type selectParts struct {
	fromTable string
	fields    []string
	distinct  bool
	where     []*where
	groupBy   []string
	orderBy   []*orderBy
	limit     int
	offset    int
}

// From initializes builder directly to table name
func From(table string) SelectInterface {
	return &selectParts{
		fromTable: table,
		distinct:  false,
	}
}

// Select initializes builder with the selected fields
func Select(fields ...string) SelectInterface {
	sp := &selectParts{distinct: false}

	return sp.Select(fields...)
}

// Select does append fields to selectParts struct
func (s *selectParts) Select(fields ...string) SelectInterface {
	s.fields = append(s.fields, fields...)

	return s
}

// Distinct is choice to add Distinct clause or not
func (s *selectParts) Distinct(distinct bool) SelectInterface {
	s.distinct = distinct

	return s
}

// From does set table to select
func (s *selectParts) From(table string) SelectInterface {
	s.fromTable = table

	return s
}

// Where does override the Where clause to clause
func (s *selectParts) Where(fieldComparison string, value interface{}) SelectInterface {
	s.where = []*where{&where{
		fieldComparison: fieldComparison,
		value:           value,
	}}

	return s
}

// AndWhere does add AND to where clause
func (s *selectParts) AndWhere(fieldComparison string, value interface{}) SelectInterface {
	s.where = append(s.where, &where{
		fieldComparison: fieldComparison,
		value:           value,
		logicalOperator: "AND",
	})

	return s
}

// OrWhere does add OR to where clause
func (s *selectParts) OrWhere(fieldComparison string, value interface{}) SelectInterface {
	s.where = append(s.where, &where{
		fieldComparison: fieldComparison,
		value:           value,
		logicalOperator: "OR",
	})

	return s
}

// GroupBy does append to GROUP BY clause
func (s *selectParts) GroupBy(field string) SelectInterface {
	s.groupBy = append(s.groupBy, field)

	return s
}

// OrderBy does append to GROUP BY clause
func (s *selectParts) OrderBy(field string, sort string) SelectInterface {
	s.orderBy = append(s.orderBy, &orderBy{field: field, sort: sort})

	return s
}

// GetBindingParameters return all values that set from Where clause
func (s *selectParts) GetBindingParameters() []interface{} {
	var parameters []interface{}
	for _, where := range s.where {
		parameters = append(parameters, where.value)
	}

	return parameters
}

// ToSQL return the SQL string result
func (s *selectParts) ToSQL() string {
	var sql bytes.Buffer
	sql.WriteString("SELECT ")

	if s.distinct {
		sql.WriteString("DISTINCT ")
	}

	if len(s.fields) == 0 {
		sql.WriteString("*")
	} else {
		sql.WriteString(strings.Join(s.fields, ", "))
	}

	if s.fromTable != "" {
		sql.WriteString(" FROM ")
		sql.WriteString(s.fromTable)
	}

	if len(s.where) > 0 {
		sql.WriteString(" WHERE")
		for i, where := range s.where {
			sql.WriteString(" ")
			if i == 0 {
				sql.WriteString(where.String(false))
			} else {
				sql.WriteString(where.String(true))
			}
		}
	}

	if len(s.groupBy) > 0 {
		sql.WriteString(" GROUP BY ")
		sql.WriteString(strings.Join(s.groupBy, ", "))
	}

	if len(s.orderBy) > 0 {
		var orderBy []string
		for _, order := range s.orderBy {
			orderBy = append(orderBy, order.String())
		}

		sql.WriteString(" ORDER BY ")
		sql.WriteString(strings.Join(orderBy, ", "))
	}

	return sql.String()
}
