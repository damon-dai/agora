# agora

声网相关接口

1、生成rtc token
```go
        agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6fc6"),
		agora.WithAppCertificate("ebda58b1d7a1408590"),
		agora.WithAppKey("1fe5550d1a1d43ab8185"),
		agora.WithAppSecret("c678927b78664a2db76"),
	)
	
	token, err := client.GenerateRtcToken("10001", 1001)
```

2、获取频道列表
```go
        agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	channel, err := client.GetChannelList()
```

3、获取频道用户列表
```go
        agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	channelUsers, err := client.GetChannelUsers("10002")
```

4、创建禁用规则
```go
        agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	ruleId, err := client.CreateKickingRule("10002", 0)
```

5、删除禁用规则
```go
        agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	rule, err := client.DeleteKickingRule(1001)
```

6、获取规则列表
```go
        agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	rule, err := client.GetKickingRule()
```