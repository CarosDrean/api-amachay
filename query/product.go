package query

import "github.com/CarosDrean/api-amachay/models"

var product = models.TableDB{
	Name:   "dbo.PRODUCT",
	Fields: []string{"Id", "IdCategory", "Name", "Description", "Price", "Stock"},
}

var Product = models.QueryDB{
	"get":    {Q: "select " + fieldString(product.Fields) + " from " + product.Name + " where " + product.Fields[0] + " =%s;"},
	"list":   {Q: "select " + fieldString(product.Fields) + " from " + product.Name + ";"},
	"insert": {Q: "insert into " + product.Name + "(" + fieldStringInsert(product.Fields) + ") values (" + valuesString(product.Fields) + ");"},
	"update": {Q: "update " + product.Name + " set " + updatesString(product.Fields) + " where " + product.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + product.Name + " where " + product.Fields[0] + " = @ID"},
}
