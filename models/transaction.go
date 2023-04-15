package models

import (
	"bpjs-mono/structs"
	"bpjs-mono/system"
	"encoding/json"
	"fmt"
	"time"
)

func CreateTransaction(req structs.Transaction) (map[string]interface{}, error){
	/* Insert Master Transaction */
	Mtrx := map[string]interface{}{
		"RequestId" : req.RequestId,
	}

	reqMtrx := system.CreatedTimeNow(Mtrx)
	trx := ModelsDb.Table("MTransaction").Create(reqMtrx)

	/* Insert Detail Transaction */
	for _, val := range req.TransactionRecords {
		TrxRecord := map[string]interface{}{
			"RequestId": req.RequestId,
			"Customer": val.Customer,
			"Quantity": val.Quantity,
			"Price": val.Price,
		}


		reqTrxRecord := system.CreatedTimeNow(TrxRecord)
		ModelsDb.Table("TransactionRecord").Create(reqTrxRecord)
	}

	/* Get Response */
	rsp, _ := GetDetailTransaction(req.RequestId)

	return rsp, trx.Error
}

func GetDetailTransaction(RequestId int)(map[string]interface{}, error){
	var TransactionRecord []map[string]interface{}

	var totalItems int64
	tx := ModelsDb.Table("TransactionRecord")
	tx.Where("RequestId = ?", RequestId)
	tx.Count(&totalItems)
	tx.Find(&TransactionRecord)

	Response := map[string]interface{}{
		"request_id": RequestId,
		"data": TransactionRecord,
		"total_items": totalItems,
	}

	return Response, tx.Error
}

func GetListTransaction(req structs.TransactionFilter)([]map[string]interface{}, error){
	var Response []map[string]interface{}
	tx := ModelsDb.Table("MTransaction").Select("Id, RequestId, CreatedAt")

	if req.Limit > 0{
		tx.Limit(req.Limit)
	}
	if req.Offset > 0{
		tx.Offset(req.Offset)
	}

	var filter map[string]interface{}
	json.Unmarshal([]byte(req.Filter), &filter)
	/* FILTER */
	for index, element := range filter {
		_, index = system.RemoveCharacters(index)
		value := fmt.Sprintf("%v", element)
		tx.Where(index+" LIKE ?", "%"+value+"%")
	}

	tx.Find(&Response)

	for _, value := range Response{
		value["CreatedAt"] = value["CreatedAt"].(time.Time).Format("2006-01-02 15:04:05")
	}

	return Response, tx.Error
}

func CreateTransactionJSON(req structs.Transaction) (map[string]interface{}, error){
	RecordData, _ := json.Marshal(req.TransactionRecords)

	TrxPayload := map[string]interface{}{
		"RequestId": req.RequestId,
		"Records": RecordData,
	}
	
	reqTrxPayload := system.CreatedTimeNow(TrxPayload)

	trx := ModelsDb.Table("TransactionJson").Create(reqTrxPayload)

	// /* Get Response */
	rsp, _ := GetDetailTransactionJson(req.RequestId)

	return rsp, trx.Error
}

func GetDetailTransactionJson(RequestId int)(map[string]interface{}, error){
	var Transaction map[string]interface{}
	var UnmarslahRecord []map[string]interface{}

	tx := ModelsDb.Table("TransactionJson")
	tx.Where("RequestId = ?", RequestId)
	tx.Find(&Transaction)

	var total_items int
	if Transaction != nil{
		json.Unmarshal([]byte(Transaction["Records"].(string)), &UnmarslahRecord)
		total_items = len(UnmarslahRecord)
	}


	Response := map[string]interface{}{
		"request_id": RequestId,
		"data": UnmarslahRecord,
		"total_items": total_items,
	}

	return Response, tx.Error
}

func GetListTransactionJson(req structs.TransactionFilter)([]map[string]interface{}, error){
	var Response []map[string]interface{}
	tx := ModelsDb.Table("TransactionJson").Select("Id, RequestId, CreatedAt")

	if req.Limit > 0{
		tx.Limit(req.Limit)
	}
	if req.Offset > 0{
		tx.Offset(req.Offset)
	}

	var filter map[string]interface{}
	json.Unmarshal([]byte(req.Filter), &filter)
	/* FILTER */
	for index, element := range filter {
		_, index = system.RemoveCharacters(index)
		value := fmt.Sprintf("%v", element)
		tx.Where(index+" LIKE ?", "%"+value+"%")
	}

	tx.Find(&Response)

	for _, value := range Response{
		value["CreatedAt"] = value["CreatedAt"].(time.Time).Format("2006-01-02 15:04:05")
	}

	return Response, tx.Error
}