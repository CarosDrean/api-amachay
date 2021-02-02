package query

import "github.com/CarosDrean/api-amachay/models"

var brand = models.TableDB{
	Name:   "dbo.BRAND",
	Fields: []string{"Id", "Name"},
}

var Brand = models.QueryDB{
	"get":    {Q: "select " + fieldString(brand.Fields) + " from " + brand.Name + " where " + brand.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(brand.Fields) + " from " + brand.Name + ";"},
	"insert": {Q: "insert into " + brand.Name + " (" + fieldStringInsert(brand.Fields) + ") values (" + valuesString(brand.Fields) + ");"},
	"update": {Q: "update " + brand.Name + " set " + updatesString(brand.Fields) + " where " + brand.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + brand.Name + " where " + brand.Fields[0] + " = @ID"},
}

