# MLine

## Team

### Name

0x22E

### Project

Web editor for reproducible pipelines

### Members

- Ksenia Poliakova
- Oleg Sidorenkov
- Anton Timonin
- Egor Timonin
- Tsvetkova Maria Andreevna

## Requirements

### Summary

Web editor for reproducible pipelines designed for creating, editing, validating specifications with a library of predefined and custom types that follow an ML experiment domain model defined as a semantic network represented in YAML. The editor could be used as an extension in VS Code or as Web App (WPA) served from website.

### Stakeholders

- Developers
- DevOps/SRE-engineers
- Q/A-engineers
- Other tech specialists and managers who interact with CI/CD
- Investors
- Society

### Expected Needs

- Delivary releases
- Rollback releases
- Validate and test releases

### Features

- **Store specifications:** allows users save specification either by their own or via customizable autosave.
- **Edit specifications:** built-in editor allows users to easily edit pipelines specifications without need to leave website or IDE.
- **Specifications validation:** automatically checks job specifications for syntax errors, helps validate pipeline scripts, ensuring they are error-proof without need to execute.
- **Code hints**: includes code highlighting, autocompletion and prompts (e. g. entity name automatic suggestions) and highly customizable linting.

### Constraints

- **Security**
The system must provide secure access to specifications. It will be achived via authentication, authorization and ensuring the security of job runners and other infrastructure components. Moreover, we embrace whitelist approach.

- **Maintainability**
The system should be easy to extend with new features and integrations. It will be achieved with multilayered architectury approach of writing code. Also, every entity will have separate non-intersecting functionality which will make system more flexible.

- **Usability**
The platform must support multiple interfaces: VS Code extension and web UI. It lets user to operate efficiently. The UI must be friendly, so user with low experience of using CI/CD can navigate easily. The will also provide a web-editor with auto-completion hints.

- **Reliability**
  - **Availability**: The service have to provide a high level of SLA (e. g. 99.9%) and be reachable/available as much time as possible.
  - **Recoverability**: The system must be able to restore to a functional state after failures within not longer period than. Regular back ups of configurations, results of pipelines and jobs must be provided, too.

- **Observability**
For maintainers debug logging and tracing must be provided. For users, container logs should be provided

## Architecture

### Draft

![Architecture](diagrams/draft-architecture.drawio.svg)
