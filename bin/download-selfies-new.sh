#!/usr/bin/env bash

# Download only selfies not existing already in given directory.

outDir="selfless"
if [ $# -gt 0 ]; then
    outDir="$1"
fi
mkdir -p $outDir

while read -r line
do
  if  [[ $line =~ href=* ]]&&[[ $line =~ png ]]  ;
    then
      png=$(echo "$line"| sed 's/<a href="//'|sed 's/">.*//')

        #only get it if not here yet
      if [ -f "$outDir/$png" ];
		    then
 		    echo "File $png exists."
		  else
			   #link it home
   			 curl "http://isaacardis.com/selfies/$png" >$outDir/"$png"
		  fi
  fi
done < <(curl -s "http://isaacardis.com/selfies/?C=M;O=D")