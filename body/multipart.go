package body

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/tansoz/request/readable"
)

// multipart boundary
func Boundary() string {
	return strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(time.Now().Unix()+rand.Int63()))), "=", "")
}

// multipart body in a request
func MultipartBody(fields []Field, boundary string, headers map[string][]string) Body {

	if headers == nil {
		headers = make(map[string][]string)
	}

	if boundary == "" {
		boundary = Boundary()
	}

	headers["Content-Type"] = []string{"multipart/form-data; boundary=" + boundary}

	readables := make([]readable.Readable, len(fields)*2+1)
	p := 0
	for _, f := range fields {
		readables[p] = readable.NewStringReadable("--" + boundary + "\r\n")
		readables[p+1] = f
		p += 2
	}

	if p > 0 {
		readables[p] = readable.NewStringReadable("--" + boundary + "--")
	}

	return RawBody(readable.NewListReadable(readables), headers)
}

func SimpleMultipartBody(fields ...Field) Body {

	return MultipartBody(fields, "", nil)
}
