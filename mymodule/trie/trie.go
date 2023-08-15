package trie

import "fmt"

/*
	Implementation of a trie.

	- Need to make a constructor
	- Insert Operation
	- Search Operation

*/

type Node struct {
	// attributes: val (letter), list of children, and a bit that signifies if at end of a word
	val byte
	children []Node
	end_of_word bool
}


type Trie struct {
	// k-ary search tree used to store words that are easily searchable
	// initialize a root node as a list of children nodes
	root Node
	max_word_len int
}

/*
func Constructor() Trie {
	// instantiates a Trie



}
*/

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
	return false
}



func PrintDebug() {
	n1 := Node{
		val: 'b',
		children: []Node{},
		end_of_word: false,
	}

	n2 := Node{
		val: 'c',
		children: []Node{},
		end_of_word: false,
	}

	n3 := Node{
		val: 'd',
		children: []Node{},
		end_of_word: false,
	}

	n4 := Node{
		val: 'a',
		children: []Node{n1, n2, n3},
		end_of_word: false,
	}

	fmt.Println(n4.val, n4.children, n4.end_of_word)
}
