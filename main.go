package main

import "fmt"


func main() {
	var lim int
 	//int var e
	e := 0.0
	fmt.Scanln(&lim)
	i := 0
	aux := 1

	for i <= lim {
		for l := 0; l <= i; l++ {
			if l == 0{
				aux = 1
			} else{
				aux *= l
			}
		}
		if aux != 0 {
			e += (1.0 / float64(aux))
			aux = 1
			fmt.Println(e)
		}
		
		i++
	}

}
