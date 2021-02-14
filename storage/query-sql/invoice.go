package query_sql

import "github.com/CarosDrean/api-amachay/models"

var invoice = models.TableDB{
	Name:   "dbo.INVOICE",
	Fields: []string{"Id", "Name", "Code", "Date", "IdImage", "IdProvider"},
}

var Invoice = models.QueryDB{
	"get":    {Q: "select " + fieldString(invoice.Fields) + " from " + invoice.Name + " where " + invoice.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(invoice.Fields) + " from " + invoice.Name + ";"},
	"insert": {Q: "insert into " + invoice.Name + " (" + fieldStringInsert(invoice.Fields) + ") values (" + valuesString(invoice.Fields) + ");"},
	"update": {Q: "update " + invoice.Name + " set " + updatesString(invoice.Fields) + " where " + invoice.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + invoice.Name + " where " + invoice.Fields[0] + " = @ID"},
}

