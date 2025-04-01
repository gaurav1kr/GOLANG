
# Structure of a Simple Golang-based Kubernetes Controller

This document explains the components and flow of a basic Golang-based Kubernetes controller, including CRD, Informer, and the Reconciliation loop.

---

## 1️⃣ Custom Resource Definition (CRD)

A CRD extends Kubernetes with a custom resource.

### 🛠️ Example: `MyApp` CRD
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: myapps.example.com
spec:
  group: example.com
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                replicas:
                  type: integer
  scope: Namespaced
  names:
    plural: myapps
    singular: myapp
    kind: MyApp
```

---

## 2️⃣ Go Code Structure

### 📁 Folder Layout
```
my-controller/
├── main.go
├── api/
│   └── v1/
│       └── myapp_types.go
├── controllers/
│   └── myapp_controller.go
```

---

## 3️⃣ Informer and Event Handlers

Informers watch for changes to Kubernetes resources and trigger event handlers.

```go
informer := factory.Core().V1().Pods().Informer()
informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
    AddFunc: onAdd,
    UpdateFunc: onUpdate,
    DeleteFunc: onDelete,
})
```

---

## 4️⃣ Controller / Reconcile Loop

The Reconcile function ensures the cluster's actual state matches the desired state.

```go
func (r *MyAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var myapp examplev1.MyApp
    if err := r.Get(ctx, req.NamespacedName, &myapp); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // Ensure desired state, e.g., Deployment has correct replicas

    return ctrl.Result{RequeueAfter: time.Minute}, nil
}
```

---

## 5️⃣ Controller Manager Setup

Registers and starts the controller.

```go
func main() {
    mgr, _ := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
        Scheme: scheme,
    })

    _ = (&controllers.MyAppReconciler{
        Client: mgr.GetClient(),
        Scheme: mgr.GetScheme(),
    }).SetupWithManager(mgr)

    mgr.Start(ctrl.SetupSignalHandler())
}
```

---

## ✅ Summary Flow

```text
1. CRD defines your custom resource (e.g., MyApp)
2. Informer watches for changes to MyApp objects
3. EventHandler triggers the Reconcile loop
4. Reconciler checks current state vs desired state
5. Takes action: create/update/delete other resources
```

---

## ⚙️ Tools You Can Use

| Tool              | Purpose                                 |
|-------------------|-----------------------------------------|
| `kubebuilder`     | Scaffold CRDs, controllers, reconcilers |
| `client-go`       | Kubernetes API client                   |
| `controller-runtime` | Simplified controller library       |

