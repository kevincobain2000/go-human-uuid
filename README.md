<!-- <p align="center">
  <img alt="go-human-uuid" alt="human readable uuid in GO" src="https://imgur.com/fHfULta.png" width="160">
</p> -->

<p align="center">
  Generate Human Readable UUID in GO.
</p>

---

**UUID:** Generate human readable UUID in GO.

![go-build-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-human-uuid&type=go-build-time&branch=master)
![go-test-runtime](https://coveritup.app/badge?org=kevincobain2000&repo=go-human-uuid&type=go-test-runtime&branch=master)
![coverage](https://coveritup.app/badge?org=kevincobain2000&repo=go-human-uuid&type=coverage&branch=master)
![go-binary-size](https://coveritup.app/badge?org=kevincobain2000&repo=go-human-uuid&type=go-binary-size&branch=master)
![go-mod-dependencies](https://coveritup.app/badge?org=kevincobain2000&repo=go-human-uuid&type=go-mod-dependencies&branch=master)

![go-build-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-human-uuid&type=go-build-time&output=svg&width=160&height=160&branch=master&line=fill)
![go-test-runtime](https://coveritup.app/chart?org=kevincobain2000&repo=go-human-uuid&type=go-test-runtime&output=svg&width=160&height=160&branch=master)
![coverage](https://coveritup.app/chart?org=kevincobain2000&repo=go-human-uuid&type=coverage&output=svg&width=160&height=160&branch=master&line=fill)
![go-binary-size](https://coveritup.app/chart?org=kevincobain2000&repo=go-human-uuid&type=go-binary-size&output=svg&width=160&height=160&branch=master)
![go-mod-dependencies](https://coveritup.app/chart?org=kevincobain2000&repo=go-human-uuid&type=go-mod-dependencies&output=svg&width=160&height=160&branch=master&line=fill)




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
