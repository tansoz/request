package readable

import (
	"io"
)

type listReadable struct {
	list    []Readable
	readpos int
	len     int64
}

func NewListReadable(readables []Readable) Readable {
	readable := new(listReadable)

	readable.list = readables
	readable.len = -1

	return readable
}
func (l *listReadable) Read(p []byte) (int, error) {
	var num int
	var err error
	ll := len(l.list)
	if ll > 0 && l.readpos < ll {
		if l.list[l.readpos] == nil {
			l.readpos += 1
			return l.Read(p)
		}
		num, err = l.list[l.readpos].Read(p)
		if err == io.EOF {
			err = nil
			l.list[l.readpos].Reset()
			l.readpos += 1
			if num == 0 {
				return l.Read(p)
			}
		}
	}
	if l.readpos >= ll {
		err = io.EOF
	}
	return num, err
}
func (l *listReadable) Len() int64 {
	if l.len == -1 {
		l.len = 0
		for _, readable := range l.list {
			if readable != nil {
				l.len += readable.Len()
			}
		}
	}
	return l.len
}
func (l listReadable) String() string {
	s := ""

	for _, readable := range l.list {
		if readable != nil {
			s += readable.String()
		}
	}

	return s
}
func (l *listReadable) Reset() {
	for _, readable := range l.list {
		if readable != nil {
			readable.Reset()
		}
	}
}
