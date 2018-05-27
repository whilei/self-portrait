# Just images and gifs.
# Push to remote host for websiting and storage. 
#rsync -avz --update ~/dev/isaacardis/app/images/selfies/ freya:~/isaacardis/images/selfies/
#rsync -avz --update ~/dev/isaacardis/app/images/gifs/ freya:~/isaacardis/images/gifs/
#rsync -avz --update ~/dev/isaacardis/app/scripts/services/selfielist.js freya:~/isaacardis/scripts/services/selfielist.js

rsync -avz --update ~/Pictures/self-portraits/ freya:~/isaacardis.com/selfies/ 

