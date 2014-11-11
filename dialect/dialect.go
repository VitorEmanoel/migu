package dialect

type Dialect interface {
	ColumnType(name string, size uint64, autoIncrement bool) (typ string, null, autoIncrementable bool)
	Quote(s string) string
	AutoIncrement() string
}
