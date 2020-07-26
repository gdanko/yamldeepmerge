package main

import (
	"fmt"
	"io/ioutil"
	"yamldeepmerge/yamldeepmerge"
)

func readFile(filename string) ([]byte) {
	contents, _ := ioutil.ReadFile(filename)
	return contents
}

func main() {
	baseFile := ("/Users/gdanko/go/src/merge/nginx_base.yml")
	envFile  := ("/Users/gdanko/go/src/merge/nginx_env.yml")

	baseYaml := readFile(baseFile)
	envYaml := readFile(envFile)
	derivedYaml := deep_merge.DeepMergeMapSliceOut(baseYaml, envYaml)
	//fmt.Println(string(derivedYaml))
	fmt.Println(derivedYaml)
}
