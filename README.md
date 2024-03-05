<!--
SPDX-FileCopyrightText: 2023 Deutsche Telekom AG

SPDX-License-Identifier: CC0-1.0    
-->

<p align="center">
  <img src="./docs/img/probe_logo.png" alt="Vortex logo" width="200px" height="200px">
  <h1 align="center">Probe</h1>
</p>

## Overview
Probe is a tool for performing Horizon SSE E2E tests 📡 

## Running Probe
Before running Probe you might want to generate a configuration file instead of using environment variables.
If you want to run Probe from source replace `probe` with `go run .`.

### Generate a new configuration
```bash
probe init
```

### Run Probe
```bash
probe start --template template.json
```
### Using a template
To define the content of the event you are sending to Horizon, you have to provide a template in the form of a JSON file.
The `{{ EventId }}` placeholder is automatically replaced with a fresh UUID.

**A template could look as follows**:
```json
{
  "id":"{{ EventId }}",
  "source":"http://apihost/some/path/resource/1234",
  "specversion":"1.0",
  "type":"my.example.event.v1",
  "datacontenttype":"application/json",
  "dataref":"http://apihost/some/api/v1/resource/1234",
  "data":{
    "hello": "world"
  },
  "dataschema":"http://apihost/schema/definition.json"
}
```

### Behaviour
The success status of the application can be determined by checking the exit code. 
If an error occurred or the configured thresholds have been violated the application will exit with
exit-code **1** instead of **0**.

## Configuration
All the configuration properties can be set in your `config.yml` or with an environment variable. 

| Path                         | Variable                           | Type   | Default                                    | Description                                     |
|------------------------------|------------------------------------|--------|--------------------------------------------|-------------------------------------------------|
| logLevel                     | PROBE_LOGLEVEL                     | string | info                                       | Defines the log-level.                          |
| publishing.endpoint          | PROBE_PUBLISHING_ENDPOINT          | string | https://horizon.example.com/events         | Horizon endpoint for publishing events.         |
| publishing.oidc.url          | PROBE_PUBLISHING_OIDC_URL          | string | https://oidc.example.com/                  | OIDC token endpoint of the idp.                 |
| publishing.oidc.clientId     | PROBE_PUBLISHING_OIDC_CLIENTID     | string | client-id                                  | The client-id that is passed to the idp.        |
| publishing.oidc.clientSecret | PROBE_PUBLISHING_OIDC_CLIENTSECRET | string | client-secret                              | The client-secret that is passed to the idp.    |
| consuming.endpoint           | PROBE_CONSUMING_ENDPOINT           | string | https://horizon.example.com/sse/somesubid  | Horizon endpoint for retrieving events via sse. |
| consuming.oidc.url           | PROBE_CONSUMING_OIDC_URL           | string | https://oidc.example.com/                  | OIDC token endpoint of the idp.                 |
| consuming.oidc.clientId      | PROBE_CONSUMING_OIDC_CLIENTID      | string | client-id                                  | The client-id that is passed to the idp.        |
| consuming.oidc.clientSecret  | PROBE_CONSUMING_OIDC_CLIENTSECRET  | string | client-secret                              | The client-secret that is passed to the idp.    |

## Docker
Probe is supposed to run in a Docker container.
If you run Probe a container with the default command you can pass arguments through the following arguments.

| Argument        | Variable                | Default       |
|-----------------|-------------------------|---------------|
| --message-count | PROBE_ARG_MESSAGE_COUNT | 3             |
| --timeout       | PROBE_ARG_TIMEOUT       | 30s           |
| --max-latency   | PROBE_ARG_MAX_LATENCY   | 5s            |
| --template      | PROBE_ARG_TEMPLATE      | template.json |

## Code of Conduct
This project has adopted the [Contributor Covenant](https://www.contributor-covenant.org/) in version 2.1 as our code of conduct. Please see the details in our [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md). All contributors must abide by the code of conduct.

By participating in this project, you agree to abide by its [Code of Conduct](./CODE_OF_CONDUCT.md) at all times.

## Licensing
This project follows the [REUSE standard for software licensing](https://reuse.software/).
Each file contains copyright and license information, and license texts can be found in the [LICENSES](./LICENSES) folder. For more information visit https://reuse.software/.

### REUSE
For a comprehensive guide on how to use REUSE for licensing in this repository, visit https://telekom.github.io/reuse-template/.   
A brief summary follows below:

The [reuse tool](https://github.com/fsfe/reuse-tool) can be used to verify and establish compliance when new files are added. 

For more information on the reuse tool visit https://github.com/fsfe/reuse-tool.