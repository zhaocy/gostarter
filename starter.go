package gostarter

//资源上下文
type StarterContext ConfigBase

type Starter interface {
	Init(ctx StarterContext)
	Setup(ctx StarterContext)
	Start(ctx StarterContext)
}

//注册器
type starterRegister struct {
	starters []Starter
}

func NewRegister() {

}

type StarterBase struct {
}

func (s StarterBase) Init(ctx StarterContext) {}
func (s StarterBase) Setup()                  {}
func (s StarterBase) Start()                  {}
