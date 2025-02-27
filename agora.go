/**
 * @Author: Damon
 * @Date: 2025/2/27 14:30
 * @Description: 声网相关接口
 **/

package agora

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
	"github.com/tidwall/gjson"
	"time"
)

type AgoraOptions struct {
	appId                    string // 项目appid
	appCertificate           string // 项目app证书
	appKey                   string // 客户key
	appSecret                string // 客户密钥
	tokenExpirationInSeconds uint32 // token有效时间：s
}

const (
	GET_CHANNEL_LIST_URL  = "https://api.sd-rtn.com/dev/v1/channel/%s"         // 获取项目下的频道列表
	GET_CHANNEL_USERS_URL = "https://api.sd-rtn.com/dev/v1/channel/user/%s/%s" // 获取指定频道下的用户列表
	KICKING_RULE_URL      = "https://api.sd-rtn.com/dev/v1/kicking-rule"       // 规则管理
)

type Option func(option *AgoraOptions)

func NewAgoraClient(opts ...Option) *AgoraOptions {
	options := defaultRequestOptions() // 默认的请求选项

	for _, optFunc := range opts {
		optFunc(options)
	}

	return &AgoraOptions{
		appId:                    options.appId,
		appCertificate:           options.appCertificate,
		appKey:                   options.appKey,
		appSecret:                options.appSecret,
		tokenExpirationInSeconds: options.tokenExpirationInSeconds,
	}
}

// GenerateRtcToken 生成声网RTC Token
func (a *AgoraOptions) GenerateRtcToken(channelName string, uid uint32) (string, error) {
	if len(a.appId) == 0 || len(a.appCertificate) == 0 || len(a.appKey) == 0 || len(a.appSecret) == 0 {
		return "", errors.New("初始化配置失败")
	}
	// Token 的有效时间，单位秒
	tokenExpirationInSeconds := a.tokenExpirationInSeconds
	// 所有的权限的有效时间，单位秒，声网建议你将该参数和 Token 的有效时间设为一致
	privilegeExpirationInSeconds := a.tokenExpirationInSeconds
	fmt.Println("time:", a.tokenExpirationInSeconds)
	// 生成 Token
	result, err := rtctokenbuilder.BuildTokenWithUid(a.appId, a.appCertificate, channelName, uid, rtctokenbuilder.RolePublisher,
		tokenExpirationInSeconds, privilegeExpirationInSeconds)
	if err != nil {
		return "", err
	}

	return result, nil
}

// GetChannelList 获取项目的频道列表
func (a *AgoraOptions) GetChannelList() (string, error) {
	if len(a.appId) == 0 || len(a.appCertificate) == 0 || len(a.appKey) == 0 || len(a.appSecret) == 0 {
		return "", errors.New("初始化配置失败")
	}
	resp, err := HttpGet(fmt.Sprintf(GET_CHANNEL_LIST_URL, a.appId), map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", base64Encode(fmt.Sprintf("%s:%s", a.appKey, a.appSecret))),
	}, nil, 3*time.Second)

	return resp, err
}

// GetChannelUsers 获取频道中的用户列表
func (a *AgoraOptions) GetChannelUsers(channelName string) ([]int, error) {
	if len(a.appId) == 0 || len(a.appCertificate) == 0 || len(a.appKey) == 0 || len(a.appSecret) == 0 {
		return []int{}, errors.New("初始化配置失败")
	}
	resp, err := HttpGet(fmt.Sprintf(GET_CHANNEL_USERS_URL, a.appId, channelName), map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", base64Encode(fmt.Sprintf("%s:%s", a.appKey, a.appSecret))),
	}, nil, 3*time.Second)
	if err != nil {
		return []int{}, err
	}

	type ChannelUserResponse struct {
		Code    int  `json:"code"`
		Success bool `json:"success"`
		Data    struct {
			ChannelExist bool  `json:"channel_exist"`
			Users        []int `json:"users"`
		} `json:"data"`
	}

	var response ChannelUserResponse
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}

	if !response.Data.ChannelExist {
		return []int{}, nil
	}

	return response.Data.Users, nil
}

