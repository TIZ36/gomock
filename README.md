# gamemock


## 单体架构 20230510
```
- base // 会用到的一些基础模块（暂时自己封装，为了学习；上生产可以找替代的成熟项目），Cache，FuncProxy
- config // 配置文件(yaml)
- app
  - internal // 内部模块
    - http // http接口
      - handler/ // http接口处理函数
      - middleware/ // http中间件
      - router.go // http路由
    - svr // 业务逻辑
    - model // 数据库模型  
  - main.go // 入口
  - app_ctx.go // app上下文
  - config.go // 配置文件
- types // 一些公共的类型定义
- doc // 文档
  - http // http接口文档
    - interface.md
- README.md // 项目说明
```

## types
> 面向接口的开发，首先对types感觉需要有一些规则

- types
  - consts.go // 常量定义
  - err.go // 错误定义
  - api.go    // 接口类型定义 （http接口）
  - svc.go    // 服务类型定义 （业务类型 + 数据库model类型，这两个部分感觉没有必要分开，因为业务类型和数据库model类型是一一对应的）