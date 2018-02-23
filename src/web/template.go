package web

import (
    "net/http"
    "html/template"
    "log"
)

func Render(writer http.ResponseWriter, templateName string, templateString string, context interface{})  {
    t := template.New(templateName)
    tpl, err := t.Parse(templateString)
    if err != nil {
        log.Println("render error, template:", templateString)
        return
    }
    tpl.Execute(writer, context)
}
