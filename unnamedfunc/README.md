# find unnamed function - 名前無しの引数だけで構成された関数の使いみちを知りたい

## 目的

名前無しの引数だけで構成された関数の使いみちを知りたい。そこで、標準パッケージ内での使用例を探す。

## 背景

言語仕様を見ていたら、名前無しの引数だけで構成された関数が宣言できることを知りました！

```go!
// https://go.dev/play/p/QVqCWysth5Q
package main

// 名前無しの引数だけで構成された関数は宣言できる(https://go.dev/ref/spec#Function_types)
func f(int, int) {}

func main() {}
```

が、この記述の使いみちが分かりません。そして、それに興味があります！ どなたがご存じないでしょうか？

興味をまとめると...

- Q1. 使いみちがわからない！ 別に困っているわけではないけれど)
  - シグネチャーを壊さないまま実質的な引数を減らす用途はありそう？
  - `//go:linkname`を使ったパターンはある様子
- Q2. `f(3, 4)`とcallされたときに`f`内で引数を受け取る方法はある？(たぶんなさそうだけど)

## 結果

```
$ go vet -vettool=unnamedfunc std            
# runtime/debug
/usr/local/go/src/runtime/debug/stubs.go:12:1: readGCStats 名前無し引数を使った関数
/usr/local/go/src/runtime/debug/stubs.go:14:1: setMaxStack 名前無し引数を使った関数
/usr/local/go/src/runtime/debug/stubs.go:15:1: setGCPercent 名前無し引数を使った関数
/usr/local/go/src/runtime/debug/stubs.go:16:1: setPanicOnFault 名前無し引数を使った関数
/usr/local/go/src/runtime/debug/stubs.go:17:1: setMaxThreads 名前無し引数を使った関数
/usr/local/go/src/runtime/debug/stubs.go:18:1: setMemoryLimit 名前無し引数を使った関数
# crypto/ecdsa
/usr/local/go/src/crypto/ecdsa/notboring.go:11:1: boringPublicKey 名前無し引数を使った関数
/usr/local/go/src/crypto/ecdsa/notboring.go:14:1: boringPrivateKey 名前無し引数を使った関数
# crypto/internal/boring
/usr/local/go/src/crypto/internal/boring/notboring.go:44:1: SHA1 名前無し引数を使った関数
/usr/local/go/src/crypto/internal/boring/notboring.go:45:1: SHA224 名前無し引数を使った関数
/usr/local/go/src/crypto/internal/boring/notboring.go:46:1: SHA256 名前無し引数を使った関数
/usr/local/go/src/crypto/internal/boring/notboring.go:47:1: SHA384 名前無し引数を使った関数
/usr/local/go/src/crypto/internal/boring/notboring.go:48:1: SHA512 名前無し引数を使った関数
/usr/local/go/src/crypto/internal/boring/notboring.go:53:1: NewGCMTLS 名前無し引数を使った関数
# crypto/internal/boring/bcache
/usr/local/go/src/crypto/internal/boring/bcache/cache.go:41:1: registerCache 名前無し引数を使った関数
# crypto/rsa
/usr/local/go/src/crypto/rsa/notboring.go:11:1: boringPublicKey 名前無し引数を使った関数
/usr/local/go/src/crypto/rsa/notboring.go:14:1: boringPrivateKey 名前無し引数を使った関数
# go/printer_test
/usr/local/go/src/go/printer/example_test.go:19:1: Test 名前無し引数を使った関数
# html/template
/usr/local/go/src/html/template/exec_test.go:721:1: vfunc 名前無し引数を使った関数
# internal/reflectlite
/usr/local/go/src/internal/reflectlite/value.go:324:1: chanlen 名前無し引数を使った関数
/usr/local/go/src/internal/reflectlite/value.go:325:1: maplen 名前無し引数を使った関数
/usr/local/go/src/internal/reflectlite/value.go:390:1: unsafe_New 名前無し引数を使った関数
# os/signal
/usr/local/go/src/os/signal/signal_unix.go:15:1: signal_disable 名前無し引数を使った関数
/usr/local/go/src/os/signal/signal_unix.go:16:1: signal_enable 名前無し引数を使った関数
/usr/local/go/src/os/signal/signal_unix.go:17:1: signal_ignore 名前無し引数を使った関数
/usr/local/go/src/os/signal/signal_unix.go:18:1: signal_ignored 名前無し引数を使った関数
# os/signal/internal/pty
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b665/_cgo_gotypes.go:19:1: _Cgo_use 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b665/_cgo_gotypes.go:38:1: _cgo_runtime_cgocall 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b665/_cgo_gotypes.go:41:1: _cgoCheckPointer 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b665/_cgo_gotypes.go:44:1: _cgoCheckResult 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b665/_cgo_gotypes.go:91:1: _cgo_runtime_gostring 名前無し引数を使った関数
# os/exec_test
/usr/local/go/src/os/exec/exec_posix_test.go:29:1: cmdPwd 名前無し引数を使った関数
/usr/local/go/src/os/exec/exec_test.go:213:1: cmdPipeTest 名前無し引数を使った関数
/usr/local/go/src/os/exec/exec_test.go:234:1: cmdStdinClose 名前無し引数を使った関数
/usr/local/go/src/os/exec/exec_test.go:279:1: cmdStderrFail 名前無し引数を使った関数
# os/user
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b666/_cgo_gotypes.go:19:1: _Cgo_use 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b666/_cgo_gotypes.go:76:1: _cgo_runtime_cgocall 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b666/_cgo_gotypes.go:79:1: _cgoCheckPointer 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b666/_cgo_gotypes.go:82:1: _cgoCheckResult 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b666/_cgo_gotypes.go:88:1: _cgo_runtime_gostring 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b666/_cgo_gotypes.go:229:1: runtime_throw 名前無し引数を使った関数
# plugin
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b673/_cgo_gotypes.go:19:1: _Cgo_use 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b673/_cgo_gotypes.go:40:1: _cgo_runtime_cgocall 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b673/_cgo_gotypes.go:43:1: _cgoCheckPointer 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b673/_cgo_gotypes.go:46:1: _cgoCheckResult 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b673/_cgo_gotypes.go:51:1: _cgo_runtime_gostring 名前無し引数を使った関数
# reflect
/usr/local/go/src/reflect/value.go:2827:1: rselect 名前無し引数を使った関数
/usr/local/go/src/reflect/value.go:2972:1: unsafe_New 名前無し引数を使った関数
/usr/local/go/src/reflect/value.go:2973:1: unsafe_NewArray 名前無し引数を使った関数
/usr/local/go/src/reflect/export_test.go:91:1: gcbits 名前無し引数を使った関数
# runtime/cgo
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b780/_cgo_gotypes.go:16:1: _Cgo_use 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b780/_cgo_gotypes.go:23:1: _cgo_runtime_cgocall 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b780/_cgo_gotypes.go:26:1: _cgoCheckPointer 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b780/_cgo_gotypes.go:29:1: _cgoCheckResult 名前無し引数を使った関数
# runtime/metrics
/usr/local/go/src/runtime/metrics/sample.go:25:1: runtime_readMetrics 名前無し引数を使った関数
# sync
/usr/local/go/src/sync/mutex.go:20:1: throw 名前無し引数を使った関数
/usr/local/go/src/sync/mutex.go:21:1: fatal 名前無し引数を使った関数
# net
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b602/_cgo_gotypes.go:19:1: _Cgo_use 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b602/_cgo_gotypes.go:69:1: _cgo_runtime_cgocall 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b602/_cgo_gotypes.go:72:1: _cgoCheckPointer 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b602/_cgo_gotypes.go:75:1: _cgoCheckResult 名前無し引数を使った関数
/var/folders/xw/4p268kr944q0t190f6ssm5kr0000gn/T/go-build1482300308/b602/_cgo_gotypes.go:132:1: _cgo_runtime_gostring 名前無し引数を使った関数
# time
/usr/local/go/src/time/sleep.go:41:1: startTimer 名前無し引数を使った関数
/usr/local/go/src/time/sleep.go:42:1: stopTimer 名前無し引数を使った関数
/usr/local/go/src/time/sleep.go:43:1: resetTimer 名前無し引数を使った関数
# text/template
/usr/local/go/src/text/template/exec_test.go:729:1: vfunc 名前無し引数を使った関数
# time/tzdata
/usr/local/go/src/time/tzdata/tzdata.go:34:1: registerLoadFromEmbeddedTZData 名前無し引数を使った関数
# net/http_test
/usr/local/go/src/net/http/fs_test.go:1279:1: TestLinuxSendfileChild 名前無し引数を使った関数
# runtime
/usr/local/go/src/runtime/cgo.go:45:1: cgoUse 名前無し引数を使った関数
/usr/local/go/src/runtime/lock_sema.go:300:1: beforeIdle 名前無し引数を使った関数
/usr/local/go/src/runtime/os_darwin.go:94:1: sigNoteSetup 名前無し引数を使った関数
/usr/local/go/src/runtime/os_darwin.go:114:1: sigNoteWakeup 名前無し引数を使った関数
/usr/local/go/src/runtime/os_darwin.go:120:1: sigNoteSleep 名前無し引数を使った関数
/usr/local/go/src/runtime/stubs.go:299:1: goexit 名前無し引数を使った関数
# time_test
/usr/local/go/src/time/example_test.go:152:1: handle 名前無し引数を使った関数
```

