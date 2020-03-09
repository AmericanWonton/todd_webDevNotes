package main

import (
	"log"
	"os"
	"text/template"
)

/* Here is our templates in one neat place to execute later in main */
var testTemplate *template.Template

func init() {
	testTemplate = template.Must(template.ParseFiles("tpl.gohtml"))
}

/* Here are some structs we can pass into our html */
type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func main() {
	/* You can only pass in one piece of data into this...but we can make a big ol' struct to pass a lot.
	In this example, the passed in value is both of those, (42), but we can use a struct to pass in
	a BUNCH of stuff.*/
	/* Test 1 */
	/*
		err := testTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42) //42 is our passed in data.
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/* Test 2: passing in slice of string */
	/*
		sages := []string{"ghandi", "MLK", "Buddha", "Jesus"}
		err := testTemplate.Execute(os.Stdout, sages)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/* Test 3: Passing in slice for index/element */
	/*
		sages := []string{"ghandi", "MLK", "Buddha", "Jesus"}
		err := testTemplate.Execute(os.Stdout, sages)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/* Test 4: Using a map */
	/*
		sages := map[string]string{
			"India":     "Ghandi",
			"America":   "MLK",
			"Mediatate": "Buddha",
			"Singing":   "Michael Jackson",
		}
		err := testTemplate.Execute(os.Stdout, sages)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/* Test 5: Passing in a struct */
	/*
		johnAppleBoy := sage{
			Name:  "JohnnyAppleBoy",
			Motto: "Get this bread and get them apples.",
		}
		err := testTemplate.Execute(os.Stdout, johnAppleBoy)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/* Test 6: Passing in a slice of struct */
	/*
		buddha := sage{
			Name:  "Buddha",
			Motto: "The belief of no beliefs",
		}

		gandhi := sage{
			Name:  "Gandhi",
			Motto: "Be the change",
		}

		mlk := sage{
			Name:  "Martin Luther King",
			Motto: "Hatred never ceases with hatred but with love alone is healed.",
		}

		jesus := sage{
			Name:  "Jesus",
			Motto: "Love all",
		}

		muhammad := sage{
			Name:  "Muhammad",
			Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
		}

		sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

		err := testTemplate.Execute(os.Stdout, sages)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/* Test 7: Passing a slice of Structs */
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := testTemplate.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
