# OnePassword Records

All you need is to remember only ONE password to access all other recorded apps' passwords.

Run locally, nothing to do with the network. All processes happen in memory. No record for your entry password, other recorded passwords are encrypted with your entry password (by AES-CBC).

## Note

1. If you don't use the default filename to save the records, make sure you type the whole name (including name extension) that observes your system's naming rules (and don't include spaces).
1. Before operating the records, make sure your entry password is correct by checking its hash in a `[INFO]` hint.
1. Your username and password should not include space (will be truncated at space).
1. If a password is shown empty, there must be something wrong with your encoded password (e.g., being modified outside), just update it or delete it.
1. You can send your recording file to others or transfer to a different machine, as long as they run the same program and get your entry password.

## Todo

1. Generator for ramdom passwords
2. Modify the entry password (and choose whether to make a new copy or overwrite the original file)

------

只需要记住一个口令，就可以读取其他应用的密码。（不同应用的密码类型和格式往往不一样，导致密码也不一致，容易记混的同时也容易遗忘。另外多个应用共用一套密码也会有一定的安全问题。只记一个自己熟悉的口令，就可以不用再管其他的密码了。）

本地运行，与网络没有任何关联。所有过程都是在内存中发生的，入口口令不会被记录，其他被记录的密码都通过入口口令和AES-CBC进行对称加密，因此整个过程都是安全的。

## 提示

1. 如果没有使用程序内默认的文件名来存储密码记录，请确保自己输入的文件名是完整的（包括文件后缀的扩展名），且文件名遵循系统的命名规则（不要带空格）。
1. 在对记录做后续的具体操作前，可以通过一条`[INFO]`提示里的口令哈希值来确认输入的入口口令是正确的。
1. 设置的用户和密码不应带空格（否则会在空格处截断）。
1. 如果一个展示的密码为空，那在记录文件中加密的密码一定存在某些问题（比如在程序外被修改了）。此时只需要更新或者删除该条目即可。
1. 可以将记录文件发给其他人或转移到另一台机器上，只需要他们运行了相同的程序并持有你的入口口令，就可以读取文件内容，实现共享。
