##Assemblies

Assemblies used the source code for Prime finding from Project Euler 3. After a brief review the method MillerRabin644 was chosen.

##Testing

Testing was simply ensuring that the prime finding method would go up to prime 10001. Also, since OIES is a thing, I took the liberty of looking up the desired prime number on that database and verifying that my search function returns the correct number.

Unfortunately, a goroutine driven piece of code didn't put out the right prime number. Simply changing to a non-goroutine version of the miller rabin test fixed that. The moral of the story is: always use Unit Tests over verifiable input! The more the better!

##Solution
An exectutable program that takes an argument n.
If no argument is passed prints out the solution.