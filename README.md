tenv - Teleport Version Manager
===============================

`tenv` is a wrapper around Teleport binaries to make it easy to install and switch between versions.

Installation
------------

```
# Download and compile the binary
go get github.com/kbence/tenv/cmd/...

# Create symlinks next to the binary
tenv link
```

Usage
-----

`tenv` wraps the Teleport binaries to make is easy to switch between versions using either a selected version, or a version specified in an environment variable.

### Install a specific version of Teleport

```
tenv install <version>
```

`tenv` downloads binaries to `$HOME/.tenv/versions/<version>/bin/`.

### Select the Teleport version

To redirect calls, you'll need to select the current version using the following command:

```
tenv use <version>
```

`tenv` stores this information at `$HOME/.tenv/selected-version`

### Select version via environment variable

The `TELEPORT_VERSION` environment variable can be used to override the currently selected version of Teleport:

```
$ tenv use 10.2.6

$ teleport version
Teleport v10.2.6 git:v10.2.6-0-g46438b451 go1.18.6

$ TELEPORT_VERSION=9.0.0 teleport version
Teleport v9.0.0 git:v9.0.0-0-g1fa8857aa go1.17.7
```

### Switch between `tsh` profiles

Use the `select-profile` subcommand to switch between `tsh` profiles (use the domain name you use to connect to your Teleport cluster):

```
tenv select-profile myteleport.example.com
```

It saves the provided domain name to `~/.tsh/current-profile`, so Teleport commands will use that cluster further on.

### Configuration

You can create a confiuration file in `$HOME/.tenv/config.yaml` to automate switching between versions depending on the currently selected cluster and/or current directory:

```
$ cat ~/.tenv/config.yaml
rules:
  - version: 10.1.1
    match:
      - currentDir: /home/myuser/projects/work-project
      - currentDir: /home/myuser/projects/home-project
        currentProfile: myteleport.example.com
```

Items are checked against the values separately, and the first matching rule will take precedence.
