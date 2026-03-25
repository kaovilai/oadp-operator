[![edge-catalog](https://github.com/kaovilai/oadp-operator/actions/workflows/edge-catalog.yml/badge.svg)](https://github.com/kaovilai/oadp-operator/actions/workflows/edge-catalog.yml)

Catalog source for installing latest commits from [OADP Operator](https://github.com/openshift/oadp-operator). Not supported.

Install catalogsource
```sh
oc apply -f https://raw.githubusercontent.com/kaovilai/oadp-operator/edge/catalogsource.yaml
```

Then filter Sources by `Tiger’s OADP Edge`

![image](https://user-images.githubusercontent.com/11228024/183729203-031b49fb-c056-4601-bd2f-2dba0d9aa248.png)

or on OpenShift Web Console
1. Administration
2. Cluster Settings
3. Configuration
4. OperatorHub
5. Sources
6. Create CatalogSource
   1. name: `oadp-edge`
   2. display name: `Tiger’s OADP Edge`
   3. publisher: `Tiger Kaovilai`
   4. image: `ghcr.io/kaovilai/oadp-operator-catalog:edge`
   5. Availability: *Cluster-wide CatalogSource*
   6. *Create*
7. You should now be able to install latest OADP Operator.

## Channels
Multiple nightly channels are available for different OADP versions:
- `dev` - Latest from oadp-dev branch
- `stable-1.6` - Latest from oadp-1.6 branch
- `stable-1.5` - Latest from oadp-1.5 branch
- `stable-1.4` - Latest from oadp-1.4 branch
- `stable-1.3` - Latest from oadp-1.3 branch
- `stable-1.2` - Latest from oadp-1.2 branch
- `stable-1.1` - Latest from oadp-1.1 branch
- `stable-1.0` - Latest from oadp-1.0 branch

Select your desired channel when installing the operator to get nightlies from that version branch.

Images are tagged with git ref from openshift/oadp-operator used to build them.

Available catalog tags:
- `edge` receive automatic updates (hopefully) - contains all channels (dev, stable-1.6, stable-1.5, etc).
- `dev` - Latest dev channel only
- `stable-1.6` - Latest stable-1.6 channel only
- `stable-1.5` - Latest stable-1.5 channel only
- `stable-1.4` - Latest stable-1.4 channel only
- `stable-1.3` - Latest stable-1.3 channel only
- `stable-1.2` - Latest stable-1.2 channel only
- `stable-1.1` - Latest stable-1.1 channel only
- `stable-1.0` - Latest stable-1.0 channel only
- `$(git rev-parse upstream/<branch>)` pin to a specific upstream commit from any branch.

Version numbering:
Each channel uses a unique version prefix to distinguish builds:
- dev channel: 99.0.YYYYMMDDHHMM
- stable-1.6 channel: 1.60.YYYYMMDDHHMM
- stable-1.5 channel: 1.50.YYYYMMDDHHMM
- stable-1.4 channel: 1.40.YYYYMMDDHHMM
- stable-1.3 channel: 1.30.YYYYMMDDHHMM
- stable-1.2 channel: 1.20.YYYYMMDDHHMM
- stable-1.1 channel: 1.10.YYYYMMDDHHMM
- stable-1.0 channel: 1.00.YYYYMMDDHHMM

Where YYYYMMDDHHMM is year, month, date, hour, minutes (UTC)

Icon credits:
<a href="https://commons.wikimedia.org/wiki/File:Tiger_passant_guardant.svg">User:Hellerick</a>, <a href="https://creativecommons.org/licenses/by-sa/4.0">CC BY-SA 4.0</a>, via Wikimedia Commons
