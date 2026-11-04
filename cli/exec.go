package cli

import "regexp"

const (
	// 子进程里承载密钥的环境变量默认名；用 --env-name 可覆盖。
	defaultSecretEnvName = "KEYGET_SECRET"
	itemEnvName          = "KEYGET_ITEM"
	fdEnvName            = "KEYGET_SECRET_FD"
)

var envNamePattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
