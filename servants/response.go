package servants

import "net/http"

type result struct {
	Rsp interface{} `json:"rsp,omitempty"`
}

type response struct {
	Args result `json:"args"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
}

func okResp(data interface{}) *response {
	return &response{
		Args: result{
			Rsp: data,
		},
		Code: http.StatusOK,
		Msg:  "success",
		Ret:  0,
	}
}

func errResp(err error) *response {
	return &response{
		Args: result{},
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
		Ret:  0,
	}
}
