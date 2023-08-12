
func _title_creator (){
_arguments \
    '--text[Text to render]:text:_files' \
    '--characters[Characters to render]:characters:_files' \
    '--resolution[Text resolution]:resolution:' \
    '--aspect[Character height to width]:aspect:' \
    '--font[Font filename]:font:_files' \
    '--size[Font size in points]:size:' \
    '--max-width[Maximum width to render]:max_width:' \
    '--mode[Render mode]:mode:' \
    '--allow-inverted[Use inverted characters]:allow_inverted:(true false)' \
    '--load[Load saved character map]:load:_files' \
    '--save[Save character map]:save:_files' \
    '--output[Save output to file]:output:_files' \
    '*: :_files'
}

compdef _title_creator title_creator
