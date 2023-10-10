package marketdata

type GetContractSizeParams struct {
	InstrumentName string `json:"instrument_name"`
}

type GetContractSizeResponse struct {
	ContractSize int `json:"contract_size"`
}
