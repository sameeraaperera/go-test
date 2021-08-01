package datasource

type DataSource interface {
	Value(key string) (string, error)
}
