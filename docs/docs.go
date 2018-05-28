// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-05-28 14:50:22.533799409 -0700 PDT m=+0.070776123

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Health check",
        "title": "Blocks api",
        "contact": {
            "name": "Ivan Santos",
            "url": "https://ivansantos.me",
            "email": "pingme@ivansantos.me"
        },
        "license": {},
        "version": "1.0"
    },
    "host": "localhost:3000",
    "basePath": "/",
    "paths": {
        "/_health": {
            "get": {
                "description": "Health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Health check"
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
