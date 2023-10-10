package marketdata

type GetContractSizeParams struct {
	InstrumentName string `json:"instrument_name"`
}

type GetContractSizeResponse struct {
	ContractSize float64 `json:"contract_size"`
}
