# Certificate

When using client certificate authentication, we can generate certificates manually using **`easyrsa`** or **`openssl`**.

### openssl

`openssl` can also be use to manually generate certificates for your cluster.

1. Set `MASTER_IP` ENV to Kubernetes Master Node IP

2. Generate a ca.key with 2048bit:
```
openssl genrsa -out ca.key 2048
```

3. According to the ca.key generate a ca.crt
```
openssl req -x509 -new -nodes -key ca.key -subj "/CN=${MASTER_IP}" -days 10000 -out ca.crt
# Set ICINGA_CA_CERT env to it
set ICINGA_CA_CERT (base64 ca.crt -w 0)
```

4. Generate a server.key with 2048bit
```
openssl genrsa -out server.key 2048
# Set ICINGA_SERVER_KEY env to it
set ICINGA_SERVER_KEY (base64 server.key -w 0)
```

5. According to the server.key generate a server.csr:
```
openssl req -new -key server.key -subj "/CN=${MASTER_IP}" -out server.csr
```

6. According to the ca.key, ca.crt and server.csr generate the server.crt:
```
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 10000
# Set ICINGA_SERVER_CERT env to it
set ICINGA_SERVER_CERT (base64 server.crt -w 0)
```

### easyrsa

`easyrsa` can be used to manually generate certificates for your cluster.

1. Download, unpack, and initialize the patched version of easyrsa3.
```
curl -L -O https://storage.googleapis.com/kubernetes-release/easy-rsa/easy-rsa.tar.gz
tar xzf easy-rsa.tar.gz
cd easy-rsa-master/easyrsa3
./easyrsa init-pki
```

2. Set `MASTER_IP` ENV to Kubernetes Master Node IP

3. Generate a CA. (--batch set automatic mode. --req-cn default CN to use.)
```
./easyrsa --batch "--req-cn=${MASTER_IP}@`date +%s`" build-ca nopass
```

4. Generate server certificate and key
```
./easyrsa --subject-alt-name="IP:${MASTER_IP}" build-server-full kubernetes-master nopass
```

5. Set ENV to use in Secret
```
# Set ICINGA_CA_CERT
set ICINGA_CA_CERT (base64 pki/ca.crt -w 0)
# Set ICINGA_SERVER_KEY
set ICINGA_SERVER_KEY (base64 pki/private/kubernetes-master.key -w 0)
# Set ICINGA_SERVER_CERT
set ICINGA_SERVER_CERT (base64 pki/issued/kubernetes-master.crt -w 0)
```


## Acknowledgement

This documentation is adapted from [kubernetes.io]((https://kubernetes.io/docs/admin/authentication/#appendix)). 
