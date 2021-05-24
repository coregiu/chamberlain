package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func Test_should_successfully_when_load_yaml_file(t *testing.T) {
	workPath, _ := os.Getwd()
	defer deleteYamlFileOfCurrentDir(workPath)
	copyYamlFileToCurrentDir(workPath)

	testConfig := new(ChamberlainConfig)
	result := testConfig.loadConfigFromFile()
	if result != nil {
		t.Error(result.Error())
	}

	if strings.Compare(testConfig.BlogWorkPath, "/giu/chamberlain/books") != 0 {
		t.Error("failed to parse work path : " + testConfig.BlogWorkPath)
	}
	if testConfig.LogConfig.LogLevel != 1 || strings.Compare(testConfig.LogConfig.Path, "/var/chamberlain.log") != 0 {
		t.Error("failed to parse log config")
	}

	if strings.Compare(testConfig.DatabaseConfig.Host, "database") != 0 ||
		strings.Compare(testConfig.DatabaseConfig.Database, "chamberlain") != 0 ||
		strings.Compare(testConfig.DatabaseConfig.Username, "root") != 0 ||
		strings.Compare(testConfig.DatabaseConfig.Password, "199527") != 0 ||
		testConfig.DatabaseConfig.Port != 3306 {
		t.Error("failed to parse database config")
	}
	if len(*testConfig.BlogRepositories) != 2 {
		t.Error("failed to parse blog repository")
	}
	fmt.Println((*testConfig.BlogRepositories)[0])
	fmt.Println((*testConfig.BlogRepositories)[1])
	aConfig := (*testConfig.BlogRepositories)[0]
	if strings.Compare(aConfig.RepoPath, "https://github.com/coregiu/philosophy.git") != 0 ||
		strings.Compare(aConfig.RepoName, "Philosophy") != 0 {
		t.Error("failed to parse blog repository detail value")
	}
}

func copyYamlFileToCurrentDir(workPath string) {
	commandStr := "cd " + workPath + " && cp ../chamberlain.yml ./"
	cmd := exec.Command("/bin/sh", "-c", commandStr)
	_ = cmd.Run()
}

func deleteYamlFileOfCurrentDir(workPath string)  {
	commandStr := "cd " + workPath + " && rm -rf chamberlain.yml"
	cmd := exec.Command("/bin/sh", "-c", commandStr)
	_ = cmd.Run()
}