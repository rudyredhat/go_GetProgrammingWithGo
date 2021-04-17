// while copying of maps, impacts both the value
// so inshort maps are not copied
package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	// delete map key:val
	delete(planets, "Earth")
	fmt.Println(planets)
}

// map[Earth:whoops Mars:Sector ZZ9]
// map[Earth:whoops Mars:Sector ZZ9]
// map[Mars:Sector ZZ9]
