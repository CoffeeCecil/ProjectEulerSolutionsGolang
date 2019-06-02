##Project Motivation

So I'm working at the DMV. I have somehow lucked into working with some 
swell people with a low traffic flow. Since I'm a greeter, that means I 
have to say hi to a few hundred people per day and then find something to 
do while no one is there. Project Euler has been really good for filling the time.  

When I come home I code. I'm coding in Golang after testing out Rust for
a few weeks. Rust just wasn't effective for attacking the kind of problems 
I was working with, and didn't have a runtime speedup compared to Golang 
that justified me using it. I'm not using Python not because of the speed 
(one could make a relative argument that 'well, some smart guy can program this in C and it will run a thousand times faster!') but because in Python, you shouldn't optimize.

Now I know that sounds like a weird thing to say, but if you try an make 
an algorithm run more optimal, you're going to have to hit the 
Python abstraction with a baseball bat until you get to the raw bits. If you 
shift a bit on an integer, you probably are going to run a gauntlet of object 
code which normally just works under the conditions assumed by the interpreter.
And might very well have access to abstractions you can't reach under a C api.

Whereas in Golang, you are given the tools to mess with low level abstractions, 
with a tuple and multiple return syntax that makes life a lot simpler than C.
You can also potentially expose these algorithms to threading and parallelism,
which can sometimes yield faster solutions.

##Instructions

Where assemblies exist, use 'go test . -v' against the files in 
the source folder to print out their testing results. To get at 
the benchmarks use 'go test -benchmem -bench .' to get a readout 
of the speed and memory used by the project.

Where solutions exist, simply execute the program in the bin after you've 
built it on your system in the normal way, without arguments. This
will print out the solution on the command line.

##Testmarking and Benching

If you run a go program and time it you run into the problem that the program 
itself usually takes awhile to start up. So I'm benchmarking based off of the 
Golang built in test suite. I'm also including an exectutable that prints out
the a verified answer for the project.
 