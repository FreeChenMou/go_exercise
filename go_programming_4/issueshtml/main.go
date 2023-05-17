// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"go_code/go_exercise/go_programming_4/issues/github/search"
	"log"
	"net/http"
)

//!+template
import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

//!-template

//exercise4.14 通过查询保存信息
func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form["q"]
	result, err := search.SearchIssues(url)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
