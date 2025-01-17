package main

import(
	"fmt"
	"bufio"
	"os"
	"slices"
)

var a []string

func addItem(){
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("What do you need to do? ")
	scanner.Scan()
	line := scanner.Text()

	if slices.Index(a, line) == -1{
		a = append(a, line)
		fmt.Println("Added", line)
		fmt.Println("Your list now has", len(a), "items!")
	} else {
		fmt.Println(line, "is already in your todo list!")
	}
}

func printList(){
	fmt.Println("Here are the items on your todo list:")
	for _, item := range a{
		fmt.Println(slices.Index(a, item)+1, ":", item)
	}
}

func completeItem(){
	printList()
	fmt.Print("Which item did you complete? (Please input a number): ")

	var input int
	fmt.Scanln(&input)

	if input > len(a) || input <= 0{
		fmt.Println("Sorry, item number", input, "was not found in your todo list.")
	} else {
		slices.Replace(a, input-1, input, a[input-1]+" (Completed!)") 
	}
}

func deleteItem(){
	printList()
	fmt.Print("What would you like to delete? (Please input a number): ")

	var input int
	fmt.Scanln(&input)

	if input > len(a) || input <= 0{
		fmt.Println("Sorry, item number", input, "was not found in your todo list.")
	} else {
		slices.Delete(a, input-1, input) 
	}

	a = a[0:len(a)-1]

	fmt.Println("Item number", input, "was removed from your todo list.")
}

func helpMenu(){
	fmt.Println("1. Display List:")
	fmt.Println("----------------")
	fmt.Println("Displays all tasks currently on your todo list.")
	fmt.Println("")
	fmt.Println("2. Add Item:")
	fmt.Println("------------")
	fmt.Println("Allows you to add an item to your list.")
	fmt.Println("")
	fmt.Println("3. Mark Item Complete:")
	fmt.Println("----------------------")
	fmt.Println("Marks an item in your list as ''(Completed!)''")
	fmt.Println("")
	fmt.Println("4. Delete Item:")
	fmt.Println("---------------")
	fmt.Println("Removes an item from the list, regardless of completion status.")
	fmt.Println("")
	fmt.Println("5. Display Help Menu:")
	fmt.Println("---------------------")
	fmt.Println("Explains the function of all menu options.")
	fmt.Println("")
	fmt.Println("6. Exit Program:")
	fmt.Println("----------------")
	fmt.Println("Closes the program. Your todo list will not be saved!")
}

func menu() int{
	fmt.Println("")
	fmt.Println("Todo List Menu:")
	fmt.Println("---------------")
	fmt.Println("1. Display List")
	fmt.Println("2. Add Item")
	fmt.Println("3. Mark Item Complete")
	fmt.Println("4. Delete Item")
	fmt.Println("5. Display Help Menu")
	fmt.Println("6. Exit Program")
	fmt.Println("")

	fmt.Print("What would you like to do? (Input the number before the command to execute that command.) ")
	
	var input int
	fmt.Scanln(&input)
	
	// validate input
	for input < 1 || input > 6 {
		fmt.Print("Sorry, ", input, " doesn't coorespond with a menu command. Please try again: ")
		fmt.Scanln(&input)
	}
	return input
}

func main(){
	var m int
	fmt.Println("Welcome to your todo list!")
	// for loop to prevent program from closing after early
	for m != 6 {
		m = menu()
		switch m {
		case 1:
			printList()
		case 2:
			addItem()
		case 3:
			completeItem()
		case 4:
			deleteItem()
		case 5:
			helpMenu()
		case 6:
			//empty statement to prevent error message from appearing when exiting
		default:
			fmt.Println("This is an error message! You shouldn't be seeing this!")
		}
	}
	fmt.Println("Goodbye!")
}