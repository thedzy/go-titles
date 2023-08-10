# title_creator

Print in large font in your terminal/console \
Complete rewrites

This was an attempt to write th python version of title_creator in Go.  The idea was to improve the speed.  And yes, it is faster, it's not as reliable to render with all fonts.

``` console
  ,,,_    _gMMp,                 _qqq,   ,qqq
 MMMMMM   %MMMMMk                MMMMM  @MMMMB
 @MMMM"    MMMM"                 MMMMW  @MMMM
 @MMMM    ]MMMM                  MMMMB  @MMMM
 MMMMM    1MMMM                  MMMMB  @MMMM
 MMMMH    1MMMM       _gMMMMp,   MMMMB  @MMMM     ,pMMMMpq_
 @MMMH    dMMMM      pMMMMMMMMp  MMMMB  1MMMM    gMMMMMMMMMp
 @MMMM,qqpMMMMMp    MMMMP  %MMM  MMMMB  1MMMM   gMMM@   @MMMp
@MMMMMMMMMMMMMMMM  gMMMK   @MMM  MMMMB  1MMMM  _MMMM     MMMM
"MMMMMMMWWPMMMM"   MMMM   gMMM"  MMMMB  1MMMM  gMMMH     MMMM
 @MMMM     MMMM    MMMM_gMMMH`   MMMMB  1MMMM  MMMMB     MMMM
 @MMMM     MMMM    MMMMMMMH"     MMMMB  jMMMM  @MMMH    qMMMM
 @MMMM     MMMM    MMMMW"`  ,pp_ MMMMB  ]MMMM  'MMMM    @MMMW
 @MMMM    @MMMM    MMMMk   ,MMMM MMMMB  ]MMMM   @MMMp  gMMMM
 @MMMM    MMMMM@    MMMMMMMMMMH` MMMMH  gMMMM,   @MMMMMMMMM`
 MMMMM      `""      "MMMMMWP"   MMMMM  @MMMM#    "MMMMMM"

,,,,,,,,,         =,        ,,,,,,,,                                    ,,===[]                ,===[[                 ,=
 ''[]][          ,]]           ][]'                                        [[]}                  {[[[                []]
   []]]          ]]]]          [[                                          [][[                  ][[]                [[[
    []]         [}][],         ]/                                          [[[[                  ]][]                []]
    []][       {] '[[]        []             ,,,,              ,    ,,     {][[             ,,,  ]]]]                {]]
     ][[      =[   \[]]      ,]'         ,=='''''=]]=     ===[[[ ,=][]]    {[[          ,=='''=[][[[]                 []
     [[[,    ,]'    []]\     [}        ,]['       '[]],      [][=/'''[[    {]]       ,[]'       '{[]]                 [[
     ']]]    [/      [[],   ][        ,][[         ']]]      [[[           {[]      ,]['         ][[]                 []
      ][[,  [}       ']][   [/        [[]           ][[]     ][[           {[]      ]]]          ][[]                 []
      '[[] {]         {[[] {[         ][[[          ][]}     [[]           {[][     []]          ]]]]                 ='
       ][[=[           [[],]/         [][]          []]      ]][           []][     [[][         []]]
       ]][]'           '[[[[           [][=        ,[]'      [][           [[[}     '[]]=       ={]]]                ,==
        ][}             \]]'            \[[],    ,=]'       ,[[[,          [[[]       =]][]====' ][][,               [][


```

## What?

### title_creator2.py
title_creator is a single use script that will print a title in a window or to file

