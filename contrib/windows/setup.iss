
[Setup]
AppName=Golden Point
AppVersion=1.2.15
DefaultDirName={autopf}\Golden Point
DefaultGroupName=Golden Point

[Components]
Name: "main"; Description: "Main Files"; Types: full compact custom; Flags: fixed

[Tasks]
Name: desktopicon; Description: "Create a &desktop icon"; GroupDescription: "Additional icons:"; Components: main

[Files]
Source: "..\golden-windows-386.exe"; DestDir: "{app}"
Source: "..\golden-windows-amd64.exe"; DestDir: "{app}"

[Icons]
Name: "{group}\Golden Point"; Filename: "{app}\golden-windows-386.exe"; Components: main; Tasks: desktopicon
Name: "{group}\Golden Point"; Filename: "{app}\golden-windows-amd64.exe"; Components: main; Tasks: desktopicon
