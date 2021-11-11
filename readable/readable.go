package readable

type Readable interface {
	Read([]byte) (int, error)
	Len() int64
	String() string
	Reset()
}
