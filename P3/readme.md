##Assemblies

The biggest issue with making assemblies is that prime finding is such a rich
topic that one can easily get sidetracked. Finding primes was put into
its own folder and inspired an intermission or two while building this. Hence,
why it has taken so long to write what is really a rather simple program.

After a lot of work and some experiments with coroutines that failed to produce
a much faster version of the Miller-Rabin test, that was accurate, the two
contending solutions attacked the problem in the same way, by factoring out
the primes in the solution.

The code for Miller Rabin relies on searching a special array of primes
that will cover all int64 numbers. Since I'm not in college and this
is kind of an art project, I'm not to ashamed to say I got that from
https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test.

##Testing

Testing is mainly consumed by a number of failing projects that used 
less efficient, brute force algorithms or that used Goroutines.
Goroutines, at minimimum tripled the runtime of the task no matter how
I seemed to arrange them. Testing integers for primality with goroutines
also fowled up the miller rabin test in this case, leading to at
least one instance where I spotted an inaccurate number.

What ended up being most efficient was to scan and factor using
the Sieve of Erasthones. There's a stress test
that I used to determine a boundary condition in version, and the results from
it made it into the goroutine. After that condition, Miller Rabin is used.

Update 11/19/2019
I was reviewing the source code for this project. In an embarassing case of
operator fatigue I didn't write the correct project description. Because I did
actually manage to figure out an ugly bit of code using mutex's and locks
to correctly operate the Miller Rabin test, apparently up to the value given for this test. However, my assertion that I had not written an accurate test turned out to be accurate (see Project Euler 7). Mores the pitty because using channels and goroutines does speed up the code significantly for this project.
I've updated the code here to use the most accurate version of the miller rabin test.

##Solution
An exectutable program that takes an argument n.
If no argument is passed prints out the solution.