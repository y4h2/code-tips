package builder

import "fmt"

func ExampleActorBuilder() {
	ab := NewAngelBuilder()

	actor := ConstructActor(ab)

	fmt.Println(actor.GetSex())
	fmt.Println(actor.GetFace())
	fmt.Println(actor.GetCostume())
	fmt.Println(actor.GetHairstyle())

	// Output:
	// female
	// beautiful
	// shorts
	// golden
}
