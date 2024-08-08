package scriptParser

type NodeType int

const (
	Global NodeType = iota
    NewBatch
	BatchContents
	Batch
	File
)
