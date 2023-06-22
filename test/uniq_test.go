package uniq

import (
	"log"
	"testing"

	cp "github.com/UedaTakeyuki/compare"
	"local.packages/uniq"
)

func Test_01(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	// int
	a := []int{0, 1, 2, 0, 1}
	log.Println("a", a)
	b := uniq.Uniq[int](a)
	log.Println("b", b)
	cp.Compare(t, len(b), 3)
	cp.Compare(t, b[0], 0)
	cp.Compare(t, b[1], 1)
	cp.Compare(t, b[2], 2)
}

func Test_02(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	// string
	a := []string{"alpha", "beta", "gamma", "alpha", "beta"}
	log.Println("a", a)
	b := uniq.Uniq[string](a)
	log.Println("b", b)
	cp.Compare(t, len(b), 3)
	cp.Compare(t, b[0], "alpha")
	cp.Compare(t, b[1], "beta")
	cp.Compare(t, b[2], "gamma")
}

func Test_version(t *testing.T) {
	cp.Compare(t, uniq.Version(), "1.0.0")
}
