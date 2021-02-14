package query_sql

import "github.com/CarosDrean/api-amachay/models"

var warehouse = models.TableDB{
	Name:   "dbo.WAREHOUSE",
	Fields: []string{"Id", "Name", "Address", "State"},
}

var Warehouse = models.QueryDB{
	"get":    {Q: "select " + fieldString(warehouse.Fields) + " from " + warehouse.Name + " where " + warehouse.Fields[0] + " =%s;"},
	"list":   {Q: "select " + fieldString(warehouse.Fields) + " from " + warehouse.Name + ";"},
	"insert": {Q: "insert into " + warehouse.Name + "(" + fieldStringInsert(warehouse.Fields) + ") values (" + valuesString(warehouse.Fields) + ");"},
	"update": {Q: "update " + warehouse.Name + " set " + updatesString(warehouse.Fields) + " where " + warehouse.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + warehouse.Name + " where " + warehouse.Fields[0] + " = @ID"},
}

