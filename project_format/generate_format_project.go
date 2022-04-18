package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TotalProjectFormat struct {
	ProjectName string           `json:"project_name"`
	NextLevel   []*ProjectFormat `json:"next_level"`
}

type ProjectFormat struct {
	CurrentLevelName []string         `json:"level_name"`
	NextLevel        []*ProjectFormat `json:"next_level"`
}

var temp = "{\"project_name\":\"TestProject\",\"next_level\":[{\"level_name\":[\"resource\"],\"next_level\":[{\"level_name\":[\"sql\",\"script\"]}]},{\"level_name\":[\"config\"],\"next_level\":[{\"level_name\":[\"debug\"],\"next_level\":[{\"level_name\":[\"config.json\"]}]},{\"level_name\":[\"release\"],\"next_level\":[{\"level_name\":[\"config.json\"]}]}]},{\"level_name\":[\"main\"],\"next_level\":[{\"level_name\":[\"main.go\",\"service\",\"manager\",\"entity\",\"util\",\"biz\"],\"next_level\":[{\"level_name\":[\"repo\"],\"next_level\":[{\"level_name\":[\"cao\",\"dao\"]}]}]}]},{\"level_name\":[\"tools\"]}]}"

var filePath = flag.String("f", "", "input the file value.")
var projectName = flag.String("p", "SeaProjectDefault", "input the project name.")
var outPath = flag.String("o", getCurrentPath(), "input the output file path ")

func main() {
	flag.Parse()
	baseProject := new(TotalProjectFormat)

	var defaultFormat []byte
	var baseOutPath string

	if filePath == nil || *filePath == "" {
		defaultFormat = []byte(temp)
	} else {
		defaultFormat = ReadInfoFrom(*filePath)
	}
	err := json.Unmarshal(defaultFormat, baseProject)
	if err != nil {
		panic(err)
	}

	if outPath == nil || *outPath == "" {
		baseOutPath = getCurrentPath()
	} else {
		baseOutPath = *outPath
	}
	GenerateTheProjectFormat(baseOutPath, baseProject, *projectName)
}

func ReadInfoFrom(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return file
}

func GenerateTheProjectFormat(basePath string, format *TotalProjectFormat, defaultProjectName string) {
	if format == nil {
		log.Println(format)
		panic("current TotalProjectFormat has no info")
	}
	if defaultProjectName != "" {
		format.ProjectName = defaultProjectName
	}
	CreateDirector(basePath, format.ProjectName)
	for _, v := range format.NextLevel {
		GenerateTheDirectory(basePath+"/"+format.ProjectName, v)
	}
}

func GenerateTheDirectory(basePath string, format *ProjectFormat) {
	if format == nil || format.CurrentLevelName == nil || len(format.CurrentLevelName) == 0 {
		log.Println(format)
		panic("current has no info")
	}
	CreateDirector(basePath, format.CurrentLevelName...)
	for _, v := range format.NextLevel {
		if len(format.CurrentLevelName) > 1 {
			GenerateTheDirectory(basePath, v)
		} else {
			GenerateTheDirectory(basePath+"/"+format.CurrentLevelName[0], v)
		}
	}
}

func CreateDirector(basePath string, val ...string) {
	if val == nil {
		log.Println("current val is nil. no directory has been create")
	}
	for _, v := range val {
		filePath := strings.Join([]string{basePath, v}, "/")
		if fileInfo, err := os.Stat(filePath); err != nil {
			if os.IsNotExist(err) {
				if strings.Contains(v, ".go") {
					_, fileCreateErr := os.Create(filePath)
					if fileCreateErr != nil {
						log.Printf("file[%s] create has err[%v]\n", v, fileCreateErr)
					}
					continue
				}
				createErr := os.Mkdir(filePath, os.ModePerm)
				if createErr != nil {
					log.Printf("dir[%s] create has err[%v]\n", v, createErr)
				}
			} else {
				log.Printf("dir[%s] info read has err[%v]\n", v, err)
			}
		} else {
			if fileInfo.IsDir() {
				log.Printf("dir[%s] has exist\n", v)
			}
		}
	}
}

func getCurrentPath() string {
	val, err := os.Executable()
	if err != nil {
		panic(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(val))
	return res
}
