# codedx-toml2json

This repository stores the toml2json program, which you can use to convert TOML to JSON in your pwsh-based Code Dx add-in scripts.

You can run the following toml2json command:

```
./toml2json -tomlFile ./request.toml -jsonFile ./request.json
```

to convert the following request.toml file:

```
[context]
target = ''                                   # the URL where the scan starts

[scanOptions]
runActiveScan = false                         # the decision to run an active scan (when true)

[reportOptions]
minRiskThreshold = 0                          # the minimum risk code for report findings
minConfThreshold = 0                          # the minimum confidence for report findings

[authentication]
type = "none"                                 # the authentication type: none, formAuthentication, or scriptAuthentication
```

to the following JSON in request.json:

```
{
  "authentication": {
    "type": "none"
  },
  "context": {
    "target": ""
  },
  "reportoptions": {
    "minconfthreshold": 0,
    "minriskthreshold": 0
  },
  "scanoptions": {
    "runactivescan": false
  }
}
```

Here's an example that shows how to read the JSON data in pwsh:

```
$   pwsh
PS> Get-Content ./request.json | ConvertFrom-Json

authentication context    reportoptions                             scanoptions
-------------- -------    -------------                             -----------
@{type=none}   @{target=} @{minconfthreshold=0; minriskthreshold=0} @{runactivescan=False}
```
