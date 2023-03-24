package main

// You cannot call the main directly because it creates a wrong number of arguments
func ExampleMain() {
	main()
	// Output:
	// Hi there!
}

func ExampleSayHi_name() {
	SayHi()
	SayHi("Rob")
	SayHi("Doris")

	// Output:
	// Hi there!
	// Hi Rob!
	// Hi Doris!

}
