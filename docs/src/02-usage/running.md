
# Running the Application

To deploy the Sterling Data Exchange components:

## 1. Prepare Your Configuration files

You can use the configuration found [here](https://github.com/Knickkennedy/sterling-data-exchange-deployer/blob/main/examples/jobs/example-job-run.yaml) to deploy Sterling Data Exchange products without any local dependencies.

## 2. Save the changes you've made

## 3. Apply the configuration to the cluster

You can either:

```bash
oc login <url> --username=myuser --password=mypass
oc apply -f <filename>.yaml
```

or if you're on the Openshift Web UI, you can click the + symbol next to your username in the upper-right hand corner and then paste the entire yaml file in.

---

For advanced usage and customization, see the [Configuration](https://knickkennedy.github.io/sterling-data-exchange-deployer/02-usage/configuration) section.

To better understand the kubernetes job and necessary permissions, see [Reference](job-reference.md)
