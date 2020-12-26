package query

import "github.com/CarosDrean/api-amachay/models"

var person = models.TableDB{
	Name:   "dbo.PERSON",
	Fields: []string{"Id", "Name", "LastName", "Cel", "Phone", "Address", "Dni", "Mail"},
}

var Person = models.QueryDB{
	"get":    {Q: "select " + fieldString(person.Fields) + " from " + person.Name + " where " + person.Fields[0] + " =%s;"},
	"list":   {Q: "select " + fieldString(person.Fields) + " from " + person.Name + ";"},
	"insert": {Q: "insert into " + person.Name + "(" + fieldStringInsert(person.Fields) + ") values (" + valuesString(person.Fields) + ");"},
	"update": {Q: "update " + person.Name + " set " + updatesString(person.Fields) + " where " + person.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + person.Name + " where " + person.Fields[0] + " = @ID"},
}
