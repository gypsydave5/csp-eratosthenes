# Eratosthenes with CSP

According to [this Golang blog post][goblog], channels in Go are based on C. A.
Hoare's [Communicating Sequential Processes][QSP]. That same blog post
references [a brief history of CSP][hist], which outlines an agorithm for
calculating prime numbers using the [sieve of Eratosthenes][sieve] and QSP
channels.

This is an implementation of that algorithm in Go.

[QSP]:http://www.usingcsp.com/
[goblog]:https://blog.golang.org/share-memory-by-communicating
[hist]:https://swtch.com/~rsc/thread/
[sieve]:https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes