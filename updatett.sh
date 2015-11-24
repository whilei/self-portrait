# copies files from a current working directory into icloud textastic 

#!/bin/sh

## Variables
foldertocopy=$1
app_name=$2

# iCloud Textastic folder
destination="/Users/ia/Library/Mobile Documents/M6HJR9W95L~com~textasticapp~textastic/Documents/"

## Functions

copier () {
	local name="$1" 
	local folder="$2"
	rsync -avz --update "$folder"/ "$destination""$name"
	echo "Updated ""$foldertocopy"" as ""$name"" into iCloud Textastic dir at ""$destination""$name""."
}

copier "$app_name" "$foldertocopy"
