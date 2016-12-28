package api

type CommonResponse struct {
  Code      string  `json:"code"`
  Msg       string  `json:"msg"`
  SubCode   string  `json:"sub_code"`
  SubMsg    string  `json:"sub_msg"`
  Sign      string  `json:"sign"`
}

func (this *CommonResponse) IsSuccess() bool {
  return this.SubCode == ""
}
