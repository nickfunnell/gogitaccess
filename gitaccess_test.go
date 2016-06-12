package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortTags(t *testing.T) {
	tag1 := Tag{Name: "0.0.1", Date: parseTime("Tue, 24 May 2016 13:09:00 +0100")}
	tag2 := Tag{Name: "0.0.3", Date: parseTime("Tue, 24 May 2016 14:14:00 +0100")}
	tag3 := Tag{Name: "0.0.2", Date: parseTime("Tue, 24 May 2016 14:09:00 +0100")}
	tags := []Tag{tag1, tag2, tag3}
	//Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone

	dateOrderedTags := []Tag{tag1, tag3, tag2}

	fmt.Println(tags)

	sort.Sort(ByDate(tags))
	for i, tag := range tags {
		if tag != dateOrderedTags[i] {
			t.Errorf("tags don't match")
		}
	}
	fmt.Println(tags)

}

func TestFetchTagsFromGit(t *testing.T) {
	repo := GitRepo{RepoLocation: "/Users/nick/work/src/github.com/nickfunnell/testrepo"}
	fmt.Println(repo.fetchTagsFromGit())
}
