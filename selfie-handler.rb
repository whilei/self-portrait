selfies = Dir['/Users/ia/Pictures/self-portraits/*']
selfies_string = selfies.to_s
selfies_string_gsubbed = selfies_string.gsub("/Users/ia/Pictures/self-portraits/", "")
intro = "'use strict'; angular.module('isArtApp').constant('SelfieList', {SELFIES:"
outro = "});"
string_to_write = intro+selfies_string_gsubbed+outro

File.write('/Users/ia/dev/isaacardis/app/scripts/services/selfielist.js', string_to_write)
