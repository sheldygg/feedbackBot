package helpers

import (
	"encoding/json"
	"feedbackBot/src/config"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type BotVerification struct {
	BotUserId         int64  `json:"bot_user_id"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
	CustomDescription string `json:"custom_description"`
}

type ChatWithBotVerification struct {
	gotgbot.Chat
	BotVerification *BotVerification `json:"bot_verification,omitempty"`
}

// GetChatWithBotVerification calls getChat and preserves custom response fields.
func GetChatWithBotVerification(bot *gotgbot.Bot, chatId int64, opts *gotgbot.GetChatOpts) (*ChatWithBotVerification, error) {
	params := map[string]string{
		"chat_id": strconv.FormatInt(chatId, 10),
	}

	var reqOpts *gotgbot.RequestOpts
	if opts != nil {
		reqOpts = opts.RequestOpts
	}

	raw, err := bot.Request("getChat", params, nil, reqOpts)
	if err != nil {
		return nil, err
	}

	var baseChat gotgbot.Chat
	if err := json.Unmarshal(raw, &baseChat); err != nil {
		return nil, err
	}

	var extra struct {
		BotVerification *BotVerification `json:"bot_verification,omitempty"`
	}
	if err := json.Unmarshal(raw, &extra); err != nil {
		return nil, err
	}

	return &ChatWithBotVerification{
		Chat:            baseChat,
		BotVerification: extra.BotVerification,
	}, nil
}

func HasRequiredBotVerification(bot *gotgbot.Bot, userId int64) (bool, error) {
	chat, err := GetChatWithBotVerification(bot, userId, &gotgbot.GetChatOpts{})
	if err != nil {
		return false, err
	}

	return chat.BotVerification != nil &&
		chat.BotVerification.BotUserId == config.CurrentConfig.Verification.BotUserId, nil
}
