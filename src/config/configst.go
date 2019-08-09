package config

type tomlConfig struct {
	PostPath string
	Host     hostsvr
	DB       database
	FileCfg  filecfg
	LogCfg   logcfg
	PageCfg  pagecfg
	WebSite  website
}

type hostsvr struct {
	Server string
	Port   int
}

type database struct {
	Enable bool
	DBType string
	DBName string
}

type filecfg struct {
	AutoRefresh bool
	RefreshFreq int //seconds
	IgnoreFile  []string
	UseFilePool bool
}

type logcfg struct {
	Enable        bool
	ShowInConsole bool
	LogPath       string
}

type pagecfg struct {
	MaxItemPerPage int
}

type website struct {
	WebTitle string
}
