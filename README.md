## Take a selfie as a cron job.
Keepin it candid:
Cron calls selfie script every 6 minutes.
Selfie script only executes half the time and pauses randomly between 1 second and 12 minutes.

All selfies are available [here](http://isaacardis.com/selfies/?C=M;O=D).

### Dependencies
```shell
$ brew install imagesnap # self-portrait.sh
$ brew install imagemagick # self-ish.sh
$ brew install opencv # detect.go
$ go get -v github.com/lazywei/go-opencv # detect.og
```
Average face I make over the course of about a year:
![Face average beginning (~201511xx) ->
20161212](./face_averages/self-ish-recurfxfit.png)

Average picture of the same course (not auto-face-detect-cropped):
![Original average beginning (~201511xx) ->
20161215](./original_averages/self-ish-recurfxfit.png)

