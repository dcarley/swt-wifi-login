# swt-wifi-login

A command-line utility to log in to the [WiFi on South West Trains
services][swt] without using the graphical interface.

[swt]: https://www.southwesttrains.co.uk/news/freewifi/

## Usage

Fetch, build, and install:
```
go install github.com/dcarley/swt-wifi-login
```

Login:
```
swt-wifi-login <username> <password>
```

If you're using Mac OS X, you can disable the graphical prompt with:
```
sudo defaults write \
  /Library/Preferences/SystemConfiguration/com.apple.captive.control \
  Active -boolean false
```
