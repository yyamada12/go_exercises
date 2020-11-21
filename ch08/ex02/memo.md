# FTP

## ドキュメント

- RFC
  http://srgia.com/docs/rfc959j.html

- @IT インターネット・プロトコル詳説
  https://www.atmarkit.co.jp/ait/articles/0107/17/news002.html

## TODO

- [x] USER
- [x] PORT
- [x] TYPE
- [x] MODE
- [ ] STRU
- [ ] RETR
- [ ] STOR
- [ ] NOOP
- [ ] QUIT

## 最小の実装

型 - ASCII Non-print
モード - ストリーム
構造 - ファイル、レコード
コマンド - USER, QUIT, PORT,
TYPE, MODE, STRU,(デフォルト値のためのもの)
RETR, STOR,
NOOP.

### コード

500 (文法エラー、そのコマンドは認識されない)
コマンド行が長すぎるといったエラーもこれに含まれて良い。
501 (パラメータや引数の文法エラー)
504 (そのコマンドに対するその引数は実装されていない)
421 (サービスは利用不可であり、コントロール接続は閉じられようとしている)
シャットダウンしなければならないことをサービスが分かっている場合、
これは任意のコマンドに対するリプライとなりうる。
530 (ログインしていない)

### USER

USER <SP> <ユーザー名> <CRLF>

USER
230: ログインした
530: This is a private system - No anonymous login
500, 501, 421
331, 332: パスワード、課金情報が必要

### QUIT

QUIT <CRLF>
221: サービスはコントロール接続を閉じようとしている
500

### PORT

PORT <SP> <ホスト-ポート> <CRLF>
200: OK
500, 501, 421, 530

200 PORT command successful
501 Syntax error in IP address
530 You aren't logged in
501 Active mode is disabled
501 Sorry, but I won't connect to ports < 1024

### TYPE

TYPE <SP> <型コード> <CRLF>
<型コード> ::= A [<sp> <型コード>]
| E [<sp> <型コード>]
| I
| L <sp> <バイトサイズ>
200: OK
500, 501, 504, 421, 530

### MODE

MODE <SP> <モードコード> <CRLF>
<モードコード> ::= S | B | C

### STRU

STRU <SP> <構造コード> <CRLF>
<構造コード> ::= F | R | P
200
500, 501, 504, 421, 530

### RETR

RETR <SP> <パス名> <CRLF>
125: データ接続はすでに開かれている　転送が開始される
150: ファイル状態は正常　データ接続を開こうとしている
(110): 再開マーカーリプライ
226: データ接続は閉じられようとしている　要求されたファイル操作は成功した(例えばファイル転送や中断)
250: 要求されたファイル操作は正常に完了した
425: データ接続を開けない
426: 接続は閉じられた　転送が中断された
451: 要求された操作は中止された　処理中にローカルエラー発生
450: 要求されたファイル操作は実行されない
550: 要求された操作は実行されない
500, 501, 421, 530

### STOR

STOR <SP> <パス名> <CRLF>
125: データ接続はすでに開かれている　転送が開始される
150: ファイル状態は正常　データ接続を開こうとしている
(110): 再開マーカーリプライ
226: データ接続は閉じられようとしている　要求されたファイル操作は成功した(例えばファイル転送や中断)
250: 要求されたファイル操作は正常に完了した
425: データ接続を開けない
426: 接続は閉じられた　転送が中断された
451: 要求された操作は中止された　処理中にローカルエラー発生
551: 要求された操作は中止された　ページタイプが不明である
552: 要求されたファイル操作は中止された　現在のディレクトリやデータセットのための保存領域を超過した
532: ファイルを保存するには課金情報が必要
450: 要求されたファイル操作は実行されない
452: 要求された操作は実行されない　システムに十分な空き領域がない
553: 要求された操作は実行されない　許可されないファイル名
500, 501, 421, 530

### NOOP

NOOP <CRLF>
200
500 421

## データ接続

### 通常 アクティブ状態

サーバー　 → クライアント

サーバーのポートは 20
クライアントのポートをサーバーに伝える必要あり
これは PORT コマンドホスト・ポートを伝える
これができていないと 425 No data connection
接続できない

### パッシブ

クライアント → サーバー
サーバーのポートはランダムに割り当て、クライアントに伝えて、そのポートに繋いでもらう。
これは PASV コマンドに対してホスト・ポートを返答することで行う。

## local で動かす

```
docker start ftpd_server
docker logs -f ftpd_server
ftp test@localhost
```

参考: https://www.sukerou.com/2019/07/dockerftp.html

## 図 2

                  コントロール   +----------------+ コントロール
                    +--------->| ユーザー側 FTP   |<----------+
                    |          | ユーザー側 PI    |           |
                    |          |      "C"       |           |
                    V          +----------------+           V
            +----------------+                        +----------------+
            | サーバー側 FTP   |      データ接続          | サーバー側 FTP   |
            |      "A"       |<---------------------->|      "B"        |
            +----------------+ ポート(A)    ポート(B)   +----------------+


                ＼    ／
       A - ASCII |    | N - Non-print
                 |-><-| T - Telnet 書式制御文字
       E - EBCDIC|    | C - 改行制御 (ASA)
                ／    ＼
       I - Image

       L <バイト幅> -ローカルバイトのバイト幅

## const

