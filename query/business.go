package query

import "github.com/CarosDrean/api-amachay/models"

var business = models.TableDB{
	Name:   "dbo.BUSINESS",
	Fields: []string{"Id", "Name", "RUC", "Address", "Cel", "Phone", "Mail"},
}

var Business = models.QueryDB{
	"get":    {Q: "select " + fieldString(business.Fields) + " from " + business.Name + " where " + business.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(business.Fields) + " from " + business.Name + ";"},
	"insert": {Q: "insert into " + business.Name + " (" + fieldStringInsert(business.Fields) + ") values (" + valuesString(business.Fields) + ");"},
	"update": {Q: "update " + business.Name + " set " + updatesString(business.Fields) + " where " + business.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + business.Name + " where " + business.Fields[0] + " = @ID"},
}
