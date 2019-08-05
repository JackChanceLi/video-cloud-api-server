package defs

type LiveRoomCondition struct {
	Lid string `json:"lid"`
	Condition int `json:"condition"`
	ConditionType int `json:"condition_type"`
	Price float32 `json:"price"`
	Duration int `json:"duration"`
	TryToSee int `json:"try_to_see"`
	Email string `json:"email"`
	VerificcationCode string `json:"verificcation_code"`
	WhiteList []string `json:"white_list"`
}

type DataForCondition struct {
	Code int `json:"code"`
	Data LiveRoomCondition `json:"data"`
	Msg Message `json:"msg"`
}

