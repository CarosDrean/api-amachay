package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func checkError(err error, operation string, ctx string){
	if err != nil {
		log.Println("Error al: " + operation + " en: " + ctx)
		log.Println(err)
	}
}

func returnErr(w http.ResponseWriter, err error, operation string) {
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al %s, error: %s", operation, err.Error()))
}
