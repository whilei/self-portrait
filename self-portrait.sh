#!/bin/bash

$p=mt_rand(1,2); # execute 50% of the time
if ( $p != 1) {
	exit;
} else {

$s=mt_rand(1, 720); # wait between 1s and 12m
sleep($s);
imagesnap ~/Pictures/self-portraits/$(date +%y%m%d)-$(date +%H%M).png

}