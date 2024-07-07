package magentgo

type ErrorResponse struct {
	Message string `json:"message"`
	Errors  []struct {
		Message    string `json:"message"`
		Parameters struct {
			FieldName string `json:"fieldName"`
		} `json:"parameters"`
	} `json:"errors"`
	Trace string `json:"trace"`
}