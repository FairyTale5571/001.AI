package tg

import (
	"github.com/srevinsaju/gofer/platforms/matrix"
	"github.com/srevinsaju/gofer/platforms/telegram"
	"github.com/srevinsaju/gofer/types"
)

var telegramListeners = types.Listeners{
	File:        telegram.SendFile,
	Message:     telegram.SendMessage,
	Misc:        telegram.SendMisc,
	Photo:       telegram.SendPhoto,
	EditMessage: telegram.SendEdit,
}

var matrixListeners = types.Listeners{
	File:        matrix.SendFile,
	Message:     matrix.SendMessage,
	Misc:        matrix.SendMisc,
	Photo:       matrix.SendPhoto,
	EditMessage: matrix.SendEdit,
}
