package main

import (
  "github.com/ant0ine/go-json-rest/rest"
  "strings"
  "log"
  "fmt"
  "os"
  "os/exec"
  "net/http"
)

func main() {
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {

    conf := os.Getenv("CONFIG")
    app  := os.Getenv("APP")

    w.Header().Set("Content-Type", "text/plain")
    cmd := exec.Command("transporter", "run", "--config", conf, app)
    outs, err := cmd.CombinedOutput()

    out_cmds := fmt.Sprintf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
    w.(http.ResponseWriter).Write([]byte( out_cmds ))

    if err != nil {
      out_errs := fmt.Sprintf("==> Error: %s\n", err.Error())
      w.(http.ResponseWriter).Write([]byte( out_errs ))
    } else {
      out_outs := fmt.Sprintf("==> Output: %s\n", string(outs))
      w.(http.ResponseWriter).Write([]byte( out_outs ))
    }

  }))
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
