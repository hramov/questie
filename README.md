# Questie
### Telegram bot for making text quests
### Based on finite state machine Golang implementation
To add new quest
- create file inside <code>internal/quests</code> folder and implement method <pre><code>InitSteps() // example - birthday.go</code></pre>
- add new quest title in <code>config.go</code>
- add new case statement in <code>main.go</code>
- try it on your Telegram bot (don't forget to create .env file with <code>TG_TOKEN</code> variable)