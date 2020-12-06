package db

import "database/sql"

var DB *sql.DB

type queryConfig struct {
	Name string
	Q    string
}

type TableDB struct {
	Name   string
	Fields []string
}

var user = TableDB{
	Name:   "dbo.Users",
	Fields: []string{"IdUser", "Username", "password", "rol", "IdPerson"},
}

var QuerySystemUser = map[string]*queryConfig{
	"getUserName": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[2] + " = %s;"},
	"get": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[0] + " = %s;"},
	"updatePassword": {Q: "update " + user.Name + " set v_Password = @Password where " + user.Fields[0] + " = %s;"},
	"list": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + ";"},
	"insert": {Q: "insert into "+ user.Name + "("+ fieldStringInsert(user.Fields)+  ") values (" +valuesString(user.Fields) + ");"},
	"update": {Q: "update " + user.Name + " set " + updatesString(user.Fields) + " where " + user.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + user.Name + " where " + user.Fields[0] + " = @ID"},
}

var category = TableDB{
	Name:   "dbo.Category",
	Fields: []string{"idCategory", "Name"},
}

var queryCategory = map[string]*queryConfig{
	"get":    {Q: "select " + fieldString(category.Fields) + " from " + category.Name + " where " + category.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(category.Fields) + " from " + category.Name + ";"},
	"insert": {Q: "insert into " + category.Name + " (" + fieldStringInsert(category.Fields) + ") values (" + valuesString(category.Fields) + ");"},
	"update": {Q: "update " + category.Name + " set " + updatesString(category.Fields) + " where " + category.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + category.Name + " where " + category.Fields[0] + " = @ID"},
}

var person = TableDB{
	Name: "dbo.Person",
	Fields: []string{"IdPerson", "Name", "LastName", "Phone", "Dni", "Adress", "mail"},
}

var queryPerson = map[string]*queryConfig{
	"get": {Q: "select " + fieldString(person.Fields) + " from " + person.Name + " where " + person.Fields[0] + " =%s;"},
	"list": {Q: "select " + fieldString(person.Fields) + " from " +person.Name + ";"},
	"insert": {Q: "insert into " + person.Name + "("+ fieldStringInsert(person.Fields) + ") values (" + valuesString(person.Fields) + ");"},
	"update": {Q: "update " + person.Name + " set " + updatesString(person.Fields) + " where " + person.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + person.Name + " where " + person.Fields[0] + " = @ID"},

}

var product = TableDB{
	Name: "dbo.Product",
	Fields: []string{"IdProduct", "Name" ,"description","Price","Stock","IdCategory"},
}

var QueryProduct = map[string]*queryConfig{
	"get": {Q: "select " + fieldString(product.Fields) + " from " + product.Name + " where " + product.Fields[0] + " =%s;"},
	"list": {Q: "select " + fieldString(product.Fields) + " from " +product.Name + ";"},
	"insert": {Q: "insert into " + product.Name + "("+ fieldStringInsert(product.Fields) + ") values (" + valuesString(product.Fields) + ");"},
	"update": {Q: "update " + product.Name + " set " + updatesString(product.Fields) + " where " + product.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + product.Name + " where " + product.Fields[0] + " = @ID"},

}

var input =  TableDB{
	Name: "dbo.Input",
	Fields: []string{"IdInput", "Date", "Quantity", "IdProduct"},
}

var QueryInput = map[string]*queryConfig{
	"get": {Q: "select " + fieldString(input.Fields) + " from " + input.Name + " where " + input.Fields[0] + " =%s;"},
	"list": {Q: "select " + fieldString(input.Fields) + " from "+ input.Name + ";"},
	"insert": {Q: "insert into " + input.Name + "("+ fieldStringInsert(input.Fields) + ") values (" + valuesString(input.Fields) + ");"},
	"update": {Q: "update " + input.Name + " set " + updatesString(input.Fields) + " where " + input.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + input.Name + " where " + input.Fields[0] + " = @ID"},
}