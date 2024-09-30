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

- **Store configurations** - After user set up configuration, they can save it for future. So after that users can refer to this configurations either for launching pipelines or for creating new ones.
- **Edit configurations** - Project provides easy to use built-in editor that allows users to quickly modify and save updates for pipelines configurations. In addition editor helps users to write configurations correctly with autocompletion and hints to save users time.
- **Validation of configuration** - Automatical checking job configurations for syntax errors is included and it helps to validate pipeline scripts, ensuring they are error-free before execution. If validation of configuration file is not passed, user will be informed about that.
- **VS Code extension** - Project provides users special VS Code extension for launching piplines exactly from VS Code. This option can help users not being distracted by visiting web site to track results, but have a quick visual demonstration of the resuls of tests, etc.
- **Web application** - There is also web version, for users who wants to get comprehensive information about launched jobs. Web application can suggest user-friendly interface, where logs of each stage can be accessed, the website is completed with up-to-date technology Progressive Web Application.


**Run pipelines that are described in the configurations** - allows users to execute defined pipelines from Configuration Storage, leveraging robust scheduling and triggering capabilities.

**Execute multi-stage tasks:** users can decompose tasks into smaller ones and launch multi-staged pipelines that have parts relying on each other.

**Logging status information:** possibility to see logs to identify problems in cases of faults, track tests passing or adjusting stages for particular purposes.

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
