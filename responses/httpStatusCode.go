package responses

const (
	ResCodeSuccess      = 20001
	ResCodeParamInvalid = 20003
)

var msg = map[int]string{
	ResCodeSuccess:      "success",
	ResCodeParamInvalid: "Email is invalid",
}
