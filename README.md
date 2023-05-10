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