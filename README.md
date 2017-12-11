# User Provided Service Manager
[![Build Status](https://travis-ci.org/bjurgess1/cf-ups-manager.svg?branch=master)](https://travis-ci.org/bjurgess1/cf-ups-manager)
## Introduction
Cloud Foundry CLI plugin to create and update User Provided Services across several Cloud Foundry spaces from a ups-manifest file.

## Installation
You can download the plugin from the CF Community Repository:
```
cf add-plugin-repo CF-Community https://plugins.cloudfoundry.org
cf install-plugin ups-manager -r CF-Community
```
Or, if you have go installed:
```
go get -u github.com/bjurgess/cf-ups-manager
cf install-plugin $GOTPATH/bin/cf-ups-manager
```
Additionally, you can download the latest binary and palce the binary on your path. Then run:
```
cf install-plugin [path/to/cf-ups-manager]
```

## Usage
First, login and set your target to the space you wish to update. Then, run:
```
cf ups-manager -f ups-manifest.yml [-u LIST_OF_USER_PROVIDED_SERVICES]
```

* -f: Path to the user-provided-service manifest yaml file
* -u: Optional ***comma separated*** list of user provided services to deploy from the manifest

ups-manifest.yml example:
```yaml
spaces:
- name: dev
  user-provided-services:
  - name: UPS1
    credentials:
      Credential1: 1
      Credential2: 2
      Credential3: 3
  - name: UPS2
    credentials:
      Credential1: 1
      Credential2: 2
      Credential3: 3
- name: qa
  user-provided-services:
  - name: UPSqa1
    credentials:
      Credential1: 1
      Credential2: 2
      Credential3: 3
- name: stage
  user-provided-services:
```

### Manifest Requirements
* Currently, a UPS must have a list of credentials associated with it
* Does not work with syslog urls yet
* Does not work with route-services yet 
* The name of the space ***must*** match the name of the space in cloud foundry
* You must have your target set to the space you wish to deploy
