package main

import (
	"log"
	"os"
	"text/template" //Need to make sure it's NOT html/template...you use that later.
)

/* This is for the 'initial parse', to parse all our templates at once.
After this, you would just need to execute them all in Main! */
/*
var templateUltimate *template.Template

func init() {
	templateUltimate = template.Must(template.ParseGlob("templates/*"))
}
*/

/* main func */
func main() {

	/* This is for our initial parse glob above...it's the most efficent way to read files into a template */

	/* Give the template the name of the file, (OR MANY FILES), to parse. It returns the pointer to
	a template and an error. This will have ALL the templates you parsed in a container.*/
	template, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	/* Here we are taking that pointer above and executing it to standard out,
	( with no data passed in, that's nil) */
	err = template.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	/* You can also passed created files into this temlate and execute them afterwards */
	newFile, err2 := os.Create("index.html")
	if err2 != nil {
		log.Fatalln(err2)
	}
	defer newFile.Close()

	err2 = template.Execute(newFile, nil)
	if err2 != nil {
		log.Fatalln(err2)
	}

	/* We already have that template pointer created, why not parse more files into that container?
	Just make sure you execute the specific template in that pointer! (vespa or the other ones...) */
	template, err2 = template.ParseFiles("two.gmao", "vespa.gmao")
	if err2 != nil {
		log.Fatalln(err2)
	}

	err2 = template.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err2 != nil {
		log.Fatalln(err2)
	}
	err2 = template.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err2 != nil {
		log.Fatalln(err2)
	}

	/* You can also use ParseGlob,(which Todd argues is neater. It gets all the file extensions in the folder) */
	/*
		template2, err3 := template.ParseGlob("templates/*")
		if err != nil {
			log.Fatalln(err3)
		}

		err3 = template2.Execute(os.Stdout, nil)
		if err3 != nil {
			log.Fatalln(err3)
		}

		err3 = template2.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}

		err3 = template2.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
		if err3 != nil {
			log.Fatalln(err3)
		}

		err3 = template2.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
		if err3 != nil {
			log.Fatalln(err3)
		}
	*/

	/* We can also use this to make sure all of the files are parsed effeciently in our folder
	by calling this function. You just want to execute the templates once!*/

}
