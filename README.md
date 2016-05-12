# HTTP and SOCKS proxy

A simple HTTP and SOCKS proxy based on combining the following two go libraries:

* go-socks5 - https://github.com/armon/go-socks5
* goproxy - https://github.com/elazarl/goproxy

This application just runs the two proxies in the same process with command line flags to controll them.

    $ ./http-socks-proxy -h
    Usage of ./http-socks-proxy:
      -hhost string
        	The host/ip on which the HTTP proxy listens for connections (default "127.0.0.1")
      -hport int
        	The port on which the HTTP proxy listens for connections (default 2223)
      -shost string
        	The host/ip on which the SOCKS proxy listens for connections (default "127.0.0.1")
      -sport int
        	The port on which the SOCKS proxy listens for connections (default 2222)
      -verbose
        	If verbose logging should be done


## Use case

This can be used to access services running behind a firewall through a reverse SSH tunnel. E.g.

Computer A is behind a firewall. Start the proxy on Computer A then ssh from Computer A to Computer B and configure remote tunnels for the two ports:

    ssh -R 2222:localhost:2222 -R 2223:localhost:2223 user@computerb

After this the HTTP and SOCKS proxies can be accessed on Computer B on port 2222 and 2223.

Now one can configure a web browser on Computer B to use the SOCKS or HTTP proxy and access resources only accessable from Computer A.

### Configure Subversion to use a HTTP proxy

Edit **~/.subversion/servers** and set the *http-proxy-host* and *http-proxy-port* configuration parameters in the 'global' section:

    http-proxy-host = 127.0.0.1
    http-proxy-port = 2223    

### Configure Apache Ant to use a HTTP proxy

    $ ANT_OPTS="-Dhttp.proxyHost=127.0.0.1 -Dhttp.proxyPort=2223" ant






## Cross-compile for Windows

Cross-compile for windows using the following command:

    $ GOOS=windows GOARCH=386 go build -o http-socks-proxy.exe main.go
    
