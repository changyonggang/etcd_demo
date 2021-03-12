package utils

import (
	"fmt"
	"os"
)

/*
@Desc :
@Time : 2020/6/11 2:04 下午
@Author : Chang yg
@File : utils
*/

func GetNodeId() string {
	return fmt.Sprintf("%d", os.Getpid())
}