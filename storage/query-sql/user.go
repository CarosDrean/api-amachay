package query_sql

import "github.com/CarosDrean/api-amachay/models"

var user = models.TableDB{
	Name:   "dbo.USERS",
	Fields: []string{"Id", "IdPerson", "UserName", "Password", "Rol", "IdWarehouse"},
}

var SystemUser = models.QueryDB{
	"getUserName": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[2] + " = '%s';"},
	"get":         {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[0] + " = %s;"},
	"list":        {Q: "select " + fieldString(user.Fields) + " from " + user.Name + ";"},
	"insert":      {Q: "insert into " + user.Name + "(" + fieldStringInsert(user.Fields) + ") values (" + valuesString(user.Fields) + ");"},
	"update":      {Q: "update " + user.Name + " set " + updatesString(user.Fields) + " where " + user.Fields[0] + " = @ID;"},
	"delete":      {Q: "delete from " + user.Name + " where " + user.Fields[0] + " = @ID"},
}
