bsprite-go
==========

A go package for creating bsprites

##Usage

Example of creating sprites from all .jpg within a folder and then serving them on a URL `/images`

  package main
  
  import (
  	"net/http"
  	"log"
  	"github.com/mtharrison/bsprite-go"
  )
  
  func main() {
  
  	err, sprite := bsprite.Make("./images/*.jpg")
  
  	if err != nil {
  		log.Fatal(err)
  	}
  
  	http.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
  
  		for i, v := range sprite.Headers() {
  			w.Header().Set(i, v)	
  		}
  
  		w.Header().Set("Content-type", "application/octet-stream")
  
  		w.Write(sprite.Body())
  	})
  
  	log.Fatal(http.ListenAndServe(":8080", nil))
  
  }
