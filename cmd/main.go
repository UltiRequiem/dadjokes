package main

import (
	"fmt"
	"strconv"
	"sync"

	"flag"

	chigo "github.com/UltiRequiem/chigo/pkg"
	dadjokes "github.com/UltiRequiem/dadjokes/pkg"
)

func main() {

	flag.Parse()

	times := flag.Arg(0)

	if times == "" {
		times = "1"
	}

	parsedTimes, errorParsingNumber := strconv.Atoi(times)

	if errorParsingNumber != nil {
		fmt.Println("The input was not a valid number.")
		panic(errorParsingNumber)
	}

	var wg sync.WaitGroup

	wg.Add(parsedTimes)

	for i := 0; i < parsedTimes; i++ {
		go func() {
			joke, errorGettingJoke := dadjokes.GetJoke()

			if errorGettingJoke != nil {
				fmt.Println("Probably you don't have internet connection.")
				panic(errorGettingJoke)
			}

			fmt.Println(chigo.Colorize(joke.Joke))

			wg.Done()
		}()
	}

	wg.Wait()
}
