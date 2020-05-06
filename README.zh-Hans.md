# MYDICTIONARY 演示程序

[English](./README.md)

### 1. 简介

这是一个基于Excel表格与在线内容的词典应用程序。

这是一个基于[MYDICTIONARY](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.zh-Hans.md)的跨平台应用程序。

这是一个基于命令行的桌面应用程序。

### 2. 发布

从[这里](https://github.com/zzc-tongji/mydictionary-demo/releases)获得针对不同平台的可执行文件。

### 3. 使用指南

最重要的事情先说三遍：

- **键入“*”和“回车”以退出程序。**
- **键入“*”和“回车”以退出程序。**
- **键入“*”和“回车”以退出程序。**

同时，为了更好的理解本文，请先阅读[基本信息](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.zh-Hans.md#2-%E5%9F%BA%E6%9C%AC%E4%BF%A1%E6%81%AF)。

#### 3.1. 配置文件

在使用应用程序以前，应当编辑[JSON](http://www.json.org/)格式的配置文件。

这些文件可以用纯文本编辑器进行编辑，例如“记事本”（Windows）和“文本编辑”（macOS）。

用户还可以使用[JSON Editor Online](http://jsoneditoronline.org/)来编辑它们：

1. 打开默认的配置文件。
2. 将其内容拷贝到网页的左边。
3. 点击右箭头，文件内容会被解析到右边。
4. 在右边进行编辑。
5. 点击左箭头，文件内容会被更新。
6. 将内容拷贝到配置文件。
7. 保存配置文件。

##### 3.1.1. 配置文件 "mydictionary.setting.json"

从[这里](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.zh-Hans.md#25-%E9%85%8D%E7%BD%AE)获取更多信息。

##### 3.1.2. 配置文件 "mydictionary-local-cli.setting.json"

示例：

```json
{
	"autoSaveFile": {
		"enable": true,
		"timeIntervalSecond": 600,
		"notification": true
	}
}
```

当应用程序启动时，*生词本文件*和*离线词典文件*会被拷贝到对应的*生词本*和*离线词典*（内存映像）中。当应用程序退出时，*生词本*和*离线词典*会被写回到对应的*生词本文件*和*离线词典文件*中。这种机制使得对*词条*的*记录*均在内存中完成，大大提升了运行效率。

但是这也带来一个问题：一旦程序异常退出，本次运行中*记录*的*词条*将会永久丢失。为了解决这个问题的，应用程序具有自动写回功能。

`"autoSaveFile"`用于控制自动写回功能，具有三个成员：

- 布尔型`"enable"`：如果设定为`true` ， 那么启用自动写回功能。
- 整数`"timeIntervalSecond"`：决定两次写回操作的时间间隔（秒）。
- 布尔型`"notification"`：如果设定为`true` ，应用程序会在自动写回时显示信息。

#### 3.2.启动

在启动应用程序以前，请确保所有XLSX文件未被其他应用程序（比如[Microsoft Excel](https://products.office.com/zh-cn/excel)或[WPS表格](http://www.wps.cn/product)）所使用。

如果应用程序成功启动，用户会看到：

```
(... 终端信息 ...) $ ./mydictionary-local-cli

[2017-06-06 15:24:54]

mydictionary-local

[2017-06-06 15:24:54]

(... 配置文件 mydictionary-local-cli.setting.json 的内容 ...)

[2017-06-06 15:24:54]

mydictionary

[2017-06-06 15:24:54]

(... 配置文件 mydictionary.setting.json 的内容 ...)

[2017-06-06 15:24:55]

(... 网络检查的结果 ...)

[2017-06-06 15:24:55]

ready

TIPS: press "*" and "enter" to quit at any time
```

此时，应用程序已准备完毕。

若应用程序未能正常启动，则输出错误并退出。例如：

```
[2017-06-06 17:31:47]

incorrect format of file "(... 路径 ...)/data/animal.xlsx": missing cell "Word" in row 1

[2017-06-06 17:31:47]

Quit (press "enter" to continue).
```

#### 3.3. 查询

一个典型的查询如下：

```
apple!@#
// 这是注释，终端不会输出这部分信息
// 用户输入的词汇和查询选项

============================================================
apple (advance, online, do not record)
// 词汇和查询选项
------------------------------------------------------------
* BASIC
// 查询类型，BASIC 或 ADVANCE

apple
  [fruit: 1] (0)
  // 对于生词本和本地词典, [名字: 序号] (查询计数)
  n. 苹果公司；【植】苹果；【植】苹果树
  网络 苹果电脑；美国苹果；美国苹果公司
  // 释义
  # Since the meeting takes place on Apple's property, the company has every right to make the rules.
  # 由于会议地点在苹果公司，公司完全有权利制定规则。
  // 笔记

apple
  [Merriam Webster]
  // 对于在线词典, [名字]
  noun
  : a round fruit with red, yellow, or green skin and firm white flesh

apple
  [Bing Dictionary]
  n. 苹果公司；【植】苹果；【植】苹果树
  网络 苹果电脑；美国苹果；美国苹果公司

apple
  [iCIBA Collins]
  01 N-VAR 苹果
     An apple is a round fruit with smooth green, yellow, or red skin and firm white flesh.
  02 PHRASE 心肝宝贝；掌上明珠
     If you say that someone is the apple of your eye, you mean that they are very important to you and you are extremely fond of them.
============================================================
```

##### 3.3.1. 基本查询

键入单词和“回车”，应用程序会执行*基本查询*。

从[这里](https://github.com/zzc-tongji/mydictionary/blob/master/readme/vocabulary.zh-Hans.md#2-%E6%9F%A5%E8%AF%A2)获取更多关于*基本查询*和*高级查询*的信息。

```
cat

============================================================
cat
------------------------------------------------------------
* BASIC

cat
  [animal: 1] (1)
  n. 【动】猫；【动】猫科动物；爵士乐爱好者；【船】起锚滑车
  v. 【船】把(锚)吊放在锚架上；宿娼；〈口〉呕吐；〈俚〉寻欢
  abbr. (=computerized axial tomography)【医】计算机化轴向层面X射线摄影法
  网络 卡特(Carter)；过氧化氢酶(catalase)；卡特彼勒(Caterpillar)
  # The cat sat in front of the bird cage in an agony of frustration at being so near and yet so far.
  # 猫无可奈何地坐在鸟笼前，眼看着鸟儿近在咫尺，可怎么也够不着。
============================================================

apple

============================================================
apple
------------------------------------------------------------
* BASIC

apple
  [fruit: 1] (1)
  n. 苹果公司；【植】苹果；【植】苹果树
  网络 苹果电脑；美国苹果；美国苹果公司
  # Since the meeting takes place on Apple's property, the company has every right to make the rules.
  # 由于会议地点在苹果公司，公司完全有权利制定规则。
============================================================

frustration

============================================================
frustration
------------------------------------------------------------
* BASIC

frustration
  [Merriam Webster]
  noun
  : a feeling of anger or annoyance caused by being unable to do something : the state of being frustrated
  : something that causes feelings of anger and annoyance
  : the fact of being prevented from succeeding or doing something

frustration
  [Bing Dictionary]
  n. 沮丧；受挫；挫败；懊丧
  网络 挫折；失败；挫折感

frustration
  [iCIBA Collins]
  01 VERB 使灰心;使沮丧;使愤怒
     If something frustrates you, it upsets or angers you because you are unable to do anything about the problems it creates.
  02 VERB 挫败;阻挠…的成功
     If someone or something frustrates a plan or attempt to do something, they prevent it from succeeding.
============================================================
```

##### 3.3.2. 高级查询

在键入的单词中加入`!`或`！`，应用程序会执行*高级查询*。

```
frustration!

============================================================
frustration (advance)
------------------------------------------------------------
* BASIC

frustration
  [bing-dictionary: 1] (2)
  n. 沮丧；受挫；挫败；懊丧
  网络 挫折；失败；挫折感

frustration
  [iciba-collins: 1] (2)
  01 VERB 使灰心;使沮丧;使愤怒
     If something frustrates you, it upsets or angers you because you are unable to do anything about the problems it creates.
  02 VERB 挫败;阻挠…的成功
     If someone or something frustrates a plan or attempt to do something, they prevent it from succeeding.

frustration
  [merriam-webster: 1] (2)
  noun
  : a feeling of anger or annoyance caused by being unable to do something : the state of being frustrated
  : something that causes feelings of anger and annoyance
  : the fact of being prevented from succeeding or doing something
------------------------------------------------------------
* ADVANCE

cat
  [animal: 1] (1)
  n. 【动】猫；【动】猫科动物；爵士乐爱好者；【船】起锚滑车
  v. 【船】把(锚)吊放在锚架上；宿娼；〈口〉呕吐；〈俚〉寻欢
  abbr. (=computerized axial tomography)【医】计算机化轴向层面X射线摄影法
  网络 卡特(Carter)；过氧化氢酶(catalase)；卡特彼勒(Caterpillar)
  # The cat sat in front of the bird cage in an agony of frustration at being so near and yet so far.
  # 猫无可奈何地坐在鸟笼前，眼看着鸟儿近在咫尺，可怎么也够不着。
============================================================
```

##### 3.3.3. 在线查询

在键入的单词中加入`@`，应用程序会被告知用户希望执行在线查询。

**应用程序是否真正执行在线查询还取决于*配置文件*`mydictionary.setting.json`中的`online.mode`。** 从[这里](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.zh-Hans.md#2531-mode)获取更多信息。

```
apple@

============================================================
apple (online)
------------------------------------------------------------
* BASIC

apple
  [fruit: 1] (2)
  n. 苹果公司；【植】苹果；【植】苹果树
  网络 苹果电脑；美国苹果；美国苹果公司
  # Since the meeting takes place on Apple's property, the company has every right to make the rules.
  # 由于会议地点在苹果公司，公司完全有权利制定规则。

apple
  [Merriam Webster]
  noun
  : a round fruit with red, yellow, or green skin and firm white flesh

apple
  [Bing Dictionary]
  n. 苹果公司；【植】苹果；【植】苹果树
  网络 苹果电脑；美国苹果；美国苹果公司

apple
  [iCIBA Collins]
  01 N-VAR 苹果
     An apple is a round fruit with smooth green, yellow, or red skin and firm white flesh.
  02 PHRASE 心肝宝贝；掌上明珠
     If you say that someone is the apple of your eye, you mean that they are very important to you and you are extremely fond of them.
============================================================
```

##### 3.3.4. 不要记录

在键入的单词中加入`#`，应用程序将不会*记录*结果到*生词本*和*离线词典*。

```
apple#

============================================================
apple (do not record)
------------------------------------------------------------
* BASIC

apple
  [fruit: 1] (2)
  n. 苹果公司；【植】苹果；【植】苹果树
  网络 苹果电脑；美国苹果；美国苹果公司
  # Since the meeting takes place on Apple's property, the company has every right to make the rules.
  # 由于会议地点在苹果公司，公司完全有权利制定规则。
============================================================
```

#### 3.4. 退出

**键入“\*”和“回车”以退出程序。** 否则，*生词本*和*离线词典*将不会被写回到对应的*生词本文件*和*离线词典文件*中，这意味着所有更改将会丢失。

```
*

[2017-06-06 15:45:57]

(... 写回操作的信息 ...)

[2017-06-06 15:45:57]

Quit.

(... 终端信息 ...) $
```

### 4. 交流

- [反馈](https://github.com/zzc-tongji/mydictionary-demo/issues)

### 5. 其他

- 所以代码文件是用[Atom](https://atom.io/)编写的。
- 所有".md"文件是用[Typora](http://typora.io)编写的。
- 所有".md"文件的风格是[Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown)。
- 各行以LF（Linux）结尾。

