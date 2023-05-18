package types

type DefaultParam struct {
	RecvWindow int    `url:"recvWindow,omitempty" validate:"omitempty"`
	Timestamp  int64  `url:"timestamp" validate:"required"`
	Signature  string `url:"signature,omitempty" validate:"omitempty"`
}
