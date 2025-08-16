# secret-sync-operator
TLS Secret Sync Operator automatically replicates and keeps TLS Secrets in sync across multiple namespaces. It watches a source Secret and propagates changes (like certificate rotation) to all targets, simplifying multi-namespace certificate management in Kubernetes.
