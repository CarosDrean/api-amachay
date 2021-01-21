package query

import "github.com/CarosDrean/api-amachay/models"

var provider = models.TableDB{
	Name:   "dbo.PROVIDER",
	Fields: []string{"Id", "IdBusiness", "Type"},
}

var Provider = models.QueryDB{
	"get":    {Q: "select " + fieldString(provider.Fields) + " from " + provider.Name + " where " + provider.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(provider.Fields) + " from " + provider.Name + ";"},
	"insert": {Q: "insert into " + provider.Name + " (" + fieldStringInsert(provider.Fields) + ") values (" + valuesString(provider.Fields) + ");"},
	"update": {Q: "update " + provider.Name + " set " + updatesString(provider.Fields) + " where " + provider.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + provider.Name + " where " + provider.Fields[0] + " = @ID"},
}
