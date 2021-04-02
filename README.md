<p align="center">
<img src="https://raw.githubusercontent.com/astals/NaAv/main/logo.png" />

**TLDR;** NaAV configures the system to simulate it's a malware analysis sandbox.

**TLDR2;** NaAV can check if the system is a honeypot.

**Stable Release:** None, still on development

## <ins>Why use NaAV as a defender?</ins>

Most advanced malware binaries use sandbox detection to avoid being analyzed by enterprise antimalware solutions or reversing teams.

NaAV configures the target system to make the malware think it has landed on an analysis system and in most cases induce to delete itself, stop the execution, don't download stage 2 binaries/modules, etc.

### Example Use case 2
The Check option (naav.exe --check) allows you to detect in seconds if your sandboxes or honeypods can be easily detected by binaries with anti-VM checks.

## <ins>Why use NaAV as an attacker?</ins>

The Check option (naav.exe --check) allows you to detect in seconds if you have landed on a poorly configured honeypot.

## <ins>Basic usage</ins>
naav.exe --install [configFile] -> Install, this action requires a configuration file, you can see an example on https://github.com/astals/NaAV/blob/main/config.json

naav.exe --uninstall -> Uninstall, to uninstall is recommended running 'C:\\Program Files (x86)\\NaAV\\naav.exe --uninstall' instead of using the downloaded file

naav.exe --update -> Update

naav.exe --check -> Check, this action checks your system to know how many Virtual Machine checks it passes

naav.exe --version -> Versions (installed and current binary)

naav.exe -h/--help -> Help

naav.exe -v[X] -> Verbosity level (example: naav.exe --check -v2)
<img src="https://raw.githubusercontent.com/astals/NaAv/main/vebosityLevels.png" />

## <ins>NaAV Numbers</ins>
|Functionality | Count (approx.) |
|:-------------|:-------------:|
| Guest Files | 200 |
| Guest Drivers | 15 |
| Guest Processes | 5 |
| Analysis Tools Processes | 45 |
| Network Interfaces | 10 |
| Registry Keys & ValueNames | 130 |
| Hardware Recognition Checks | 70 |

### Virtualbox Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|
| Guest Files | 40 |
| Guest Drivers | 5 |
| Guest Network Interfaces | 2 |
| Guest Processes | 2 |
| Guest Registry Keys & ValueNames | 90 |

### VMware Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|
| Guest Files | 140 |
| Guest Drivers | 10 |
| Guest Network Interfaces | 4 |
| Guest Processes| 4 |
| Guest Registry Keys & ValueNames | 40 |

### QEMU Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|

### Hyper-V Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|
| Guest Network Interfaces | 1 |

### Wine Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|


### Parallels Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|
| Guest Network Interfaces | 1 |


### Sandboxie Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|
| Files | 22 |
| Processes| 1 |

### Analysis Tools Emulation
|Functionality | Count (approx.) |
|:-------------|:-------------:|
| Fake Processes | 45 |

## <ins>Roadmap</ins>

|Functionality | Status | Version | 
|:-------------|:-------------:|:-------------:|
| Output in Json format| Backlog | x.x |
| Hardware recognition hooks | Backlog | x.x |
| Central management server | Backlog | x.x |

|Task | Status |  
|:-------------|:-------------: |
| Add more emulated systems | WIP |
| Add more items to currently emulated systems | WIP |
| Community config files repository | Not scheduled |
