manyfaced
=========

# 介绍

manyfaced是一个用Go编写的运行时动态获取最新配置项的库。

在管理和使用程序的配置的时候，我们会遇到两个问题：
1）配置源种类非常多，可以是本地的配置文件（如yaml, toml, ini等格式的文件），也可以是db/zookeeper/etcd等；
也可以通过http/rpc等方式获取。在一个应用程序中，可能会用到好几种配置源，而这些配置源的API都不一样，
给使用带来不方便。因此，需要对配置源做一种抽象，无论哪种格式的配置源(toml/yaml/ini/db/zookeeper/etcd/http/rpc),
都提供一种统一的API访问。
2）很多配置项的值都是写死在配置源中，然后一次读取入内存，当我们改变配置项的值并让它生效时，我们需要上线重启应用程序。当然，我们可以在需要
读取配置项的值的时候，每次都解析配置源，这样虽然能够保证每次都能够读取到最新的值，但是效率低下。因此，我们需要一种机制来保证
当配置源的配置项改变时，能够及时且高效地反馈到程序中。

为了解决上述两个痛点，我设立manyfaced项目。manyfaced的名字来源于HBO热播的电视剧《The Game Of Throne》中的千面之神（Many-Faced God）。
由于本人非常喜爱《The Game Of Throne》这部电视剧；且这个项目是为了让配置项动态地从配置源获取最新的值，在程序运行时表现出不同的值，展现出
“千面”的特点。所以项目取名为manyfaced。

# 特点

1. 提供了对所有配置源的抽象，将配置源抽象成两类：轮询的配置源（PolledConfigurationSource）和监视的配置源（WatchedConfigurationSource）。
所有的配置只要实现这两类的接口，就可以被manyfaced正常使用。

2. 提供了一个轮询和通知回调的框架，当配置源内容发生变化时，轮询和通知回调的框架能够及时发现配置源的变化，从而修改内存配置对象Configuration的值。

3. 提供了配置项的高效动态获取，当配置源的内容发生改变时，配置项能够较实时改变，而不需要重启应用程序，也不需要每次都读取配置源。

4. 提供了轮询方式的DB(DBConfigurationSource)、本地配置文件(FileConfigurationSource)、HTTP访问(URLConfigurationSource)的配置源的实现；
提供了监视方式的本地配置文件(WatchedFileConfigurationSource)的配置源。

# 使用

1. 设置动态配置源。

动态配置源由配置源（Source）和调度器(Scheduler)组合而成。
目前提供的动态配置源由DynamicFileConfiguration、DynamicDBConfiguration和DynamicURLConfiguration。

```
    //1s 轮询解析/dir/to/config.properties文件一次。
    dynamicConfiguration, err := manyfaced.NewDynamicFileConfiguration([]string{"/dir/to/config.properties"}, 
        time.Millisecond * 100, time.Millisecond * 100, time.Millisecond * 1000)
    if err := nil {
        //TODO: Handle this error
    }
    defer dynamicConfiguration.Close()
```

- 如果需要单独创建配置源。

示例：

```
   fileSources := source.NewFileConfigurationSource([]string{"/dir/to/config.propertiers"}, 100 * time.Millisecond)
```

目前实现的配置源有：db、localfile和http，zookeeper的配置源由于时间关系没能够来得及实现，先占坑，后续再实现。

localfile目前只实现了properties格式的文件的解析。如果要添加其他格式的文件解析，需要如下步骤：

```
//例如：这是实现Toml文件的解析的类。
//Step 1: 实现 ConfigFileParser 这个Interface
type TomlFileParser struct {}

func (p *TomlFileParser)GetParserType() string {
    return "toml"
}
    
func (p *TomlFileParser)Parse(file string) (map[string]interface{}, string) {
    //TODO: Add your code.
}

//Step 2: 将TomlFileParser 这个struct 注册到 manyfaced中。
parser.RegisterParser(&TomlFileParser{})

//这样FileConfigurationSource就可以自动解析toml文件了。
```

- 如果需要单独设置轮询调度器。

提供抽象的调度器BasePollScheduler 和一个固定周期调度执行的FixedDelayPollScheduler。
如果FixedDelayPollScheduler 无法满足你的需求，你需要实现PollScheduler接口。
一般来说，你只需要组合继承BasePollScheduler，然后实现Schedule(task func()) 和Stop()两个方法。

```
type SpecificPollScheduler struct {
    *BasePollScheduler
}

func (s *SpecificPollScheduler)Schedule(task func()) {
    //TODO: Add your code
}

func (s *SpecificPollScheduler)Stop() {
    //TODO: Add your code
}
```

2. 将动态配置源加入manyfaced。

```
manyfaced.AddNamedConfiguration(dynamicConfiguration, "config")
```

3. 读取配置源的配置项。

**正确的姿势**

```
    //定义动态属性的变量
    dynamicTimeout := manyfaced.GetIntProperty("http.timeout", 1000)
    //获取动态属性的值
    SetTimeout(dynamicTimeout.Get())
```

这样能够每次都取到的http.timeout都是最新的值。

**错误的姿势**

```
    timeout := manyfaced.GetIntProperty("http.timeout", 1000).Get()
    SetTimeout(timeout)
```

最好不要将manyfaced取到的配置项的值保存到内存变量中。因为这样可能很容易导致
当配置源改变时，也无法影响到内存变量。

# TODO
由于时间关系，未来得及提供单测、示例和说明文档。
添加单测。
添加示例。
补充说明文档。

# 联系
QQ: 525999199@qq.com