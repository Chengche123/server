## 项目结构
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
│      │  └─cmd: 启服务
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
│  └─token: token verifyer
├─interceptor: 中间件
│  └─micro: micro中间件
│      └─auth: 取uid注入ctx
├─mongo
│  ├─mgotesting
│  └─util
├─mysql
│  └─test
└─os
    └─env
</code></pre>
</details>

<details>
<summary>推荐服务 (recommend文件夹,一级目录仿open-bilibili)</summary>
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
