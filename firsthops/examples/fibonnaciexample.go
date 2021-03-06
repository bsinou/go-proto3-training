/* Forked from http://www.re-cycledair.com/playing-with-go-channels-and-go-routines to play with channels and goroutines */

package examples

import (
	"os"
	"strconv"
	"fmt"
)

type fibonacci struct {
	a       int;
	b       int;
	current int;
	goal    int;
}

func compute() {
	//Check the command line arguments for sanity.
	goal, err := checkArgs();
	if err != "" {
		fmt.Fprintf(os.Stdout, "Error: %s", err)
		os.Exit(1)
	}

	//Special case:  If the goal is 2 or less, the Fibonacci number is always
	//1.  This isn’t really worth creating channels + go routines for.
	if goal <= 2 {
		fmt.Fprintf(os.Stdout, "Fibnonacci Value % d is: %d\n", goal, 1);
	} else {
		// Construct two channels, the input channel acts as our communication between the Go routines.
		// As you can see, it can be typed as an ADT(struct), which makes passing data REALLY easy between go routines.
		// The final channel is for passing back the goal Fibonacci number to the main thread.
		input := make(chan fibonacci);
		final := make(chan int);

		//Create a new fibonacci object and set it’s values.
		x := fibonacci{ 1, 1, 2, goal}

		// Create a new Go routine with the addNumber function. We pass it the two channels that were declared earlier
		go addNumber(input, final)

		// The first addNumber() Go routine will block until it gets input from the input channel. we pass it our object
		input <- x;

		//Now, we block and wait for the Go routines (addNumber()) to finish up and send us back the final value.
		number := <-final;
		fmt.Fprintf(os.Stdout, "Fibnonacci Value % d is: %d\n", x.goal, number);
	}
}

/*
* The addNumber() function takes two channels as parameters.
* One is a channel of fibonacci, which holds information about which number to calculate and when to stop,
* and then a channel of type int, which will be used to send the final goal number back to the main thread.
*/
func addNumber(input chan fibonacci, final chan int) () {
	//Block here and wait for input from the previous thread or the main thread, depending on the situation.
	x := <-input;

	// Here we check to see if we have met our goal, if we have, we return the goal value back to the main thread
	// via the final channel.  Otherwise, we calculate the next number in the sequence.
	if x.current < x.goal {

		fib := x.a + x.b

		// This is fun because you can assign multiple values at once. x.a <- x.b and x.b <- fib
		x.a, x.b = x.b, fib;
		x.current = x.current + 1;

		// Make a new channel of type fibonacci so that we can communicate with the new Go routine that we are about to create.
		output := make(chan fibonacci);

		//Pass the new Go routine the newly created channel, as well as the
		//final channel that was passed in from the caller.
		go addNumber(output, final);

		//Send the fibonacci object out on the channel.
		output <- x;
	} else {
		//Send the goal value out on the channel back to main.
		final <- x.b;
	}
	//Go routine dies now.
}

/* The checkArgs() function is a great example of returning multiple values.  We return an integer and a string
to the caller, so handling errors is fairly straight forward. */
func checkArgs() (int, string) {
	//Fetch the command line arguments.
	args := os.Args;

	//Check the length of the arguments
	if len(args) != 2  { return -1, "Invalid number of arguments.\n" }

	//Convert the goal argument to an integer.
	goal, err := strconv.Atoi(args[1]);

	//Make sure the conversion went correctly, otherwise return failure.
	if err != nil { return -1, "Invalid argument.Argument must be an integer.\n" }

	// Make sure the user can’t go beyond the program’s limits.
	if goal > 46 { return -1, "This program only calculates up to the 46th Fibonacci number.\n"; }

	//Check the lower bound as well.
	if goal < 1 { return -1, "Invalid range. Number must be >= 1.\n"; }

	//On success, return the goal value and an empty string indicating
	//that everything is good.
	return goal, ""
}
