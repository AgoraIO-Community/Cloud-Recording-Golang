package utils

import (
	"math/rand"
	"time"

	"github.com/AgoraIO-Community/Cloud-Recording-Golang/schemas"
	"github.com/spf13/viper"
)

// GetRtcToken generates token for Agora RTC SDK
func GetRtcToken(channel string, uid int) (string, error) {
	var RtcRole Role = RolePublisher

	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp := currentTimestamp + 86400

	return BuildRTCTokenWithUID(viper.GetString("APP_ID"), viper.GetString("APP_CERTIFICATE"), channel, uint32(uid), RtcRole, expireTimestamp)
}

// GetRtmToken generates a token for Agora RTM SDK
func GetRtmToken(user string) (string, error) {

	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp := currentTimestamp + 86400

	return BuildRTMToken(viper.GetString("APP_ID"), viper.GetString("APP_CERTIFICATE"), user, RoleRtmUser, expireTimestamp)
}

// GenerateUserCredentials generates uid, rtc and rtc token
func GenerateUserCredentials(channel string) (*schemas.UserCredentials, error) {
	uid := int(rand.Uint32())
	rtcToken, err := GetRtcToken(channel, uid)
	if err != nil {
		return nil, err
	}
	return &schemas.UserCredentials{
		Rtc: rtcToken,
		UID: uid,
	}, nil
}
