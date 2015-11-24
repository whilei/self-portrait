# copies files from icloud textastic to  directory

#!/bin/sh

## Variables

app_name=$1
where_to=$2
ttdir="/Users/ia/Library/Mobile Documents/M6HJR9W95L~com~textasticapp~textastic/Documents/"

## Functions

copier () {
        local name="$1"
	local whereto="$2"
        rsync -avz --update "$ttdir""$name""/" "$whereto"/
        echo "Copied ""$name"" from iCloud Textastic dir at ""$ttdir""$name"" into ""$whereto"
}

copier "$app_name" "$where_to"
