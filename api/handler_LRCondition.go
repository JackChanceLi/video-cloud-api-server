package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-api-server/api/dbop"
	"go-api-server/api/defs"
	"go-api-server/api/session"
	"io/ioutil"
	"log"
	"net/http"
)

func UpdateLRCondition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vars := r.URL.Query()
	aid := vars.Get("aid")
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.LiveRoomCondition{}
	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)

    condition, err := dbop.UpdateLRConditionByLid(ubody.Lid, ubody.VerificationCode, ubody.Email, ubody.Condition,
    	                  ubody.ConditionType, ubody.Duration, ubody.TryToSee, ubody.Price)
    if err != nil {
    	sendErrorResponse(w, defs.ErrorDBError)
		return
	}
    roomCondition := &defs.LiveRoomCondition{}
    roomCondition.Lid = condition.Lid
    roomCondition.Condition = condition.Condition
    roomCondition.ConditionType = condition.ConditionType
    roomCondition.Price = condition.Price
    roomCondition.Duration = condition.Duration
    roomCondition.TryToSee = condition.TryToSee
    roomCondition.VerificationCode = condition.VerificationCode
    roomCondition.WhiteUserList = condition.WhiteUserList

	if resp, err := json.Marshal(roomCondition); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
	defer session.UpdateSession(aid)
}