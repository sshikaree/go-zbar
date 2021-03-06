go-zbar
=======

Go wrapper around ZBar - C bar code reader library.

## Requirements
Original zbar library required. On Ubuntu it can be installed with 
```$ sudo apt install zbar-dev```

## Installing go-zbar
``` go get -u github.com/sshikaree/go-zbar```

## Example

```go
import (
	"fmt"
	"github.com/sshikaree/go-zbar"
)

func main() {
	p := zbar.NewProcessor(1)
	defer p.Destroy()

	p.SetConfig(0, zbar.ZBAR_CFG_ENABLE, 1)
	p.Init("/dev/video0", 1)
	p.SetVisible(1)
	p.SetActive(1)
	if ok := p.ProcessOne(-1); ok < 0 {
		fmt.Println("Error occured. Exiting..")
		return
	} else if ok == 0 {
		fmt.Println("No symbols were found")
	}

	results := p.GetResults()
	if results == nil {
		return
	}
	symbol := results.SetFirstSymbol()
	if symbol == nil {
		return
	}
	fmt.Println("Symbol type:", symbol.GetType())
	fmt.Println("Symbol name:", symbol.GetName())
	fmt.Println("Symbol data:", symbol.GetData())	
}

```
## Example with DataHandler function
```go
import (
	"fmt"
	"github.com/sshikaree/go-zbar"
)

func ImgHandler(img *zbar.Image) {
	s := img.FirstSymbol()
	fmt.Println("Symbol name:", s.GetName())
	fmt.Println("Symbol data:", s.GetData())
}

func main() {
	p := zbar.NewProcessor(1)
	defer p.Destroy()

	p.SetConfig(0, zbar.ZBAR_CFG_ENABLE, 1)
	p.Init("/dev/video0", 1)

	p.SetDataHandler(ImgHandler)

	p.SetVisible(1)
	p.SetActive(1)
	p.UserWait(-1)
}
```