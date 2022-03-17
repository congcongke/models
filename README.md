# models
harbor models

```
swagger generate model -f swagger/swagger.yaml --template-dir swagger/templates --additional-initialism=CVE --additional-initialism=GC

mv models/* ./
rm -rf models
```