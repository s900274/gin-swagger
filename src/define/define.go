package define

var Cfg ServiceConfig

type ServiceConfig struct {
	LogFile        string `cfg:"logfile" toml:"logfile"`
	HttpServerPort int    `cfg:"httpserverport" toml:"httpserverport"`
	HttpServerIp   string `cfg:"httpserverip" toml:"httpserverip"`

	SqlLiteFile    string `cfg:"sqlite_file" toml:"sqlite_file"`
}