package entry

type Table struct {
	FullName string
	Schema   string
	Name     string
	GoName   string
	Columns  []*Column
}