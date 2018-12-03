package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chilakantip/avitar/log"
	"github.com/chilakantip/my_retail/env"
	"github.com/chilakantip/my_retail/mg_persist"
	"github.com/chilakantip/my_retail/pg_persist"
	"github.com/labstack/echo"
)

func assignHandlers(e *echo.Echo) {
	e.GET("/heartbeat", heartbeat)
	e.POST("/products", create)    //create
	e.GET("/products", read)       //read
	e.PUT("/products", update)     //update
	e.DELETE("/products", delete_) //delete

	return
}

//check system status
func heartbeat(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%s %s[%s] application running", env.AppName, env.Varsion, env.AppEnv))
}

func create(c echo.Context) error {
	defer c.Request().Body.Close()
	prod := product{}
	err := json.NewDecoder(c.Request().Body).Decode(&prod)
	if err != nil {
		log.Error("failed to parse product info", err)
		return apiErrProductInfoParseFailed().render(c)
	}
	err = prod.validate()
	if err != nil {
		log.Error("validation failed", err)
		return apiErrValidationFailed(err).render(c)
	}

	id, err := pg_persist.AddProduct(prod.CurrentPrise.Value, prod.CurrentPrise.CurrencyCode)
	if err != nil {
		log.Error("failed add product", err)
		return apiErrAddProductFailed().render(c)
	}

	err = mg_persist.AddProduct(id, prod.ProductDetails.Name, prod.ProductDetails.Description, prod.ProductDetails.Type)
	if err != nil {
		log.Error("failed add product", err)
		return apiErrAddProductFailed().render(c)
	}

	return c.String(http.StatusOK, fmt.Sprintf(`{"product_id":%d}`, id))

}

func read(c echo.Context) error {
	idPara := c.QueryParam("id")
	if idPara == "" {
		log.Error("product id is empty")
		return apiErrProductIDEmpty().render(c)
	}
	id, err := strconv.Atoi(idPara)
	if err != nil {
		log.Error("product id invalid")
		return apiErrInvalidProductID().render(c)
	}

	prise, code, err := pg_persist.GetProduct(id)
	if err == pg_persist.ErrNoRecords {
		return apiErrNoRecord().render(c)
	}
	if err != nil {
		log.Error("Failed to get product prise info", err)
		return apiErrGetProductFailed().render(c)
	}

	prodDet, err := mg_persist.GetProductDetails(id)
	if err == mg_persist.ErrNoRecords {
		return apiErrNoRecord().render(c)
	}
	if err != nil {
		log.Error("Failed to get product info", err)
		return apiErrGetProductFailed().render(c)
	}

	prodResp := product{ProductID: id}
	prodResp.CurrentPrise.Value, prodResp.CurrentPrise.CurrencyCode = prise, code
	prodResp.ProductDetails = prodDet

	return c.JSONPretty(http.StatusOK, prodResp, jsonIndent)
}

//its full update of product currency info and also product info
func update(c echo.Context) error {
	defer c.Request().Body.Close()
	prod := product{}
	err := json.NewDecoder(c.Request().Body).Decode(&prod)
	if err != nil {
		log.Error("failed to parse product info", err)
		return apiErrProductInfoParseFailed().render(c)
	}
	err = prod.validate()
	if err != nil {
		log.Error("validation failed", err)
		return apiErrValidationFailed(err).render(c)
	}

	id := prod.ProductID
	err = pg_persist.UpdateProduct(id, prod.CurrentPrise.Value, prod.CurrentPrise.CurrencyCode)
	if err != nil {
		log.Error("product update failed", err)
		return apiErrFailedUpdateProduct().render(c)
	}

	i := prod.ProductDetails
	err = mg_persist.UpdateProduct(id, i.Name, i.Description, i.Type)
	if err != nil {
		log.Error("product update failed", err)
		return apiErrFailedUpdateProduct().render(c)
	}

	return c.String(http.StatusOK, `{"message":"update success"}`)
}

func delete_(c echo.Context) error {
	idPara := c.QueryParam("id")
	if idPara == "" {
		log.Error("product id is empty")
		return apiErrProductIDEmpty().render(c)
	}
	id, err := strconv.Atoi(idPara)
	if err != nil {
		log.Error("product id invalid")
		return apiErrInvalidProductID().render(c)
	}

	err = pg_persist.DeleteProduct(id)
	if err == pg_persist.ErrNoRowsAffected {
		return apiErrNoRecordAffected().render(c)
	}
	if err != nil {
		log.Error("Failed to delete product", err)
		return apiErrDeleteProductFailed().render(c)
	}

	err = mg_persist.DeleteProduct(id)
	if err == mg_persist.ErrNoRowsAffected {
		return apiErrNoRecordAffected().render(c)
	}
	if err != nil {
		log.Error("Failed to delete product", err)
		return apiErrDeleteProductFailed().render(c)
	}

	return c.String(http.StatusOK, `{"message":"delete success"}`)
}
