# uniq: Make Unique items array

## requirement
go version 1.18 or later

## How to use

```
import (
	"github.com/UedaTakeyuki/uniq"
)

func main(){
  {
    a := []int{0, 1, 2, 0, 1}
    log.Println("a", a)
    b := uniq.Uniq[int](a)
    log.Println("b", b)
  }

  {
    a := []string{"alpha", "beta", "gamma", "alpha", "beta"}
    log.Println("a", a)
    b := uniq.Uniq[string](a)
    log.Println("b", b)
  }
}
```

## History
2023.06.22 v1.0.0