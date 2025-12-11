package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Circuit struct {
	id    int
	nodes []int
}

type JunctionBox struct {
	X             int
	Y             int
	Z             int
	id            int
	connectedList []int
	nodeDistances map[int]int
}

func (jb *JunctionBox) StraightLineDistance(jb2 *JunctionBox) int {
	xDiff := AbsInt(jb.X - jb2.X)
	yDiff := AbsInt(jb.Y - jb2.Y)
	zDiff := AbsInt(jb.Z - jb2.Z)
	return (xDiff * xDiff) + (yDiff * yDiff) + (zDiff * zDiff)
}

func (jb *JunctionBox) AddConnection(nn int) {
	jb.connectedList = append(jb.connectedList, nn)
}

func day08_1(input string) {
	fmt.Println("Day 8, Part 1")

	// create an empty array/slice to store all Junction Boxes and Circuits
	var junctionBoxes []*JunctionBox
	// var circuits []*Circuit

	// parse the input to get the x,y,z coordinates
	// build Junction Box structs
	// put those structs into the array
	for i, line := range strings.Split(input, "\n") {
		x, y, z := getCoordinates(line)
		// fmt.Println(x, y, z)
		jb := &JunctionBox{X: x, Y: y, Z: z, id: i, nodeDistances: make(map[int]int)}
		junctionBoxes = append(junctionBoxes, jb)
	}

	// for each Junction Box in the list, compute its
	// straight-line distance to the other boxes
	// store the distances in a k,v store with k=ID and v=distance to tht ID
	for _, currentJB := range junctionBoxes {
		for _, jb := range junctionBoxes {
			if currentJB.id != jb.id {
				distanceBetweenJBs := currentJB.StraightLineDistance(jb)
				currentJB.nodeDistances[jb.id] = distanceBetweenJBs
			}
		}
	}

	//find nearest neighbor and append that to the connectedList
	for i, _ := range junctionBoxes {
		for _, jb := range junctionBoxes {
			_, nn, _ := findNearestNeighbor(jb)
			jb.AddConnection(nn)
			// remove the nearestNeighbor from the nodeDistances Map since we don't want keep matching on it
			delete(jb.nodeDistances, nn)
			fmt.Println(jb)
		}
		fmt.Println()
		i++
	}

	// create Circuits based on the distances of Neighbors
}

func findNearestNeighbor(jb *JunctionBox) (int, int, int) {
	var nearestNeighborDistance int
	var nearestNeighborID int
	first := true
	for k, v := range jb.nodeDistances {
		if first {
			nearestNeighborID = k
			nearestNeighborDistance = v
			first = false
			continue
		}
		if v < nearestNeighborDistance {
			nearestNeighborID = k
			nearestNeighborDistance = v
		}
	}
	// fmt.Println(jb.id, nearestNeighborID, nearestNeighborDistance)
	return jb.id, nearestNeighborID, nearestNeighborDistance
}

func getCoordinates(line string) (int, int, int) {
	parts := strings.Split(line, ",")
	// fmt.Println(parts)
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return x, y, z
}

func AbsInt(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
