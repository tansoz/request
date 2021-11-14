package readable

import (
	"fmt"
	"os"
)

type fileReadable struct {
	file *os.File
	len  int64
}

// get file readable.Readable by filename
func NewFileReadable(filepath string, offset int64) Readable {
	readable := new(fileReadable)

	if file, err := os.OpenFile(filepath, os.O_RDONLY, 0111); err == nil {
		readable.file = file
		fileInfo, _ := file.Stat()
		readable.len = fileInfo.Size()
		if offset > 0 {
			if _, err := file.Seek(offset, 0); err != nil {
				panic(err)
			}
			readable.len -= offset
		} else if offset >= readable.len {
			readable.len = 0
		}
		if readable.len < 0 {
			readable.len = 0
		}
	} else {
		panic(err)
	}

	return readable
}
func (f fileReadable) Read(p []byte) (int, error) {
	return f.file.Read(p)
}
func (f fileReadable) Len() int64 {
	return f.len
}
func (f fileReadable) String() string {
	return fmt.Sprintf("[BINARY_DATA(%d)]", f.len)
}
func (f fileReadable) Reset() {
	f.file.Close()
}
