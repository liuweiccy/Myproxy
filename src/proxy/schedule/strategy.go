package schedule


// 定义链接的选择策略
type Strategy interface {
    Init()
    Choose(client string, servers []string) string
}

const (
    PollName = "poll"
)

var registry = make(map[string]Strategy)

func init()  {
    registry[PollName] = new(Poll)
}

func GetStrategy(name string) Strategy {
    return registry[name]
}
