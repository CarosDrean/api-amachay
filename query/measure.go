package query

import "github.com/CarosDrean/api-amachay/models"

var measure = models.TableDB{
	Name:   "dbo.MEASURE",
	Fields: []string{"Id", "Name"},
}

var Measure = models.QueryDB{
	"get":    {Q: "select " + fieldString(measure.Fields) + " from " + measure.Name + " where " + measure.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(measure.Fields) + " from " + measure.Name + ";"},
	"insert": {Q: "insert into " + measure.Name + " (" + fieldStringInsert(measure.Fields) + ") values (" + valuesString(measure.Fields) + ");"},
	"update": {Q: "update " + measure.Name + " set " + updatesString(measure.Fields) + " where " + measure.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + measure.Name + " where " + measure.Fields[0] + " = @ID"},
}
