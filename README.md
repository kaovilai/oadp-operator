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
   2. display name: `Tiger's OADP Edge`
   3. publisher: `Tiger Kaovilai`
   4. image: `ghcr.io/kaovilai/oadp-operator-catalog:edge`
   5. Availability: *Cluster-wide CatalogSource*
   6. *Create*
7. You should now be able to install latest OADP Operator.

Images are tagged with git ref from openshift/oadp-operator used to build them.

Available catalog tags:
- `edge` receive automatic updates (hopefully).
- `$(git rev-parse upstream/master)` pin to a specific upstream/master commit.

GitHub Actions is set to refresh available images every 30 minutes.
Star this repo to trigger a refresh.

Version numbering:
0.99.YYYYMMDDHHMM, year month date hour minutes (UTC)

Icon credits:
<a href="https://commons.wikimedia.org/wiki/File:Tiger_passant_guardant.svg">User:Hellerick</a>, <a href="https://creativecommons.org/licenses/by-sa/4.0">CC BY-SA 4.0</a>, via Wikimedia Commons
