package sqlr

// Scanner is an abstraction of a type which can scan his values onto
// variables passed in dest. May be used for sql.Row and sql.Rows
type Scanner interface {
	Scan(dest ...interface{}) error
}
