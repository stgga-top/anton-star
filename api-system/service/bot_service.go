package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/data/display"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/bilibili"
	"github.com/axiangcoding/ax-web/service/bot"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/axiangcoding/ax-web/tool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"hash/crc32"
	"strconv"
	"time"
)

// DrawNumber 抽一个数字
func DrawNumber(id int64, now time.Time) int32 {
	date := now.Format("2006-01-02")
	sprintf := fmt.Sprintf("%d+%s", id, date)
	hash := crc32.ChecksumIEEE([]byte(sprintf))
	return rand.New(rand.NewSource(uint64(hash))).Int31n(101)
}

func NumberBasedResponse(number int32, template int) string {
	if number == 0 {
		return bot.SelectStaticMessage(template).LuckResp.Is0
	} else if number <= 30 {
		return bot.SelectStaticMessage(template).LuckResp.Between0130
	} else if number <= 50 {
		return bot.SelectStaticMessage(template).LuckResp.Between3050
	} else if number <= 70 {
		return bot.SelectStaticMessage(template).LuckResp.Between5070
	} else if number <= 80 {
		return bot.SelectStaticMessage(template).LuckResp.Between7080
	} else if number <= 95 {
		return bot.SelectStaticMessage(template).LuckResp.Between8095
	} else if number < 100 {
		return bot.SelectStaticMessage(template).LuckResp.Between95100
	} else {
		return bot.SelectStaticMessage(template).LuckResp.Is100
	}
}

// QueryWTGamerProfile 查询系统中已有的玩家的游戏资料。如果资料不存在，则调用爬虫爬取
func QueryWTGamerProfile(nickname string, sendForm cqhttp.SendGroupMsgForm) (*string, *display.GameUser, error) {
	find, err := FindGameProfile(nickname)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if missionId, err := RefreshWTUserInfo(nickname, sendForm); err != nil {
			return nil, nil, err
		} else {
			return missionId, nil, nil
		}
	} else {
		user := find.ToDisplayGameUser()
		return nil, &user, nil
	}
}

func RefreshWTUserInfo(nickname string, sendForm cqhttp.SendGroupMsgForm) (*string, error) {
	missionId := uuid.NewString()
	form := ScheduleForm{
		SendForm: sendForm,
		Nick:     nickname,
	}
	if err := SubmitMissionWithDetail(missionId, table.MissionTypeUserInfo, form); err != nil {
		return nil, err
	}
	tool.GoWithRecover(func() {
		if err := GetUserInfoFromWarThunder(missionId, nickname); err != nil {
			logging.Warn("start crawler failed. ", err)
		}
	})
	return &missionId, nil
}

func GetBiliBiliRoomInfo(roomId int64) (*bilibili.RoomInfoResp, error) {
	client := resty.New().SetTimeout(time.Second * 10)
	var roomInfo bilibili.RoomInfoResp
	url := "https://api.live.bilibili.com/room/v1/Room/get_info"
	resp, err := client.R().SetQueryParam("room_id", strconv.FormatInt(roomId, 10)).
		SetResult(&roomInfo).
		Get(url)
	if err != nil {
		logging.Warn(err)
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("response status code error")
	}
	return &roomInfo, err
}

func DoActionQuery(retMsgForm *cqhttp.SendGroupMsgForm, value string, fullMsg bool) {
	if !IsValidNickname(value) {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.NotValidNickname
		return
	}
	mId, user, err := QueryWTGamerProfile(value, *retMsgForm)
	if err != nil {
		logging.Warnf("query WT gamer profile error. %s", err)
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.CanNotRefresh
	}
	if mId != nil {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.QueryIsRunning
		tool.GoWithRecover(func() {
			if err := WaitForCrawlerFinished(*mId); err != nil {
				logging.Warnf("wait for callback error. %s", err)
			}
		})
	} else {
		if fullMsg {
			retMsgForm.Message = user.ToFriendlyFullString()
		} else {
			retMsgForm.Message = user.ToFriendlyShortString()
		}
	}
}

func DoActionRefresh(retMsgForm *cqhttp.SendGroupMsgForm, value string) {
	if !IsValidNickname(value) {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.NotValidNickname
		return
	}
	if !CanBeRefresh(value) {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TooShortToRefresh
		return
	}
	missionId, err := RefreshWTUserInfo(value, *retMsgForm)
	if err != nil {
		logging.Warn("refresh WT gamer profile error. ", err)
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.CanNotRefresh
	}
	retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.QueryIsRunning
	tool.GoWithRecover(func() {
		if err := WaitForCrawlerFinished(*missionId); err != nil {
			logging.Warnf("wait for callback error. %s", err)
		}
	})
}

func DoActionDrawCard(retMsgForm *cqhttp.SendGroupMsgForm, value string, id int64) {
	number := DrawNumber(id, time.Now().In(time.FixedZone("CST", 8*3600)))
	retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.DrawCard, number)
}

func DoActionLuck(retMsgForm *cqhttp.SendGroupMsgForm, value string, id int64) {
	number := DrawNumber(id, time.Now().In(time.FixedZone("CST", 8*3600)))
	retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Luck, number, NumberBasedResponse(number, retMsgForm.MessageTemplate))
}
