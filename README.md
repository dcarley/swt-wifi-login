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

The portal doesn't always work first time so the utility will keep retrying:
```
2017/10/06 08:41:49 login failed: Get http://swt.passengerwifi.com/cws/?password=REDACTED&rq=login&username=REDACTED: net/http: timeout awaiting response headers
2017/10/06 08:41:49 sleeping for: 5s
2017/10/06 08:42:02 login successful!
```

If you're using Mac OS X, you can disable the graphical prompt with:
```
sudo defaults write \
  /Library/Preferences/SystemConfiguration/com.apple.captive.control \
  Active -boolean false
```