``` console
▇▀▀▇▀▀▊ ▆                            ■▎      ■▊         ▞▀▜                ▝▋                                 ▆▏                               ▖  ▇               ▆▏                      ▉                                     ▇
  ▕▉    █▃■▅▖  ▄■▅       ▖■▗▖ ▗▖  ▅  ▄▏ ▗▖■▅  ▊ ▗▖     ▐▋  ▁▖■▅▖ ▗▅ ▗▖     ▄▖ ▗▖  ▅▏ ▄▃■▅▖▄■▅  ▅▄■▅▖  ▖■▄  ▗■■▟▏     ▁▖■▅▖ ▅▖ ▗▖ ▄■▅▖ ▄▄▅▖    ■▍  █▃■▅▖  ▖■▅      ▜▏ ▗▖■▖ ▗■ ▅▖ ▅  ▗▖     ▉▄▅▄  ▄▄▅ ▃■■▄  ▅ ▗▅  ▅ ▄▃■■▖      ▄■▗█  ▗▖■▄  ▗■▗▄
   ▉    ▜▏ ▐▋ ▟▘■▝▘     ▉  ▐▌ ▐▌  █  ▜▏ ▊  ▔ ▕▊▆▔      ▐▋ ▕▊   █  ▝▆▘      ▕▋ ▕▋  ▟▎ ▐▍  ▉  ▜▎ ▜▍ ▕▊ ▟■■▝▘▐▋  ▜▎     ▉   █ ▝▙ ▞ ▐▍■▝▀ ▐▋      ▐▍  █  ▐▌ ▟ ■▝▘     ▟▏ ▝▃▕▊   ▟▀  ▝▌ ▞      ▉  ▜▌ ▜▏ ▐▌  ▐▋ ▜▖▗█▖▗▘ ▜▍  ▊     ▗▍  █ ▐▋  ▕▉ ▜▁▁▉
  ▕▉    ▟▏ ▐▋ ▝▙▂▂▖     ▜▂▁▟▌ ▐▙▁▃█▂ ▟▎ ▜▄▂▄▘▕▊▝▚▃     ▐▋  ▜▂ ▂▛ ▁▞▔▚▖     ▕▊ ▕▙▁▂▟▂ ▐▍  ▉  ▟▎ ▐▍▁▃▘ ▜▄▂▂▌▕▙▂▁▟▂     ▜▖ ▂▛  ▝▆▘ ▝▙▂▂▞ ▐▋      ▐▍▁ █  ▐▌ ▜▄▂▂▖     ▟▎ ▉▁▕▊▁▗▟▁▁▂  ▜▞       ▉ ▁▟▘ ▟▏ ▝▙▁ ▄▘  ▜▛ █▘  ▟▍  ▉     ▝▙▁▂█▂ ▚▁ ▗▛ ▟▄▄▄
  ▔▔▔  ▔▔▔ ▔▔▔ ▔▔▔       ▔▔▐▌  ▔▔ ▔ ▔▔▔  ▔▔  ▔▔▔▔▔     ▔▔▔  ▔▔▔  ▔▔ ▔▔    ▄▂▋  ▔▔▔▔  ▔▔ ▔▔▔ ▔▔ ▟▍▔▔   ▔▔▔   ▔▔▔       ▔▔▔    ▔    ▔▔  ▔▔▔      ▔ ▔▔▔ ▔▔  ▔▔▔     ▔▔▔ ▔▔ ▔ ▔▔▔▔▔ ▄▞        ▔▔▔  ▔▔▔   ▔▔       ▔   ▔▔ ▔▔▔     ▔▔▔▔   ▔▔▔ ▕▙▂▁▃
▗▆▇▇▇▇▆▖▟█▏                          ▐█▎      ▐█▍        ▆█▋                 ■▊                                  █▉                                 ▗▅ ▕█▊              ▟█                      ▐█▌                                     ▕█▋            ▂
▀▔▕█▛ ▀ ██▅▆▅  ▄▆▆▆▖    ▗▅▆▆▆▖ ▅▅ ▅▅ ▗▅▏ ▗▅▆▆▖▟█▏▗▅▎   ▗▆█▙▅ ▄▆▆▅▖▕▅▖ ▅▅     ▅▅ ▗▅▏▐▆▖ ▇▇▅▅▅▄▅▅▖ ▆▇▆▆▅  ▄▆▆▅▖ ▗▆▆██     ▄▆▆▅▖▕▅▖ ▕▆▖ ▄▆▆▅▖▗▅▄▆▆   ▕▅██▅▐█▙▆▆▖  ▅▆▆▅     █▊ ▅▆▆▅▖ ▗▆▅▆▅▗▅▖ ▗▅▏   ▐█▆▆▆▖ ▅▅▅▆▖▗▅▆▆▄ ▗▅▏▐▅▍▕▆▖▗▇▙▅▅▖     ▄▆▆█▊ ▗▅▆▆▄  ▄▆▆▇▛
  ▐█▌   █▉ ▐█▌▗█▁▃█▉    █▋ ▕█▌▐█▍ ▐█▎▟█ ▗█▘ ▝ ▟█▂█▛     ▐█▍ ▟▛▔ ▜█ ▜█▄█▘     ▜█ ▟▉  █▊▕█▛▔▜█▎▔█▉ █▊ ▔█▌▟▉▁▃█▋▗█▘ ▜█    ▗█▔ ▐█▏█▋ ▕█▍▟▉▁▃█▋▐█▛▔▀    ▕█▌ ▐█▌ ▜█ █▛▂▄█▍    █▋ ▀▁▃█▉  ▔▟█▀▕█▙ ▟█▏   ▐█▎ ▜█▏█▉▔▀▘█▛ ▔▜▋▜█▏▐█▍▕█▌▟█▘▔█▊    ▟█▘ █▉▐█▘ ▔█▌▐█▎ ▟▉
  ▟█▌   █▊ ▐█▌▜█▛▀▔     ▝█▅▄█▌▐█▍ ▐█▎▟█▏▜█    ▟█▛▜▇▖    ▐█▍ █▊  ▟█▏▗███▖     ▟█▏█▉  █▊▐█▌ ▐█▎ █▉ █▙▄▟▛▘██▀▀▔ ▜█  ▟█    ▜▉  ▐█▍▜█▁▟█ ██▀▀▔ ▐█▎      ▐█▌ ▐█▍ ▟█▕██▀▀▔     █▊ ▟▛▀▜█ ▗█▛▘  ▝▜▇█▘    ▐█▎ ▟█▏█▋  ▕█▍  ▟▉▝█▌▟█▙▗█▏▟█  █▉    █▉  █▉▐█▎  █▋ ███▛▘
  ▜█▊   █▉ ▟█▍▝█▆▅▅▛     ▔▀▐█▌▕█▙▅██▏▜█▎▝█▇▆▇▘▜█▎▕█▉    ▐█▋ ▝█▇▇█▛▐█▍ ▜█▏   ▆█▛ ▜█▅▟█▋▕█▋ ▟█▏ █▉ █▊▝▔  ▝█▆▅▆▌▝██▆██    ▝█▇▇█▛  ▜█▛▘ ▝█▆▅▆▌▐█▌      ▕█▙ ▐█▍ █▉ ▜█▅▅▆▘    █▉ ▜█▅██▏██▅▅▅▐▇█▛▘     ▐██▇█▘ █▉   ▜█▇▇█▘ ▜█▛▝██▘ ▜█▏▕█▊    ▝█▇▆█▊ ▜█▇██▘▜█▆▇▇▇
▟▇▇█▊     ▁▁▁                 ▁    ▁     ▁▁                             ▁▂   ▁▁  ▁▁             ▁ ▁▁▁  ▁       ▁▁   ▁  ▁▁          ▁                ▁   ▁               ▁      ▁▁        ▂
▀▜▛▔ ▇▏▟▎█▛▛▘    ▅█▙ ▟▏ ▜▍▇ ▅██▌█ ▟▌    ▟▛▀▘▅█▙ ▇▄ ▟▘    ▅▆▇■▗▍ ▐▋▐▙ ▗█▐█▀█▐█▛▛▏▆█▜█▏    ▟█▙ ▖ ▟▘█▛▛▘▆▛▜▌    ▟███▕▊ █▐██▀    ▐▌   ▗█▏■██▍▕▇▖▗▌    ▗▛▜▋▅█▜▊▗▇█▖▐▋   ▇▖▆▁ ▟▏   ▗▇███▏▟█▙ ▗▛▜▋
 ▟▌  █▄█▏█▆▌    ▕▊▔█▏█  ▐▍▉▐▛ ▝▘█▟▘     █▅▆▕▊ ▜▎▝██▋      ▕▋ ▟▎ ▐▋▟█▄██▐▙▆▛▐█▆▎ ▉  ▟▎   ▐▌ █▐█▃▛▕█▆▖▕█▄█▘     █▎ ▕▊▅▉▐▙▆▎    ▐▍  ▗▊█▍ ▟▀  ▝▜▛     ▜██▅▜▙▟▘▟▎▕▊▟▋▟▎ ▜▌██▅█     ▉  ▟▐▋ ▜▏▉ ▅▙
 ▜▋  █▔▜▎█▅▅▖   ▕███ ▜▆▅█▕▉▝▙▅▄▖██▅     ▉▔  █▆▛  ▟▜█▎    ▐▆▉ ▝▇▅▟▘▟▛▀▘▜▕▊▔ ▐█▅▅▐█▅▇▛    ▐█▆▛ ▜█ ▕█▅▅▖▜█▅      █▎ ▐▊▔▉▐█▅▅    ▐▙▅▅▟▛▜▌▟█▆▆▏ ▟▘     ▟▙▅▘▐█▙▖▝▇▇▘▝███▆█▐▋▝▜█    ▐█▅▇▛▕█▆▛▕█▆██
 ▔▔  ▔ ▝▘▔▔▔     ▔▔▔  ▔▔  ▔ ▔▔▔ ▔ ▔     ▔    ▔     ▔      ▔▔  ▔▔▔ ▔   ▔ ▔   ▔▔▔ ▔▔▔      ▔▔   ▔  ▔▔▔ ▔ ▔      ▔   ▔ ▀ ▔▔▔     ▔▔▔▔  ▔▔▔▔▔  ▔      ▔▔   ▔▔  ▔▔  ▔ ▔▔  ▔  ▀     ▔▔▔  ▔▔   ▔▔▟
▅▅▆▆▅▐█▏                       ▆▆     ▐█▎      ▗▇▊              ▗▆▖                             ▟▊                            ▂▂ ▟▊            ▟▊                    ▟▊                                  █▌
 ▟█▔ ▟█▃▃  ▂▃▃      ▃▃▂▃ ▃▖ ▃▃ ▃▔ ▂▃▃ ▟█ ▃▖    █▊ ▁▃▃▁ ▃▃ ▁▃    ▗▃ ▃▖ ▃▃ ▃▂▃▂▁▃▂ ▃▃▃▃  ▁▃▃▁  ▃▃▂█▌    ▃▄▃  ▃  ▃▃ ▁▃▃▁ ▗▃▃▖   ▃█▙▕█▙▃▖ ▁▃▃▂    ▕█▍ ▃▃▂▃ ▃▃▃▃ ▃  ▂▃   ▕█▙▃▃  ▃▃▃ ▃▃▃  ▃  ▃▃  ▃ ▃▃▃▂    ▁▃▃▐█▏▁▃▃▂  ▂▃▂▃
 █▊  █▛▔█▍▟█▄█▙   ▗█▛▀▜█▕█▍ █▌▐█▏▟▛▀▀▘███▀    ▐█▀▟█▀▀█▌ ▜▇▛▘    ▟▊▕█▍ █▌▐█▀▜█▀▜▉▕█▀▀▜▊▗████▎▟▛▀▜█▎   ▇▛▀▜█ ▜▙▗█▘▗██▟█▎▟█▀▘   ▜█▀▐█▀▜▉▗██▄█▖   ▐█▕█▛▀▜█ ▀██▘ ▜▙▗█▘   ▐█▀▀█▋▟█▀▘█▛▀▜█ ▜▋▟██▃█▘▐█▀▜█   ▗█▀▀█▉▗█▀▀█▋▟▛▀▜█
▕█▌ ▐█▎▐█▏▜█▄▟▖   ▐█▙▄█▋▕█▙▅█▏▟▊ ▜▙▄▄▕█▛█▄    ▗█ ▜█▅▅█▘▄█▀█▖    █▌▕█▅▟█ ▟▊ ▟▊ ▟▊▗█▅▅█▘▝█▄▄▅▔█▙▄▟▉    ▜▙▄▟▛ ▝█▛▘ ▝█▄▄▅ █▌     ▟▊ ▟▊ █▋▝█▙▄▅▔   ▟▉▐█▙▄█▊▗▟█▄▄ ▕█▛▘    ▟█▅▅█▘█▋ ▕█▙▄▟▛ ▐█▛ ▜█▘ ▟█ ▟▊   ▝█▅▄█▋▝█▅▅█▘▜█▄▟█
 ▔   ▔  ▔  ▔▔▔     ▔▔▔█▍  ▔▔  ▔▔  ▔▔▔ ▔ ▔▔    ▔▔  ▔▔▔  ▔  ▔▔   ▐█▏  ▔▔  ▔▔ ▔▔ ▔▔▟▊▔▔▔  ▔▔▔   ▔▔▔▔     ▔▔▔   ▔    ▔▔▔  ▔      ▔▔ ▔▔ ▔   ▔▔     ▔▔ ▔▔▔▔▔▔▔▔▔▔ ▟▛      ▔▔▔▔  ▔    ▔▔▔   ▔  ▔   ▔▔ ▔▔    ▔▔▔▔   ▔▔  ▅▅▃█▋
▆■■▇▆■▆▏■▇▍                             ▗▆▖       ▆▇▏        ▅▉▀▇▏                ▅▆                                     ■▇▌                                  ▗  ■▇▊               ▆▇                        ▆▇                                         ■▆▊
▀ ▐█▋ ▝▘▐█▍▃▄▂  ▁▃▄▃      ▂▄▄▂▃ ▁▂▖ ▁▃▖ ▁█▏  ▃▄▄▁ ▜█▏▗▄▃    ▄█▙▃  ▃▄▄▂ ▗▄▄▖▗▄▃    ▂▙ ▁▂▖ ▁▃▖ ▂▃▁▃▄▁▂▄▃  ▂▃▂▄▃   ▂▄▄▂  ▁▃▄▐█▌     ▂▄▄▃  ▄▃▃ ▃▄ ▁▃▄▃  ▁▃▂▄▃    ▄█▄▖ █▉▂▄▃   ▃▄▃▁     ██  ▂▄▄▃  ▗▄▃▄▃▗▃▄▖ ▃▄    ██▃▄▃  ▁▃▂▄▃  ▃▄▄▂ ▃▄▄ ▃▄▃ ▃▃▁▂▖▂▄▃      ▃▄▃█▉  ▂▄▄▃   ▂▄▄▃▂
  ▐█▋   ▐█▛▔▜█▏▗█▌▁█▙    ▟█▘ ██ ▐█▋ ▐█▌ ▜█▎▕█▛ ▝▀ ██▃▟▀     ▐█▋ ▗█▛ ▝█▋ ▝█▙▟▀    ▔██ ▐█▋ ▐█▌ ▜█▎▔█▉▔▐█▌ ▜█▎▔▜█▏▟█▁▁█▋▗█▛ ▐█▌    ▟█  ██ ▝█▌ ▟▘▗█▌▁██ ▜█▛▔▀    ▜█▎  █▉▔▐█▋ ▟▉▁▐█▖    ██ ▝▀▘▂█▊ ▝▔▟█▀ ▝█▌ ▟▘    ██▔ ██ ▜█▛▔▀▗█▛ ▝█▋ █▉ ▐█▋ ▟ ▕█▛▔▐█▌    ▟█ ▔█▉ ▟█▘ ▜█▏▗█▎ ▜█
  ▐█▋   ▐█▍ ▟█▏▜█▍▔▔▔    ██  ██ ▕█▋ ▐█▌ ▜█▎▐█▋    ██▀█▙     ▕█▋ ▐█▌ ▕█▉  ▟▜█▖     ▜█ ▕█▋ ▐█▌ ▜█▎ █▉ ▐█▌ ▜█▎ ▟█▎██▔▔▔▔▐█▌ ▐█▌    █▉  ▟█▎ ▜█▅▘ ▝█▌▔▔▔ ▐█▌      ▜█▏  █▊ ▐█▋▕█▉▔▔▔▔    ██ ▗▇▛▔█▉  ▟█▘   ▜█▅▌     ██  ██▏▐█▌  ▐█▌ ▕█▉ ▝█▙▛▝█▄▍ ▕█▋ ▐█▌    █▉  █▉ ██  ▟█▍▝█▙▄█▀
  ■█■▖  ▟█▙ ■█▎ ▀■▅■▀    ▝▜▙■██  ▜█■■█■ ▟█■ ▀▜▅■▀▗██▏▝█■▖   ■█▊▖ ▀■▄▟▀▘▗■▌▐██■    ▟█▏ ▜█■■█■ ■█▎▗█▉▖■█▙ ▟█■▅█▀ ▝▀■▅■▘ ▀█▅■█■    ▝▜■▅▛▘   ▜▛   ▀■▅■▀ ▟█▙      ▝█▅■▗█▉ ▐█■ ▝▜■▅■▘   ▗██▖▝█■■▜▉▖■█▙■■▏  ▜▛      ▜▀■▅▛▘ ▟█▙   ▀■▄▟▀▘  ▜█▏ ▜▛  ■█▊ ■█▙    ▝▜▅■█▉ ▝▀▙▅■▀ ▝█▆▆▅▅
                             ██                                                ▗▄ █▛                    ▟█▎                                                                                       ▗▖ ▛                                                             ▇▌  ▔█
▆▆▆▆▆▆▆▆▐▇▋                                 ▐█▌       ▕▇▋         ▗▇█▀▍                    ▜█▎                                         ▟█▏                                     ▗▅▏  ▇▇▏                ▕▇▊                           ▇▇▏                                              ▇▊
  ▔█▉   ▐█▙▅▆▆▅  ▄▆■▇▆▖     ▗▅▆▆▅▆▌ ▗▆▏ ▕▆▆ ▗▆▖ ▗▅▆▇▆▖▐█▋ ▗▆■    ▗▆█▆▆ ▄▆▇▇▆▄ ■▆▅ ▗▆▆     ▆▆▆▏ ▆▆  ▗▆▍ ▆▆▄▆▆▅▃▅▇▆▖ ▗▆▄▆▇▆▖  ▅▆■▆▅  ▄▆▆▆▟█▏     ▄▆▇▇▆▄ ▆▆   ▆▆ ▗▆■▇▆▄ ▗▆▃▅▆▏   ▗██▆▆▏██▄▆▆▆▖ ▁▅▆■▆▅     ▕█▊ ▗▆■▇▆▅ ▕▆▆▆▆▆▋■▆▖  ▆▆     ▜█▄▆▇▆▖ ▆▆▃▅▆ ▄▆▇▇▆▖ ▆▆  ▆▆▖ ▗▆▏▗▆▄▅▆▆▖      ▅▆▇▅█▊  ▅▆▇▆▅▖ ▗▅▆▆▅▆
   █▉   ▐█▋  ▜█▏▗█▙▄▄██▏    ██▔ ▐█▋ ▜█▏ ▕█▊ ▐█▍▕█▉   ▔▐█▊▟█▀      ▜█▍ ▐█▌  ▜█▍ ▝█▇█▛       ▜█▎ █▊  ▝█▍ ██▔ ▜█▌ ▕█▊ ▝█▍  ▜█▏▟█▄▄▟█▋▗█▛  ▜█▏    ▗█▌  ▜█▍▝█▙ ▟█▘▐█▙▄▄▟█▎▝█▛▔▔     ██▏  ██▏ ▐█▋ ██▄▄▟█▋    ▕█▊  ▃▄▄▟█▍   ▅█▀  ██ ▗█▘     ▜█▏ ▝█▉ ██▀▔▔▗█▍  ▜█▎▝█▌▗█▜▙▕█▛ ▜█▍ ▔█▉     ▟█▘  █▊ ▟█▘  ██ ██▔ ▐█
   █▉   ▐█▋  ▜█▏▝█▋▔▔▔▔     ██▁ ▐█▋ ▜█▎ ▕█▊ ▐█▍▕█▉   ▁▐██▀█▙▖     ▟█▍ ▐█▙  ▟█▍ ▗▇▛█▙       ▜█▏ █▉  ▟█▍ ██▏ ▐█▍  █▊ ▐█▍ ▁██▏▜█▔▔▔▔▔▐█▌  ▟█▏    ▝█▌  ▟█▍ ▝█▙█▘ ▐█▙▔▔▔▔ ▝█▍       ▜█▏  ██▏ ▐█▊ ██▔▔▔▔▔    ▕█▊ ▟█▔▔▐█▍ ▄█▛▘   ▝█▙█▛      ██▏ ▗█▉ ▜█▏  ▝█▌  ▟█▎ ▜██▍ ███▏ ▜█▍  █▉     ▜█▖ ▁█▉ ▜█▖  ██ ██▂▁▂█
   █▉   ▐▜▋  ▝█▏ ▝▀█▇■▀     ▝▀█▀▜█▋ ▝▜█▛▀█▊ ▝█▍ ▝▀▜■▛▘▕▜▋ ▝▜■     ▝█▍  ▝▀▇▇▀▀ ■■▀ ▝▜■      ▟█▏ ▝▜█▀▀■▍ ▜▜  ▝█▘  ■▊ ▟█▛■█▀▘  ▀▜■■▀▘ ▀▜█▀▜█▏     ▀▜▇▇▀▘   ▝■▀   ▝▀▜▇■▀▏▝■▍       ▝▜█■▏▝▛▏ ▕▜▋ ▔▀▜▇■▀▘    ▕■▊ ▝▜■■▀█▍▐█▜■■▀▋  ▜█▛       ▀▛▀■█▀  ▝▜▏   ▀▜▇■▀▘  ▕▀▀  ▝■▘  ▝█▎  ▀▉     ▝▀█▛▀▛▊  ▀▜■█▀▘ ▝▀▀▀▜█
                                ▐█▌                                                      ▆▆█▛                      ▝█▍                                                                                                     ██                                                                    ▐▇▆▆█▛
▁▃▄▅▇▖▟▊                     ▗▅▖    ▗▆▎       ▗▇▛              ▄                           ▄▄                           ▁▁ ▐█▎          ▕▇▌                    ▗▄                                 ▄▄
▀▜█▍▔ ▟▊▁▃▁  ▃▃     ▃▃▃▃ ▄ ▂▁▐█▏ ▃▃ ▐█▎ ▃▂    ▟▊▃▁▁▃▃ ▁▃ ▗▄   ▕█▎▃▃ ▂ ▃▃▗▄▂ ▁  ▄▂▃▁  ▃▃▁ ▁▃▟▉    ▁▃▃ ▃▂  ▃▃ ▂▃▂ ▂ ▃▄    ██▄▐█▏▂▃  ▂▃▃   ▕█▍ ▁▃▃  ▁▁▂▃▃▃▃  ▃▂   ▜▉▃▂ ▁▂ ▄▄ ▃▃▁▂▃  ▃  ▗▄▂▃  ▂     ▁▃▟▉ ▁▃▃  ▁▃▂▃
 ▐█▍  ███▜█ ▟▛▜▊   ▟█▀█▋▐█ █▊▐█▏▟▛█▋▐█▄▆█▀   ■█▛▀▚█▀▜▙▝█▙█▘   ▐█▎█▌▐█▎█▋██▋▟█▋▐█▛▜█▏▟▛▜▉▗█▀█▉   ▕█▀▜█▝█▖▗█▛▟▛▜█▐██▛▔   ■█▊▀▐██▛█▌▗█▀█▎  ▕█▌▟▛▀█▌ ▀▀▜█▛▜█▎▗█▛   ▜█▛█▌▜██▀▔▟▛▀█▟█▏▐█▋ ▟▛▐█▎▇█▊   ▗█▀█▉▕█▀▜█▕█▀██
 ▐█▍  ██▘▐█▐█▅▛▘   █▉ █▋▜█ █▋▐█▕█▍▐▂▐██▉▁     ▟▊ ▐█ ▐█▏▟█▌    ▐█▎█▌▐█▏██▘██▛█▊▐█▌▐█▎█▙▛▘▜█ █▊   ▐█▏▐█▏▜▙▟▛ █▙█▘▐█▘      █▊ ▐█▛ ▟▊█▙▇▀   ▕█▍▄▇▀█▊  ▟█▘  ▜▙█▛    ▜█ █▊▜█▘ ▕█▍▕█▍█▙█▛█▅█▏▐██▘█▉   ▜█ █▉▐█▏▐█▐█ ▟█
 ▐█▍  ▜▊ ▐█ ▜▙▅▉   ▝█▆█▋▝█▅█▘▐█▏▜▙▟▀▐█▛▀█▆    ▜▊ ▝█▄▟▛▗█▀█▖   ▐█▏▜▙▟▛ ▝▛ ██▏█▊▕██▟▛ ▜▙▅█▝█▅█▉   ▕█▙▟▛ ▐██▏ ▜▙▄█▐█▎      █▊ ▐█▎ █▌▝█▄█▍  ▕█▌█▙▟▜█▗▇█▅▆▇  █▉     ▟█▙█▌▟█   ▜▙▟▛ ▐█▉ ▜█▌ ▐█▋ █▉   ▝█▅█▉▕█▙▟▛▕█▆▛█
  ▔▔  ▔▔     ▔▔      ▔▟▋ ▔▔▔  ▔  ▔▔  ▔   ▔    ▔▔   ▔▔  ▔ ▔▔  ▗▟▛  ▔▔     ▔▘ ▔▔▐█▎▔   ▔▔   ▔▔▔     ▔▔   ▔▔   ▔▔  ▔       ▔▔  ▔     ▔▔     ▔▔▔▔  ▔ ▔▔▔   ▕█▍     ▔▔▔▔  ▔    ▔▔   ▔▔ ▔▔   ▀▘        ▔▔▔  ▔▔  ▗▅▄█
▗■ ▅▅ ▅▖▐▛                       ▃      ▜▌       ▄▀▜            ▗▖                             ▇                             ▗▌           ▇▘                  ▕▇                                 ▇
   ▉    ▟ ▁    ▁      ▁ ▁        ▔   ▁  ▛ ▁     ▟▘  ▁           ▝▔       ▁  ▁ ▁   ▗     ▁    ▁▟▘     ▁   ▁     ▁  ▁ ▁    ▗▖  ▛ ▁         ▗▛    ▂  ▁   ▁       ▟▘▁  ▁ ▁   ▁             ▁       ▁▟▘  ▁    ▁
  ▟▘   ▗▙▝▜▋ ▅▚▛    ▄▀▔█▘■█▘▗▛ ▝█▘ ▞▔▀ ▟▎▁▛▏   ▜▛▔▗▛▔▉ ▀▙▖▀▘   ▝▉ ▝█ ▗▛ ▀█▖▚▊▞▜▋ ▐▛▔▇ ▗▞▃▋▗▅▀▐▛    ▄▀▔▊ ▐▛ ▟▎▗▀▟▘▝▉▞▀   ▕▛▘ ▟▍▀▜▍▗▞▚▛    ▟  ▖▀▜▘▕▛▀▛▘▝▜▎▟▏   ▗▋▀▐▋▝▜▎▀▏▆▘▜▍ ▜▘ ▞ ▐▌▝█▍▀█▏   ▗▅▀▐▛ ▄▀▔▊ ▆▘▟▎
 ▗▛   ▗▛▔ ▟ ▟▘▔    ▟▘▂▟▌ ▗▘▖█▚ ▟▘ ▟▏▁ ▗▀▜▔     ▟▏▐▛ ▟▘▁▗▜▖▁    ▟▘ ▟▘▖█▂ ▟▘ ▟▘▗▛▁ ▟ ▗▛▕▊▔▁▕▛ ▄█▂   ▐▛ ▟▘ ▟▘▗▘▗▛▔  ▟▘     ▟▚▖▗▛ ▗▛ ▉▔▁    ▟▘▃▟▁▅▛▁▁▅▘ ▖ ▟▎▘   ▕▉ ▗▛ ▗▛  ▟▘▗▛ ▗▛▗▟▎▁▞ ▗▛ ▟▘   ▕▛ ▄█▂▐▛ ▟▘▝▙■▀
▝▀▀   ▝ ▝▀  ▝▀▔    ▝▀▔▉  ▀▔ ▀▔ ▀▔ ▝▀▔ ▀  ▀    ▗▌  ▀▀▔ ▀▔ ▀▘   ▗▛  ▀▔▝▀▔ ▘ ▝▘ ▝▀ ▟▀▀▔  ▀▀▔ ▀▘ ▀▔    ▀▀▔  ▀▀   ▀▀  ▘      ▀▔ ▀ ▝▘  ▀▀▔    ▀▘ ▀▔▝▀ ▘▔▀▀▔ ▟▘     ▀▀▔  ▝   ▝▀▘  ▝▀▔▝▀▔  ▀  ▀▔    ▀▘ ▀▔ ▀▀▔ ▃■▅▂
                    ▁▟▌                     ▅▄▞             ▗▃▞                ▃▊                                                                  ▗▄▝▔                                              ▝▙▂▃▛
▅■■▅■■▅ ▇▌                         ▕▅       ▐▇        ▅▀▀▎              ▅▖                                ▐▇                                 ▐▇             ▜▋                      ▇▌                                    ▇▌
▘ ▐▉  ▘ █▁▃▂   ▂▃▂     ▁▃▂▂ ▂▂  ▂  ▂   ▁▂▃  ▟▌ ▂▂   ▁▟▌▁ ▂▃▂  ▁▃  ▂▂   ▂▃ ▂▃  ▁  ▂▂▁▃▂ ▂▂  ▂▂▁▂   ▁▂▃   ▂▃▟▌     ▂▃▁ ▂▃  ▁▖  ▂▃  ▂▁▁▂   ▁▟▍▁ ▟▌▂▃   ▁▃▂     ▟▎  ▂▃▂  ▂▂▂▂▂▂▁  ▁     █▁▃  ▂▃▁▂  ▂▃▂ ▁▃  ▁   ▂ ▃▁▂▃      ▁▂▃█▏ ▁▂▂    ▂▃▂
  ▟▌   ▐▛▔▔█▏▗▇▔▃▛    ▟▘ ▟▌ ▟▋  █  ▟▍ ▟▘ ▝▎ █▂▞▘    ▕█▔▔▟▘ ▝▙  ▜▙▞▀    ▐▊ ▐▊  █▎ ▟▛▔▐▉▔▔█▎ ▟▛▔▐▉ ▟▀▂▟▘▗▛▔ █▏   ▗▛▔ ▜▌ ▜▎ ▟▘▗▛▁▟▋ ▟▍▔▘   ▐▊▔▔ █▔▔▜▌ ▟▘▂▟▘   ▐▉ ▗▛▔▕█ ▕▘▁▟▀ ▔█▏▕▛    ▐▛▔▝█ ▕█▔▀ ▇▘ ▐▙ ▜▋ ▟▌ ▟▘ █▀▔▜▋    ▟▀ ▐▋ ▟▛▔▔█▎▗▛▔ █
 ▕█▏   ▟▎ ▐▊ ▜▋▔▔    ▐▊ ▁█  █▎ ▟▋ ▕█ ▐▊    ▐▋▝▙     ▐▋ ▐▉  ▗▛ ▁▄▜▌     ▟▍ ▟▍ ▐▉  █▏ ▟▌ ▕▉  █▏ ▟▛▐█▔▔  █▍ ▐▋    ▜▌  ▟▘ ▝▙▞▘ █▎▔  ▕█      ▟▍  ▗▋  █▏▐█▔▔     ▟▍ █▍ ▟▋  ▄▛▘▗  ▝▋▞     ▟▎ ▗▛ ▟▌  ▐▊  ▐▛ ▕▉▟▀▊▟▘ ▐▉  ▟▎   ▐▉  ▟▎ █▎  █▏▟▌ ▐▊
 ▀▀▘   ▀  ▝▀  ▀▀▀▔    ▀▀▜▋  ▝▀▀▝▀ ▝▀  ▀▀▀▔ ▝▘ ▝▀    ▟▏  ▝▀■▀ ▝▀  ▀▘   ▗▛  ▝▀▀▝▀ ▕▀  ▀▘ ▝▀ ▐▛■▀▔  ▝▀▀▘ ▝▀▀▝▀     ▀■▀▘  ▝▀   ▝▀▀▀ ▝▘      ▀▀▘ ▀▘  ▀▘ ▀▀▀▘    ▀▘ ▝▀▀▝▀ ▀▀▀▀▀  ▕▛      ▀■▀▘  ▀    ▝▀▀▀   ▀  ▀   ▝▘  ▀▘    ▀▀▔▀▘ ▔▀■▀▔ ▝▀▀█▘
                       ▃█▍                       ▕▅▄▘                ▗▀                  ▗▟▙                                                                              ▄▛                                                     ▕▅▄■▘
▟▇▇█▊     ▁▁▁                 ▁    ▁     ▁▁                             ▁▂   ▁▁  ▁▁             ▁ ▁▁▁  ▁       ▁▁   ▁  ▁▁          ▁                ▁   ▁                      ▁▁        ▂
▀▜▛▔▔▇▏▟▎██▀▘    ▅█▙ ▟▏ ▜▍▇ ▅██▌█ ▟▌    ▟▛▀▘▅█▙ ▇▄ ▟▘    ▅▆▇■▗▍ ▐▋▐▙ ▗█▐█▀█▐██▀▕▆█▜█▏    ▟█▙ ▖ ▟▘██▀▘▆█▜▌    ▟███ ▊ █▐█■▀    ▐▌   ▗█▏■██▍▕▇▖▗▍    ▗▛▜▌▅█▜▊▗▇█▖▐▋   ▇▖▆▁ ▟▏   ▗▇█▜█▏▟█▙ ▗▛▜▋
 ▟▋  █▄█▏█▆▖    ▕▊▔█▏█  ▐▍▉▐▛ ▝▘█▟▘     █▅▆▕▊ ▜▎▝██▋      ▕▋ ▟▎ ▐▋▟█▄██▐▙▆▛▐█▆▎ ▉  ▟▎   ▐▌ █▐█▃▛▕█▆▖▕█▄█▘     █▎ ▕▙▅▉▐▙▆▎    ▐▍  ▗▉█▍ ▟▀  ▝▜▛     ▜██▅▜▙▟▘▟▎ ▊▟▋▟▎ ▜▌██▅█     ▉  ▟▐▋ ▜▏▉ ▅▙
 ▟▋  █▔▜▎█▅▅▖   ▕███ ▜▆▅█▕█▝▙▄▄▖██▅     ▉▔  █▆▛  ▟▜█▎    ▐▅▉ ▝▇▅▟▘▟▛▀▘▜▕▊▔ ▐█▅▅▐█▅▆▛    ▐█▆▛ ▜█ ▕█▅▅▖▟█▅      █▎ ▐▊▔▉▐█▅▅    ▐▙▆▅▟▛▜▌▟▙▆▆▏ ▟▘     ▟▙▅▘▐█▙▖▝▇▇▘▝███▆█▐▋▝▜█    ▐█▅▇▛▕█▆▛▕█▆██
 ▔▔  ▔ ▝▘▔▔▔     ▔▔▔  ▔▔▔ ▔ ▔▔▔ ▔ ▔     ▔    ▔     ▔      ▔▔  ▔▔▔ ▔   ▔ ▔   ▔▔▔ ▔▔▔      ▔▔   ▔  ▔▔▔ ▔▔▔      ▔   ▔ ▀ ▔▔▔     ▔▔▔▔  ▔▔▔▔▔  ▔      ▔▔   ▔▔  ▔▔  ▔ ▔▔  ▔  ▀     ▔▔▔  ▔▔   ▔▔▟
```

