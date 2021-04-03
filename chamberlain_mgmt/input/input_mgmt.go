package input

type Input struct {
	InputTime   uint32
	Year        uint16
	Month       uint8
	Type        string
	Base        float32
	AllInput    float32
	Tax         float32
	Actual      float32
	Description string
}

type InputMgmt interface {
	AddInput() error
	UpdateInput() error
	DeleteInput() error
	GetInput(year uint8, month uint8) ([]Input, error)
	GetInputByMonth(year uint8, month uint8) ([]Input, error)
	GetInputByYear(year uint8, month uint8) ([]Input, error)
	GetInputByType(year uint8, month uint8) ([]Input, error)
}
