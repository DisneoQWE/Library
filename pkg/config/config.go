package config

import "github.com/spf13/viper"

// swagger: param Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
	DBPort string `mapstructure:"DB_PORT"`
}

func LoadConfig() (*Config, error) {
	config := new(Config)
	v := viper.New()
	v.AutomaticEnv()
	//
	err := v.BindEnv("PORT")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("DB_HOST")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("DB_USER")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("DB_PASS")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("DB_NAME")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("DB_PORT")
	if err != nil {
		return nil, err
	}
	err = v.Unmarshal(&config)
	if err != nil {
		//log.Println("Error to unmarshal config", err)
		return nil, err
	}
	return config, nil
}
