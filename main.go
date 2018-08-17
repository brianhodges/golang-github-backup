package main

import (
	"encoding/json"
	"fmt"
	"golang-github-backup/pkg/repository"
	"golang-github-backup/pkg/util"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const PAGE_CNT = "100"

func main() {
	
	fmt.Println()
	
	if len(os.Args) <= 1 {
		fmt.Println("You must enter a github user as a parameter.")
		fmt.Println("Ex. go run main.go brianhodges")
		os.Exit(1)
	}
	
	username := os.Args[1]
	fmt.Println(username)
	url := "https://api.github.com/users/" + username + "/repos?per_page=" + PAGE_CNT

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	util.Check(err)

	res, getErr := client.Do(req)
	util.Check(getErr)

	body, readErr := ioutil.ReadAll(res.Body)
	util.Check(readErr)

	repos := []repository.Repository{}
	jsonErr := json.Unmarshal(body, &repos)
	util.Check(jsonErr)

	os.RemoveAll("./Repos")

	for _, repo := range repos {
		fmt.Println("Cloning: " + repo.FullName + "...")
		_, err := git.PlainClone("./Repos/"+repo.Name, false, &git.CloneOptions{
			URL:               repo.Url,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		util.Check(err)
	}

	fmt.Println()
	fmt.Println("Done.")
	fmt.Println(strconv.Itoa(len(repos)) + " repositories cloned.")
}
