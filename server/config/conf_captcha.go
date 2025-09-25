package config

type Captcha struct {
	Height   int     `json:"height" yaml:"height"`
	Width    int     `json:"width" yaml:"width"`
	Length   int     `json:"length" yaml:"length"`
	MaxSkew  float64 `json:"maxSkew" yaml:"max_skew"`
	DotCount int     `json:"dotCount" yaml:"dot_count"`
}
