package main

import (
	"bytes"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.Static("/css", "./public/css")
  //r.LoadHTMLGlob("templates/*.html")

  tmpl := make(map[string]*template.Template)

  funcMap := template.FuncMap{
    "call_template": func(name string) (ret string) {
      buf := bytes.NewBuffer([]byte{})
      err := tmpl["index.html"].ExecuteTemplate(buf, name, nil)
      if err != nil { panic(err) }
      ret = buf.String()
      return;
    },
	}

  tmpl["index.html"] = template.Must(
    template.New("index").
      Funcs(funcMap).
      ParseFiles("templates/index.html"))

  r.GET("/", func(c *gin.Context) {
    c.Status(http.StatusOK)
    tmpl["index.html"].ExecuteTemplate(c.Writer, "index", gin.H{
      "is_fullpage": true,
      "pagename"   : "home",
    })
  })

  r.GET("/home", func(c *gin.Context) {
    c.Status(http.StatusOK)
    if c.GetHeader("x-partial") == "true" {
      tmpl["index.html"].ExecuteTemplate(c.Writer, "home", nil)
    } else {
      tmpl["index.html"].ExecuteTemplate(c.Writer, "index", gin.H{
        "is_fullpage": true,
        "pagename"   : "home",
      })
    }
  })

  r.GET("/about", func(c *gin.Context) {
    c.Status(http.StatusOK)
    if c.GetHeader("x-partial") == "true" {
      tmpl["index.html"].ExecuteTemplate(c.Writer, "about", nil)
    } else {
      tmpl["index.html"].ExecuteTemplate(c.Writer, "index", gin.H{
        "is_fullpage": true,
        "pagename"   : "about",
      })
    }
  })

  r.GET("/contacts", func(c *gin.Context) {
    c.Status(http.StatusOK)
    if c.GetHeader("x-partial") == "true" {
      tmpl["index.html"].ExecuteTemplate(c.Writer, "contacts", nil)
    } else {
      tmpl["index.html"].ExecuteTemplate(c.Writer, "index", gin.H{
        "is_fullpage": true,
        "pagename"   : "contacts",
      })
    }
  })

  r.Run(":3000")
}