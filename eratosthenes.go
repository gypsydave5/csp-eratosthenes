// Package cspEratosthenes implements the Sieve of Eratosthenes
// (https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
// using the Communicating Sequential Processes example given in the
// 'Bell Labs and CSP Threads' article by Russ Cox
// https://swtch.com/~rsc/thread/
package cspEratosthenes

// Eratosthenes calculates the first n prime numbers.
//
// In this version, the function that processes the first channel dynamically
// generates any extra channels needed to fill up the silce of primes.
func Eratosthenes(n int) (primes []int) {
	var primesChannel = make(chan int)
	var firstChan = make(chan int)

	go accumulate(primesChannel, &primes)
	go process(firstChan, primesChannel)

	for i := 2; len(primes) < n; i++ {
		firstChan <- i
	}
	return primes
}

// This time process only takes the input slice, and only generates an output
// slice when it needs one. Process is called recursively on the output slice,
// again using a go routine
//
// As per https://swtch.com/~rsc/thread/sieve.gif
func process(input chan int, primes chan int) {
	var output chan int
	var p = <-input

	primes <- p

	for {
		x := <-input

		if x%p != 0 {
			if output == nil {
				output = make(chan int)
				go process(output, primes)
			}
			output <- x
		}
	}
}

func accumulate(pchan chan int, primes *[]int) {
	for p := range pchan {
		*primes = append(*primes, p)
	}
}
