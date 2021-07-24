**程序员的mac一些小技巧**

### 小问题

**brew install卡**
暂时禁止自动更新，关闭窗口后失效

``` BASH
export HOMEBREW_NO_AUTO_UPDATE=true
```
### 命令替代

**tree命令**

``` BASH
alias tree="find . -print | sed -e 's;[^/]*/;|____;g;s;____|; |;g'"
```