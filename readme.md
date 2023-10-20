# msi meestic go profile toggler

toggle between profiles for meestic - msi delta 15

https://github.com/Koromix/meestic

/etc/meestic.ini - location of profiles


Example of file with 4 profiles - white dim, mid, max, off...


```
# The profile set with the "Default" will run when the GUI starts
# Read the commented out section at the end for possible profile values

Default = Disabled

[White Dim]
Mode = Static
Intensity = 1
Colors = White

[White Mid]
Mode = Static
Intensity = 5
Colors = White

[White Max]
Mode = Static
Intensity = 10      
Colors = White

[Disabled]
Mode = Disabled

```

Build and run


```
go build .
./meestic
```