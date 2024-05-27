package options

type TailFileOptions struct {
	Location  *SeekInfo `mapstructure:"location"`
	ReOpen    bool      `mapstructure:"reopen"`
	MustExist bool      `mapstructure:"must-exist"`
	Poll      bool      `mapstructure:"poll"`
	Follow    bool      `mapstructure:"follow"`
}

type SeekInfo struct {
	Offset int64 `mapstructure:"offset"`
	Whence int   `mapstructure:"whence"`
}

func NewTailFileOptions() *TailFileOptions {
	return &TailFileOptions{
		Location: &SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      false,
		Follow:    false,
	}
}
