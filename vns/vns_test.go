package vns

import (
	"container/heap"
	// "fmt"
	"log"
	"math/rand"
	"reflect"
	"testing"
)

func TestRuleHeap(t *testing.T) {
	cases := []struct {
		insert []struct {
			r   uint32
			fit float64
		}
		want []struct {
			r   uint32
			fit float64
		}
	}{
		{
			// insert
			[]struct {
				r   uint32
				fit float64
			}{{0, 1.9}, {1, 3.5}, {2, 93.1}, {93, 2.0}},
			// want
			[]struct {
				r   uint32
				fit float64
			}{{2, 93.1}, {1, 3.5}, {93, 2.0}, {0, 1.9}},
		},
	}
	for _, c := range cases {
		h := new(RuleHeap)
		heap.Init(h)
		for _, el := range c.insert {
			heap.Push(h, el)
		}
		got := make([]struct {
			r   uint32
			fit float64
		}, len(c.insert))
		for i := 0; h.Len() > 0; i++ {
			got[i] = heap.Pop(h).(struct {
				r   uint32
				fit float64
			})
		}
		if !reflect.DeepEqual(got, c.want) {
			log.Fatalf("want %v\n, but got %v", c.want, got)
		}
	}
}

func TestFlip(t *testing.T) {
	cases := []struct {
		r    uint32
		p    int
		want uint32
	}{
		{0b0, 0, 0b1},
		{0b1, 0, 0b0},
		{0b10011100, 3, 0b10010100},
	}
	for _, c := range cases {
		if got, err := flip(c.r, c.p); err != nil {
			log.Fatal(err)
		} else if got != c.want {
			log.Fatalf("want %v, but was %v", c.want, got)
		}

	}
}

func TestFlipIsOwnInverse(t *testing.T) {
	cases := 10
	for i := 0; i < cases; i++ {
		r := rand.Uint32()
		p := rand.Intn(32)
		f, err := flip(r, p)
		if err != nil {
			log.Fatal(err)
		}
		got, err := flip(f, p)
		if err != nil {
			log.Fatal(err)
		}
		want := r
		if got != want {
			log.Fatalf("want %v, but got %v", want, got)
		}
	}
}

func TestFlipError(t *testing.T) {
	// We are assuming that const RADIUS = 2
	cases := []struct {
		r    uint32
		p    int
		want string
	}{
		{0b0, 32, "flip at position 32 outside interval [0, 31]"},
		{0b0, -1, "flip at position -1 outside interval [0, 31]"},
	}
	for _, c := range cases {
		if _, err := flip(c.r, c.p); err == nil {
			log.Fatal("want non-nil error, but was nil")
		} else if err.Error() != c.want {
			log.Fatalf("want %#v, but got %#v", c.want, err.Error())
		}
	}
}

func TestFlipN(t *testing.T) {

}

func TestPow(t *testing.T) {
	cases := []struct {
		n, m uint
		want uint64
	}{
		{0, 1, 0}, {1, 0, 1}, {10, 0, 1}, {2, 2, 4}, {3, 7, 2187},
	}
	for _, c := range cases {
		if got, err := pow(c.n, c.m); err != nil {
			log.Fatal(err)
		} else if got != c.want {
			log.Fatalf("want %v, but got %v", c.want, got)
		}
	}
}

func TestPowError(t *testing.T) {
	cases := []struct {
		n, m uint
		want string
	}{
		{0, 0, "pow(0, 0) is undefined"},
	}
	for _, c := range cases {
		if _, err := pow(c.n, c.m); err == nil {
			log.Fatal("want error but was nil")
		} else if err.Error() != c.want {
			log.Fatalf("want error message %#v, but got %#v", c.want, err.Error())
		}
	}
}

func TestGetNeighborhood(t *testing.T) {

}

func TestGenAllRules(t *testing.T) {

}

func TestCountAlive(t *testing.T) {
	cases := []struct {
		config []bool
		want   int
	}{
		{[]bool{}, 0},
		{[]bool{true}, 1},
		{[]bool{false}, 0},
		{[]bool{true, false, false, false, true, true, false}, 3},
	}
	for _, c := range cases {
		if got := CountAlive(c.config); got != c.want {
			log.Fatalf("want %v, but got %v", c.want, got)
		}
	}
}

func TestCalcFitness(t *testing.T) {

}

func TestFitness(t *testing.T) {

}