type KickingRuleRequest struct {
	AppId      string   `json:"appid"`
	Cname      string   `json:"cname"`
	UID        uint32   `json:"uid"`
	Time       int      `json:"time"` // 踢出时间（秒）
	Privileges []string `json:"privileges"`
}

// CreateKickingRule 创建禁用踢人规则
func (a *AgoraOptions) CreateKickingRule(channelName string, uid int) (int64, error) {
	if len(a.appId) == 0 || len(a.appCertificate) == 0 || len(a.appKey) == 0 || len(a.appSecret) == 0 {
		return 0, errors.New("初始化配置失败")
	}

	// 设置请求头
	authorization := fmt.Sprintf("Basic %s", base64Encode(fmt.Sprintf("%s:%s", a.appKey, a.appSecret)))
	params := map[string]interface{}{
		"appid": a.appId,
		"cname": channelName,
		"time":  1000, // 分钟
		"privileges": []string{
			"join_channel",  // 加入频道
			"publish_audio", // 发送音频流
			"publish_video", // 发送视频流
		},
	}

	if uid > 0 {
		params["uid"] = uid
	}
	resp, err := HttpRequest(KICKING_RULE_URL, "POST", map[string]string{
		"Authorization": authorization,
	}, params, 3*time.Second)
	if err != nil {
		return 0, err
	}

	return gjson.Get(resp, "id").Int(), nil
}

// DeleteKickingRule 删除规则
func (a *AgoraOptions) DeleteKickingRule(ruleId int64) (string, error) {
	if len(a.appId) == 0 || len(a.appCertificate) == 0 || len(a.appKey) == 0 || len(a.appSecret) == 0 {
		return "", errors.New("初始化配置失败")
	}

	authorization := fmt.Sprintf("Basic %s", base64Encode(fmt.Sprintf("%s:%s", a.appKey, a.appSecret)))
	resp, err := HttpRequest(KICKING_RULE_URL, "DELETE", map[string]string{
		"Authorization": authorization,
	}, map[string]interface{}{
		"appid": a.appId,
		"id":    ruleId,
	}, 3*time.Second)
	if err != nil {
		return "failed", err
	}

	return gjson.Get(resp, "status").String(), nil
}

// GetKickingRule 获取规则列表
func (a *AgoraOptions) GetKickingRule() (string, error) {
	if len(a.appId) == 0 || len(a.appCertificate) == 0 || len(a.appKey) == 0 || len(a.appSecret) == 0 {
		return "", errors.New("初始化配置失败")
	}

	authorization := fmt.Sprintf("Basic %s", base64Encode(fmt.Sprintf("%s:%s", a.appKey, a.appSecret)))
	resp, err := HttpGet(KICKING_RULE_URL, map[string]string{
		"Authorization": authorization,
	}, map[string]interface{}{
		"appid": a.appId,
	}, 3*time.Second)

	return resp, err
}

// base64Encode 对字符串进行Base64编码
func base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (a *AgoraOptions) WithAppId(appId string) Option {
	return func(option *AgoraOptions) {
		option.appId = appId
	}
}

func (a *AgoraOptions) WithAppCertificate(appCertificate string) Option {
	return func(option *AgoraOptions) {
		option.appCertificate = appCertificate
	}
}

func (a *AgoraOptions) WithAppKey(appKey string) Option {
	return func(option *AgoraOptions) {
		option.appKey = appKey
	}
}

func (a *AgoraOptions) WithAppSecret(appSecret string) Option {
	return func(option *AgoraOptions) {
		option.appSecret = appSecret
	}
}

func (a *AgoraOptions) WithTokenExpirationInSeconds(seconds uint32) Option {
	return func(option *AgoraOptions) {
		option.tokenExpirationInSeconds = seconds
	}
}

func defaultRequestOptions() *AgoraOptions {
	return &AgoraOptions{ // 默认请求选项
		tokenExpirationInSeconds: 3600, // 默认1小时过期
	}
}
