package query

import "github.com/CarosDrean/api-amachay/models"

var movement = models.TableDB{
	Name:   "dbo.MOVEMENT",
	Fields: []string{"Id", "IdProduct", "IdWareHouse", "DateTime", "Quantity", "Type", "IdUser", "IdClient"},
}

var Movement = models.QueryDB{
	"get":   {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " + movement.Fields[0] + " =%s;"},
	"list":  {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + ";"},
	"stock": {Q: "select sum(Quantity) as stock from MOVEMENT where IdWareHouse = %d and IdProduct = %d;"},
	"listWarehouseId": {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " +
		movement.Fields[2] + " = %s;"},
	"listWarehouseFilter": {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " +
		movement.Fields[2] + " = %s and " + movement.Fields[5] + " = '%s' " +
		"and CAST(DateTime as date) >= CAST('%s' as date) and CAST(DateTime as date) <= CAST('%s' as date) " +
		"order by Id desc;"},
	"insert": {Q: "insert into " + movement.Name + "(" + fieldStringInsert(movement.Fields) + ") values (" + valuesString(movement.Fields) + ");"},
	"update": {Q: "update " + movement.Name + " set " + updatesString(movement.Fields) + " where " + movement.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + movement.Name + " where " + movement.Fields[0] + " = @ID"},
}
