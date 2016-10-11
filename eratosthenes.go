// Package eratosthenes implements the Sieve of Eratosthenes
// (https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
// using the Communicating Sequential Processes example given in the
// 'Bell Labs and CSP Threads' article by Russ Cox
// https://swtch.com/~rsc/thread/
package cspEratosthenes

// Eratosthenes calculates the first n prime numbers.
func Eratosthenes(n int) (primes []int) {
	//'n+1' as we need an additional channel to put rejected ints in.
	var pchan = make(chan int)
	var chans = make([]chan int, n+1)

	// we need to make each chan, else they are 'nil'.
	for i := range chans {
		chans[i] = make(chan int)
	}

	// kick off the processing logic between the channels with go routines.
	for i := 0; i < len(chans)-1; i++ {
		go process(chans[i], chans[i+1], pchan)
	}

	// accumulate the primes into the slice
	go accumulate(pchan, &primes)

	// keep putting ints into the first channel until we have a full primes slice.
	for i := 2; len(primes) != n; i++ {
		chans[0] <- i
	}

	return primes
}

// process picks up the first number from the left hand channel, adds it to the
// primes slice, and then checks whether every subsequent number from the left
// channel is divisible by the first prime. If it's not it sends it to the right
// channel
//
// As per https://swtch.com/~rsc/thread/sieve.gif
func process(left chan int, right chan int, primes chan int) {
	p := <-left
	primes <- p

	for {
		x := <-left
		if x%p != 0 {
			right <- x
		}
	}
}

func accumulate(pchan chan int, primes *[]int) {
	for p := range pchan {
		*primes = append(*primes, p)
	}
}
