package config

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_Should_Be_Set_Struct_Value_When_Reflect_Invoke_ChamberlainConfig(t *testing.T) {
	config := new(ChamberlainConfig)
	configReflect := reflect.ValueOf(config)
	logConfig := configReflect.Elem().FieldByName("LogConfig")
	fmt.Println(logConfig.Type())

	reflectLogObject := reflect.New(logConfig.Type().Elem())
	reflectLogObject.Elem().FieldByName("Path").SetString("/var/test.log")
	fmt.Println(reflectLogObject)

	logConfig.Set(reflectLogObject)
	if strings.Compare(config.LogConfig.Path, "/var/test.log") != 0 {
		t.Error("failed to set work path")
	}
}

func Test_Should_Be_Set_Slice_Value_When_Reflect_Invoke_ChamberlainConfig(t *testing.T) {
	config := new(ChamberlainConfig)
	configReflect := reflect.ValueOf(config)

	reposConfig := configReflect.Elem().FieldByName("BlogRepositories")
	reposConfig = reflect.New(reposConfig.Type().Elem())
	fmt.Println("---------------")
	fmt.Println(reposConfig)
	fmt.Println(reposConfig.Type())
	fmt.Println(reposConfig.Type().Elem())
	fmt.Println(reposConfig.Type().Elem().Elem())
	fmt.Println(reposConfig.Type().Elem().Elem().Elem())
	fmt.Println(reposConfig.Elem())
	fmt.Println(reposConfig.Elem().Type())
	fmt.Println(reposConfig.Elem().Type().Elem())
	fmt.Println(reposConfig.Elem().Type().Elem().Elem())
	fmt.Println("---------------")
	//aRepoPtr := reflect.New(reposConfig.Type().Elem().Elem())
	aRepo := reflect.New(reposConfig.Type().Elem().Elem().Elem())
	repoEle := aRepo.Elem()
	fmt.Println(repoEle.Type())
	repoEle.FieldByName("RepoPath").SetString("https://github.com/coregiu")
	repoEle.FieldByName("RepoName").SetString("coregiu")
	//aRepoPtr.Set(reflect.ValueOf(&(repoEle.Interface())))
	reposConfigValue := reflect.Append(reposConfig.Elem(), aRepo)
	fmt.Println("---------!!!---------")
	fmt.Println(reposConfigValue.Index(0))
	reposConfig.Elem().Set(reposConfigValue)
	fmt.Println(reposConfig.Elem().Index(0))
	configReflect.Elem().FieldByName("BlogRepositories").Set(reposConfig)
	fmt.Println(len(*config.BlogRepositories))
	fmt.Println(*config.BlogRepositories)
	if len(*config.BlogRepositories) != 1 {
		t.Error("failed to set slice value")
	}
}

func Test_Should_Successfully_When_Parse_Object(t *testing.T) {
	content := "LogConfig:\n  Path: /var/chamberlain.log\n  # DEBUG: 0 , INFO: 1, WARN: 2, ERROR: 3\n  LogLevel: 1\nDatabaseConfig:\n  Host: database\n  Port: 3306\n  Database: chamberlain\n  Username: root\n  Password: 199527\nBlogWorkPath: /giu/chamberlain/books"
	contentArr := strings.Split(content, "\n")
	testConfig := new(ChamberlainConfig)
	result := readFileObject(&contentArr, 0, 0, reflect.ValueOf(testConfig))
	if -1 == result {
		t.Error("failed to parse yaml object")
	}

	fmt.Println("----------------")
	fmt.Println(testConfig.LogConfig)
	fmt.Println(testConfig.DatabaseConfig)
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
}

func Test_Should_Successfully_When_Parse_Array(t *testing.T) {
	content := "BlogRepositories:\n  - RepoPath: https://github.com/coregiu/philosophy.git\n    RepoName: Philosophy\n  - RepoPath: https://github.com/coregiu/summary.git\n    RepoName: Technology"
	contentArr := strings.Split(content, "\n")
	testConfig := new(ChamberlainConfig)
	result := readFileObject(&contentArr, 0, 0, reflect.ValueOf(testConfig))
	if -1 == result {
		t.Error("failed to parse yaml object")
	}

	fmt.Println("----------------")
	fmt.Println(testConfig.BlogRepositories)
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

func Test_should_successfully_when_load_yaml_file(t *testing.T) {
	testConfig := new(ChamberlainConfig)
	result := testConfig.loadConfigFromFile()
	if result != nil {
		t.Error(result.Error())
	}
}