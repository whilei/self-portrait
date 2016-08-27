#!/bin/bash

#creating a vid of the selfies
outDir="selfless"
mkdir -p $outDir



#download available selfies
while read -r line
do
  if  [[ $line =~ href=* ]]&&[[ $line =~ png ]]  ;
then
    png=$(echo "$line"| sed 's/<a href="//'|sed 's/">.*//')
    #only get it if not here yet
    if [ -f "$png" ];
		then
 		    echo "File $png exists."
		else
   			curl "http://isaacardis.com/selfies/$png" >$outDir/"$png"
		fi
    echo "$png"
fi
done < <(curl -s "http://isaacardis.com/selfies/?C=M;O=D")

