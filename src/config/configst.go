package config

type tomlConfig struct {
	PostPath  string
	Host      hostsvr
	DB        database
	Cache     cache
	FileCfg   filecfg
	LogCfg    logcfg
	PageCfg   pagecfg
	WebSite   website
	CookieCfg cookiecfg
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

type cache struct {
	Enable    bool
	CacheType string
}

type filecfg struct {
	AutoRefresh bool
	RefreshFreq int //seconds
	IncludeFile []string
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
	PassWord string
}

type cookiecfg struct {
	CookieName string
	MaxAge     int
}
