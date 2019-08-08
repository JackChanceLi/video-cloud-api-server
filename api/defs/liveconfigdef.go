package defs

type LiveRoomAllConfig struct {
	//Lid string `json:"lid"`
	//直播预约
	Qorder int    `json:"qorder"`
	Prepic string `json:"pre_pic"`
	//直播条件
	Condition        int      `json:"condition"`
	ConditionType    int      `json:"condition_type"`
	Price            float32  `json:"price"`
	Duration         int      `json:"duration"`
	TryToSee         int      `json:"try_to_see"`
	Email            string   `json:"email"`
	VerificationCode string   `json:"verification_code"`
	WhiteUserList    []string `json:"white_user_list"`
	//直播配置
	LivePic       string `json:"live_pic"`
	Danmu         int    `json:"danmu"`
	Chat          int    `json:"chat"`
	Share         int    `json:"share"`
	ShareText     string `json:"share_text"`
	Advertisement int    `json:"advertisement"`
	AdJumpUrl     string `json:"ad_jump_url"`
	AdPicUrl      string `json:"ad_pic_url"`
	AdText        string `json:"ad_text"`
	//直播安全设置
	Logo             int    `json:"logo"`
	LogoUrl          string `json:"logo_url"`
	LogoPosition     int    `json:"logo_position"`
	LogoTransparency int    `json:"logo_transparency"`
	Lamp             int    `json:"lamp"`
	LampType         int    `json:"lamp_type"`
	LampText         string `json:"lamp_text"`
	LampFontSize     int    `json:"lamp_font_size"`
	LampTransparency int    `json:"lamp_transparency"`
	//直播质量设置
	Delay         int   `json:"delay"`
	Transcode     int   `json:"transcode"`
	TranscodeType []int `json:"transcode_type"`
	Record        int   `json:"record"`
	RecordType    int   `json:"record_type"`
	//直播权限安全设置
	WhiteSiteList string  `json:"white_site_list"`
	BlackSiteList string  `json:"black_site_list"`
	LiveRoomInfo  LiveRoomIdentity `json:"live_room_info"`
}

type LiveRoomCondition struct {
	Lid              string   `json:"lid"`
	Condition        int      `json:"condition"`
	ConditionType    int      `json:"condition_type"`
	Price            float32  `json:"price"`
	Duration         int      `json:"duration"`
	TryToSee         int      `json:"try_to_see"`
	Email            string   `json:"email"`
	VerificationCode string   `json:"verification_code"`
	WhiteUserList    []string `json:"white_user_list"`

}

type DataForCondition struct {
	Code int `json:"code"`
	Data LiveRoomCondition `json:"data"`
	Msg Message `json:"msg"`
}

type DataForAllConfig struct {
	Code int `json:"code"`
	Data LiveRoomAllConfig `json:"data"`
	Msg Message `json:"msg"`
}

