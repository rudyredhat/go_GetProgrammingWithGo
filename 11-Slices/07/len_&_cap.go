// find length and cap of the slice
// append function affects the length and capacity, check output
package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	// affect the append func
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarf1", dwarfs1)
	dwarfs2 := append(dwarfs1, "Orcus")
	dump("dwarf2", dwarfs2)
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarf3", dwarfs3)

}

// dwarfs: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
// dwarfs[1:2]: length 1, capacity 4 [Pluto]
// dwarf1: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
// dwarf2: length 6, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus]
// dwarf3: length 9, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus Salacia Quaoar Sedna]
