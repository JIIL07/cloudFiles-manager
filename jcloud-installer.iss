[Setup]
AppName=Jcloud Client
AppVersion=0.0.1
DefaultDirName={commonpf}\JcloudClient
DefaultGroupName=Jcloud Client
OutputDir=userdocs:Output
OutputBaseFilename=JcloudClientSetup
Compression=lzma
SolidCompression=yes
SetupIconFile=installer-icon.ico

[Files]
Source: "cmd/cloud/cloud.exe"; DestDir: "{app}/bin"; Flags: ignoreversion
Source: "config/*.yaml"; DestDir: "{app}/config"; Flags: ignoreversion
Source: "internal/client/*.*"; DestDir: "{app}/src"; Flags: ignoreversion recursesubdirs createallsubdirs

[Run]
Filename: "{app}\bin\cloud.exe"; Description: "Run Jcloud Client"; Flags: nowait postinstall skipifsilent

[Tasks]
Name: "addtopath"; Description: "Add Jcloud Client to PATH"; GroupDescription: "Additional icons:"; Flags: unchecked

[Registry]
Root: HKCU; Subkey: "Environment"; ValueType: expandsz; ValueName: "PATH"; ValueData: "{olddata};{app}/bin"; Tasks: addtopath
