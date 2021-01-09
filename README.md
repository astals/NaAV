# <ins>NaAV - Not an AntiVirus</ins>
**TLDR;** NaAV makes the system simulate it's a malware analysis sandbox.

**Stable Release:** None, still on development

## <ins>Why use NaAV?</ins>
Most advanced malware binaries use sandbox detection to avoid being analyzed by enterprise antimalware solutions or reversing teams.

NaAV configures the target system to make the malware think it has landed on an analysis system and in most cases induce to delete itself, stop the execution, don't download stage 2 binaries/modules, etc. 

## <ins>Basic usage</ins>
naav.exe -i [configFile] -> Install, this acction requires a configuration file, you can see an example on https://github.com/astals/NaAV
naav.exe -u -> Uninstall
naav.exe -c -> Check, this action checks your system in order to know how many Virtual Machine checks it passes
naav.exe -h -> Help

## <ins>Functionalities</ins>
### VMware 
|Functionality | Status | Version |
|:-------------|:-------------:|:-------------:|
| Fake C:\\WINDOWS\\system32\\drivers\\vmhgfs.sys driver | Implemented | 0.1 |
| Fake C:\\WINDOWS\\system32\\drivers\\vmmouse.sys driver | Implemented | 0.1 |
| Fake guest Network interface | WIP | 0.1 |

### Virtualbox 
|Functionality | Status | Version |
|:-------------|:-------------:|:-------------:|
| Fake drivers | WIP | 0.1 |
| Fake guest Network interface | WIP | 0.1 |

### Analysis tools 
|Functionality | Status | Version |
|:-------------|:-------------:|:-------------:|
| Fake processes | WIP | 0.1 |

