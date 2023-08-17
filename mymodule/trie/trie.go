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


func (this *Node) IsPresent(val byte) *Node {
	// checks if a value is already present in the children of a node
	// if the value is present --> return a REFERENCE to the child node
	// if the value is absent  --> return nil

	length := len(this.children)
	if length == 0 {
		return nil
	}
	i := 0
	for i < length {
		if this.children[i].val == val {
			return &this.children[i]
		}
		i+=1
	} 

	return nil
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
	if length > this.max_word_len {
		this.max_word_len = length
	}
	cur := &this.root
	var end_of_word bool


	i := 0
	for i < length {
		
		if i == length - 1 {
			end_of_word = true
		} else {
			end_of_word = false
		}
		
		// check if the letter is already present 
		present := cur.IsPresent(word[i])
		if present != nil {
			// need to think about the case where this is the end of the word
			// advance to the correct child node
			cur = present

			if end_of_word {
				// change end_of_word attribute to true
				cur.end_of_word = true
			}
			i+=1 // advance to next iteration
		}
		// if no children for cur --> just append the letter
		if (len(cur.children) == 0) || present == nil {
			cur.children =  append(cur.children, 
				Node{
					val: word[i],
					children: []Node{},
					end_of_word: end_of_word, 
				},
			)
		}
		cur = &(cur.children[0])
		// fmt.Println(i)
		i+=1
	}
	




	// remember to update the 'end_of_word' attribute on final letter of word

}


func (this *Trie) Search(word string) bool {
	/*
	params: 
		- this: a Trie data structure
		- word: the string to be searched

	* determines whether or not a string is present in the Trie
		* true - present
		* false - absent
	*/

	// first check if the length of the word exceeds the max len word in the trie
	if len(word) > this.max_word_len {
		return false
	}



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
	//t.Insert("work")

	fmt.Println(t.root.children[0].IsPresent('r'))

	fmt.Println(t)
}
