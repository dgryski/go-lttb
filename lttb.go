// Package lttb implements the Largest-Triangle-Three-Buckets algorithm for downsampling points
/*

The downsampled data maintains the visual characteristics of the original line
using considerably fewer data points.

This is a translation of the javascript code at
    https://github.com/sveinn-steinarsson/flot-downsample/
*/
package lttb

import (
	"math"
)

// Point is a point on a line
type Point struct {
	X float64
	Y float64
}

// LTTB down-samples the data to contain only threshold number of points that
// have the same visual shape as the original data
func LTTB(data []Point, threshold int) []Point {

	if threshold >= len(data) || threshold == 0 {
		return data // Nothing to do
	}

	sampled := make([]Point, 0, threshold)

	// Bucket size. Leave room for start and end data points
	every := float64(len(data)-2) / float64(threshold-2)

	sampled = append(sampled, data[0]) // Always add the first point

	bucketStart := 1
	bucketCenter := int(math.Floor(every)) + 1

	var a int

	for i := 0; i < threshold-2; i++ {

		bucketEnd := int(math.Floor(float64(i+2)*every)) + 1

		// Calculate point average for next bucket (containing c)
		avgRangeStart := bucketCenter
		avgRangeEnd := bucketEnd

		if avgRangeEnd >= len(data) {
			avgRangeEnd = len(data)
		}

		avgRangeLength := float64(avgRangeEnd - avgRangeStart)

		var avgX, avgY float64
		for ; avgRangeStart < avgRangeEnd; avgRangeStart++ {
			avgX += data[avgRangeStart].X
			avgY += data[avgRangeStart].Y
		}
		avgX /= avgRangeLength
		avgY /= avgRangeLength

		// Get the range for this bucket
		rangeOffs := bucketStart
		rangeTo := bucketCenter

		// Point a
		pointAX := data[a].X
		pointAY := data[a].Y

		maxArea := -1.0

		var nextA int
		for ; rangeOffs < rangeTo; rangeOffs++ {
			// Calculate triangle area over three buckets
			area := (pointAX-avgX)*(data[rangeOffs].Y-pointAY) - (pointAX-data[rangeOffs].X)*(avgY-pointAY)
			// We only care about the relative area here.
			// Calling math.Abs() is slower than squaring
			area *= area
			if area > maxArea {
				maxArea = area
				nextA = rangeOffs // Next a is this b
			}
		}

		sampled = append(sampled, data[nextA]) // Pick this point from the bucket
		a = nextA                              // This a is the next a (chosen b)

		bucketStart = bucketCenter
		bucketCenter = bucketEnd
	}

	sampled = append(sampled, data[len(data)-1]) // Always add last

	return sampled
}
