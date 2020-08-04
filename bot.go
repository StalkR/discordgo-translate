/*
Package translate is a Discord bot library to translate messages with Google Translate API.

  import translate "github.com/StalkR/discordgo-translate"

Discord bot setup:

 - create a new application: https://discordapp.com/developers/applications/me
 - get client ID, required to add bot to server below
 - click add bot
 - get token, required to configure the bot
 - add bot to server: https://discordapp.com/oauth2/authorize?&client_id=CLIENT_ID&scope=bot&permissions=0

Add a Google Translate API key and start the bot.

To translate a message into a given language, react to it with the desired
language flag. The source language is automatically detected.
*/
package translate

import (
  "fmt"

  "github.com/bwmarrin/discordgo"
)

// A Bot represents a Discord Translate bot.
type Bot struct {
  token  string
  apiKey string

  session       *discordgo.Session
  removeHandler func()
}

// New creates a new Discord Translate bot.
// Requires a Discord bot token and a Google Translate API key.
func New(token string, apiKey string) *Bot {
  return &Bot{
    token:  token,
    apiKey: apiKey,
  }
}

// Start starts the bot, connecting it to Discord.
func (b *Bot) Start() error {
  s, err := discordgo.New("Bot " + b.token)
  if err != nil {
    return fmt.Errorf("error creating session: %v", err)
  }
  if err := b.initialize(s); err != nil {
    s.Close()
    return err
  }
  b.session = s // for b.Close()
  return nil
}

func (b *Bot) initialize(s *discordgo.Session) error {
  b.removeHandler = s.AddHandler(b.handleReaction)
  if err := s.Open(); err != nil {
    return fmt.Errorf("error opening connection: %v,", err)
  }
  return nil
}

// Close closes the bot: handler uninstalled and Discord session closed.
// Bot can be started again.
func (b *Bot) Close() error {
  b.removeHandler()
  b.removeHandler = nil
  err := b.session.Close()
  b.session = nil
  return err
}
