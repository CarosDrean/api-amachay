package query

import "github.com/CarosDrean/api-amachay/models"

var product = models.TableDB{
	Name:   "dbo.PRODUCT",
	Fields: []string{"Id", "IdCategory", "Name", "Description", "Price", "Stock", "Perishable"},
}

var Product = models.QueryDB{
	"get":    {Q: "select " + fieldString(product.Fields) + " from " + product.Name + " where " + product.Fields[0] + " =%s;"},
	"list":   {Q: "select " + fieldString(product.Fields) + " from " + product.Name + ";"},
	"insert": {Q: "insert into " + product.Name + "(" + fieldStringInsert(product.Fields) + ") values (" + valuesString(product.Fields) + ");"},
	"update": {Q: "update " + product.Name + " set " + updatesString(product.Fields) + " where " + product.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + product.Name + " where " + product.Fields[0] + " = @ID"},
	"getProductWarehouse": {Q: "select p.Name, p.Description, p.Price, p.Stock,p.Perishable, pw.Ignore from PRODUCT p " +
		"inner join PRODUCT_WAREHOUSE pw on p.Id  = pw.IdProduct " +
		"inner join WAREHOUSE w on pw.IdWarehouse = w.Id " +
		"where w.Id = '%s' " +
		"group by p.Name, p.Description, p.Price, p.Stock, p.Perishable, pw.Ignore"},

	"getAllNoIgnore": {Q: "select p.Id, p.IdCategory, p.Name, p.Price, sum(Quantity) as stock, m.Name,  pm.Unity as Unidades, " +
		"p.Description, p.Perishable, pw.Ignore, c.Name, m.Id, pm.Id, pm.MinAlert from PRODUCT p " +
		"left join PRODUCT_WAREHOUSE pw on p.Id  = pw.IdProduct " +
	    "left join PRODUCT_MEASURE pm on p.Id = pm.IdProduct " +
		"inner join MEASURE m on pm.IdMeasure = m.Id " +
		"left join MOVEMENT mo on p.Id = mo.IdProduct " +
		"inner join CATEGORY c on p.IdCategory = c.Id " +
		"where (pw.Ignore = 0 or pw.Ignore is null) and (mo.IdWareHouse = '%s' or mo.Id is null) " +
		"group by p.Id, p.IdCategory, p.Name, p.Price, stock, m.Name, pm.Unity, p.Description, p.Perishable, pw.Ignore, c.Name, m.Id, pm.Id, pm.MinAlert"},

	"getAll": {Q: "select p.Id, p.IdCategory, p.Name, p.Price, sum(Quantity) as stock, m.Name,  pm.Unity as Unidades, " +
		"p.Description, p.Perishable, pw.Ignore, c.Name, m.Id, pm.Id, pm.MinAlert from PRODUCT p " +
		"left join PRODUCT_WAREHOUSE pw on p.Id  = pw.IdProduct " +
		"left join PRODUCT_MEASURE pm on p.Id = pm.IdProduct " +
		"inner join MEASURE m on pm.IdMeasure = m.Id " +
		"left join MOVEMENT mo on p.Id = mo.IdProduct " +
		"inner join CATEGORY c on p.IdCategory = c.Id " +
		"where mo.IdWareHouse = '%s' or mo.Id is null " +
		"group by p.Id, p.IdCategory, p.Name, p.Price, stock, m.Name, pm.Unity, p.Description, p.Perishable, pw.Ignore, c.Name, m.Id, pm.Id, pm.MinAlert"},
}
