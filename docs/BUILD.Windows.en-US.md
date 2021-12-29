# Golden Point build instruction

## Prepare environment

 1) Install **TDM-GCC** compiler tools. You may download it at address https://jmeubank.github.io/tdm-gcc/download/
 2) Install **GitSCM** system https://git-scm.com/downloads

## Build with PowerShell script

 1) Invoke PowerShell script in source code directory (example, directory C:\Golden\src)

    C:\Golden\src> powershell -executionpolicy RemoteSigned -file "build-windows.ps1"

## Build manually 

 1) (Optional) Prepare Golang environemnt for target platform
 2) Execute next commands
 
    C:\Golden\src> go generate
    C:\Golden\src> go build
