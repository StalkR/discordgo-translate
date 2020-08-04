package translate

import (
  "log"
  "os"
  "os/signal"
  "syscall"
)

func ExampleBot() {
  b := New("<Discord bot token>", "<Google Translate API token>")
  if err := b.Start(); err != nil {
    log.Fatal(err)
  }
  defer b.Close()

  log.Printf("Discord Translate bot running - press CTRL-C to exit")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc
}
