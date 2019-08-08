package dbop

import (
	"fmt"
	"go-api-server/api/defs"
)

func GetAllConfigByLid(lid string) (*defs.LiveRoomAllConfig, error) {
	//获取预约界面信息
	roomIntro, err := RetrieveLRIntroByLid(lid)
	if err != nil {
		fmt.Printf("Error of retrieve liveroom_intro:%v", err)
		return nil, err
	}
	roomAllConfig := &defs.LiveRoomAllConfig{}
	roomAllConfig.Prepic = roomIntro.Prepic
	roomAllConfig.Qorder = roomIntro.Qorder

	//获取直播页面设置信息
    roomConfig, err := RetrieveLRConfigByLid(lid)
    if err != nil {
		fmt.Printf("Error of retrieve liveroom_config:%v", err)
		return nil, err
	}
    roomAllConfig.LivePic = roomConfig.LivePic
    roomAllConfig.Danmu = roomConfig.Danmu
    roomAllConfig.Chat = roomConfig.Chat
    roomAllConfig.Share = roomConfig.Share
    roomAllConfig.ShareText = roomConfig.ShareText
    roomAllConfig.Advertisement = roomConfig.Advertisement
    roomAllConfig.AdJumpUrl = roomConfig.AdJumpUrl
    roomAllConfig.AdPicUrl = roomConfig.AdPicUrl
    roomAllConfig.AdText = roomConfig.AdText

    //获取观看条件设置，暂未完成此处功能，后续添加
    roomCondition, err := RetrieveLRConditionByLid(lid)
    if err != nil {
		fmt.Printf("Error of retrieve liveroom_condition:%v", err)
		return nil, err
	}
    fmt.Println(roomCondition.Condition)
    roomAllConfig.Condition = roomCondition.Condition
    roomAllConfig.ConditionType = roomCondition.ConditionType
    roomAllConfig.Price = roomCondition.Price
    roomAllConfig.Duration = roomCondition.Duration
    roomAllConfig.TryToSee = roomCondition.TryToSee
    roomAllConfig.VerificationCode = roomCondition.VerificationCode
    roomAllConfig.WhiteUserList = roomCondition.WhiteUserList

    //获取直播服务设置
    roomQuality, err := RetrieveLRQualityByLid(lid)
	if err != nil {
		fmt.Printf("Error of retrieve liveroom_quality:%v", err)
		return nil, err
	}
    roomAllConfig.Delay = roomQuality.Delay
    roomAllConfig.Transcode = roomQuality.Transcode
    roomAllConfig.TranscodeType = roomQuality.TranscodeType
    roomAllConfig.Record = roomQuality.Record
    roomAllConfig.RecordType = roomQuality.RecordType

    //获取版本安全
    roomSafe, err := RetrieveLRSafeByLid(lid)
	if err != nil {
		fmt.Printf("Error of retrieve liveroom_safe:%v", err)
		return nil, err
	}
    roomAllConfig.Logo = roomSafe.Logo
    roomAllConfig.LogoPosition = roomSafe.LogoPosition
    roomAllConfig.LogoTransparency = roomSafe.LogoTransparency
    roomAllConfig.LogoUrl = roomSafe.LogoUrl
    roomAllConfig.Lamp = roomSafe.Lamp
    roomAllConfig.LampFontSize = roomSafe.LampFontSize
    roomAllConfig.LampText = roomSafe.LampText
    roomAllConfig.LampTransparency = roomSafe.LampTransparency
    roomAllConfig.LampType = roomSafe.LampType

    //获取权限安全信息
    roomWhiteList, err := RetrieveLRAuthSafeWhiteList(lid)
	if err != nil {
		fmt.Printf("Error of retrieve liveroom_auth_whitelist:%v", err)
		return nil, err
	}
    for i:= 0; i < len(roomWhiteList); i++ {
    	roomAllConfig.WhiteSiteList += roomWhiteList[i].Website + ";"
	}

	roomBlackList, err := RetrieveLRAuthSafeBlackList(lid)
	if err != nil {
		fmt.Printf("Error of retrieve liveroom_auth_blacklist:%v", err)
		return nil, err
	}
	for i:= 0; i < len(roomBlackList); i++ {
		roomAllConfig.BlackSiteList += roomBlackList[i].Website + ";"
	}

	//获取基本信息
	roomInfo, err := RetrieveLiveRoomByLid(lid)
	if err != nil {
		fmt.Printf("Error of retrieve liveroom by lid:%v", err)
		return nil, err
	}
	roomAllConfig.LiveRoomInfo.Lid = roomInfo.Lid
	roomAllConfig.LiveRoomInfo.Cid = roomInfo.Cid
	roomAllConfig.LiveRoomInfo.Name = roomInfo.Name
	roomAllConfig.LiveRoomInfo.Kind = roomInfo.Kind
	roomAllConfig.LiveRoomInfo.Size = roomInfo.Size
	roomAllConfig.LiveRoomInfo.StartTime = roomInfo.StartTime
	roomAllConfig.LiveRoomInfo.EndTime = roomInfo.EndTime
	roomAllConfig.LiveRoomInfo.PushUrl = roomInfo.PushUrl
	roomAllConfig.LiveRoomInfo.PullHlsUrl = roomInfo.PullHlsUrl
	roomAllConfig.LiveRoomInfo.PullRtmpUrl = roomInfo.PullRtmpUrl
	roomAllConfig.LiveRoomInfo.PullHttpFlvUrl = roomInfo.PullHttpFlvUrl
	roomAllConfig.LiveRoomInfo.DisplayUrl = roomInfo.DisplayUrl
	roomAllConfig.LiveRoomInfo.Status = roomInfo.Status
	roomAllConfig.LiveRoomInfo.Permission = roomInfo.Permission
	roomAllConfig.LiveRoomInfo.CreateTime = roomInfo.CreateTime
	roomAllConfig.LiveRoomInfo.PictureUrl = roomInfo.PictureUrl
	return roomAllConfig, nil
}
