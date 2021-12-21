# shorthike

A portable golang utility for performing simple TCP checks and logging the results

Famous last words --- "It's just a short hike...it won't take very long at all"

Find out just how hard the journey will be before departing

# How to use shorthike

Best run inside a docker container. Check if your container or pod can connect to its dependency

# Environment Variables

TCPHEALTH_HOST_01 - either the hostname or ip address of the host (or service) you wish to connect to
TCPHEALTH_PORT_01 - the TCP port on that host (or service)

...you should be able to run as many number of checks as needed by incrementing the number above...

TCPHEALTH_RUN_INTERVAL - how often you'd like the check performed, measured in golang time.ParseDuration()

TCPHEALTH_CONNECT_TIMEOUT - how long the TCP connection attempt should try before failing

LOG_LEVEL=DEBUG - debug logging turned on
