# GOtusernames
Generate possible AD usernames from names like John Doe > J.Doe or JDoe

## Installation:
```
~/programming/golang/src/gotyourusername ❯ git clone https://github.com/T4TCH3R/GOtusernames.git
~/programming/golang/src/gotyourusername ❯ cd GOtusernames
~/programming/golang/src/gotyourusername/GOtusernames ❯ go build -o GOtusernames GOtusernames.go
```

## Example user file:
```
~/programming/golang/src/gotyourusername ❯ cat testfile/users.txt
Ham Brook
James Phelps
Keely Lyons
Dax Santiago
Sierra Frye
Kyla Stewart
Kaiara Spencer
Dave Simpson
Ben Thompson
Chris Stewart
```

## Example:

### Help menu
```
~/programming/golang/src/gotyourusername ❯ ./GOtusernames -h
Usage of ./GOtusernames:
  -f string
        Please select the file with first and lastnames like: John Doe
  -o string
        Output the new domain usernames to a file.
```
        
### output generated usernames to your screen
```
~/programming/golang/src/gotyourusername ❯ ./GOtusernames -f testfile/users.txt 
H.Brook 
HamBrook
HBrook    
Ham.Brook    
Ham.B   
Ham            
Brook    
J.Phelps 
JamesPhelps
JPhelps   
James.Phelps 
James.P 
James        
Phelps 
......
```

### Output the generated users to a file
```
~/programming/golang/src/gotyourusername ❯ ./GOtusernames -f testfile/users.txt -o output.txt
~/programming/golang/src/gotyourusername ❯ cat output.txt                                     
H.Brook                                                                                       
HamBrook                                       
HBrook                                         
Ham.Brook                                                                                     
Ham.B                                                                                         
Ham                                                                                           
Brook                                                                                         
J.Phelps                                       
JamesPhelps                                    
JPhelps                                        
James.Phelps                                   
James.P                                        
James                                          
Phelps
......
```



