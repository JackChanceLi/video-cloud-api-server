package defs

var AuthFir = []string {"c_admin", "d_admin", "e_admin", "c_live", "d_live", "e_live"}
var AuthSed = []string {"c_live", "d_live", "e_live"}
var AuthTrd = []string {"e_live"}

type AdminAuth struct {
	Name string   `json:"name"`
	Auth []string `json:"auth"`
}

var (
	SuperAdmin = AdminAuth{Name: "SA", Auth: AuthFir}
	SeniorAdmin = AdminAuth{Name: "HA", Auth: AuthSed}
	JuniorAdmin =AdminAuth{Name: "JA", Auth: AuthTrd}
)