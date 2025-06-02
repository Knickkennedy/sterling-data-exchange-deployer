
# Running the Application

To deploy the Sterling Data Exchange components:

## 1. Prepare Your Configuration files

You can use the configuration found [here](https://github.com/Knickkennedy/sterling-data-exchange-deployer/blob/main/examples/jobs/example-job-run.yaml) to deploy Sterling Data Exchange products without any local dependencies. If you want a minimal way to run the product, then the only changes you'll need to make are to add your IBM Entitlement Key to the secret in the example-job-run.yaml and to select which products you wish to deploy in the global.yml section of the configmap. The examples should work as-is.

## 2. Save the changes you've made

## 3. Apply the configuration to the cluster

You can either:

```bash
oc login <url> --username=myuser --password=mypass
oc apply -f <filename>.yaml
```

or if you're on the Openshift Web UI, you can click the + symbol next to your username in the upper-right hand corner and then paste the entire yaml file in.

---

For advanced usage and customization, see the [Configuration](configuration/index.md) section.

To better understand the kubernetes job and necessary permissions, see [Reference](job-reference.md)
