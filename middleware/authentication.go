package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/CarosDrean/api-amachay/constants"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/storage/mssql"
	"github.com/CarosDrean/api-amachay/storage/query-sql"
	"github.com/CarosDrean/api-amachay/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "no permitido"})
		}

		return f(c)
	}
}

func LoginEcho(c echo.Context) error {
	data := models.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := utils.NewResponse(utils.Error, "structura no valida", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	bUser, bPassword, idUser := validationLogin(&data)
	if !bUser || !bPassword {
		message := "¡No existe Usuario!"
		if bUser {
			message = "¡Contraseña Incorrecta!"
		}
		resp := utils.NewResponse(utils.Error, message, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := GenerateJWT(assemblyClaimUser(idUser))
	if err != nil {
		resp := utils.NewResponse(utils.Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	dataToken := map[string]string{"token": token}
	resp := utils.NewResponse(utils.Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func assemblyClaimUser(idUser string) models.ClaimUser {
	user, _ := mssql.UserDB{Ctx: "Auth", Query: query_sql.SystemUser}.Get(idUser)
	return models.ClaimUser{ID: idUser, Role: user.Role}
}

// 1er true: usuario existe, 2do true password correcto, string: id de usuario
func validationLogin(data *models.Login) (bool, bool, string) {
	stateLogin, id := mssql.ValidateSystemUserLogin(data.User, data.Password)
	switch stateLogin {
	case constants.Accept:
		return true, true, id
	case constants.NotFound:
		return false, false, ""
	case constants.InvalidCredentials:
		return true, false, ""
	default:
		return false, false, ""
	}
}

func CheckSecurity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Authorization", r.Header.Get("x-token"))
		token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					_, _ = fmt.Fprintln(w, "Su token ha expirado")
					return
				case jwt.ValidationErrorSignatureInvalid:
					_, _ = fmt.Fprintln(w, "Su firma de token no coincide")
					return
				default:
					_, _ = fmt.Fprintln(w, "Su token no es valido def")
					return
				}
			default:
				log.Println(err)
				_, _ = fmt.Fprintln(w, "Su token no es valido fin def")
				return
			}
		}

		if token.Valid {
			//fmt.Fprintf(w, "Bienvenido al sistema")
			next(w, r)
		} else {
			_, _ = fmt.Fprintf(w, "Su token no es valido fin")
			return
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	login := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error al leer el usuario %s\n", err)
		return
	}

	stateLogin, id := mssql.ValidateSystemUserLogin(login.User, login.Password)

	switch stateLogin {
	case constants.Accept:
		systemUser, _ := mssql.UserDB{Ctx: "Auth", Query: query_sql.SystemUser}.Get(id)
		userResult := models.ClaimUser{ID: id, Role: systemUser.Role}
		token, _ := GenerateJWT(userResult)
		result := models.ResponseToken{Token: token}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			fmt.Println(w, "Error al generar el json")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonResult)
	case constants.NotFound:
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "¡No existe Usuario!")
	case constants.InvalidCredentials:
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = fmt.Fprintf(w, "¡Contraseña Incorrecta!")
		break
	}
}
