package main

import "fmt"

// type aliases
type (
	Name          string
	Breed         int
	Gender        int
	NameToDogFunc func(Name) Dog
)

// define possible breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

// define possible genders
const (
	Male Gender = iota
	Female
)

var (
	maleHavaneseSpawner = DogSpawner(Havanese, Male)
	femalePoodleSpawner = DogSpawner(Poodle, Female)
)

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

func DogSpawnerCurry(breed Breed) func(Gender) NameToDogFunc {
	return func(gender Gender) NameToDogFunc {
		return func(name Name) Dog {
			return Dog{
				Breed:  breed,
				Gender: gender,
				Name:   name,
			}
		}
	}
}

func DogSpawner(breed Breed, gender Gender) NameToDogFunc {
	return func(n Name) Dog {
		return Dog{
			Breed:  breed,
			Gender: gender,
			Name:   n,
		}
	}
}

func main() {
	bucky := DogSpawnerCurry(Havanese)(Male)("Bucky")
	fmt.Println(bucky)

	havaneseSpawner := DogSpawnerCurry(Havanese)
	rocky := havaneseSpawner(Male)("Rocky")
	fmt.Println(rocky)

	femaleHavanese := havaneseSpawner(Female)
	lola := femaleHavanese("Lola")
	dotty := femaleHavanese("Dotty")
	fmt.Println(lola)
	fmt.Println(dotty)

	rocky = maleHavaneseSpawner("rocky")
	tipsy := femalePoodleSpawner("tipsy")
	fmt.Println(rocky)
	fmt.Println(tipsy)

}

func createDogsWithoutPartialApplication() {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}

	rocky := Dog{
		Name:   "Rocky",
		Breed:  Havanese,
		Gender: Male,
	}

	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}
	_ = bucky
	_ = rocky
	_ = tipsy
}
