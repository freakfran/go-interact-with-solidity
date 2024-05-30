package cmdUtils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-cmd/cmd"
	"github.com/gookit/slog"
	"os"
	"strings"
)

// Execute
//
//	@Description: 执行cmd命令,从项目根路径执行
//	@param cmdStr
//	@return error
func Execute(cmdStr string) error {
	slog.Infof("Execute cmd: %s", cmdStr)
	args := strings.Split(cmdStr, " ")
	newCmd := cmd.NewCmd(args[0], args[1:]...)
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	if strings.ContainsAny(dir, "/test") {
		newCmd.Dir = strings.Replace(dir, "/test", "", 1)
	}
	log.Info(newCmd.Dir)
	status := <-newCmd.Start()
	if len(status.Stdout) > 0 {
		slog.Info(status.Stdout)
	}
	if len(status.Stderr) > 0 {
		return fmt.Errorf("execute cmd error: %v", status.Stderr)
	}
	return status.Error
}
