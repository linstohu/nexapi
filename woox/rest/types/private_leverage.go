package types

type UpdateLeverageSettingParam struct {
	Leverage int64 `url:"leverage" validate:"required,oneof=1 2 3 4 5 10 15 20"`
}
