package main

import "fmt"

func playWithAsterik(n int) {

	for i := 1; i <= n; i++ {
		for space := 1; space <= n-i; space++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}

}

func main() {

	playWithAsterik(5)

	/*

	       *

	      * *

	     * * *

	    * * * *

	   * * * * *

	*/

}
