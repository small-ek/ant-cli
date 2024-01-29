package template

func Config(name string) string {
	return `#系统配置
[system]
#启动地址
address = ":9000"
#权限配置路径
rbac_path = "./config/rbac.conf"
#是否开启跨域
cors = true
#是否开发调试模式
debug = true

#权限配置文件目录
[casbin]
path = "rbac.conf"

#接口请求日志
[log]
#路径
path = "./log/ant.log"
#格式 json、console
format = "console"
#日志服务名称
service_name = "antgo"
#日志输出等级 all、info、warn、error、debug、dpanic、panic、fatal
level = "all"
#是否输出控制台
console = true
#是否开启日志
switch = true
#文件最长保存时间(天)
max_age = 30
#分割大小(MB)
max_size = 1
#保留30个备份(个)
max_backups = 300
#是否需要压缩
compress = false

#数据库设置1
[[connections]]
#数据库名称
name = "mysql1"
#数据库类型
type = "mysql"
#服务器地址
hostname = "127.0.0.1"
#服务器端口
port = "3306"
#数据库用户名
username = "root"
#数据库密码
password = "root"
#数据库名
database = "test"
#数据库连接参数
params = "charset=utf8mb4&parseTime=True&loc=Local"
#是否开启日志
log = true

#阿里云配置
[oss]
key_id = ""
key_secret = ""
endpoint = ""
bucket = ""

#Redis配置
[[redis]]
name="redis1"
address = "localhost:6379"
password = ""
db = 0

#邮箱警报发送
[emaill]
switch = true
to = ['56494565@qq.com']
from = '56494565@qq.com'
host = 'smtp.qq.com:25'
secret = 'fdtshicbbvybbiic'

#Json web token
[jwt]
private_key = ""

public_key = ""
`
}
