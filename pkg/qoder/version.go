package qoder

// 编译时通过 -ldflags 注入
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)
