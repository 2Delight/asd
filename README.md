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
- Validation of configuration
- Edit configurations (?)
- Run pipelines that are described in the configurations
- Execute multi-stage tasks
- Logging status information
- Pipelines rely on each other
- VS Code extension
- Web application 


TODO: Look in more requirement-specific constraints 
### Constraints

- **Scalability**  
The system should support both horizontal and vertical scaling to handle more increasing workloads and a growing number of simultaneously running pipelines.
  
- **Reliability**   
The service have to be accessible and functioning with a high level of availability (e.g., 99.9%), including failover strategies.
  
- **Security**   
The system must provide secure access control to configurations, jobs and pipelines, including authentication, authorization and ensuring the security of job runners and other infrastructure components.

- **Recoverability**   
The CI/CD platform must restore to a functional state after failures, including regular back ups of configurations, pipelines and jobs. After failing master node system should give master's rights to one of the slave nodes. After slave's disaster system should continue work correctly

- **Maintainability**   
The system should be easy to extend and update without a significant downtime or complete system restart.

- **Usability**    
The platform should be accessible in VS Code and a Progressive Web App (PWA) served directly from a web site. The system interface should be easy to managing pipelines, configuring jobs and monitoring results. 

- **Consistency**  
The system must guarantee that jobs and pipelines run consistently across different environments and runners. Pipeline management across multiple runners should be seamless and transparent.

- **Traceability**  
The platfrom must provide linking between job runners and pipelines, including audit logs and tracking changes in configurations.


## Architecture

### Draft

![Architecture](diagrams/draft-architecture.drawio.svg)
