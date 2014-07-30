bsprite-go
==========

**Experimental**

[See original post](matt-harrison.com/bsprites-a-new-way-of-serving-combined-resources-with-arraybuffer-and-data-uris/)

A go package for creating bsprites (binary sprites). BSprites are a combined set of web resources (anything that accepts a data URI) transferred over a single HTTP request which contain all the metadata required to parse them into the separate resources on the client side. They only work in browsers that support Data URIs and Typed arrays. You can mix and match different mime types in the same sprite 'package', e.g. svg, jpg, mp3 can be combined into a single named resource package.

##Public methods

**ServeHTTP()** Satisfies the http.Handler interface so a Sprite object can be passed
into the `http` package as a handler for a route

    func (sprite Sprite) ServeHTTP(w http.ResponseWriter, r *http.Request)

**Make()** Generates a bsprite object from one or more glob patterns

    func Make(globs ...string) (err error, sprite Sprite)
    
**Body()** Returns the combined generated sprite data   

    func (sprite Sprite) Body() []byte

**Headers()** Returns the headers to be served with the sprite  
    
    func (sprite Sprite) Headers() map[string]string

##Usage

Example of creating sprites from all `.jpg`s within an `img` folder and then serving them on a URL `/images`

    package main

    import (
        "net/http"
        "log"
        "github.com/mtharrison/bsprite-go"
    )

    func main() {

        err, sprite := bsprite.Make("./img/*.jpg")

        if err != nil {
            log.Fatal(err)
        }

        http.Handle("/images", sprite)

        log.Fatal(http.ListenAndServe(":8080", nil))

    }