```

const (
	StatusRestartMarkerReply    = "110"
	StatusNotReady              = "120"
	StatusTransferStarting      = "125"
	StatusOpeningDataConnection = "150"

	StatusOK                    = "200"
	StatusNotImplemented        = "202"
	StatusSystem                = "211"
	StatusDirectory             = "212"
	StatusFile                  = "213"
	StatusHelp                  = "214"
	StatusNameSystemType        = "215"
	StatusReadyForNewUser       = "220"
	StatusNoTransfer            = "225"
	StatusClosingDataConnection = "226"
	StatusPassiveMode           = "227"
	StatusLoggedIn              = "230"
	StatusActionCompleted       = "250"
	StatusCreated               = "257"

	StatusNeedPassword = "331"
	StatusNeedAccount  = "332"
	StatusPending      = "350"

	StatusNotAvailable             = "421"
	StatusCannotOpenDataConnection = "425"
	StatusTransferAborted          = "426"
	StatusFileBusy                 = "450"
	StatusLocalError               = "451"
	StatusInsufficientStorageSpace = "452"

	StatusInvalidCommand               = "500"
	StatusInvalidParameter             = "501"
	StatusCommandNotImplemented        = "502"
	StatusBadCommandSequence           = "503"
	StatusParameterNotMatchedToCommand = "504"
	StatusNotLoggedIn                  = "530"
	StatusNeedAccountForStoring        = "532"
	StatusFileUnavailable              = "550"
	StatusUnknownPageType              = "551"
	StatusExceededStorageAllocation    = "552"
	StatusFileNameNotAllowed           = "553"
)
```

## wire shark

```
220---------- Welcome to Pure-FTPd [privsep] [TLS] ----------
220-You are user number 1 of 5 allowed.
220-Local time is now 00:20. Server port: 21.
220-This is a private system - No anonymous login
220-IPv6 connections are also welcome on this server.
220 You will be disconnected after 15 minutes of inactivity.
USER test
331 User test OK. Password required
PASS test
230 OK. Current directory is /
SYST
215 UNIX Type: L8
FEAT
211-Extensions supported:
 EPRT
 IDLE
 MDTM
 SIZE
 MFMT
 REST STREAM
 MLST type*;size*;sizd*;modify*;UNIX.mode*;UNIX.uid*;UNIX.gid*;unique*;
 MLSD
 AUTH TLS
 PBSZ
 PROT
 UTF8
 TVFS
 ESTA
 PASV
 EPSV
 SPSV

211 End.
EPSV
229 Extended Passive mode OK (|||30006|)
LIST
150 Accepted data connection
226-Options: -l
226 0 matches total
```

```
220---------- Welcome to Pure-FTPd [privsep] [TLS] ----------
220-You are user number 1 of 5 allowed.
220-Local time is now 11:16. Server port: 21.
220-This is a private system - No anonymous login
220-IPv6 connections are also welcome on this server.
220 You will be disconnected after 15 minutes of inactivity.
user
530 This is a private system - No anonymous login

220---------- Welcome to Pure-FTPd [privsep] [TLS] ----------
220-You are user number 1 of 5 allowed.
220-Local time is now 11:16. Server port: 21.
220-This is a private system - No anonymous login
220-IPv6 connections are also welcome on this server.
220 You will be disconnected after 15 minutes of inactivity.
user tes
331 User tes OK. Password required
pass hoge
530 Login authentication failed
user test
331 User test OK. Password required
pass test
230 OK. Current directory is /
list
425 No data connection
type
501-Missing argument
501-A(scii) I(mage) L(ocal)
501 TYPE is now ASCII
type a
200 TYPE is now ASCII
type i
200 TYPE is now 8-bit binary
type a
200 TYPE is now ASCII
mode
501 Missing argument
mode s
200 S OK
mode b
504 Please use S(tream) mode
mode c
504 Please use S(tream) mode
```

```
220---------- Welcome to Pure-FTPd [privsep] [TLS] ----------
220-You are user number 1 of 5 allowed.
220-Local time is now 11:22. Server port: 21.
220-This is a private system - No anonymous login
220-IPv6 connections are also welcome on this server.
220 You will be disconnected after 15 minutes of inactivity.
USER test
331 User test OK. Password required
PASS test
230 OK. Current directory is /
SYST
215 UNIX Type: L8
FEAT
211-Extensions supported:
 EPRT
 IDLE
 MDTM
 SIZE
 MFMT
 REST STREAM
 MLST type*;size*;sizd*;modify*;UNIX.mode*;UNIX.uid*;UNIX.gid*;unique*;
 MLSD
 AUTH TLS
 PBSZ
 PROT
 UTF8
 TVFS
 ESTA
 PASV
 EPSV
 SPSV

211 End.
TYPE I
200 TYPE is now 8-bit binary
EPSV
229 Extended Passive mode OK (|||30003|)
STOR README.md
150 Accepted data connection
226-File successfully transferred
226 0.003 seconds (measured here), 416.06 Kbytes per second
TYPE A
200 TYPE is now ASCII
EPSV
229 Extended Passive mode OK (|||30003|)
LIST
150 Accepted data connection
226-Options: -l
226 1 matches total
TYPE I
200 TYPE is now 8-bit binary
SIZE README.md
213 1089
EPSV
229 Extended Passive mode OK (|||30008|)

RETR README.md
150 Accepted data connection
226-File successfully transferred
226 0.000 seconds (measured here), 7.15 Mbytes per second

MDTM README.md
213 20201117112224
SIZE hoge.txt
550 Can't check for file existence
EPSV
229 Extended Passive mode OK (|||30003|)
RETR hoge.txt
550 Can't open hoge.txt: No such file or directory
```
