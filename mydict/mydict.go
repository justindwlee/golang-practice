package mydict

import (
	"errors"
)

//Dictionary type
type Dictionary map[string]string

var (
	errWordNotFound = errors.New("word not found")
	errWordExists = errors.New("word already exists")
	errCantUpdate = errors.New("cannot update non-existing word")
)

//search for a word
func (d Dictionary) Search(word string) (string, error){
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "",errWordNotFound
}

//Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errWordNotFound {
		d[word] = def;
	} else if err == nil {
		return errWordExists
	}
	// switch err {
	// case errWordNotFound:
	// 	d[word] = def;
	// case nil:
	// 	return errWordExists
	// }
	return nil;
}

//Update a word in the dictionary
func (d Dictionary) Update(word, newDef string) error {
	_, err := d.Search(word)
	if err == errWordNotFound {
		return errCantUpdate
	} else if err == nil {
		d[word] = newDef
	}
	return nil
}

//Delete a word in the dictionary 
func (d Dictionary) Delete(word string) {
	delete(d, word)
}