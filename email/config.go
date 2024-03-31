package email

const (
	HostQQMail = "smtp.qq.com"
	PortQQMail = 587

	HostExmail = "smtp.exmail.qq.com"
	PortExmail = 465
)

type Config struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}
