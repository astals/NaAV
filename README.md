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
|Functionality | Count (aprox.) |
|:-------------|:-------------:|
| Guest Files | 180 |
| Guest Drivers | 15 |
| Guest Processes & Services | TODO |
| Analysis Tools Processes | TODO |
| Network Interfaces | 10 |
| Registry Keys & ValueNames | 130 |
| Hardware Recognition Checks | TODO |

### Virtualbox Emulation
|Functionality | Count (aprox.) |
|:-------------:|:-------------:|
| Guest Files | 40 |
| Guest Drivers | 5 |
| Guest Network Interfaces | 2 |
| Guest Processes & Services | |
| Guest Registry Keys & ValueNames | 90 |

### VMware Emulation
|Functionality | Count (aprox.) |
|:-------------|:-------------:|
| Guest Files | 140 |
| Guest Drivers | 10 |
| Guest Network Interfaces | 4 |
| Guest Processes & Services| |
| Guest Registry Keys & ValueNames | 40 |

### QEMU Emulation
|Functionality | Count (aprox.) |
|:-------------:|:-------------:|
| Guest Files | |
| Guest Drivers | |
| Guest Network Interfaces | |
| Guest Processes & Services| |
| Guest Registry Keys & ValueNames | |

### Hyper-V Emulation
|Functionality | Count (aprox.) |
|:-------------:|:-------------:|
| Guest Files | |
| Guest Drivers | |
| Guest Network Interfaces | 1 |
| Guest Processes & Services| |
| Guest Registry Keys & ValueNames | |

### Wine Emulation
|Functionality | Count (aprox.) |
|:-------------:|:-------------:|
| Guest Files | |
| Guest Drivers | |
| Guest Network Interfaces | |
| Guest Processes & Services| |
| Guest Registry Keys & ValueNames | |

### Parallels Emulation
|Functionality | Count (aprox.) |
|:-------------:|:-------------:|
| Guest Files | |
| Guest Drivers | |
| Guest Network Interfaces | 1 |
| Guest Processes & Services| |
| Guest Registry Keys & ValueNames | |

### Analysis Tools Emulation
|Functionality | Count (aprox.) |
|:-------------|:-------------:|
| Fake Processes | |



## <ins>Functionalities only availabe on Check</ins>

|Functionality | Count | 
|:-------------|:-------------:|
| Hardware Recognition | 0 |


## <ins>Roadmap</ins>

|Functionality | Status | Version | 
|:-------------|:-------------:|:-------------:|
| Hardware Recognition | WIP | 0.2 |
