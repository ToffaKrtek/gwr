package main

import (
	"math/rand"
)

const (
	widht  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, widht)
	}
	return u
}

func (u Universe) Seed() {
	for i := 0; i < (widht * height / 4); i++ {
		u.Set(rand.Intn(widht), rand.Intn(height), true)
	}
}

func (u Universe) Set(x, y int, b bool) {
	u[x][y] = b
}

func (u Universe) Alive(x, y int) bool {
	x = (x + widht) % widht
	y = (y + height) % height
	return u[x][y]
}

func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func Step(a, b Universe) {
	a, b = b, a
}
