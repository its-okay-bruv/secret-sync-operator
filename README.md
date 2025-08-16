# TLS Secret Sync Operator

The **TLS Secret Sync Operator** is a lightweight Kubernetes operator that ensures a single TLS Secret is automatically **replicated and kept in sync across multiple namespaces**.  

In multi-tenant clusters (like **Gardener Shoots**) it is common for several applications or ingress controllers to require the same TLS certificate. Instead of manually copying or scripting, this operator automates the process and guarantees that all targets stay up to date.  

---

## ✨ Features
- Single source of truth for TLS certificates  
- Automatic propagation of TLS Secret changes (e.g., after renewal)  
- Cross-namespace synchronization with optional annotation/label copying  
- Clean-up of target Secrets on CR deletion (via finalizers)  
- Status reporting per target namespace  

---

## 🚀 Example

```yaml
apiVersion: certsync.yourdomain.io/v1alpha1
kind: TLSSecretSync
metadata:
  name: prod-cert-sync
  namespace: cert-admin
spec:
  sourceRef:
    namespace: cert-admin
    name: tls-wildcard-prod
  targets:
    namespaces:
      - default
      - payments
      - web
  copyAnnotations: true
  copyLabels: true
  pruneOnDelete: true
