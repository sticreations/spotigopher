#Spotigopher
===
##Overview

Spotigopher is a Spotify API written based on the dbus.
It is right now tested on Linux.

##Usage

```
import "github.com/sticreations/spotigopher/spotigopher"
client := spotigopher.NewClient()
// Get Information
information, err := client.GetInfo()
// PlayPause
client.PlayPause()
```
