package system

import (
	"fmt"
)

// 编译相关的信息
var (
	Name         string // Name
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	BuildName    string // 编译程序名称
	GitCommitID  string // git 最后的提交 commit
	GitBranch    string // git branch
	GoVersion    string // golang的版本
	HostName     string // 编译机器
	Company      string // 公司名称
	Project      string // 项目名称
)

// Version 显示版本信息
func version() {
	fmt.Printf("%s", Icon)
	fmt.Printf("************************************************************\n")
	fmt.Printf("* build name:      %s\n", BuildName)
	fmt.Printf("* build version:   %s\n", BuildVersion)
	fmt.Printf("* build time:      %s\n", BuildTime)
	fmt.Printf("* go version:  	   %s\n", GoVersion)
	fmt.Printf("* git commit:      %s\n", GitCommitID)
	fmt.Printf("* git branch:      %s\n", GitBranch)
	fmt.Printf("* host name:       %s\n", HostName)
	fmt.Printf("* company:         %s\n", Company)
	fmt.Printf("* project:         %s\n", Project)
	fmt.Printf("************************************************************\n")
}
