package env

import "os"

const (
	DevMode  = "development"
	GrayMode = "gray"
	ProdMode = "production"
)

var (
	hxsEnv = os.Getenv("hxsenv")
	App    = &struct {
		TokenKey, DevSign, Env            string
		IsDevMode, IsGrayMode, IsProdMode bool
		SignTimeExpiration                int64
	}{
		Env:                hxsEnv,
		TokenKey:           "xxxxxxxxxxx",
		DevSign:            "abcdefghijk",
		IsDevMode:          hxsEnv == DevMode,
		IsGrayMode:         hxsEnv == GrayMode,
		IsProdMode:         hxsEnv == ProdMode,
		SignTimeExpiration: 60,
	}
)
