package types

type Symbols struct {
	Data []string `json:"data"`
}

type SymbolParams struct {
	Symbol string `url:"symbol,omitempty" validate:"omitempty"`
}
