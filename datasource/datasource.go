package datasource

type DataSource interface {
	Value(key string) (string, error)
	Store(key string, value string) error
}
