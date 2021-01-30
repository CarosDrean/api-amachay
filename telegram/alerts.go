package telegram

import (
	"fmt"
	"github.com/CarosDrean/api-amachay/utils"
	"log"
	"net/http"
)

func AlertStock(warehouse string, product string, measure string, stock int) {
	config, err := utils.GetConfigurationTelegram()
	if err != nil {
		log.Println(err)
	}
	parseMode := "&parse_mode=markdown"
	message := fmt.Sprintf("Al Almacen *%s* se le esta agotando el stock de *%s*, stock actual: *%d %s*",
		warehouse, product, stock, measure)
	if stock == 0 {
		message = fmt.Sprintf("El Almacen *%s* se quedo sin stock de *%s*",
			warehouse, product)
	}
	resp, err := http.Get(fmt.Sprintf(endpoint + message + parseMode, config.Token, config.ChatId))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
