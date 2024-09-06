# MLine

## Requirements

### Summary

The presented software is a CI/CD platform that allows user to launch complex pipelines in containers both from VSCode and web interface

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

![Architecture](diagrams/architecture.drawio.svg)
