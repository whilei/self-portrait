# Da whole dang ding.
#rsync -avz --update ~/dev/isaacardis/app/ areteh.co:isaacardis/

# Just images and gifs.
rsync -avz --update ~/dev/isaacardis/app/images/selfies/ areteh.co:isaacardis/images/selfies/
rsync -avz --update ~/dev/isaacardis/app/images/gifs/ areteh.co:isaacardis/images/gifs/
rsync -avz --update ~/dev/isaacardis/app/scripts/services/selfielist.js areteh.co:isaacardis/scripts/services/selfielist.js

# Bower components.
rsync -avz --update ~/dev/isaacardis/bower_components/ areteh.co:isaacardis/bower_components/

