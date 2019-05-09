#### a wraper of http package to send request<br>
 `surport chainable usage`<br>

#### Examples
firstly, import it, we can give a breif name of this package, like "hc"

```go
import (
	hc "github.com/DingHub/su4go/httpclient"
	"net/http"
)
```

```go
hc.New().Get("http://127.0.0.1:8888/test").Do(func(response *http.Response, body []byte, err error) {
	// do something with response
	})
```
```
type Person struct {
    Age int
    Name string
}

p := Person{Age: 27, Name: "Tom"}

hc.New().POST("http://127.0.0.1:8888/test").
	Header("some key", "some value").
	ContentType(hc.ContentTypeJson).
	Body(p).
	Do(func(response *http.Response, body []byte, err error) {
	// do something with response
	})
```
