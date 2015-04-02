### Todo app in [Go](http://golang.org) and [AngularJS](https://angularjs.org/)

#### Intention
To demonstrate how to develop Go web application along with Angularjs framework. In this demo, I extract Angularjs code from [TodoMVC](http://todomvc.com/), implement back-end code in [Go](www.golang.org).

#### How to run this app
* Run `go get github.com/pilu/fresh` to install the `fresh` commandline utility, with `fresh`, you don’t need to shutdown and restart the Go app under development. It watches all of files in your project directory, restart Go app whenever you save source file, LIFE SAVER.
* Run `go get` to install required libraries, you can find all of them in below `Resource` section.
* [Install NodeJS](https://github.com/joyent/node/wiki/Installation), then run `cd public && npm install` to install AngularJS and two of it's third-party modules: `angular-route` and `angular-mock`.
* Run `fresh`, then happy coding with your favorite editor.
* To make your life much happier as a front-end developer, you can install [LiveReload](http://livereload.com/), it costs you $9.99

#### Anything else
* Back-end code tests is missing, which is TOP priority for now!!
* Currently there is a workaround for the feature `Clear Completed Todos`, the todo item will to be deleted once at a time, it should be a performance hit for large dataset.
* Security hole. The parameters from front-end is passed to data CRUD without further processing. It’s ok for a demo project, but DO NOT DO THIS in production code.
* Dependency management. I haven’t find a good dependency management library/utility for now, [Glide](https://github.com/Masterminds/glide) is a good option, I’m trying it.
* Currently this project is implemented with a bunch of Go web libraries, we can implemente it again with Go web frameworks, such as [Martini](http://martini.codegangsta.io/) or [Beego](http://beego.me/), you could find a better one for your requirement.
* JS minification and static files caching, etc., although it's not imported for a demo project.

#### Resources:
* [TodoMVC](http://todomvc.com/) It contains all kinds of front-end frameworks/libraries implementation of Todo application.
* [negroni]("github.com/codegangsta/negroni") Go middleware library
* [gorilla/mux]("github.com/gorilla/mux") Go routing library
* [blackfriday]("github.com/russross/blackfriday") Generate html from Markdown
* [go-sqlite3]("github.com/mattn/go-sqlite3") Go sqlite3 DB driver
* [fresh](https://github.com/pilu/fresh) Auto loading go code
* [Go by Example](http://gobyexample.com/)
* [Make a RESTful JSON API in Go](http://thenewstack.io/make-a-restful-json-api-go/)
