package sylvanpkg

// A golang package to provide common sylvan forest growth model functions
//
// by David R. Larsen, Professor Emeritus
// The School of Natural Resources
// Univerity of Missouri
//
// September 2, 2021
//

import (
	"fmt"
)

// Plot descriptions information
type Plot struct {
	Units   string // An integer unit used for all dimensions default "centimeter"
	Origin  string // This is the 0,0 point  "CT" as center, or "SW" SW corner
	Xmin    int64  // Minimum X value
	Xmax    int64  // Minimum X value
	Ymin    int64  // Minimum Y value
	Ymax    int64  // Maximum Y value
	Elevmin int64  // Minimum Elevation value
	Elevmax int64  // Maximum Elevation value
}

// Tree location a identification destriptors
type Tree struct {
	Id    int64  //Tree id usually a number
	GX    int64  //Tree x coordinate
	GY    int64  //Tree y coordinate
	GElev int64  //Tree elevation coordinate
	Spp   string //Tree species code
	//Section []slice
}

// Radial section information for a specific tree
// one slice assume a circle crown, additional section allow irregularities
type Section struct {
	Azi float64 //Azimuth of the section
	GR  int64   //Ground radius of the tree stem 2*GR is the ground diameter
	BHR int64   //Breast height radius of the tree stem 2*BHR is the diameter at breast height DBH
	BHt int64   //Breast height for the tree stem,  height of BHR measurement in Elevation
	Tht int64   //Tree height in Elevation
}

// Plot definition structure
type Generation struct {
	Ptype     string   //Plot type "plantation" or "natural"
	Ppattern  []string //Planting pattern by species
	Species   []string //List of species in plot
	Spdensity []int64  //List of stem density by species
	Hardcore  int64    //Space around each tree in which no tree can be added
}

//Create a forest plot
func createplot() string {
	return fmt.Sprintf("%s", "fred")
}

//Write data to a json file

//Read data from a json file
