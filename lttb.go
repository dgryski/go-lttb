// Package lttb implements the Largest-Triangle-Three-Buckets algorithm for downsampling points
/*

The downsampled data maintains the visual characteristics of the original line
using considerably fewer data points.

This is a translation of the javascript code at
    https://github.com/sveinn-steinarsson/flot-downsample/
*/
package lttb

import "math"

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

	var sampled []Point

	// Bucket size. Leave room for start and end data points
	every := float64(len(data)-2) / float64(threshold-2)

	sampled = append(sampled, data[0]) // Always add the first point

	var a int

	for i := 0; i < threshold-2; i++ {

		// Calculate point average for next bucket (containing c)
		avg_range_start := int(math.Floor(float64(i+1)*every) + 1)
		avg_range_end := int(math.Floor(float64(i+2)*every) + 1)

		if avg_range_end >= len(data) {
			avg_range_end = len(data)
		}

		avg_range_length := float64(avg_range_end - avg_range_start)

		var avg_x, avg_y float64
		for ; avg_range_start < avg_range_end; avg_range_start++ {
			avg_x += data[avg_range_start].X
			avg_y += data[avg_range_start].Y
		}
		avg_x /= avg_range_length
		avg_y /= avg_range_length

		// Get the range for this bucket
		range_offs := int(math.Floor(float64(i+0)*every) + 1)
		range_to := int(math.Floor(float64(i+1)*every) + 1)

		// Point a
		point_a_x := data[a].X
		point_a_y := data[a].Y

		var max_area float64
		var area float64
		var max_area_point Point

		var next_a int
		for ; range_offs < range_to; range_offs++ {
			// Calculate triangle area over three buckets
			area = math.Abs((point_a_x-avg_x)*(data[range_offs].Y-point_a_y)-
				(point_a_x-data[range_offs].X)*(avg_y-point_a_y)) * 0.5
			if area > max_area {
				max_area = area
				max_area_point = data[range_offs]
				next_a = range_offs // Next a is this b
			}
		}

		sampled = append(sampled, max_area_point) // Pick this point from the bucket
		a = next_a                                // This a is the next a (chosen b)
	}

	sampled = append(sampled, data[len(data)-1]) // Always add last

	return sampled
}
