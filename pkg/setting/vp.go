package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func Parse() (*Setting, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("configs/")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{v}, nil
}
