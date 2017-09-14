package main

import (
        "fmt"
        "io"
        "log"
        "net/http"
        "os"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
                os.Exit(1)
        }

        for 3 > 0 {
            response, err := http.Get(os.Args[1])
        }

        if err != nil {
                log.Fatal(err)
        } else {
                defer response.Body.Close()
                _, err := io.Copy(os.Stdout, response.Body)
                if err != nil {
                        log.Fatal(err)
                }
        }
}
