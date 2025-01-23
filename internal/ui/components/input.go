package components

type Input struct {
	value  string
	cursor int
}

func NewInput() Input {
	return Input{}
}

func (i Input) Value() string {
	return i.value
}

func (i *Input) HandleInput(str string) {
	i.value += str
	i.cursor += len(str)
}

func (i *Input) Backspace() {
	if len(i.value) > 0 {
		i.value = i.value[:len(i.value)-1]
		i.cursor--
	}
}

func (i *Input) Clear() {
	i.value = ""
	i.cursor = 0
}
