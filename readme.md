To run these solutions, use go -test against the files in the source folder.

So I'm working at the DMV. I have somehow lucked into working with some swell people with a low traffic flow. Since I'm a greeter, that means I have to say hi to about 400 people per day and then find something to do with the half hour in between. Project Euler is really good for filling the time.  

When I come home I'll code. I'm coding in Golang after testing out Rust for a few weeks. Rust just wasn't effective for attacking the kind of problems I was working with, and didn't have a runtime speedup compared to Golang that justified me using it. I'm not using Python not because of the speed (one could make a relative argument that 'well, some smart guy can program this in C and it will run a thousand times faster!') but because in Python, you shouldn't optimize.

Now I know that sounds like a weird thing to say, but if you try an make an algorithm run more optimal, you're going to have to hit the Python abstraction with a baseball bat until you get
to the raw bits. If you shift a bit on an integer, you probably are going to run a gauntlet of object code which is gauranteed to just work under the conditions assumed by the interpreter.
Which at the end of the day, means you probably aren't being nearly as clever as you think you are.

 Whereas in Golang, they're right in front you to play with, with a tuple and multiple return
 syntax that makes life a lot simpler. And if I really go crazy built in threading and parallelism routines.
 
 If you run a go program and time it you run into the problem that the program itself usually takes awhile to start up. So I'm benchmarking based off of the Golang built in test suite.
 So there isn't an executable, and unless a program in Euler does something real special their
 won't be an executable.