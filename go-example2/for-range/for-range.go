package main

func main() {
	names := []string{"John", "Jane", "Joe", "June"}

	for i, name := range names {
		println(i, name)
	}
}
