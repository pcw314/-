package start

import (
	"context"
	"fmt"
	"gitee.com/xygfm/authorization/conf"
	"gitee.com/xygfm/authorization/global"
	"gitee.com/xygfm/authorization/ioc"
	"gitee.com/xygfm/authorization/logger"
	"gitee.com/xygfm/authorization/protocol"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var Cmd = &cobra.Command{
	Use:   "start",
	Short: "启动项目",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := conf.LoadConfigFromYaml("./config.yaml")
		if err != nil {
			panic(err)
		}
		// 初始化全局库
		global.Init()
		// 初始化ioc中Controller对象
		err = ioc.InitControllerObjects()
		if err != nil {
			panic(err)
		}
		// 构造一个Http对象
		http := protocol.NewHttp()
		go func() {
			// 启动Http服务
			httpErr := http.Start()
			if httpErr != nil {
				panic(httpErr)
			}
		}()
		// 如果有关闭的信号量，进行一下停止操作
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		s := <-ch
		logger.Info(fmt.Sprintf("接收到关闭信号: %s", s))
		// 获取上下文
		ctx := context.Background()
		// 优雅关闭Http服务
		err = http.Stop(ctx)
		if err != nil {
			panic(err)
		}
	},
}
