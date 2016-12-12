# Take a selfie as a cron job.
Keepin it candid:
Cron calls selfie script every 6 minutes.
Selfie script only executes half the time and pauses randomly between 1 second and 12 minutes.

All selfies are available [here](http://isaacardis.com/selfies/?C=M;O=D).

> [randomness reference](http://www.nightbluefruit.com/blog/2009/03/run-a-cron-job-at-multiple-random-times/)

> selfie script uses [imagesnap](http://iharder.sourceforge.net/current/macosx/imagesnap/)

### Dependencies
```shell
$ brew install imagesnap # self-portrait.sh
$ brew install imagemagick # self-ish.sh
$ brew install opencv # detect.go
$ go get -v github.com/lazywei/go-opencv # detect.og
```
