package template

func Config(name string) string {
	return `#系统配置
[system]
#启动地址
address = ":9000"
#是否开启跨域
cors = false
#是否开发调试模式
debug = true


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
max_age = 180
#分割大小(MB)
max_size = 10
#保留30个备份(个)
max_backups = 5000
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
username = ""
#数据库密码
password = ""
#数据库名
database = ""
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
name=""
address = ""
password = ""
db = 0

#邮箱
[emaill]
switch = true
to = ['']
from = ''
host = ''
secret = ''

#Json web token
[jwt]
private_key = ""

public_key = ""
`
}
