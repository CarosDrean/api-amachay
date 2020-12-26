package models
// este documnto es para lo s model que no son tablas de BD

type queryConfig struct {
	Name string
	Q    string
}

type TableDB struct {
	Name   string
	Fields []string
}

type QueryDB map[string]*queryConfig
