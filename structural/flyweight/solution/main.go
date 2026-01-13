package main

import "fmt"

type CharacterFlyweight interface {
	Draw(x, y int)
}

type Character struct {
	char  rune
	font  string
	size  int
	color string
}

func (c *Character) Draw(x, y int) {
	fmt.Printf(
		"Drawing '%c' at (%d,%d) with font=%s size=%d color=%s\n",
		c.char, x, y, c.font, c.size, c.color,
	)
}

type CharacterFactory struct {
	cache map[string]CharacterFlyweight
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{
		cache: make(map[string]CharacterFlyweight),
	}
}

func (f *CharacterFactory) GetCharacter(
	char rune,
	font string,
	size int,
	color string,
) CharacterFlyweight {

	key := fmt.Sprintf("%c-%s-%d-%s", char, font, size, color)

	if flyweight, ok := f.cache[key]; ok {
		return flyweight
	}

	flyweight := &Character{
		char:  char,
		font:  font,
		size:  size,
		color: color,
	}

	f.cache[key] = flyweight
	return flyweight
}

func main() {
	factory := NewCharacterFactory()

	c1 := factory.GetCharacter('A', "Arial", 12, "Black")
	c2 := factory.GetCharacter('A', "Arial", 12, "Black")
	c3 := factory.GetCharacter('A', "Times", 12, "Black")

	fmt.Println(c1 == c2) // ✅ true (shared flyweight)
	fmt.Println(c1 == c3) // ❌ false (different intrinsic state)

	c1.Draw(10, 10)
	c2.Draw(20, 20)
	c3.Draw(30, 30)
}
