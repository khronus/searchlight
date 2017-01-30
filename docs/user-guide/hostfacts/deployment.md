# Deployment Guide

This guide will walk you through how to deploy hostfacts service in kubernetes node.

### Deploy Hostfacts

Write `hostfacts.service` file in __systemd directory__ in your kubernetes node.

##### systemd directory
###### Ubuntu
```sh
/lib/systemd/system
```
###### RedHat
```sh
/usr/lib/systemd/system
```


##### `hostfacts.service`
```ini
[Unit]
Description=Provide host facts

[Service]
ExecStart=/usr/bin/hostfacts
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

If you want to set authentication in `hostfacts`, set one of the following
###### Basic Auth
```sh
# Use ENV
# Add Environment in hostfacts.service under [Service] section
Environment=HOSTFACTS_AUTH_USERNAME="<token>"
Environment=HOSTFACTS_AUTH_PASSWORD="<password>"

# You can also pass flags for basic auth.
# Modify ExecStart in [Service] section
ExecStart=/usr/bin/hostfacts --username="<username>" --password="<password>"
```
###### Token
```sh
# Use ENV
# Add Environment in hostfacts.service under [Service] section
Environment=HOSTFACTS_AUTH_TOKEN="<token>"

# You can also pass flag for token.
# Modify ExecStart in [Service] section
ExecStart=/usr/bin/hostfacts --token="<toekn>"
```

If you want to set SSL certificate, do following

1. Generate certificates and key. See process [here](../icinga2/certificate.md).
2. Use flags to pass file directory

    ```sh
    # Modify ExecStart in [Service] section
    ExecStart=/usr/bin/hostfacts --caCertFile="<path to ca cert file>" --certFile="<path to server cert file>" --keyFile="<path to server key file>"
    ```

You can ignore SSL when Kubernetes is running in private network like GCE, AWS.

> __Note:__ Modify `ExecStart` in `hostfacts.service`


### Add `hostfacts` binary

Download `hostfacts` and add binary in `/usr/bin`
```sh
curl -G  https://storage.googleapis.com/appscode-dev/binaries/hostfacts/0.3.0/hostfacts-linux-amd64 -o /usr/bin/hostfacts
```

##### Start Service
```sh
# Configure to be automatically started at boot time
systemctl enable hostfacts

# Start service
systemctl start hostfacts
```
