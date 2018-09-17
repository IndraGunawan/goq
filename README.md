# goq

[![Build Status](https://travis-ci.org/IndraGunawan/goq.svg?branch=master)](https://travis-ci.org/IndraGunawan/goq)

Super simple query builder

This is super simple query builder library to generate SQL string by using OOP styles.

## Setup
- Go to `$GOPATH/src/github.com/IndraGunawan`
- Clone this repository `git clone git@github.com:IndraGunawan/goq.git`
- Install [`dep`](https://golang.github.io/dep/)
- Run `dep ensure` to fetch the dependency

## Test
Run `go test ./...`

## Usage
```go
import "github.com/IndraGunawan/goq"

// Basic Select
sql := goq.From("mytable").ToSQL() // SELECT * FROM mytable
sql = goq.Select("id").From("mytable") // SELECT id FROM mytable

// Complete
sql := goq.
    From("mytable").
    Where("name = ?", "myname").
    OrWhere("name = ?", "yourname").
    OrderBy("id", "desc")

sql.ToSQL() // SELECT * FROM Mytable WHERE name = ? OR name = ? ORDER BY ID DESC
sql.GetBindingParameters() // []interface{"myname", "yourname"}
```

## License

This library is available as open source under the terms of the [MIT License](LICENSE).
