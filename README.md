# title_creator

## Print in large font in your terminal/console

This was an attempt to rewrite the python version of title_creator in Go. The idea was to improve the speed. And yes, it
is much faster but with caveats, see [State](#state).

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

### title_creator.go

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

### Usage

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
        maximum width to render
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

### Example

```
title_creator -font /Library/Fonts/CooperBlackStd.otf -resolution 64 -allow-inverted=false -text "Numbers" -characters="1234567890#%$^*.,-+<> "

  ,,,,,        ,,,,,
<%%%%%%0,     %%%%%%%>                                           ,,#%%%%>
  %%%%%%%%%     %%%0                                             %%%%%%%
   %%%%%%%%%%,  %%%>  ,,+4##%  ,,+##%%    ,,+#%  ,4#%%,  ,4#%%,    %%%%% ,+#%%,      ,+#%%%+,    ,,+##% ,4%%,   ,4%%%4#%,
   %%%%%%%%%%%%#%%%   70%%%%%  %%%%%%%  #%%%%%%0%%%%%%%%%%%%%%%%   %%%%%%%%%%%%%,  ,%%%0770%%%, %%%%%%%%%%%%%  0%%%%0%%%%,
   %%%%7%%%%%%%%%%%     %%%%%   4%%%%%   %%%%%%7 3%%%%%#  %%%%%%   %%%%%7 7%%%%%% 2%%%%%,,#%%%%  3%%%%%0%%%%#  %%%%%%, 77
   %%%1   %%%%%%%%%     %%%%%   4%%%%%   4%%%%%   %%%%%7  %%%%%%   %%%%%   3%%%%% %%%%%0%%%%%%#  4%%%%%        7%%%%%%%%,
   %%%8     %%%%%%%     %%%%%,  8%%%%%   4%%%%%   %%%%%1  %%%%%%   %%%%%,  ,%%%%# %%%%%0,   ,,,  4%%%%%       ,%,7%%%%%%%,
 <%%%%%%      7%%%%     %%%%%%%%%%%%%%% ,%%%%%%, #%%%%%% ,%%%%%%, ,%%%%%%,,%%%%7   %%%%%%%%%%%7 <%%%%%%%%     3%%%,#%%%%%
 7#%%%#7        7%#      7#%%#7  %%##7  7#%%%##7 9#%%%#7 79#%%##      7#%%%#7        79%%%#77   7#%%%%##7      7###%%%#7
```

## Why?

I like starting my programs with a nice title.
**So why another?** I wanted to see if go could improve the performance over python.
**Was it fast?** Yes. Much

## How?

1. It takes the ***<span style="color: yellow;">text</span>*** and renders it to an image using the
   ***<span style="color: yellow;">font</span>***  and the ***<span style="color: yellow;">size</span>***
2. It scales the image to no bigger than **<span style="color: yellow;">max-width</span>**.
3. It scales the image to the inverse of ***<span style="color: yellow;">aspect</span>*** so that when drawn to screen
   it matches the original image.
4. It crops the image into chucks of ***<span style="color: yellow;">resolution</span>*** x
   ***<span style="color: yellow;">resolution</span>*** and puts it into a 4D array.
5. It then takes the ***<span style="color: yellow;">characters</span>*** and maps each one to a 2D array of brightness
   values, and optionally ***<span style="color: yellow;">save</span>s*** or
   ***<span style="color: yellow;">load</span>s*** this
6. It then goes through 2 dimensions of the 4D array and compares ,using the
   ***<span style="color: yellow;">mode</span>***, to see which is the best match, and optionally
   ***<span style="color: yellow;">allow-inverted</span>*** versions of the character
7. It draws each matching character to screen or ***<span style="color: yellow;">output</span>s*** to file

## What are the modes?
- Mode 0 (default) : Mean Squared Error
- Mode 1 : Sum of absolute differences in matrix
- Mode 2-4 : Absolute differences in matrix with neighbouring values
- Mode 10-23: Contrasted values

## Improvements?

Add more font types

## State

- Some fonts it does not render some characters, for example, the letter 'n' might not appear when using myfont.ttf??
- Unicode characters beyond 1byte (255 or \u00FF) do not render, this matters for title and for characters used in the
  character map, but will write to screen
    - Workaround: Use the python script to create json maps and use those in the go program. It will allow you to use
      those character to render but still not be in the title, for that stick the python version
- Does not work with all fonts.

## Notes/Troubleshooting:
### “title_creator” cannot be opened because the developer cannot be verified.
xattr -d com.apple.quarantine /path/to/title_creator
### zsh: permission denied: /path/to/title_creator
chmod a+x /path/to/title_creator
### Block characters to not line up, or words are "wavey"
Windows only renders block characters aligned in the console and no where else \
Make sure you are using a fixed width font
### Empty squares appear
Choose a font that supports the character you choose the render with

## New

### 1.0

Create a title and save/load the character map created

### 1.1
- Fixed bad ascii code in default character set
- Support multiline
- Better cropping
- Inverted mode
- New render modes
- Code cleanup

### 1.1.1
- Printing first character only in json key, this allows for duplicate keys, example ascii-art.json