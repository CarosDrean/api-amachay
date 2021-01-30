# API AMACHAY

Api Amachay es el que se encarga de toda la gestion de datos.

# Instalacion

```
go get
```

# Configuracion

Para las configuracionde Base de Datos debera crear el archivo **configuration.json** con los siguientes campos:

```json
{
  "engine": "mssql",
  "server": "DREAN",
  "port": "1433",
  "user": "sa",
  "password": "123456",
  "database": "AMACHAY"
}
```

Para las configuracinoes de alertas por telegram debera crear el archivo **configuration-telegram.json** con los siguienntes campos:

```json
{
  "token": "tokenbot",
  "chatId": "groupid"
}
```

# Compilacion

Para compilar el proyecto use:

```
go build
```
Para que el proyecto pueda funcionar correctamente no olvide el archivo **configuration.json** y **configuration-telegram.json**