### How to stop mtxagent ###

First, edit the cron tab-file (with `sudo crontab -e` and comment out the cronjob for `mtxagent` and save (as you would with any file in Vim (with `<esc>:x`). Then, open up activity monitor and kill the `mtxagent` process.

### How to correct symbolic link issues in packrat ###

When installing new packages using `packrat` (particularly when installing packages from Bioconductor) there are symbolic link issues with the `lib-R` directory. This is pretty easily dealt with by running the following commadn andn restarting the R session.

```unlink("packrat/lib-R", recursive = TRUE)```
