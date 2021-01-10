<p align="center">
<img src="https://raw.githubusercontent.com/astals/NaAv/main/logo.png" />
          
**TLDR;** NaAV makes the system simulate it's a malware analysis sandbox.

**Stable Release:** None, still on development

## <ins>Why use NaAV as a defender?</ins>

Most advanced malware binaries use sandbox detection to avoid being analyzed by enterprise antimalware solutions or reversing teams.

NaAV configures the target system to make the malware think it has landed on an analysis system and in most cases induce to delete itself, stop the execution, don't download stage 2 binaries/modules, etc.

### Example Use case 2
The Check option (naav.exe -c) alows you to detect in seconds if your sandboxes or honeypods can be easily detected by binaries with anti-VM checks.

## <ins>Why use NaAV as an attacker?</ins>

The Check option (naav.exe -c) alows you to detect in seconds if you have landed on a poorly configured Honeypot.

## <ins>Basic usage</ins>
naav.exe --install [configFile] -> Install, this acction requires a configuration file, you can see an example on https://github.com/astals/NaAV/blob/main/config.json

naav.exe --uninstall -> Uninstall

naav.exe --update -> Uninstall

naav.exe --check -> Check, this action checks your system in order to know how many Virtual Machine checks it passes

naav.exe -v/--version -> Version

naav.exe -h/--help -> Help

## <ins>Functionalities</ins>
### VMware Emulation
|Functionality | Status | Version |
|:-------------|:-------------:|:-------------:|
| Fake C:\\WINDOWS\\system32\\drivers\\vmhgfs.sys driver | Implemented | 0.1 |
| Fake C:\\WINDOWS\\system32\\drivers\\vmmouse.sys driver | Implemented | 0.1 |
| Fake guest Network interface | WIP | 0.1 |

### Virtualbox Emulation
|Functionality | Status | Version |
|:-------------|:-------------:|:-------------:|
|Fake VirtualBox Guest Additions Files (C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\*)| Implemented | 0.1 |
| Fake drivers | WIP | 0.1 |
| Fake guest Network interface | WIP | 0.1 |

### Analysis Tools Emulation
|Functionality | Status | Version |
|:-------------|:-------------:|:-------------:|
| Fake processes (~40, defined in config file) | WIP | 0.1 |

