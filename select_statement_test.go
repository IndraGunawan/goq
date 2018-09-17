package goq_test

import (
	"testing"

	"github.com/IndraGunawan/goq"
	"github.com/stretchr/testify/assert"
)

func TestBasicSelect(t *testing.T) {
	builder := goq.From("mytable")
	assert.Equal(t, "SELECT * FROM mytable", builder.ToSQL())

	builder = goq.Select("id", "name").From("mytable")
	assert.Equal(t, "SELECT id, name FROM mytable", builder.ToSQL())

	assert.Len(t, builder.GetBindingParameters(), 0)
}

func TestSelectDistinct(t *testing.T) {
	builder := goq.From("mytable").Distinct(true)
	assert.Equal(t, "SELECT DISTINCT * FROM mytable", builder.ToSQL())
}

func TestSelectWithWhereCondition(t *testing.T) {
	builder := goq.From("mytable").Where("id = ?", 1)
	assert.Equal(t, "SELECT * FROM mytable WHERE id = ?", builder.ToSQL())

	builder = goq.From("mytable").AndWhere("id = ?", 1)
	assert.Equal(t, "SELECT * FROM mytable WHERE id = ?", builder.ToSQL())

	builder = goq.From("mytable").Where("id = ?", 1).Where("name = ?", "myname")
	assert.Equal(t, "SELECT * FROM mytable WHERE name = ?", builder.ToSQL())

	builder = goq.From("mytable").Where("id = ?", 1).AndWhere("name = ?", "myname")
	assert.Equal(t, "SELECT * FROM mytable WHERE id = ? AND name = ?", builder.ToSQL())

	builder = goq.From("mytable").Where("id = ?", 1).OrWhere("name = ?", "myname")
	assert.Equal(t, "SELECT * FROM mytable WHERE id = ? OR name = ?", builder.ToSQL())

	assert.Len(t, builder.GetBindingParameters(), 2)
}

func TestSelectWithGroupBy(t *testing.T) {
	builder := goq.From("mytable").GroupBy("name")
	assert.Equal(t, "SELECT * FROM mytable GROUP BY name", builder.ToSQL())

	builder = goq.From("mytable").GroupBy("name").GroupBy("last_name")
	assert.Equal(t, "SELECT * FROM mytable GROUP BY name, last_name", builder.ToSQL())
}

func TestSelectWithOrderBy(t *testing.T) {
	builder := goq.Select("*").From("mytable").OrderBy("id", "ASC")
	assert.Equal(t, "SELECT * FROM mytable ORDER BY id ASC", builder.ToSQL())

	builder = goq.Select("*").From("mytable").OrderBy("id", "ASC").OrderBy("name", "")
	assert.Equal(t, "SELECT * FROM mytable ORDER BY id ASC, name", builder.ToSQL())
}
