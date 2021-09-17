# toyLang
## 问答
采用 *Input* *Output* 关键字

列表作为输入，自动生成
```
智齿科技提供智能全客服解决方案，并支持多种部署方式。产品包含：Que

Que
智能客服机器人
++Que：请问你想问哪些方面的信息 EndQue
++++Ans：
我是谁
++Ans：哒哒哒哒
EndQue
```

## 内部转接
## 
==注意，是解释执行==

采用 *采用全局配置* Model 字段内

## 配置我猜你想问

## 配置

Logo ，Name， 



## 版本列表
### 0.01版——进行配置
```
Model 字段
同时，应该实现一个类Shell，能够记住机器人配置的内容。
需要实现一些内置的命令，比如print,默认一个环境下只有一个机器人实例，所以print出的结果就是这个机器人。


理想的环境是
```
### 0.1版——单纯的对话

```
Que，QueEnd//EndQue,不用QueEnd
Ans，AnsEnd
Que
Ans
```

### 0.2版——只有数字记忆性
```
能够根据Ans的内容进行回复，比如0，1
QueRe

QueList(Auto) //
Que QueEnd //默认回复为这个
Que QueEnd //当用户回答0的时候，回复此问题
Que QueEnd //当用户回答1的时候，回复此问题


QueListEnd
QueReEnd

```

### 0.3版——可绑定数字记忆
```
能够根据Ans的内容进行回复，比如0，1
QueRe

QueList(Human)
Que(Default) QueEnd//默认回复为这个
Ans，AnsEnd //下方的省略Ans条
Que QueEnd Case CaseEnd//当用户回答case内的内容时 的时候，回复此问题
Que QueEnd //当用户回答1的时候，回复此问题


QueListEnd
QueReEnd

```

### 0.4版——可自定义答案列表
```
能够根据Ans的内容进行回复，比如0，1
$Human=ans1,ans2,ans3,ans4 $ //通过这种方式定义一个Human变量，在QueList里自动绑定每个Case
QueRe

QueList($Human) 默认提供数字和字母（大写，小写，大小写）
Que(Default) QueEnd //默认回复为这个
Que QueEnd Case CaseEnd//当用户回答case内的内容时 的时候，回复此问题
Que QueEnd //当用户回答1的时候，回复此问题


QueListEnd
QueReEnd

```

### 0.5版——可在自定义答案列表的时候，格式化问题列表，比如用顿号隔开而不是每行一个
```
能够根据Ans的内容进行回复，比如0，1
$Human=ans1,ans2,ans3,ans4 $ //通过这种方式定义一个Human变量，在QueList里自动绑定每个Case
QueRe

QueList($Human,p) 默认提供数字和字母（大写，小写，大小写）,通过p把两个分开
Que(Default) QueEnd //默认回复为这个
Que QueEnd Case CaseEnd//当用户回答case内的内容时 的时候，回复此问题
Que QueEnd //当用户回答1的时候，回复此问题


QueListEnd
QueReEnd

```