package controllers

import "log"

func checkError(err error, operation string, ctx string){
	if err != nil {
		log.Println("Error al: " + operation + " en: " + ctx)
		log.Println(err)
	}
}
