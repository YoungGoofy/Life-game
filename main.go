package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	world := make(Universe, height)
	for index := range world {
		world[index] = make([]bool, width)
	}
	return world
}

func (u Universe) Show() {
	fmt.Print("\033[H\033[2J", u.String())
}

func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

func (u Universe) Set(x, y int, flag bool) {
	u[y][x] = flag
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Neighbours(x, y int) int {
	var count int
	for i := y - 1; i < y+2; i++ {
		for j := x - 1; j < x+2; j++ {
			if i == y && j == x {
				continue
			} else {
				if u.Alive(j, i) == true {
					count++
				}
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	neighboursCount := u.Neighbours(y, x)
	if u.Alive(y, x) == true {
		if neighboursCount < 2 {
			return false
		}
		if 1 < neighboursCount && neighboursCount < 4 {
			return true
		}
	} else {
		if neighboursCount == 3 {
			return true
		}
	}
	return false
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = '-'
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}

	return string(buf)
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
	}
}
