#Hide&Seek🎮

![](preview.png)

**_You can install the game using insructions below_**

Create a dir

```
$ go get github.com/YarikRevich/Hide-Seek-with-Guns
```

Go to sources
```
$ cd $(go env GOPATH)/src/github.com/YarikRevich/Hide-Seek-with-Guns
```

Use Bazel

```
$ sudo bazel build :deps
$ sudo bazel build --action_env=USER=$USER :service_env
$ sudo bazel build //cmd:hide-seek-client 
```

Run and enjoy the game👌

```
$ HideSeek
```

Use 

```
$ HideSeek --help
```

To see all available flags

😊 If you want to contribute you definetily should use [debug cli](DEBUG.md)


