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

selflessMovie=$outDir/selfless.mp4


#now upload that sucker to youtube... check out https://github.com/tokland/youtube-upload
#cd $outDir
#wget https://github.com/tokland/youtube-upload/archive/master.zip
# unzip master.zip
# cd youtube-upload-master
# sudo python setup.py install

#below has some menial work to upload

$outDir/bin/youtube-upload --title="selfless" --client-secrets $outDir/client_secrets.json $selflessMovie
