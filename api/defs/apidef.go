package defs


type UserIdentity struct {
	UserName string `json:"user_name"` //注册用户名
	Password   string `json:"password"` //注册用户的密码
	Email    string `json: "email"` //注册用户的邮箱
	//Role    int  `json:"role"`  //表示用户权限，1为管理员，2位普通用户
}

type DataForUser struct {
	SessionID string `json:"session_id"`
	User UserInformation `json:"user"`
}



type UserInformation struct {
	Aid string `json:"aid"`
	Cid string `json:"cid"`
	Name string `json:"name"`
	Email string `json:"email"`
	Auth string `json:"auth"`
}

type Message struct {
	Error string `json:"name"`
	ErrorCode string `json:"error_code"`
}

type SignedUp struct{
	Code    int `json:"code"` //session 验证是否成功
	Data    DataForUser   `json:"data"`
	Msg Message `json:"msg"`
}


type Register struct {
	Success bool `json:"success"`
	Username string `json:"user_name"`
}
//session处理字段
type Session struct{
	Uid string  //session对用的用户名
	TTL      int64  //session的有效期
}

type LiveRoom struct {
	Code  int `json:"code"`
	Data  DataForLiveRoom `json:"data"`
	Msg Message `json:"msg"`
}

type LiveRoomList struct {
	Code int `json:"code"`
	Data [] LiveRoomIdentity `json:"data"`
	Msg Message `json:"msg"`
}


type DataForLiveRoom struct {
	LiveRoomInfo LiveRoomIdentity `json:"live_room"`
}

type LiveRoomIdentity struct {
	Aid string `json:"aid"`
	Lid string `json:"lid"`
	Cid string `json:"cid"`
	Name string  `json:"name"`
	Kind int `json:"kind"`
	Size int `json:"size"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	PushUrl string `json:"push_url"`
	PullHlsUrl string `json:"pull_hls_url"`
	PullRtmpUrl string `json:"pull_rtmp_url"`
	PullHttpFlvUrl string `json:"pull_http_flv_url"`
	DisplayUrl string `json:"display_url"`
	Status int `json:"status"`
	Permission string `json:"permission"`
	CreateTime string `json:"create_time"`
}


type Resourse struct {
	Code  int `json:"code"`
	Data  DataForResourse `json:"data"`
	Msg Message `json:"msg"`
}

type ResourseList struct {
	Code int `json:"code"`
	Data [] ResourseIdentity `json:"data"`
	Msg Message `json:"msg"`
}


type DataForResourse struct {
	ResourseInfo ResourseIdentity `json:"resourse"`
}

type ResourseIdentity struct {
	Rid string `json:"rid"`
	Aid string `json:"aid"`
	Cid string `json:"cid"`
	Name string  `json:"name"`
	Rtype string `json:"rtype"`
	Size float64 `json:"size"`
	Label string `json:"label"`
	Time string `json:"time"`
}

var (
	EmptyUser = UserInformation{Cid:"", Name:"", Email:"", Auth:""}
	//EmptySignedUp = SignedUp{Success:false, SessionId:"", Cid:"", Name:"", Email:"", Auth:""}

)