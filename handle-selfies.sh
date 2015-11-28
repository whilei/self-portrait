# Update website source from original selfies. 
gifme ~/Pictures/self-portraits/*.png -o ~/dev/isaacardis/app/images/gifs/da-gif.gif -w 280 
rsync -avz --update ~/Pictures/self-portraits/ ~/dev/isaacardis/app/images/selfies

# Create angular constant from all file names. 
ruby ~/dev/self-portrait/selfie-handler.rb

# resize all pngs in website images to max 140x140. 
# (originals still at full size in Pictures/)
# mogrify -resize 140x140 ~/dev/isaacardis/app/images/*.png 
