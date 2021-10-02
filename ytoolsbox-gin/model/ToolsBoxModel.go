package model

import "gorm.io/gorm"

type Tool struct {
	gorm.Model
	// 工具类型：脚本，容器
	ToolType string `form:"toolType" json:"toolType" gorm:"column:toolType" binding:"required"`
	// 镜像名称
	ToolDockerImageName string `form:"toolDockerImageName" json:"toolDockerImageName" gorm:"column:toolDockerImageName"`
	// 工具名称
	ToolName string `form:"toolName" json:"toolName" gorm:"column:toolName" binding:"required"`
	// 工具运行需要的参数
	ToolOptions string `form:"toolOptions" json:"toolOptions" gorm:"column:toolOptions"`
	// 工具最终执行指令
	ToolRunCMD string `form:"toolRunCMD" json:"toolRunCMD" gorm:"column:toolRunCMD"`
	// 工具描述
	ToolDesc string `form:"toolDesc" json:"toolDesc" gorm:"column:toolDesc;default:'这个工具很神秘,没有留下介绍'"`
	// 工具执行位置
	ToolExecuteLocation string `form:"toolExecuteLocation" json:"toolExecuteLocation" gorm:"column:toolExecuteLocation" binding:"required"`
	// 工具评分
	ToolRate float32 `form:"toolRate" json:"toolRate" gorm:"column:toolRate"`
	//  工具评分人数
	ToolRateCount int8 `form:"toolRateCount" json:"toolRateCount" gorm:"column:toolRateCount"`
	// 远程ssh IP
	ToolRemoteIP string `form:"toolRemoteIP" json:"toolRemoteIP" gorm:"column:toolRemoteIP"`
	// 远程ssh Port
	ToolRemoteSSH_Port string `form:"toolRemoteSSH_Port" json:"toolRemoteSSH_Port" gorm:"column:toolRemoteSSH_Port"`
	// 远程ssh Account
	ToolRemoteSSH_Account string `form:"toolRemoteSSH_Account" json:"toolRemoteSSH_Account" gorm:"column:toolRemoteSSH_Account"`
	// 远程ssh Password
	ToolRemoteSSH_Password string `form:"toolRemoteSSH_Password" json:"toolRemoteSSH_Password" gorm:"column:toolRemoteSSH_Password"`

	// 工具作者
	ToolAuthor string `form:"toolAuthor" json:"toolAuthor" gorm:"column:toolAuthor" binding:"required"`
	// 作者联系方式
	ToolAuthorMobile string `form:"toolAuthorMobile" json:"toolAuthorMobile" gorm:"column:toolAuthorMobile"`
}
