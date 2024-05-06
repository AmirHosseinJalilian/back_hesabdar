package main

import (
	"github.com/AmirHosseinJalilian/back_hesabdar/services/pepole"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation"
	"github.com/AmirHosseinJalilian/back_hesabdar/services/sale_factor_confirmation_details"
)

func main() {

	sale_factor_confirmation_details.SaleFactorConfirmationDetails()
	sale_factor_confirmation.SaleFactorConfirmation()
	pepole.Pepole()
}
