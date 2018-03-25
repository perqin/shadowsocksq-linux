package ssq

import "github.com/perqin/shadowsocksq-linux/internal/app/ssq/sqlite"

type DatabaseManager interface {
    Start()
    Stop()
}

type databaseManagerImpl struct {
}

var databaseManager DatabaseManager

func init() {
    databaseManager = &databaseManagerImpl{}
}

func GetDatabaseManager() DatabaseManager {
    return databaseManager
}

func (m *databaseManagerImpl) Start() {
    sqlite.Open()
}

func (m *databaseManagerImpl) Stop() {
    sqlite.Close()
}
