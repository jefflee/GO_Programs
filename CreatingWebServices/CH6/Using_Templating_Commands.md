# Using Templating


## Demo: Templating Basics

```
package main

import (
	"html/template"
	"os"
)

type BlogPost struct {
	Title   string
	Content string
}

func main() {
	post := BlogPost{"First Post", "This is the blog post content section"}
	tmpl, err := template.New("blog-tmpl").Parse("<h1>{{.Title}}</h1><div><p>{{.Content}}</p></div>")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, post)
	if err != nil {
		panic(err)
	}
}

```
