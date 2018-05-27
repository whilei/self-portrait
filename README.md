## Take a selfie as a cron job.
Keepin it candid:
Cron calls selfie script every 6 minutes.
Selfie script only executes half the time and pauses randomly between 1 second and 12 minutes.

All selfies are available [here](http://isaacardis.com/selfies/?C=M;O=D).

## Calendar

Average face-detected-and-cropped images by month since cron installed in November 2015.

| YY-MM | avg image (recursive avg fit) |
| --- | --- |
| 15-11 | ./calendar/15-11.avg-recurfxfit.png |
| 15-12 | ./calendar/15-12.avg-recurfxfit.png |
| 16-01 | ./calendar/16-01.avg-recurfxfit.png |
| 16-02 | ./calendar/16-02.avg-recurfxfit.png |
| 16-03 | ./calendar/16-03.avg-recurfxfit.png |
| 16-04 | ./calendar/16-04.avg-recurfxfit.png |
| 16-05 | ./calendar/16-05.avg-recurfxfit.png |
| 16-06 | ./calendar/16-06.avg-recurfxfit.png |
| 16-07 | ./calendar/16-07.avg-recurfxfit.png |
| 16-08 | ./calendar/16-08.avg-recurfxfit.png |
| 16-09 | ./calendar/16-09.avg-recurfxfit.png |
| 16-10 | ./calendar/16-10.avg-recurfxfit.png |
| 16-11 | ./calendar/16-11.avg-recurfxfit.png |
| 16-12 | ./calendar/16-12.avg-recurfxfit.png |
| 17-01 | ./calendar/17-01.avg-recurfxfit.png |


### Dependencies
```shell
$ brew install imagesnap # self-portrait.sh
$ brew install imagemagick # self-ish.sh
$ brew install opencv # detect.go
$ go get -v github.com/hybridgroup/gocv # detect.og
# Use or reference ./build.face-detector.sh to build bin.
```
