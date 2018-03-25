package main

import (
    "github.com/perqin/shadowsocksq-linux/internal/app/ssq"
    "os"
    "os/signal"
    "syscall"
    "log"
)

func main() {
    dm := ssq.GetDatabaseManager()
    dm.Start()
    cm := ssq.GetClientsManager()
    cm.Start()

    // Block
    sig := make(chan os.Signal)
    signal.Notify(sig, syscall.SIGINT)
    <-sig
    log.Println("Received interrupt signal, shutting down gracefully...")

    cm.Stop()
    dm.Stop()
}
