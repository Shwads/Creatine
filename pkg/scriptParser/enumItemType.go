package scriptParser

type ItemType int

const (
	BatchName ItemType = iota
	FileName
	Parentheses
)
