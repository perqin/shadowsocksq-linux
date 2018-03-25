package profile

type Type int

const (
    Shadowsocks = iota
    ShadowsocksR
)

type Profile struct {
    Id   int
    Type Type
    Host string
}

func (p *Profile) AsArgs() []string {
    var args []string
    //args = append(args, "-s", p.Host)
    args = append(args, "-c", "/etc/shadowsocksrr/ssq-test.json")
    return args
}
