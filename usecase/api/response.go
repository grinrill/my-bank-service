package api

type response struct {
	ok     bool			`json:"ok"`
	error  string		`json:"error,omitempty"`
	result interface{}	`json:"result,omitempty"`
}

func okResponse(result interface{}) response {
	return response{
		ok: true,
		error: "",
		result: result,
	}
}

func errorResponse(err error) response {
	return response{
		ok: false,
		error: err.Error(),
		result: nil,
	}
}