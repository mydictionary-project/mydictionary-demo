# MYDICTIONARY Demo

[简体中文](./README.zh-Hans.md)

### 1. Introduction

It is a cross-platform CLI desktop application of excel-based and online dictionaries with [MYDICTIONARY](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.md).

### 2. Release

Get executable files for different platforms from [here](https://github.com/zzc-tongji/mydictionary-demo/releases).

### 3. Manual

The most important thing should be emphasized 3 time at first:

- **Press "\*" and "enter" to quit.**
- **Press "\*" and "enter" to quit.**
- **Press "\*" and "enter" to quit.**

Also, for better comprehension, please read [basic information](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.md#2-basic-information) at first.

#### 3.1. Configure File

Before using the application, configuration files in [JSON](http://www.json.org/) should be edited.

Such files could be modified by any editors of plain text, such as "Notepad" (for Windows) and "TextEdit" (for macOS).

Users could also use [JSON Editor Online](http://jsoneditoronline.org/) to edit them:

1. Open the default configuration file.
2. Copy the content to the left part of the webpage.
3. Click the right arrow, the content will be parsed in the right part.
4. Edit things in the right part.
5. Click the left arrow, the content will be updated.
6. Copy the content to the configuration file.
7. Save the configuration file.

##### 3.1.1. Configure File "mydictionary.setting.json"

Getting further information from [here](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.md#24-configuration).

##### 3.1.2.  Configure File "mydictionary-local-cli.setting.json"

Here is an example:

```json
{
	"autoSaveFile": {
		"enable": true,
		"timeIntervalSecond": 600,
		"notification": true
	}
}
```

When the application launches, *collection files* and *dictionary files* are copied to corresponding *collections* and *dictionaries* (RAM images). When the application exits, *collections* and *dictionaries* are written back to corresponding *collection files* and *dictionary files*. By doing this, the application can *record* *vocabularies* in RAM, which provides high running efficiency.

But here is a problem: if the application aborts, all *vocabularies* *recorded* by this run will lose forever. To prevent this, the application is able to automatically write back data to files.

`"autoSaveFile"` is designed to control the action of writing back. It has got 3 members.

- Boolean `"enable"`: If it is `true`, the action is enabled.
- Integer `"timeIntervalSecond"`: it determines the time interval (second) between two actions.
- Boolean `"notification"`: if it is `true`, the application will display the information when the action happen.

#### 3.2. Launch

Before launching the application, all *collection files* and *dictionary files* should not be used by any other applications, like [Microsoft Excel](https://products.office.com/excel) or [WPS表格](http://www.wps.cn/product).

If the application launches successfully, users will see this:

```
(... bash ...) $ ./mydictionary-local-cli

[2017-06-06 15:24:54]

mydictionary-local

[2017-06-06 15:24:54]

(... the content of file "mydictionary-local-cli.setting.json" ...)

[2017-06-06 15:24:54]

mydictionary

[2017-06-06 15:24:54]

(... the content of file "mydictionary.setting.json" ...)

[2017-06-06 15:24:55]

(... the result of network checking ...)

[2017-06-06 15:24:55]

ready

TIPS: press "*" and "enter" to quit at any time
```

At this time, the application is ready.

If there is something wrong, the application will display the reason and quit, like this:

```
[2017-06-06 17:31:47]

incorrect format of file "(... path ...)/data/animal.xlsx": missing cell "Word" in row 1

[2017-06-06 17:31:47]

Quit (press "enter" to continue).
```

#### 3.3. Query

Here is a typical query:

```
apple!@#
// the word with option users input

============================================================
apple (advance, online, do not record)
// the word with option
------------------------------------------------------------
* BASIC
// query type, "BASIC" or "ADVANCE"

apple
  [fruit: 1] (0)
  // for collection or dictionary, [name: serious number] (query counter)
  n. 苹果公司；【植】苹果；【植】苹果树
  网络 苹果电脑；美国苹果；美国苹果公司
  // define(s)
  # Since the meeting takes place on Apple's property, the company has every right to make the rules.
  # 由于会议地点在苹果公司，公司完全有权利制定规则。
  // note(s)

apple
  [Merriam Webster]
  // for service, [name]
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

##### 3.3.1. Basic Query

If users input a word and press "enter", the application will execute *basic query*.

Getting further information about *basic query* and *advance query* from [here](https://github.com/zzc-tongji/mydictionary/blob/master/readme/vocabulary.md#2-query).

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

##### 3.3.2. Advance Query

If users input a word with `!` or `！`, the application will execute *advanced query*.

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

##### 3.3.3. Online Query

If users input a word with `@`, it will let the application know that users need query the word online.

**Whether the application queries the word online also depends on `online.mode` in *configuration file* `mydictionary.setting.json`.** Getting further information from [here](https://github.com/zzc-tongji/mydictionary/blob/master/readme/main.md#2531-mode).

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

##### 3.3.4. Do Not Record

If users input a word with `#`, the application will not *record* anything to *collections* and *dictionaries*.

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

#### 3.4. Exit

**Press "\*" and "enter" to quit.** Otherwise, *collections* and *dictionaries* will not be written back to corresponding *collection files* and *dictionary files*, which means all changes will lose.

```
*

[2017-06-06 15:45:57]

(... the information of writing back ...)

[2017-06-06 15:45:57]

Quit.

(... bash ...) $
```

### 4. Communication

- [Feedback](https://github.com/zzc-tongji/mydictionary-demo/issues)
- QQ group: 727068810

![727068810](./README.picture/727068810.png)

### 5. Others

- All code files are edited by [Atom](https://atom.io/).
- All ".md" files are edited by [Typora](http://typora.io).
- The style of all ".md" files is [Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown).
- There is a LF (Linux) at the end of each line.

