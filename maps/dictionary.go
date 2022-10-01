package maps

type Dictionary map[string]string

type DictionaryErr string

const ErrNotFound DictionaryErr = "could not find the word you were looking for"
const ErrAlreadyExists DictionaryErr = "cannot add word as it already exists"
const ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
const ErrCannotDelete = DictionaryErr("cannot delete word as it does not exist")

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
}

func (d Dictionary) Update(word string, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDefinition
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrCannotDelete
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}
