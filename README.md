# swt-wifi-login

A command-line utility to log in to the [WiFi on South Western Railway][swr] (née South West Trains) without using the graphical interface.

[swr]: https://www.southwesternrailway.com/travelling-with-us/onboard/wifi

## Usage

Fetch, build, and install:
```
go get -u github.com/dcarley/swt-wifi-login
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