```console
Usage: title_creator

Create a title

optional arguments:
display:
    --text=Hello World!
        text to render (default Hello World!)
    --characters= !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}
        text to render, ignored when loading a map (default  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|})
    --resolution=16
        text to render, ignored when loading a map (default 16)
    --aspect=0.5
        character height to width (default 0.5)
    --font=/System/Library/Fonts/Monaco.ttf
        filename of the ttf/otf font (default /System/Library/Fonts/Helvetica.ttc)
    --size=12
        font size in points (default 25)
    --max-width=None
        maximium width to render
    --mode=20
        render mode (default 20)
    --allow-inverted=false
        use inverted characters, ignored when writing to file
input/output:
    --load=None
        load saved character map
    --save=None
        save character map
    --output=None
        save output to file
```

## Why?
I like starting my programs with a nice title. 
**So why another?** I wanted to see if go could improve the performance over python.
**Was it fast?** Yes.

## Improvements?
Automated way to find fonts cross-platform

## State?
- Some fonts it does not render some characters, for example, the letter 'n' might not appear when using myfont.ttf
- Unicode character rarely render, this matters for title and for characters used in the character map, but will write to screen
  - Workaround: Use the pyton script to create json maps and use those in the go program.  It will allow you to use those character to render but still not be in the title, for that stick the python version
- Does not work with all fonts.

## New
### 1.0
Create a title and save/load the character map created
