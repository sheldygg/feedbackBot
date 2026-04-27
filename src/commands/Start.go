/*
 * Start.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/config"
	"feedbackBot/src/helpers"
	"feedbackBot/src/rates"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	// Reply with a verification-dependent start message.
	if rates.Check(ctx.EffectiveChat.Id, 10) && config.CurrentConfig.Welcome.Enabled {
		isVerified, err := helpers.HasRequiredBotVerification(b, ctx.EffectiveUser.Id)
		if err != nil {
			return err
		}

		messageText := config.CurrentConfig.Verification.UnverifiedStartMessage
		if isVerified {
			messageText = config.CurrentConfig.Verification.VerifiedStartMessage
		}

		_, err = ctx.EffectiveMessage.Reply(
			b,
			messageText,
			&gotgbot.SendMessageOpts{ParseMode: "HTML"},
		)
		return err
	}

	return nil
}
