package global

// System 全局系统配置
type Sum struct {
	Web   Web   `mapstructure:"web" json:"web" yaml:"web"`
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Log   Log   `mapstructure:"log" json:"log" yaml:"log"`
	Path  Path  `mapstructure:"path" json:"path" yaml:"path"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

///////////////////////////////////////////单项配置

type Web struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Mongo struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}

type Path struct {
	File string `yaml:"file"`
}
