// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point

//Diego Armando Gracia Hinojosa
//A01229716
package main

import (
	"math"
	"math/rand"
	"os"
	"fmt"
	"strconv"
)

type Point struct{ x, y float64 }

func (point Point) X() float64 {
	return point.x
}

func (point Point) Y() float64 {
	return point.y
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	fmt.Print("- ")
	for i := range path {
		if i > 0 {
			fmt.Print(path[i-1].Distance(path[i]), " + ")
			sum += path[i-1].Distance(path[i])
		}
	}
	fmt.Print(path[len(path)-1].Distance(path[0]))
	sum += path[len(path)-1].Distance(path[0])
	fmt.Print(" = ")
	return sum
}

//!-path

func onSegment(p, q, r Point) bool {
	if q.X() <= math.Max(p.X(), r.X()) && q.X() >= math.Min(p.X(), r.X()) && q.Y() <= math.Max(p.Y(), r.Y()) && q.Y() >= math.Min(p.Y(), r.Y()) {
		return true;
	}
	return false;
}

func orientation(p, q, r Point) int {
	val  := (q.Y() - p.Y()) * (r.X() - q.X()) - (q.X() - p.X()) * (r.Y() - q.Y())
	if val == 0 {
		return 0;
	}
	if val > 0 {
		return 1;
	} else {
		return 2;
	}
}

func doIntersect(p1, q1, p2, q2 Point) bool {
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}

	return false
}

func generateRandom (Paths Path, sides int) []Point {
	for i := 0; i < sides; i++ {
                Paths[i].x = ((rand.Float64() * 200)-100)
                Paths[i].y = ((rand.Float64() * 200)-100)
                fmt.Println("- ( ",Paths[i].X(),", ",Paths[i].Y(),")")
        }
	return Paths
}

func main() {
        if len(os.Args) < 2 {
                fmt.Println("Usage: go run geometry.go <sides>")
        }
	sides, err := strconv.Atoi(os.Args[1])
	if err == nil {
        	fmt.Println("- Generating a [",sides,"] sides figure")
	}
	fmt.Println("- Figure's vertices")
	Paths := make( Path,sides)
	Paths = generateRandom(Paths, sides)

	// Check line segments intersection
	for doIntersect(Paths[0], Paths[1], Paths[2], Paths[3]) {
		Paths = generateRandom(Paths, sides)
	}
	fmt.Println("- Figure's Perimeter")
	fmt.Println(Paths.Distance())
}
