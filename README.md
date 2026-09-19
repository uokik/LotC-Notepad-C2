# LoTC-Notepad
## A lightweight Living off the Cloud Command & Control (C2) implementation

> A small proof of concept implementation that uses a cloud-hosted
> note service as a communication channel between an operator and an agent.


## Overview
LotC-Notepad demonstrates a C2 architecture where a cloud based note-taking service is used as the communication channel instead of a dedicated C2 server.

The project consists of two components:

- **Agent** - runs on the victim system and polls the cloud service for commands.
- **Operator Panel** - sends commands and retrieves command output.

Communication between the components is performed through encoded note content.

## Architecture

