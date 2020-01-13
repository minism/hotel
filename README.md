# Hotel

Hotel is a simple master/lobby server for multiplayer video games written in golang.  It's purpose is to facilitate the registration, discovery, and spawning of multiplayer game servers.

## Overview

There are two services which are deployed separately:

 - **hotel-master**: The master server is a web server + database that provides a REST API for clients to register and request game server information.
 - **hotel-spawner** (*optional*): The spawner is an RPC server which fulfills requests from the master server to spawn new game server processes. This component can be omitted if the game doesn't require realtime provisioning of servers.

The ideal relationship between these services is visualized in the following diagram:

![Architecture](https://raw.githubusercontent.com/minism/hotel/master/docs/architecture.png)

## Clients
Hotel was designed with Unity in mind as a client, but is agnostic to the game engine.

[hotel-client-unity] is a ready-made Unity client implementation that abstracts away much of the internal API calls.

## Example control flow

Below is an example request flow for a game client requesting that a new game server be spawned, and connecting to it:

![Spawn Flow](https://raw.githubusercontent.com/minism/hotel/master/docs/client-spawn-sequence.png)

## API

TODO

## Deployment

TODO
