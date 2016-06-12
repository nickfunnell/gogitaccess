package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

//
//func main() {
//    repo, err := git.OpenRepository("/Users/nick/work/src/github.com/nickfunnell/testrepo")
//    repo.NewBranchIterator()
//    //fmt.Println(repo.)
//    if err != nil {
//        fmt.Fprintln(os.Stderr, err)
//        os.Exit(1)
//    }
//
//    if repo.IsBare() {
//        fmt.Println("Yep, it's bare.")
//    } else {
//        fmt.Println("Nope. Not a bare repo.")
//    }
//}
type GitRepo struct {
	RepoLocation string
}

type Tag struct {
	Name string
	Date time.Time
}

func (t Tag) String() {
	fmt.Sprintf("%s: %s", t.Name, t.Date.Format(time.RFC1123Z))
}

type ByDate []Tag

func (d ByDate) Len() int {
	return len(d)
}
func (d ByDate) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d ByDate) Less(i, j int) bool {
	return d[i].Date.Before(d[j].Date)
}

func (r GitRepo) fetchTagsFromGit() []Tag {
	cmd := exec.Command("git", "-C", r.RepoLocation, "for-each-ref", "--sort=taggerdate", "--format", "'%(tag)'")
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			splitValues := strings.Split(scanner.Text(), ",")
			fmt.Println(splitValues)
			tag := Tag{Name: splitValues[0], Date: parseTime(splitValues[1])}
			m[tag] = struct{}{}
		}
	}()

	cmd.Start()
	cmd.Wait()
	return m
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC1123Z, timeStr)
	return t
}
