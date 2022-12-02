package _19

//
//import (
//	"fmt"
//	"math"
//)
//
//const maxBoundary float64 = 1000
//
//type vector3 struct {
//	x, y, z int
//}
//
//func zero() *vector3 {
//	return &vector3{}
//}
//
//type scanner struct {
//	id       int
//	location *vector3
//	beacons  []beacon
//}
//type beacon struct {
//	id            int
//	location      vector3
//	radiusTrusted float64
//	distances     map[int]float64
//}
//
//func (b *beacon) String() string {
//	result := fmt.Sprintf("Beacon %d: loc %v, edgeDist %v, beaconDists:%v", b.id, b.location, b.radiusTrusted, b.distances)
//	//result := fmt.Sprint(b.location.x) + ", " + fmt.Sprint(b.location.y) + ", " + fmt.Sprint(b.location.z) + " -- "
//	//result += "r: " + fmt.Sprint(b.radiusTrusted) + ", " + fmt.Sprint(len(b.distances)) + fmt.Sprint(b.distances)
//	return result
//}
//func (s scanner) String() string {
//	result := "*** Scanner " + fmt.Sprint(s.id) + " -- " + fmt.Sprint(s.location) + "\n"
//	for _, b := range s.beacons {
//		result += b.String() + "\n"
//	}
//	return result
//}
//
//func (s *scanner) cacheDistances() {
//	// init cache
//	for i := range s.beacons {
//		s.beacons[i].distances = make(map[int]float64)
//	}
//	for i, b := range s.beacons[:len(s.beacons)-1] {
//		s.beacons[i].radiusTrusted = b.distanceFromNearestBoundary()
//		for j, fromB := range s.beacons[i+1:] {
//			distance := b.distanceFrom(fromB)
//			s.beacons[i].distances[i+1+j] = distance
//			s.beacons[i+1+j].distances[i] = distance
//		}
//	}
//}
//
//func (b *beacon) distanceFromNearestBoundary() float64 {
//	dX := maxBoundary - math.Abs(float64(b.location.x))
//	dY := maxBoundary - math.Abs(float64(b.location.y))
//	dZ := maxBoundary - math.Abs(float64(b.location.z))
//	return math.Min(dX, math.Min(dY, dZ))
//}
//
//func (b *beacon) distanceFrom(from beacon) float64 {
//	deltaX := math.Abs(float64(b.location.x - from.location.x))
//	deltaY := math.Abs(float64(b.location.y - from.location.y))
//	deltaZ := math.Abs(float64(b.location.z - from.location.z))
//	return math.Round(math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2) + math.Pow(deltaZ, 2)))
//}
//
////func findAllBeacons(scanners []scanner) scanner {
////	// zero scanner will have some beacons already
////	result := scanners[0]
////	for {
////		var detected int
////
////		for _, s := range scanners {
////			if s.location == nil {
////				matches := make(map[int]beacon)
////				for _, sb := range s.beacons {
////					for _, rb := range result.beacons {
////						var match bool
////						for j, d := range sb.distances {
////							if d > sb.radiusTrusted || d > rb.radiusTrusted {
////								match = true
////								break
////							}
////							if sb.distances[j] != rb.distances[j] {
////								break
////							}
////						}
////						if match {
////							matchCount++
////						}
////					}
////				}
////			}
////		}
////
////		fmt.Println(detected, "matches found this round")
////		if detected == 0 {
////			return result
////		}
////	}
////}
