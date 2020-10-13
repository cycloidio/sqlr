# SQLR

Defines a default set of Interfaces to work with the default [database/sql](https://golang.org/pkg/database/sql/).


## Querier

The [sql.DB](https://golang.org/pkg/database/sql/#DB) and the [sql.Tx](https://golang.org/pkg/database/sql/#Tx) have the same signature when quering
so for that we have created the `sqlr.Querier` which is fulfilled for either `sql.DB` and `sql.Tx` so in some contexts you don't need to have the
specific implementation this interface can be used

## Scanner

Used to abstract the logic of Scanning an `sql.Row` and `sql.Rows` reusable:

```go
type Scanner interface {
  Scan(dest ...interface{}) error
}
```

Example usage:

```go
func scanUsers(rows *sql.Rows) ([]*user.User, error) {
	users := make([]*user.User, 0, 0)
	defer rows.Close()
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func scanUser(s sqlr.Scanner) (*user.User, error) {
	var u user.User

	err := s.Scan(
		&u.Name,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}
```
