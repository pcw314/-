package apps

// 通过导包加载impl和api中init方法，init方法将api和impl实例对象注册到ioc中，让ioc托管
import (
	_ "gitee.com/xygfm/authorization/apps/admin/api"
	_ "gitee.com/xygfm/authorization/apps/admin/impl"
	_ "gitee.com/xygfm/authorization/apps/audit/api"
	_ "gitee.com/xygfm/authorization/apps/audit/impl"
	_ "gitee.com/xygfm/authorization/apps/log/api"
	_ "gitee.com/xygfm/authorization/apps/log/impl"
	_ "gitee.com/xygfm/authorization/apps/message/api"
	_ "gitee.com/xygfm/authorization/apps/message/impl"
	_ "gitee.com/xygfm/authorization/apps/oss/api"
	_ "gitee.com/xygfm/authorization/apps/oss/impl"
	_ "gitee.com/xygfm/authorization/apps/place/api"
	_ "gitee.com/xygfm/authorization/apps/place/impl"
	_ "gitee.com/xygfm/authorization/apps/position/api"
	_ "gitee.com/xygfm/authorization/apps/position/impl"
	_ "gitee.com/xygfm/authorization/apps/user/api"
	_ "gitee.com/xygfm/authorization/apps/user/impl"
)
