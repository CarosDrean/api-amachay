package middleware

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/CarosDrean/api-amachay/constants"
	"github.com/CarosDrean/api-amachay/storage"
	"github.com/CarosDrean/api-amachay/models"
	"github.com/CarosDrean/api-amachay/query"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer")
	}

	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("No se pudo leer")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo leer")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo leer")
	}
}

func GenerateJWT(userResult models.UserResult) string {
	claims := models.Claim{
		UserResult: userResult,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "Admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return result
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
			w.WriteHeader(http.StatusAccepted)
			//fmt.Fprintf(w, "Bienvenido al sistema")
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = fmt.Fprintf(w, "Su token no es valido fin")
			return
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	var user models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error al leer el usuario %s\n", err)
		return
	}

	stateLogin, id := storage.ValidateSystemUserLogin(user.User, user.Password)

	switch stateLogin {
	case constants.Accept:
		systemUser, _ := storage.UserDB{Ctx: "Auth", Query: query.SystemUser}.Get(id)
		userResult := models.UserResult{ID: id, Role: systemUser.Role}
		token := GenerateJWT(userResult)
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
