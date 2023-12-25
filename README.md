# 开源库
    golang后端公共类

# 使用开源的类进行业务封装
###### 仅仅封装为了业务使用方便

## kafka
    库 github.com/segmentio/kafka-go
    简单封装了下， 做了下订阅消息重连， 发送消息重试

## redis
    库 github.com/redis/go-redis/v9
    官网 https://redis.uptrace.dev/guide/

## mysql
    使用 gorm 库， cmd/gen.go 生成model和query
    官网 https://gorm.io/zh_CN/docs/

## mqtt
    库 github.com/eclipse/paho.mqtt.golang
    官网 https://www.emqx.com/zh/blog/how-to-use-mqtt-in-golang
    客户端账号需在封装库外管理


