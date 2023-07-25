package trie

import "fmt"

/*
	Implementation of a trie.

	- Need to make a constructor
	- Insert Operation
	- Search Operation

*/

type Node struct {
	// attributes: value (a letter), and list of children
}


type Trie struct {
	// k-ary search tree used to store words that are easily searchable
	// initialize a root node as a list of children nodes
	
}


func Constructor() Trie {
	// instantiates a Trie

}


func (this *Trie) Insert(word string) {
	/*
	params: 
		- this: the Trie data structure
		- word: the string to be inserted
	*/
}


func (this *Trie) Search(word string) bool {
	/*
	params: 
		- this: the Trie data structure
		- word: the string to be searched

	* determines whether or not a string is present in the Trie
	*/
}



func PrintHello() {
	fmt.Println("Hello, World! This is my package speaking!")	
}
