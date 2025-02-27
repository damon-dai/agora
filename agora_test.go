package agora

import (
	"fmt"
	"testing"
)

func Test_agora_CreateKickingRule(t *testing.T) {
	agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	ruleId, err := client.CreateKickingRule("10002", 0)

	fmt.Println("ruleId:", ruleId, ", err: ", err)
}

func Test_agora_DeleteKickingRule(t *testing.T) {
	agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	rule, err := client.DeleteKickingRule(1001)

	fmt.Println("rule:", rule, ", err: ", err)
}

func Test_agora_GenerateRtcToken(t *testing.T) {
	agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
		//agora.WithTokenExpirationInSeconds(20),
	)
	token, err := client.GenerateRtcToken("10001", 1001)
	fmt.Println("token:", token, ", err: ", err)
}

func Test_agora_GetChannelList(t *testing.T) {
	agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	channel, err := client.GetChannelList()

	fmt.Println("channel:", channel, ", err: ", err)
}

func Test_agora_GetChannelUsers(t *testing.T) {
	agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	channelUsers, err := client.GetChannelUsers("10002")

	fmt.Println("channelUsers:", channelUsers, ", err: ", err)
}

func Test_agora_GetKickingRule(t *testing.T) {
	agora := agoraOptions{}
	client := NewAgoraClient(
		agora.WithAppId("9515b5d776b547b88e6"),
		agora.WithAppCertificate("ebda58b1d7a14085909"),
		agora.WithAppKey("1fe5550d1a1d43ab81"),
		agora.WithAppSecret("c678927b78664a2d"),
	)
	rule, err := client.GetKickingRule()

	fmt.Println("rule:", rule, ", err: ", err)
}
