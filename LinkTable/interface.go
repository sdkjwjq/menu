package LinkTable

const (
	SUCCESS = 0
	FAILURE = -1
)

type LinkTableNode interface {
	Next() *LinkTableNode
}

type LinkTable interface {
	GetHead() *LinkTableNode
	GetTail() *LinkTableNode
	GetLength() int
}
