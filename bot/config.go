package bot

//AddQa 添加 Q问题 A回答
func (b *bot) AddQa(q, a string) {
	b.qa[q] = a
}

func (b *bot) Run() {
	b.id = 1
}
