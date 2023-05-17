// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"go_code/go_exercise/go_programming_4/issues/github"
	"go_code/go_exercise/go_programming_4/issues/github/search"
	"log"
	"os"
	"time"
)

func main() {
	searchHandle()
}

//exercise4.10 通过命令行创建、读取、更新或者关闭github的issues
func searchHandle() {
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)
	monthAgo := now.AddDate(0, -1, 0)

	//三个切片，用来存储 不足一个月的问题，不足一年的问题，超过一年的问题。
	yearAgos := make([]*github.Issue, 0)
	monthAgos := make([]*github.Issue, 0)
	lessMonths := make([]*github.Issue, 0)

	result, err := search.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		//如果 yearAgo 比 创建时间晚，说明超过一年
		if yearAgo.After(item.CreatedAt) {
			yearAgos = append(yearAgos, item)
			//如果 monthAgo 比 创建时间晚，说明超过一月 不足一年
		} else if monthAgo.After(item.CreatedAt) {
			monthAgos = append(monthAgos, item)
			//如果 monthAgo 比 创建时间早，说明不足一月。
		} else if monthAgo.Before(item.CreatedAt) {
			lessMonths = append(lessMonths, item)
		}
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	fmt.Println("一年前")
	for _, item := range yearAgos {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)

	}
	fmt.Println("一年以内")
	for _, item := range monthAgos {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)

	}
	fmt.Println("一个月内")
	for _, item := range lessMonths {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func DownLoad() {

}
