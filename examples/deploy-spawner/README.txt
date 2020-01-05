This directory contains a small example for how you might deploy a game-specific
instance of the spawner stack which connects to an external master server.

You could use docker volumes (or some other method) for customizing the config
and game server binary which is deployed with the spawner server. The solution
use here is to have a new dockerfile which extends from the base spawner
docker file to inject the game-specific config and data.
