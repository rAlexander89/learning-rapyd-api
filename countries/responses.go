package countries

type Status struct {
	ErrorCode    string `json:"error_code"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ResponseCode string `json:"response_code"`
	OperationId  string `json:"operation_id"`
}

type GetCountriesResponse struct {
	Status Status    `json:"status"`
	Data   []Country `json:"data"`
}
