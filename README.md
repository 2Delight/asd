# MLine

## Requirements

### Summary

The framework is able to run complex multistage pipelines from docker containers and services with code adaptation within the stages. The specification of a pipeline follows a ML experiment domain model defined as a semantic network represented in YAML. A web editor would be both an extension to VS Code and a Progressive Web App (PWA) served directly from a web site. The editor would implement a visual editor of the experiment specification with a library of predefined and custom types, code completion and code linting

### Stakeholders

- Developers
- DevOps/SRE-engineers
- Q/A-engineers
- Investors
- Society

### Expected Needs

- Delivary releases
- Rollback releases
- Validate and test releases

TODO: Add more description
### Features

- Store configurations
- Edit configurations (?)
- Run pipelines that are described in the configurations
- Execute multi-stage tasks
- Logging status information
- Pipelines rely on each other
- VS Code extension
- Web application 
- Validation of configuration

TODO: Look in more requirement-specific constraints 
### Constraints

- Max memory for configuration: 2Mb
- Max X multi-stage tasks
- Container Resources RAM 1Gb, 2 CPU
- Network (?)
- Filesystem (?)
- Timeout
- Job has UUID

## Architecture

### Draft

![Architecture](diagrams/draft-architecture.drawio.svg)
