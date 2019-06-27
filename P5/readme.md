##Assemblies

My assemblies for project five were two functions that basically do the same thing:
using a modified version of the sieve of erasthones, count the factors of the number.

The assembly follows this routine:

First, create an array of booleans to the top of the number.
\[1\] Make two arrays that map factors to numbers of divisions.

Check to see if the number is factorable by a power of two.
If it is, then mark that number as not to be checked and increase divisions by two
by one.

For i := 3, i < n, i+=2:
    if i divides n, mark all multiples of i as numbers that divide n
    save the largest power of i

Finally, make a product of 1\*p_1^d_1\*...\*p_n^d_n, where p_1, d_1 is the prime factor
and highest degree that factor is raised to, as recorded in the arrays in \[1\].

The goal, aside from solving euler five, was to double check the idea that using
floats might be faster than using ints. And just from writing the code, that becomes doubtful.
The reason is that integers are used to address memory, and you're usually going to run into
a problem where integers are really, really useful (for instance, when you need to multiply a number by two a lot)

##Testing

The results generated by running go benchmarks can be summarized as follows:

|Name                      | #ran |  avg. runtime   | Bits/op       |   Allocations / op |
|:------------------------:|:----:|:---------------:|:-------------:|:------------------:|
|Benchmark_EulerFiveF64-8   |	10000000	|       162 ns/op	 |    224 B/op	 |      3 allocs/op|
|Benchmark_EulerFive-8   	|10000000	    |   141 ns/op	   |  224 B/op	     |  3 allocs/op |
|Benchmark_EulerFiveF64b-8   	|   20000	    | 75107 ns/op|	  204801 B/op	 |      3 allocs/op|
|Benchmark_EulerFiveb-8   	 |  20000	  |   71359 ns/op	 | 204801 B/op	    |   3 allocs/op|


Essentially, there's nearly no difference in the runtimes of these codes. F64 is the floating point version ran over project euler5 with the parameter 20.
The primary variant runs using simply integers. The b variant runs over the number 20000 instead of 20, to see if size makes a difference.

The results are interesting in that there's basically no advantage to not using a float where an integer could be used, except difficulty (and vice versa).

##Solution
An exectutable program that takes an argument n.
As of this project, you must specifically type in for the solution.
I doubt anyone will 'spoil' the solution anymore than I am by googling.
However, it now takes an act of individual agency to reveal the answer to the project.
Calling with no arguments prints out the answer of the project brief, which is 2520 for a value of 10.