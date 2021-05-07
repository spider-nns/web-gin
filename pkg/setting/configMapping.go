package setting

import "time"

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int
}

type AppSetting struct {
	DefaultPageNo   int
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}
type DataBaseSetting struct {
	DBType   string
	UserName string
	Password string
	Host     string
	DBName   string
	//TablePrefix string
	Charset     string
	ParseTime   bool
	MaxIdleCons int
	MaxOpenCons int
}

func (s *Setting) ReadConfig(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
