//Id 字段用于后期机器人池

//Bot 实际产生的客服机器人由此包生成
package bot

import "fmt"

//基本Bot类型
type bot struct {
	name string
	qa   map[string]string //问题库
	id   int               //用于后期的机器池
}

//生成一个机器人
func New(name string, qa map[string]string, id int) *bot {
	return &bot{name, qa, id}
}

//String 实现println接口
func (b *bot) String() string {
	return fmt.Sprintf("id: %d ,name:%s,qalist %v", b.id, b.name, b.qa)
}
