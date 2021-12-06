# toyLang

程序执行，直接运行./toyLang -h会给出帮助信息
若程序运行时报错c库的版本不匹配，可以go run main.go 直接运行，或者go build用当前机器的c库重新编译
其中./toyLang -port 2000 是指定所运行的端口，-file是指定所读取的DSL文件
遇到乱码的时候，可以-file testEnglish.Toy 读取英文DSL，从而解决CMD等不支持UTF-8编码的问题
## 软件开发文档
位于./doc/pkg 文件夹中点击index.html后，即可在浏览器中进行浏览，并提供跳转到各个函数详细界面的功能

当电脑上有go语言环境的时候，可以进入每个子目录，通过go doc 命令获得函数的帮助
也可以通过  `go get -v  golang.org/x/tools/cmd/godoc` 安装godoc，然后进入toyLang文件夹，输入godoc命令，在localhost:6060访问文档。

对软件进行更改后，可通过`go test ./...`批量运行test对每一个包都进行test，也可以` go test ./lexer`对单个包进行测试

程序调用parser，然后parser不断的调用lexer获取下一个Token，构造出AST，最后Eval函数在每一个Step上遍历，执行Speak，Listen等，并完成跳转
## 文法描述
### 自然语言描述
一个ToyLang的脚本应有多个Step 语句块构成

一个Step语句块以Step作为开始，下一个Step或者文件结束符作为终止
```
	Step complainProc
	Speak '我是投诉的回复，请您讲'
	Listen 2,4
	Exit

	Step thanks
	Speak '感谢您的使用，再见'
	Exit
```
如上，是两个Step语句，Step语句间的空行不是必须的，上面的例子只是为了美观
Step语句后面跟着一个英文的标识符，用于在后面跳转的时候确定跳转到哪一个语句

每个Step语句块内部，可包含Speak，Listen，Branch，Exit，Silence，Default语句

```
	Speak '你好'+$name  +' ,请问你想 和我聊天 '+ '还是' +'打电话，聊天请发送 聊天 打电话请发送 telephone，离开请发送byebye'
	Listen 2,3
	Branch '聊天',speakProc
	Branch 'telephone',teleProc
	Branch 'byebye',thanks
	Silence silenceProc
	Default defaultProc
```
其中Listen语句后面只能是Branch，Default，Silence等跳转语句，或者Exit退出，一个常规的Step一般是先Speak，然后Listen，接着是不同的跳转路径。

Speak语句后面是由 ' ' 包裹起来的字符串，若有加入 $的需求，则通过 + 把句子和变量以及下一个句子连接起来，Speak的语句内容支持UTF-8的所有内容，包括emoji等，但是需要考虑用户机器是否支持UTF-8编码的内容。

Listen语句后面是两个数字，中间用 ， 隔开，第一个数字是几秒开始接受输入，第二个数字是听多少秒后结束

Branch语句后面 首先一个''包裹的字符串，表示输入什么样的字符串的时候，会发生跳转，然后是一个逗号， 作为分隔符，最后是一个英文的标识符名称，表明跳转到哪一个Step

Silence和Default跟着的都是一个英文标识符，分别表明在用户没有说话，用户胡乱说话的时候，应该跳转到哪一个Step执行

值得注意的是，只有Speak的语句和Branch跳转的判定条件，能采用UTF-8编码的内容，而Step后面的标识符，跳转分支后面的标识符，变量命名，都是纯英文构成的
###　形式语言描述


终结符为 "Step","Speak","Listen","Default","Silence","Branch","Exit",+,$,  ','   , '  ,ident,rune,num ident为纯英文构成标识符，rune为UTF-8编码的字符，num为正数，上面的 单引号两边的空格只为了强调，实际上没有两边的空格, 其中逗号, 两边的单引号和空格也只是为了强调，实际上也没有单引号和空格。

非终结符为 Program Step TrueStep Eval Speak Listen Branch String Runes RunesPlus

开始符号为Program

Program -> Step
Step->TrueStep Step|TrueStep
TrueStep->"Step" ident Eval
Eval->Speak Listen Branch|Speak "Exit"|"Exit"
Speak->"Speak" String
String->'Runes'
Runes-> Runes rune|RunesPlus Runes|$ident|rune
RunesPlus->Runes +
Listen->"Listen" num , num
Branch ->"Default" ident|"Silence" ident|"Branch" 'rune',ident


