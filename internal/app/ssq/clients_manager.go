package ssq

import (
    "github.com/perqin/shadowsocksq-linux/internal/app/ssq/subscription"
    "github.com/perqin/shadowsocksq-linux/internal/app/ssq/config"
    "github.com/perqin/shadowsocksq-linux/internal/app/ssq/profile"
    "github.com/perqin/shadowsocksq-linux/internal/app/ssq/process"
)

type ClientsManager interface {
    Start()
    Stop()
}

type clientsManagerImpl struct {
    configService        config.ConfigService
    profilesService      profile.ProfilesService
    subscriptionsService subscription.SubscriptionsService
    processesService     process.ProcessesService
}

var clientsManager ClientsManager

func init() {
    clientsManager = &clientsManagerImpl{
        configService:        nil, //GetConfigService(),
        profilesService:      profile.GetProfilesService(),
        subscriptionsService: nil, //GetSubscriptionsService(),
        processesService:     process.GetProcessesService(),
    }
}

func GetClientsManager() ClientsManager {
    return clientsManager
}

func (m *clientsManagerImpl) Start() {
    profiles := m.profilesService.GetAllProfiles()
    for _, p := range profiles {
        // Start this client
        var binary string
        switch p.Type {
        case profile.Shadowsocks:
            binary = "/usr/bin/ss-local"
        case profile.ShadowsocksR:
            binary = "/usr/bin/ssrr-local"
        }
        args := p.AsArgs()
        m.processesService.StartOrRestartProcess(p.Id, binary, args)
    }
}

func (m *clientsManagerImpl) Stop() {
    m.processesService.StopAllProcesses()
}
