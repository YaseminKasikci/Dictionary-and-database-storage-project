package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/dictionary/dictionary"
)

func main() {
	//process command lines (CLI)
	action := flag.String("action", "list", "Action to perform on the dictionary")
// creation of the badger folder
 d, err := dictionary.New("./badger")
 handleErr(err)
 defer d.Close()

 // stock in action
 flag.Parse()
 // word to use for CLI & their action
 switch *action {
 case "list":
	actionList(d)
 case "add":
	actionAdd(d, flag.Args())
case "define":
 actionDefine(d, flag.Args())
case "remove":
 actionRemove(d, flag.Args())
 default:
	fmt.Printf("Unknown action: %v\n", *action)
 }
}

// list  the dictionary
func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionnary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

// add a word in dictionary
func actionAdd(d *dictionary.Dictionary, args []string) {
word := args[0]
definition := args[1]
err := d.Add(word, definition)
handleErr(err)
fmt.Printf("'%v' added to the dictionary\n", word)
}

//define a word in dictionary
func actionDefine(d *dictionary.Dictionary, args []string) {
word := args[0]
entry, err := d.Get(word)
handleErr(err)
fmt.Println(entry)
}

//delete a word in dictionary
func actionRemove(d *dictionary.Dictionary, args []string) {
word := args[0]
err := d.Remove(word)
handleErr(err)
fmt.Printf("'%v' was removed from the dictionary\n", word)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}