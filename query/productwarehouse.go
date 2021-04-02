package query

import "github.com/CarosDrean/api-amachay/models"

var productWarehouse = models.TableDB{ // usa bien los nombre para que no te confundas despues
	Name:   "dbo.PRODUCT_WAREHOUSE",
	Fields: []string{"Id", "IdProduct", "IdWarehouse", "Ignore"},
}

var ProductWarehouse = models.QueryDB{
	"insert": {Q: "insert into " + productWarehouse.Name + " (" + fieldStringInsert(productWarehouse.Fields) + ") values (" + valuesString(productWarehouse.Fields) + ");"},
	"update": {Q: "update " + productWarehouse.Name + " set " + updatesString(productWarehouse.Fields) + " where " + productWarehouse.Fields[0] + " = @ID;"},
	"get":    {Q: "select " + fieldString(productWarehouse.Fields) + " from " + productWarehouse.Name + " where " + productWarehouse.Fields[1] + " = %s " +
		"and " + productWarehouse.Fields[2] + " = %s;"},
	"updateIgnore": {Q: "update " + productWarehouse.Name + " set " + productWarehouse.Fields[3] + " = 0  where " + productWarehouse.Fields[0] + " = @ID;"},


}
