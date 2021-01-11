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

naav.exe --uninstall -> Uninstall, to uninstall is recommended running 'C:\\Program Files (x86)\\NaAV\\naav.exe -u' instead of using the downloaded file

naav.exe --update -> Update

naav.exe --check -> Check, this action checks your system in order to know how many Virtual Machine checks it passes

naav.exe -v/--version -> Versions (installed and current binary)

naav.exe -h/--help -> Help


## <ins>NaAV Numbers</ins>
*with default config

|Functionality | nº |
|:-------------|:-------------:|
| Fake Guest Drivers & Files | ~200 |
| Fake Guest Processes & Services | ~5 |
| Fake Analysis Tools Processes | ~40 |
| Fake Network Interfaces | -- |
| Fake Registry Keys | -- |

## <ins>Install & Check Functionalities</ins>
### VMware Emulation
|Functionality | Status | Version | Details|
|:-------------|:-------------:|:-------------:|:-------------|
| Fake Guest Drivers | Implemented | 0.1 | C:\\WINDOWS\\system32\\drivers\\vmhgfs.sys, C:\\WINDOWS\\system32\\drivers\\vmmouse.sys|
| Fake Guest Files | Implemented | 0.1 |	"C:\\Windows\\System32\\vm3d*, C:\\Windows\\System32\\vmGuestLib*,	C:\\Windows\\System32\\vmhgfs.dll, C:\\Windows\\System32\\VMWSU.DLL|
| Fake Guest Network Interfaces | WIP | 0.1 | |
| Fake Guest Processes & Services| WIP | 0.1 | Defined in config file |
| Fake Guest Registry Keys | WIP | 0.1 | |

### Virtualbox Emulation
|Functionality | Status | Version | Details|
|:-------------:|:-------------:|:-------------:|:-------------|
| Fake Guest Drivers | WIP | 0.1 | |
| Fake VirtualBox Guest Additions Files| Implemented | 0.1 | C:\\Program Files\\Oracle\\VirtualBox Guest Additions\\*, C:\\Windows\\System32\\VBox*  |
| Fake Guest Network Interfaces | WIP | 0.1 | |
| Fake Guest Processes & Services| WIP | 0.1 | Defined in config file |
| Fake Guest Registry Keys | WIP | 0.1 | |

### Analysis Tools Emulation
|Functionality | Status | Version | Details|:-------------|
|:-------------|:-------------:|:-------------:|
| Fake Processes | WIP | 0.1 | ~40, defined in config file |

## <ins>Functionalities only availabe on Check</ins>

|Functionality | Status | Version | Details|
|:-------------|:-------------:|:-------------:|:-------------|
| Hardware Recognition | WIP | 0.1 | CPU, RAM, HDD, etc |
