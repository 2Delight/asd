# MLine

## Requirements

### Summary

The framework is able to run complex multistage pipelines from docker containers and services with code adaptation within the stages. The specification of a pipeline follows a ML experiment domain model defined as a semantic network represented in YAML. A web editor would be both an extension to VS Code and a Progressive Web App (PWA) served directly from a web site. The editor would implement a visual editor of the experiment specification with a library of predefined and custom types, code completion and code linting.

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

- **Store configurations:** allows users save configuration for further pipeline launches.
- **Edit configurations:** built-in editor allows users to quickly change and save updates for pipelines configurations.
- **Validation of configuration:** automatically checks job configurations for syntax errors helps validate pipeline scripts, ensuring they are error-free before execution. In edit mode helps users to write configs correctly with autocompletion and hints.
- **Run pipelines that are described in the configurations:** allows users to execute defined pipelines from Configuration Storage, leveraging robust scheduling and triggering capabilities.
- **Execute multi-stage tasks:** users can decompose tasks into smaller ones and launch multi-staged pipelines that have parts relying on each other.
- **Logging status information:** possibility to see logs to identify problems in cases of faults, track tests passing or adjusting stages for particular purposes.
- **VS Code extension:** launching piplines exactly from VS Code extension can help users not being distracted by visiting web site to track results, but have a quick visual demonstration.
- **Web application:** for users who prefer using web version, platform can suggest user-friendly interface, where logs of each stage can be accessed, the website is completed with up-to-date technology Progressive Web Application.

### Constraints

- **Security**
The system must provide secure access to system data including configurations, results of jobs and pipelines. It will be achived via authentication, authorization and ensuring the security of job runners and other infrastructure components. Moreover, we embrace whitelist approach.

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
