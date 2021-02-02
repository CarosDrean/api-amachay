package query

import "github.com/CarosDrean/api-amachay/models"

var lot = models.TableDB{
	Name:   "dbo.Lot",
	Fields: []string{"Id", "Lot", "DueDate"},
}

var Lot = models.QueryDB{
	"get":    {Q: "select " + fieldString(lot.Fields) + " from " + lot.Name + " where " + lot.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(lot.Fields) + " from " + lot.Name + ";"},
	"insert": {Q: "insert into " + lot.Name + " (" + fieldStringInsert(lot.Fields) + ") values (" + valuesString(lot.Fields) + ");"},
	"update": {Q: "update " + lot.Name + " set " + updatesString(lot.Fields) + " where " + lot.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + lot.Name + " where " + lot.Fields[0] + " = @ID"},
}

