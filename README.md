Render a [helm](https://github.com/kubernetes/helm) manifest to a directory.

# Installation
```
go get -u github.com/databus23/schelm
```
# Usage:

```
helm install --dry-run --debug CHART > manifest.txt
schelm OUTPUT_DIR < manifest.txt

or

helm install --dry-run --debug CHART | schelm OUTPUT_DIR

or

helm get RELEASE manifest | schelm OUPUT_DIR

```

# Example:

```
➜ helm get eloping-saola manifest | schelm output/
2016/10/21 15:50:12 Writing output/mariadb/templates/deployment.yaml
2016/10/21 15:50:12 Writing output/mariadb/templates/pvc.yaml
2016/10/21 15:50:12 Writing output/mariadb/templates/secrets.yaml
2016/10/21 15:50:12 Writing output/mariadb/templates/svc.yaml
➜ tree output/
output/
└── mariadb
    └── templates
        ├── deployment.yaml
        ├── pvc.yaml
        ├── secrets.yaml
        └── svc.yaml

2 directories, 4 files
```
