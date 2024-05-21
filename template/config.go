package template

func Config(name string) string {
	return `#系统配置
[system]
#启动地址
address = "9001"
#是否开启跨域
cors = false
#是否开发调试模式
debug = true
#项目名称
app_name="antgo"

#接口请求日志
[log]
#路径
path = "./log/ant.log"
#输出格式 支持(json、console)
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
max_age = 180
#分割大小(MB)
max_size = 10
#保留30个备份(个)
max_backups = 5000
#是否需要压缩
compress = false

#数据库设置
[[connections]]
#数据库名称(必须唯一)
name = ""
#数据库类型支持mysql、pgsql、sqlsrv、clickhouse
type = "mysql"
#服务器地址
hostname = "127.0.0.1"
#服务器端口
port = "3306"
#数据库用户名
username = ""
#数据库密码
password = ""
#数据库名
database = ""
#数据库连接参数
params = "charset=utf8mb4&parseTime=True&loc=Local"
#是否开启日志
log = true
#设置空闲连接池中的最大连接数
max_idle_conns = 100
#设置数据库的最大打开连接数。
max_open_conns= 200
#设置连接可能被重用的最大时间(秒)。
conn_max_lifetime=300

#Redis配置
[[redis]]
name=""
address = ""
password = ""
db = 0

#邮箱
[email]
switch = true
to = ['']
from = ''
host = ''
secret = ''

#Json web token
[jwt]
#过期时间(秒)
exp=168
#私钥
private_key = ""
#公钥
public_key = ""
`
}
