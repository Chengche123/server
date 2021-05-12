### 目录结构

#### infra

<details>
<summary>grpc原生网关 (grpc-gateway文件夹)</summary>
<pre><code>
grpc网关
</code></pre>
</details>

<details>
<summary>micro网关和注册中心 (micro-gateway文件夹)</summary>
<pre><code>
├─docker-compose: etcd micro-api和micro-web打包启动,方便本地开发
│  └─etcd-micro
└─script
</code></pre>
</details>

<details>
<summary>common和util (share文件夹)</summary>
<pre><code>
├─auth
│  └─token: 验jwt
├─database
│  ├─gorm
│  ├─mongo
│  │  ├─mgotesting
│  │  └─util
│  └─mysql
│      ├─dsn
│      └─test
├─interceptor
│  └─micro: micro中间件
│      ├─auth: 从md取jwt验证,把uid注入ctx BFF用
│      └─error: 面向前端,统一错误 BFF用
├─key: 公钥
├─log
│  └─zap
├─micro
│  └─server
└─os
    └─env
</code></pre>
</details>

#### 内部RPC基础服务

<details>
<summary>推荐服务 (recommend文件夹)</summary>
<pre><code>
├─api: proto文件和pb
│  └─grpc
│      └─v1
├─cmd: 启server
│  └─grpc: 直接暴露grpc服务
├─conf: 配置文件
├─dao: 数据访问层
├─model: model
├─server: 接口暴露层
│  └─grpc
└─service: 业务逻辑层
</code></pre>
</details>

<details>
<summary>comic service(comic文件夹)</summary>
<pre><code>
├─api: 接口
│  └─grpc
│      └─v1
├─cmd: 启server
│  └─grpc
├─conf: 配置
├─dao: 数据访问层
├─model: model
├─server: 接口暴露层
│  └─grpc
└─service: 业务逻辑层
    └─grpc
</code></pre>
</details>

#### BFF聚合服务

<details>
<summary>auth service (auth文件夹)</summary>
<pre><code>
├─controler: 接口暴露层
│  ├─grpc: grpc原生接口
│  │  ├─api
│  │  │  └─gen
│  │  │      └─v1
│  │  └─cmd: 启服务
│  └─micro: micro框架接口
│      ├─api
│      │  └─gen
│      │      └─v1
│      ├─grpc: grpc接口
│      │  ├─cmd: 启服务,打镜像
│      │  └─db: 打db测试镜像
│      └─http: bff接口
│          └─cmd: 启服务
├─dao: 数据访问层
│  ├─mongo
│  └─mysql
│      ├─models
│      ├─raw
│      └─repository
├─key: 认证服务公私钥
├─service: 业务逻辑层
├─token
└─wechat
</code></pre>
</details>

