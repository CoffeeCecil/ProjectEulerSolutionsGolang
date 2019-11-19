##Assemblies

My assemblies were six functions that evaluated the sum square
 difference in a few different ways. Both used the difference of the 
 closed form formulas for the sum of squares and the square of the sum of squares.
 On performed simple multiplication so that it was written as: 
 (n\*\n\*n\*n4 - n\*n)/4 + (n\*n\*n - n)/6. Then two others that rearranged that formula to:
  n (n(n\*n-1)/4 + (n\*n-1)/6 and n(n*n-1)(n/4 + 1/6).

The last two functions where created by rewriting the first and third function to use 
the math/big library. With my math/big implimentations, I used the worst and best 
evaluating formula as my starting point. I had two versions of the best evaluating 
formula, because you can save a constant value 1 as a big.Float. If you don't want
to evaluate with this extra constant you have to rewrite the formula to be (n\*n - 1)(n\*n/4 + n/6).
It's an interesting question: which is better to avoid? Multiplications or constants?

Another thing to note is that my implementation floors and rounds, which gives the 
"fast and accurate" version of the number. Performance can be taken a step further 
by dropping all pretenses of safety and accepting floating point values with trailing
decimals. I did this with my math/big implimentations, mainly because there was no 
immediately handy rounding function I wanted to use.

##Testing

All formulas give the correct answer. But only the full simplification gives the best 
performance with float64 values. In general best float64 function ran in 7/10 of the 
time of the other formulas when benchmarked.  Otherwise, the two formulas ran in the 
same time. Considering that these functions run in between 7-11 ns/op when
benchmarked, I made the math.big version to see if there was a real difference.
These benchmarks were primarily with the number 10, but I didn't see a great difference
when I raised the number in terms of speed.

Using the math/big library the ratio difference seen between the best and worse 
function in terms of time was around 7/10. I believe the similarity is because rounding
scales the running time multiplicatively, and performing a floor function would
add a minimal constant amount of time to the algorithm before that scaling occurs.

Comparing best to second best, the ratio is near 19/20 in favor of using an extra constant.
For the big implimentation I tested using the maximum floating point 64 value and 1000**5,
and saw really similar performance. I also confirmed this by averaging numbers over runs - it
is at the point where how the cpu is ticking slightly effects the overall performance
of a run, enough to eliminate a correlation.

##Solution
An executable which takes an integer n, and then performs sums difference method on it 
using the float64 best performing function. If no argument is passed it requests an integer.