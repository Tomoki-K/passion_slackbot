# Passion Slackbot
「パッション」というと「パッションが足りません」と返してくれるだけの役立たずbot(の予定だった)

機能たち：
- 「パッション」というと「パッションが足りません」と返してくれる
- 5%の確率で違うことをいう
- 「〇〇の画像」というと画像を返してくれる
- 「申し訳ない」というと博士の画像を送ってくれる
- メッセージ履歴を削除できる

## setup
```
git clone https://github.com/Tomoki-K/passion_slackbot.git
export SLACK_PASSION_KEY="your api key"
make build
make run
```

supervisorでデーモン化してるので簡単には落ちない...はず
