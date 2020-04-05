package image

type Level string

const (
	Safe    Level = "safe"
	Info    Level = "info"
	Warning Level = "warning"
	Danger  Level = "danger"
)

func (l Level) String() string {
	return string(l)
}

func ValidateImage(img []byte) Level {
	//TODO:对接AI接口
	return Safe
}
