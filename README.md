ASCIImg
-------
Command Line ASCII Art Generator from image [png, jpeg, gif]

Usage:
--
- Download `git clone https://github.com/mineroot/asciimg.git`
- Build `go build`, Go 1.24 required
- Run `./asciimg --reversed --width=80 < /path/to/img.png`

All options:
--
`$ ./asciimg -help`
```
Usage of ./asciimg:
  -charset string
    	Charset to use for ASCII (default "@%#*+=-:. ")
  -debug
    	Print debug info
  -ratio float
    	ASCII width to height ratio (default 0.5)
  -reversed
    	Reverse ASCII
  -width int
    	ASCII width [10-240] (default 120)
```

```
                                            .:-=+++++=-:.                                           
                                        :+#@@@@@@@@@@@@@@@#+:                                       
                                     :*@@@@@@@@@@@@@@@@@@@@@@@*:                                    
                                   :#@@@%#%@@@@@@@@@@@@@@@%#%@@@#:                                  
                                  +@@@@@:   :=+=-----=++:   .@@@@@+                                 
                                 *@@@@@@:                   .@@@@@@#                                
                                =@@@@@@*                     +@@@@@@+                               
                                @@@@@@#                       *@@@@@@                               
                               .@@@@@@*                       =@@@@@@.                              
                                @@@@@@%                       #@@@@@@                               
                                *@@@@@@*                     +@@@@@@*                               
                                 %@@@%%@%+:               :=%@@@@@@%.                               
                                 .%@@+..*@@@%#-       :#%@@@@@@@@@%.                                
                                   +@@@= .=++=         #@@@@@@@@@+                                  
                                     +%@#=-:--         #@@@@@@@+.                                   
                                       :+%@@@%         #@@@%*-                                      
                                           :=-         :=:                                          
                       :=++++=-    .           :::     :::             :::                          
                     +@@@%###%%  =@@%   ===    @@@:   .@@@.            @@@-                         
                    *@@%:         =+-  :@@@.   @@@:   .@@@.            @@@-.::.                     
                    @@@-  -====- =@@% @@@@@@@= @@@#***#@@@.:@@@   #@@= @@@@@@@@@=                   
                   .@@@:  *%%@@% =@@%  :@@@.   @@@*+++*@@@.:@@@   #@@= @@@-  -@@@                   
                    %@@#    =@@% =@@%  .@@@    @@@:   .@@@.:@@@   #@@= @@@-  .@@@.                  
                     #@@@#**%@@% =@@%  .@@@+=: @@@:   .@@@..@@@+=+%@@= @@@#==#@@#                   
                      .=+**#*+=: :+++   .+***: +++.    +++  :+***+-++: ++==****=                    

```