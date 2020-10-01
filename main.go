package main

import "fmt"


func main() {
	var lim int
 	//int var e
	e := 0.0
	fmt.Scanln(&lim)
	i := 0
	aux := 1
	flag := 0

	for i <= lim {
		for l := 0; l <= i; l++ {
			flag = 0
			if l == 0{
				flag = 1
			} else{
				aux *= l
			}
		}
		if flag != 1 {
			e += (1.0 / float64(aux))
			aux = 1
			fmt.Println("Actual:",e)
		} else {
			fmt.Println("Actual: 1")
		}
		
		i++
	}

}
