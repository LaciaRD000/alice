package music

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type VcData struct { //読み上げデータ
	Connection *discordgo.VoiceConnection //音声を再生するコネクション
	ChannelID  string                     //読み上げるテキストチャンネルのID
	Queue      *[]string                  //読み上げたい音声のパスのキュー（のアドレス）
}

var vcDict = make(map[string]VcData)

func ExistsData(guildID string) (VcData, bool) {
	vc, ok := vcDict[guildID]
	return vc, ok
}

func InsertData(guildID string, connection *discordgo.VoiceConnection, channelID string) VcData {
	s := make([]string, 0, 10)
	vcDict[guildID] = VcData{
		Connection: connection,
		ChannelID:  channelID,
		Queue:      &s,
	}
	return vcDict[guildID]
}

func AppendQueue(data *VcData, url string) {
	*data.Queue = append(*data.Queue, url)
}

func DeleteData(guildID string) {
	delete(vcDict, guildID)
}

func Play(data *VcData) {
	if err := data.Connection.Speaking(true); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("music error")
	}

	defer func() {
		if err := data.Connection.Speaking(false); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("music error")
		}
	}()

	// ToDo

	*data.Queue = (*data.Queue)[1:]
	if len(*data.Queue) > 0 {
		Play(data)
	}
}
