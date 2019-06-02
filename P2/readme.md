##Assemblies

This is an odd one. FibEven was the algorithm I came up with, and shows what a
do - while loop looks like in Golang. It also uses one of my favorite features:
compact swaps. A swap in Golang (and Python) is written as a,b = b,a.
No need for a third variable or object (like in Java).

Now, I developed an odd scheme to find the next even Fibonacci number.
Basically, I observed that beyond a certain point, the third number
of the Fibonacci series will always be even.

Let Fib(n) return the nth Fibonacci number, where Fib(0) = 0, Fib(1) = 1, Fib(2) = 1, and Fib(3) = 2.

If Fib(n-2) is odd and Fib(n-1) is odd, then Fib(n) must be even because of even odd parity rules. Fib(n+1) and Fib(n+2) will be odd, because you are adding an even number odd number. The even number will be Fib(n), and the odd numbers will be Fib(n-1) and Fib(n+1) once it is generated under the standard scheme.
Observe that this condition holds for all numbers of the Fibonacci Sequence when n-2 = 1.

Now, I checked this against the internet when I was done coding it. And, well
there is an algorithm that is faster. But its odd how close they are.
The algorithm I implemented as fastest is the one layed out here:
https://www.mathblog.dk/project-euler-problem-2/

##Testing

The naive algorithm, tested at 4 million, produces a running time of about 
4-5k ns. Between my algorithm and the one Kristian Edlund came up with, mine is barely slower and both run at about 1300 ns when the input is at 4 million.

What do I mean by barely? Well, my algorithm assigns 272 B/op. Edlunds,
implimented in Golang, assigns 208 B/op. At n=160000000, there is a couple hundred nanosecond difference in Edlund's favor.

##Solution
An exectutable program that takes an argument n.
If no argument is passed prints out the solution.