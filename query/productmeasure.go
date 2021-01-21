package query

import "github.com/CarosDrean/api-amachay/models"

var productM = models.TableDB{
	Name:   "dbo.CATEGORY",
	Fields: []string{"Id", "IdProduct", "IdMeasure", "Unity"},
}

var ProductMeasure = models.QueryDB{
	"get":    {Q: "select " + fieldString(productM.Fields) + " from " + productM.Name + " where " + productM.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(productM.Fields) + " from " + productM.Name + ";"},
	"insert": {Q: "insert into " + productM.Name + " (" + fieldStringInsert(productM.Fields) + ") values (" + valuesString(productM.Fields) + ");"},
	"update": {Q: "update " + productM.Name + " set " + updatesString(productM.Fields) + " where " + productM.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + productM.Name + " where " + productM.Fields[0] + " = @ID"},
}
