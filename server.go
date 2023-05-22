/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/microsoft/vscode-remote-try-go/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	params := hello.Params{
		Name: r.URL.Query().Get("name"),
		Ip:   r.RemoteAddr,
	}
	log.Print(params.Name, params.Ip)
	user := hello.Register(params)
	io.WriteString(w, hello.Hello(user))
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
