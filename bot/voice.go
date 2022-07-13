package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"os"
	"time"
)

func joinVoice(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.Interaction.ApplicationCommandData().Options[0].ChannelValue(s)
	var err error
	voiceChannels[i.Interaction.GuildID], err = s.ChannelVoiceJoin(i.Interaction.GuildID, channel.ID, true, false)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Я подключился к каналу: " + channel.Mention() +
				"\nИспользуйте:\n" +
				"`/start-record` для начала записи\n" +
				"`/stop-record` для остановки записи\n" +
				"`/disconnect` - для отключения от канала",
		},
	})
}

func disconnectVoice(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if voiceChannels[i.Interaction.GuildID] == nil {
		fmt.Printf("channel undefined\n")
		return
	}
	voiceChannels[i.Interaction.GuildID].Disconnect()
	delete(voiceChannels, i.Interaction.GuildID)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Я отключился",
			Flags:   1 << 6,
		},
	})
}

func startRecord(s *discordgo.Session, i *discordgo.InteractionCreate) {
	respond := func(text string) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: text,
			},
		})
	}

	options := i.Interaction.ApplicationCommandData().Options
	var nameRecord string
	if options != nil {
		nameRecord = i.Interaction.ApplicationCommandData().Options[0].StringValue()
	}
	if nameRecord == "" {
		t := time.Now()
		nameRecord = fmt.Sprintf("Record %d-%d-%d %d-%d-%d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
	}
	voice := voiceChannels[i.Interaction.GuildID]
	if voice == nil {
		respond("Бот не в голосовом канале\nПодключитесь командой /join")
	}
	respond("Начинаю запись... " + nameRecord)

	go handleVoice(nameRecord, voice.OpusRecv)
}

func stopRecord(s *discordgo.Session, i *discordgo.InteractionCreate) {
	voice := voiceChannels[i.Interaction.GuildID]

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Останавливаю запись, я выгружу файл в этот же канал как только он будет готов",
		},
	})

	time.Sleep(10 * time.Second)
	close(voice.OpusRecv)

}

func createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
	return &rtp.Packet{
		Header: rtp.Header{
			Version: 2,
			// Taken from Discord voice docs
			PayloadType:    0x78,
			SequenceNumber: p.Sequence,
			Timestamp:      p.Timestamp,
			SSRC:           p.SSRC,
		},
		Payload: p.Opus,
	}
}

func handleVoice(name string, c chan *discordgo.Packet) {
	files := make(map[string]media.Writer)
	for p := range c {
		fmt.Printf("handle voice: %v\n", p)

		file, ok := files[name]
		if !ok {
			var err error
			if _, err := os.Stat(pathToRecords); os.IsNotExist(err) {
				if err := os.Mkdir(pathToRecords, 0750); err != nil {
					fmt.Printf("cant create folder\n%s", err.Error())
				}
			}
			file, err = oggwriter.New(fmt.Sprintf("%s\\%s.ogg", pathToRecords, name), 48000, 2)
			if err != nil {
				fmt.Printf("failed to create file %s.ogg, giving up on recording: %v\n", name, err)
				return
			}
			files[name] = file
		}
		// Construct pion RTP packet from DiscordGo's type.
		rtp := createPionRTPPacket(p)
		err := file.WriteRTP(rtp)
		if err != nil {
			fmt.Printf("failed to write to file %s.ogg, giving up on recording: %v\n", name, err)
		}
	}

	// Once we made it here, we're done listening for packets. Close all files
	for _, f := range files {
		f.Close()
	}
}
