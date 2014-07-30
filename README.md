bsprite-go
==========

A go package for creating bsprites

##Public methods

Generates a bsprite object from one or more glob patterns

    func Make(globs ...string) (err error, sprite Sprite)
    
Returns the combined generated sprite data   

    func (sprite Sprite) Body() []byte

Returns the headers to be served with the sprite  
    
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
