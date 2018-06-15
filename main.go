package main

import (
	"context"
	"fmt"
	"sort"
	"os"
	"log"
	"bufio"
	"strings"
	
	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

// LatestVersions returns a sorted slice with the √ version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run

	//1. Group by each Major.Minor, Remove smaller minor versions

	var versionGroup string

	HighestVersionMap := make(map[string]*semver.Version)

	semver.Sort(releases)

	for _,release := range releases {
		//Remove smaller minor versions
		if release.LessThan(*minVersion) {
			continue
		}

		//Group by each Major.Minor
		versionGroup=fmt.Sprintf("%d",release.Major)+"."+fmt.Sprintf("%d",release.Minor)

		//Find highest version on each group 
		HighestVersionMap[versionGroup]=release
	 }

	//2. Append Highest versions to versionSlice
	 for _,version := range HighestVersionMap {
		versionSlice=append(versionSlice,version)
	 }

	//3. Sort versionSlice in a descending order
	sort.Slice(versionSlice, func(i, j int) bool {
		return versionSlice[j].LessThan(*versionSlice[i])
	})

	return versionSlice
}

func printLastestVersion(account string, repoName string, minVer string){
	// Github
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.ListOptions{PerPage: 10}
	releases, _, err := client.Repositories.ListReleases(ctx, account, repoName, opt)
	if err != nil {
		//panic(err) // is this really a good way?
		log.Fatal(err)
	}
	minVersion := semver.New(minVer)	

	allReleases := make([]*semver.Version, len(releases))
	for i, release := range releases {

		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}

		allReleases[i] = semver.New(versionString)
	}

	versionSlice := LatestVersions(allReleases, minVersion)

	fmt.Printf("latest versions of %s/%s: %s \n", account, repoName, versionSlice)
}


// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {

	var fileName string = os.Args[1]
	file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	var lineCount int =0
	var line string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

		line=scanner.Text()

		//without first line
		if(lineCount !=0){
			splitLine := strings.Split(line, ",")
			repo,minVer := splitLine[0], splitLine[1]
	
			splitRepo := strings.Split(repo, "/")
			account,repoName := splitRepo[0], splitRepo[1]

			printLastestVersion(account,repoName,minVer)
		}

		lineCount=lineCount+1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


}
