package controllers

import (
	"bpjs-mono/helpers"
	"bpjs-mono/models"
	"bpjs-mono/structs"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateTransaction(e echo.Context) error {
	log.Println("Starting process Create Transaction")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Binding Request Payload to Struct */
	var req structs.Transaction
	if err := e.Bind(&req); err != nil {
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Proses Request to Database */
	rsp, tx := models.CreateTransaction(req)
	if tx != nil{
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest
	
		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Success"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func GetListTransactions(e echo.Context) error {
	log.Println("Starting process Get List Transaction")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Binding Request Payload to Struct */
	var req structs.TransactionFilter
	if err := e.Bind(&req); err != nil {
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Proses Request to Database */
	rsp, tx := models.GetListTransaction(req)
	if tx != nil{
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest
	
		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Success"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func GetDetailTransaction(e echo.Context)error{
	log.Println("Starting process Get Detail Transaction")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	RequestId := e.Param("request_id")
	ReqId, _ := strconv.Atoi(RequestId)

	/* Proses Request to Database */
	rsp, tx := models.GetDetailTransaction(ReqId)
	if tx != nil{
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest
	
		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Success"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func CreateTransactionJSON(e echo.Context) error {
	log.Println("Starting process Create Transaction Json")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Binding Request Payload to Struct */
	var req structs.Transaction
	if err := e.Bind(&req); err != nil {
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Proses Request to Database */
	rsp, tx := models.CreateTransactionJSON(req)
	if tx != nil{
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest
	
		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Success"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func GetListTransactionsJSON(e echo.Context) error {
	log.Println("Starting process Get List Transaction Json")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	/* Binding Request Payload to Struct */
	var req structs.TransactionFilter
	if err := e.Bind(&req); err != nil {
		response = &helpers.JSONResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad request",
			Data:       nil,
		}
		return e.JSON(response.StatusCode, response)
	}

	/* Proses Request to Database */
	rsp, tx := models.GetListTransactionJson(req)
	if tx != nil{
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest
	
		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Success"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}

func GetDetailTransactionJson(e echo.Context)error{
	log.Println("Starting process Get Detail Transaction")

	/* Declare Variable Response */
	response := new(helpers.JSONResponse)

	RequestId := e.Param("request_id")
	ReqId, _ := strconv.Atoi(RequestId)

	/* Proses Request to Database */
	rsp, tx := models.GetDetailTransactionJson(ReqId)
	if tx != nil{
		response.Data = ""
		response.Message = tx.Error()
		response.StatusCode = http.StatusBadRequest
	
		return e.JSON(http.StatusBadRequest, response)
	}

	response.Data = rsp
	response.Message = "Success"
	response.StatusCode = http.StatusOK

	return e.JSON(http.StatusOK, response)
}