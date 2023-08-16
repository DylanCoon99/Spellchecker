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


func Constructor() *Trie {
	// instantiates a Trie
	root := Node {
		val: ' ',
		children: []Node{},
		end_of_word: false,
	}

	trie := Trie {
		root: root,
		max_word_len: 0,
	}
	return &trie
}


func (this *Trie) Insert(word string) {
	/*
	params: 
		- this: a Trie data structure
		- word: the string to be inserted
	*/
	length := len(word)
	cur := &this.root
	var end_of_word bool

	// if the trie is empty
	if len(this.root.children) == 0 {
		i := 0
		for i < length {
			
			if i == length - 1 {
				end_of_word = true
			} else {
				end_of_word = false
			}
			
			cur.children =  append(cur.children, 
				Node{
					val: word[i],
					children: []Node{},
					end_of_word: end_of_word, 
				},
			)
			cur = &(cur.children[0])
			// fmt.Println(i)
			i += 1
		}
	}

	// if the trie is not empty
	


	// remember to update the 'end_of_word' attribute on final letter of word

}


func (this *Trie) Search(word string) bool {
	/*
	params: 
		- this: a Trie data structure
		- word: the string to be searched

	* determines whether or not a string is present in the Trie
	*/



	return false
}



func PrintDebug() {
	/*
	n1 := Node {
		val: 'b',
		children: []Node{},
		end_of_word: false,
	}

	n2 := Node {
		val: 'c',
		children: []Node{},
		end_of_word: false,
	}

	n3 := Node {
		val: 'd',
		children: []Node{},
		end_of_word: false,
	}

	n4 := Node {
		val: 'a',
		children: []Node{n1, n2, n3},
		end_of_word: false,
	}
	*/

	t := Constructor()
	/*
	t.root.children = append(t.root.children, Node{
					val: 'a',
					children: []Node{},
					end_of_word: false})
	*/
	fmt.Println(len(t.root.children))
	t.Insert("word")

	fmt.Println(t)
}
