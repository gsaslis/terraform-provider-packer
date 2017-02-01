package main

import (
	"bytes"
	"strings"
	"strconv"
	"fmt"
)

var newline []byte = []byte("\n")
var comma []byte = []byte(",")
const colon string = ":"

type Artifact struct {
	Builder string
	Region string
	Id     string
}

func (a Artifact) String() string {
	return "{builder:"+a.Builder+" region:"+a.Region +" id:"+a.Id +"}"
}

func ParseLines(output []byte) map[string][]Artifact {
	lines := bytes.Split(output, newline)

	artifacts := make(map[string][]Artifact)
	totalArtifacts := 0

	for _, line := range lines {
		words := bytes.Split(line, comma)
		commandType := strings.ToLower(string(words[2]))
		builderName := string(words[1])

		switch commandType {
		case "artifact-count":
			totalArtifacts, _ = strconv.Atoi(string(words[3]))
			artifacts[builderName] = make([]Artifact, totalArtifacts)
			fmt.Println("Found artifacts cound", totalArtifacts)
		case "artifact":
			artifactId, _ := strconv.Atoi(string(words[3]))
			subtype := strings.ToLower(string(words[4]))

			switch subtype {
			case "end":
				fmt.Println("finished artifact", artifacts[builderName][artifactId])
			case "id":
				split := strings.Split(string(words[5]), colon)
				artifacts[builderName][artifactId] = Artifact{Builder: builderName, Region:split[0], Id:split[1]}
				fmt.Println("found artifact", artifacts[builderName][artifactId])
			}
		}
	}

	fmt.Println("returning artifacts", artifacts)
	return artifacts
}
