package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Point struct {
	x int64
	y int64
}

type Segment struct {
	vertices [2]Point
}

func main() {

	dat, _ := os.ReadFile("./example.txt")
	
	strInput := string(dat)
	array := strings.Split(strInput, "\n")
	fmt.Println(array)

	result := part1(array)
	fmt.Println(fmt.Sprintf("part 1 result: %d", result))
}

func part1(array []string) (int) {
	xMax := int64(0)
	yMax := int64(0)

	var segments []Segment
	for _, s := range array {
		points := strings.Split(s, " -> ");

		if len(points) > 1 {
			// Convert data in points
			var segmentPoints = Segment{}
			for i, point := range points {
				coordinates := strings.Split(point, ",");
				xCoordinate, _ := strconv.ParseInt(coordinates[0], 10, 64)
				yCoordinate, _ := strconv.ParseInt(coordinates[1], 10, 64)

				// Compute maximal bounds of the coordinates
				if xMax < xCoordinate {
					xMax = xCoordinate
				}

				if yMax < yCoordinate {
					yMax = yCoordinate
				}

				segmentPoints.vertices[i] = Point{x: xCoordinate, y: yCoordinate}
				fmt.Println(segmentPoints.vertices[i])
			}

			if IsStraightLine(segmentPoints) {
				segments = append(segments, segmentPoints)
				fmt.Println(segmentPoints)
			}
		}
	}

	fmt.Println(segments)
	fmt.Println(xMax + 1)
	fmt.Println(yMax + 1)
	overlapCount := ComputeOverlapNumber(segments, xMax, yMax, 2)
	
	return overlapCount
}

func IsStraightLine(segment Segment) (bool) {
	return segment.vertices[0].x == segment.vertices[1].x || segment.vertices[0].y == segment.vertices[1].y
}

func ComputeOverlapNumber(segments []Segment, xMax int64, yMax int64, minOverlap int) (int) {
	overlapCount := 0
	//var metPoints map[Point]int
	var metPoints = make(map[int]int)

	for _, s := range segments {
		// Compute number of this cell
		fmt.Println(s)

		isHorizontal := true
		startingCoordinateBound := int64(0
		endingCoordinateBound := 0
		commonCoordinate := 0

		if s.vertices[0].x == s.vertices[1].x {
			commonCoordinate = s.vertices[0].x
			if s.vertices[0].y < s.vertices[1].y {
				startingCoordinateBound = s.vertices[0].y
				endingCoordinateBound = s.vertices[1].y
			} else {
				startingCoordinateBound = s.vertices[1].y
				endingCoordinateBound = s.vertices[0].y
			}
			
		} else {
			isHorizontal = false;
			commonCoordinate = s.vertices[0].y
			if s.vertices[0].x < s.vertices[1].x {
				startingCoordinateBound = s.vertices[0].x
				endingCoordinateBound = s.vertices[1].x
			} else {
				startingCoordinateBound = s.vertices[1].x
				endingCoordinateBound = s.vertices[0].x
			}
		}

		for i := startingCoordinateBound + 1; i < endingCoordinateBound + 1; i++ {

			// Add the point 
			// if metPoints[pts] == nil {
			// 	metPoints[pts] = 0
			// }
			//value = 

			metPoints[value] += 1
		}
	}

	fmt.Println(metPoints)
	return overlapCount;
}