##Assemblies

Pretty cut and dry. I believe I came up with an efficient solution but not the most 
efficient solution. The faster solution is that rather than adding the numbers and 
doing a bunch of divisibility checks you don't need to do, I wrote the differences 
between numbers to an array.

For instance, between 3 and 5, the difference is 2. So the sum of numbers can be 
determined by adding these differences together. The array you are looking at is 
defined as 3,2,1,3,1,2,3; which if nothing else is pretty looking.

This isn't the fastest solution however. Let the number we are examining be n. 
Then j is the nearest number to n divisible by 3 and k is the nearest number 
divisble by 5. Let m be the nearest number to divisible by 15. Now, if we divide 
each of these numbers by n, we get the number of factors of n of 3, 5, and 15. 
Which can be listed in increasing order.

Since these factors can be listed in a format of say, 1 to e, one can write the 
sum of all factors of 3 as the sum of integers 1 to e (good ol n\*\*2 + n /2 ), 
times the factor we are summing up. Since I'm stuck at a fulltime job, but have 
time and permission to make my own notes when there are no work tasks, a part of me
wants to say I discovered this relation. However, I'm not sure I did, I can't find a
note reference for it. Certainly, it is available online:
https://www.mathblog.dk/project-euler-problem-1/.

##Testing

I wrote a naive algorithm to confirm numbers, pretending this problem was a bit more
complicated than it is.

I wrote down these algorithms and benchmarked them so that they'd run against all n < 1000.
Running tests shows that the algorithms crank out 233168 for all solution answers.
The total running times of these solutions (determined by benchmarks). Final running time is 16ns by the benchmarks for the multiplication, 657 ns for my version, and 2310 ns for the naive edition.

##Solution
An exectutable program that takes an argument n.
If no argument is passed prints out the solution.