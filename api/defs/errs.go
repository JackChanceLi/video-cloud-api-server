package defs

type Err struct {
	Error string `json:"name"`
	ErrorCode string `json:"error_code"`

}

type ErrorResponse struct {
	ErrorCode int `json:"code"`
	Msg Message `json:"msg"`
}

//var (
//	ErrRequestBodyParseFailed = ErrorResponse{HttpSC:200, Error:Err {Error:"Request body parse failed", ErrorCode:"001" }}
//	ErrorNotAuthUser = ErrorResponse{HttpSC:200, Error:Err{Error:"User authentication failed", ErrorCode:"002"}}
//	ErrorDBError = ErrorResponse{HttpSC:200, Error:Err{Error:"DB ops failed", ErrorCode:"003"}}
//	ErrorInternalFaults = ErrorResponse{HttpSC:200, Error:Err{Error:"Internal service error", ErrorCode:"004"}}
//)

var (
	ErrRequestBodyParseFailed  = ErrorResponse{ErrorCode:402, Msg:Message{Error:"Request body parse failed", ErrorCode:"001"}}
	ErrorNotAuthUser           = ErrorResponse{ErrorCode:401, Msg:Message{Error:"User authentication failed",ErrorCode:"002"}}
	ErrorDBError               = ErrorResponse{ErrorCode:500, Msg:Message{Error:"DB ops failed", ErrorCode:"003"}}
	ErrorInternalFaults        = ErrorResponse{ErrorCode:501, Msg:Message{Error:"Internal service error", ErrorCode:"004"}}
	ErrorEmailRegistered       = ErrorResponse{ErrorCode:400, Msg:Message{Error:"Email has been registered", ErrorCode:"005"}}
	ErrRequestParamParseFailed = ErrorResponse{ErrorCode: 402, Msg:Message{Error: "Request params parse failed", ErrorCode:"006"}}
	ErrNoRowsInDB              = ErrorResponse{ErrorCode:405, Msg:Message{Error:"No result for retrieve liveroom by lid", ErrorCode:"007"}}
	ErrorNotAuthUserForRoom    = ErrorResponse{ErrorCode:403, Msg:Message{Error:"User auth failed for the liveroom", ErrorCode:"008"}}
	ErrorEmailNotRegistered    = ErrorResponse{ErrorCode:406, Msg:Message{Error:"Email has not been registered", ErrorCode:"009"}}
)

