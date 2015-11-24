# Take a selfie as a cron job. 
Keepin it candid: 
Cron calls selfie script every 6 minutes.
Selfie script only executes half the time and pauses randomly between 1s and 12m.


in PHP
<pre>
Run a cron job at multiple random times
Say you want to run a cron job 10 times (or so) in a day at random times.

Here’s my solution:

1. Create a probability test that gives a 10% probability of the outcome you want

$p = mt_rand(1,10);

2. Add in a bit of extra randomness with a short sleep

$s = mt_rand(60,300);
sleep($s);

giving:

$p = mt_rand(1,10);

if ($p != 1) {

exit;
} else {

$s = mt_rand(60,300);
sleep($s);

Your stuff here…
}

3. Run your cron job 4 times an hour, 24 hours a day

*/15 * * * * /usr/bin/php myscript.php

At this, your cron job will run 96 times per day, and execute 1 in 10 of those times, which gives you 9 to 10 executions per day at random times.

OK, its not totally random, and you can’t guarantee the number of executions, but if you haveÂ  fixed number of executions per day, that’s not really random, is it?

Nod, nod, wink, wink….
</pre>
> [http://www.nightbluefruit.com/blog/2009/03/run-a-cron-job-at-multiple-random-times/](http://www.nightbluefruit.com/blog/2009/03/run-a-cron-job-at-multiple-random-times/)

> selfie script uses [imagesnap](http://iharder.sourceforge.net/current/macosx/imagesnap/)
