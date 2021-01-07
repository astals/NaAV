# NaAV - Non an AntiVirus
**TLDR;** NaAV makes the system simulate it's a malware analysis sandbox.

## Why use NaAV?
Most advanced malware binaries use sandbox detection to avoid being analyzed by enterprise antimalware solutions or reversing teams.

NaAV configures the target system to make the malware think it has landed on an analysis system and in most cases induce to delete itself, stop the execution, don't download stage 2 binaries/modules, etc. 
