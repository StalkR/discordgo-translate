# Discord Translate bot

[![Build Status][1]][2] [![Godoc][3]][4]

A Translate bot using [discordgo][5] and [Google Translate API][8].

To translate a message into a given language, react to it with the desired
language flag. The source language is automatically detected.

## Example
```
import (
  translate "github.com/StalkR/discordgo-translate"
)

b := translate.New("<bot token>", "<Google Translate API token>")
if err := b.Start(); err != nil {
  log.Fatal(err)
}
defer b.Close()

log.Printf("Discord Translate bot running - press CTRL-C to exit")
sc := make(chan os.Signal, 1)
signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
<-sc
```

## Acknowledgements
Discord Go bindings: [bwmarrin/discordgo][5] ([doc][6]).

Inspiration: [Smoogle Translate][9], unfortunately the developers intentionally
refuse to translate messages from other bots, so I had to make my own.

## Bugs, comments, questions
Create a [new issue][7].

[1]: https://api.travis-ci.org/StalkR/discordgo-translate.png?branch=master
[2]: https://travis-ci.org/StalkR/discordgo-translate
[3]: https://godoc.org/github.com/StalkR/discordgo-translate?status.png
[4]: https://godoc.org/github.com/StalkR/discordgo-translate
[5]: https://github.com/bwmarrin/discordgo
[6]: https://godoc.org/github.com/bwmarrin/discordgo
[7]: https://github.com/StalkR/discordgo-translate/issues/new
[8]: https://cloud.google.com/translate/
[9]: https://smoogle.gg/
