1) How to stop `mtxagent`

First, edit the cron tab-file (with `sudo crontab -e` and comment out the cronjob for `mtxagent` and save (as you would with any file in Vim (with `<esc>:x`). Then, open up activity monitor and kill the `mtxagent` process.
