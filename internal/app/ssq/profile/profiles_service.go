package profile

import "github.com/perqin/shadowsocksq-linux/internal/app/ssq/sqlite"

type ProfilesService interface {
    GetAllProfiles()	[]Profile
}

type profilesServiceImpl struct {
}

var profilesService ProfilesService

func init() {
    profilesService = &profilesServiceImpl{}
}

func GetProfilesService() ProfilesService {
    return profilesService
}

func (p *profilesServiceImpl) GetAllProfiles() []Profile {
    var profiles []Profile
    db := sqlite.GetDb()
    rows, err := db.Query("SELECT * FROM profiles")
    if err != nil {
        return profiles
    }
    for rows.Next() {
        var profile Profile
        if err := rows.Scan(&profile.Id, &profile.Type, &profile.Host); err == nil {
            profiles = append(profiles, profile)
        }
    }
    return profiles
}
