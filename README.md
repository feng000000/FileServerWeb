# FileServerWeb
用作文件服务器的管理web

运行前先要设置环境变量 `FILE_SERVER_CONFIG` 为配置文件绝对路径
例如
```shell
export FILE_SERVER_CONFIG="/home/feng/Code/GoCode/FileServerWeb/config/config.toml"
```

## 功能

### 账号
- 基本信息
  - UUID
  - username
  - passwd
- 登录注册
- 用户级别
  - 0 管理员
    - 1 TB
  - 1 荣誉会员
    - 1 TB
  - 2 会员
    - 500 GB
  - 5 普通
    - 10 GB

### 实用小功能
- 返回IP

### 文件管理
- 文件上传
  - 文件路径前缀: 项目绝对路径
  - 文件存储相对路径: `data/user_files/{username}/{文件名}`
  - 表单上传文件
  - MySQL文件表
    - UUID
    - 文件存储相对路径
    - 存储日期
  - 存储时判断文件名是否重复, 重复的话文件名末尾加 "_1"
    - 若仍然重复则加 "_2", 以此类推
- 文件下载
- 查看已使用空间总大小
  - 文件大小: 数据库存

---

## 未实现的功能

### 文件管理
### 账号
- 注册增加邮箱验证
- 发邮件
- phone_number
- 文件使用容量

__管理员功能__
- 查看所有用户信息
- 所有用户使用容量
- 封禁用户
  - 封禁时间
  - 封禁IP

### 实用小功能

- base64编解码
- 格式化json
- url编解码
- 计算md5, sha256
- 时间戳转换
- 文本比对
- 饥荒开服
  - 启动
  - 更新
  - 关闭
- 语音
