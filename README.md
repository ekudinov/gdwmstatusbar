Application written on golang for dwm status bar.
This app use goroutines ,one main for output and others are data providers.
Main goroutine collect information from providers and show data on status bar.
Providers collect data from out services and keep it.
You can add self data provider and use it(see interface Provider in type.go).
Examples  are in provaders dirs.
Configurations are topbar.go and downbar.go and is very simple.

Attention:for weather provider you must have api key 
http://openweathermap.org/appid#get and 
set it "export OWM_API_KEY=you key there".
 
You must have patch http://dwm.suckless.org/patches/dualstatus for bottom and top 
status panel.

View must be like this screenshort.
![Alt text](screen.png?raw=true "Screen")

