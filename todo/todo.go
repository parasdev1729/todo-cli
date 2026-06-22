package todo

type Priority int

const (
	High Priority = 1
	Mid  Priority = 2
	Low  Priority = 3
)

type Todo struct {
	id       string
	title    string
	priority Priority // 1 - HIGH, 2 - MID, 3 - LOW
}
