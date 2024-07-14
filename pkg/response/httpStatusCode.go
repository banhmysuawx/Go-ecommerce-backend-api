package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	Error               = 500
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Param invalid",

}

func GetMsg(code int) string {
	return msg[code]
}