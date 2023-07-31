package readwritemodels

type ReadData struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type WriteData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
