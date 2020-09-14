/*
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s =½*a*t^2 + vo*t + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a,
initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement
as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function
returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which
is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a
function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import "fmt"

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64
	fmt.Printf("Enter an acceleration value:")
	fmt.Scan(&acceleration)
	fmt.Printf("Enter an initial velocity value:")
	fmt.Scan(&initialVelocity)
	fmt.Printf("Enter a displacement value:")
	fmt.Scan(&initialDisplacement)
	fmt.Printf("Enter time after displacement:")
	fmt.Scan(&time)
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("Displacement after %f seconds: %f", time, fn(time))
}

/*
The function below takes 3 float64 arguments and returns a function back which takes a single float 64 argument and
returns a float64 value. Here we are using anonymous functions and applying closure in terms of lexical scope.
Just like Java, go also allows variadic functions such as func name(... int) .
*/

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(time float64) float64 {
	// anonymous function with closure
	return func(t float64) float64 {
		return (0.5 * acceleration * t * t) + (initialVelocity * t) + initialDisplacement
	}
}
