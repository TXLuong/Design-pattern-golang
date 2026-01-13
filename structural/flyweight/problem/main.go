package main

import "fmt"

// Character is supposed to be a Flyweight
type Character struct {
	char     rune
	font     string
	size     int
	color    string
	x, y     int // ❌ Extrinsic state stored inside flyweight
}

// Factory that should reuse characters
type CharacterFactory struct {
	characters map[rune]*Character
}

// ❌ Factory creates new objects instead of reusing them
func (f *CharacterFactory) GetCharacter(
	char rune,
	font string,
	size int,
	color string,
	x int,
	y int,
) *Character {

	// ❌ No check if character already exists
	c := &Character{
		char:  char,
		font:  font,
		size:  size,
		color: color,
		x:     x,
		y:     y,
	}

	// ❌ Overwrites existing flyweights
	f.characters[char] = c

	return c
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{
		characters: make(map[rune]*Character),
	}
}

func main() {
	factory := NewCharacterFactory()

	c1 := factory.GetCharacter('A', "Arial", 12, "Black", 10, 10)
	c2 := factory.GetCharacter('A', "Arial", 12, "Black", 20, 20)

	fmt.Println(c1 == c2) // ❌ false, should be true
}
