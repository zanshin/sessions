# sessions
The sessions command line application displays the active session counts for the blue and green side
of a Wildfly cluster.

## Parameters
* service -s Name of service, including the tier
* inst -i Instance count, i.e., number of application nodes, defaults to 4

## Caveats
This application relies on knowing the host naming scheme for the Wildfly application nodes. The
lines where the `host` is defined and used, will need to be updated for other installations.

## License
MIT
