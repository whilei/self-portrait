#!/bin/bash

#goal is to create a time travel vid of the selfies
outDir="selfless"
mkdir -p $outDir



#download available selfies
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

#Now a movie maker
#- has some menial work 
#brew install ffmpeg
#-framerate is the number of frames per second...
#Not sure if the glob is globbing in perfect order yet...but looks like there are options to do that
#ffmpeg -framerate 6 -pattern_type glob -i '*.png' -c:v libx264 -pix_fmt yuv420p out.mp4
# re-scale varsion
#ffmpeg -framerate 6 -pattern_type glob -i '*.png' -c:v libx264 -pix_fmt yuv420p -vf "scale=300:300,format=yuv420p" out.mp4 

#now upload that sucker to youtube... check out https://github.com/tokland/youtube-upload
#youtube-upload --title="selfless" --client-secrets $outDir/client_secrets.json out.mp4
