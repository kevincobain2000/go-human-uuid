<!-- <p align="center">
  <img alt="go-human-uuid" alt="human readable uuid in GO" src="https://imgur.com/fHfULta.png" width="160">
</p> -->

<p align="center">
  Generate Human Readable UUID in GO.
</p>

---

**UUID:** Generate human readable UUID in GO.

## Usage

```
DARAW
XAHAJ
LAXOQ
LUPUV
VECIQ
```

**Generate with Options**

```go
import (
    g github.com/kevincobain2000/go-human-uuid/lib
)

func main() {
	options := []g.Option{
		func(opt *g.Options) error {
			opt.Length = 10
			return nil
		},
	}
    gen, err := g.NewGenerator(options...)
    gen.Generate()
}
```

### Disclaimer
