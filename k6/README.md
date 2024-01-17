# Stress Test

In the folder you can find some stress tests, written to put your implementation under load.  
The tests uses [k6.io](https://k6.io/docs/), that allow us to write `scripts` in plain Javascript that are then launched
using k6 CLI.

## Local run

In order to run locally, you need to have `k6` installed on your machine.
On MacOS you can obtain it with `brew`:

```
brew install k6
```

For other OS please refer to the [official docs](https://k6.io/docs/getting-started/installation/).

From the root folder of this repo, run
```
k6 run k6/scripts/{script_name}.js
```

You can even use docker to run the script, avoiding the installation of k6 on your machine. Please refer to
the [docs](https://k6.io/docs/getting-started/running-k6/) for more information.  
_N.B._ On MacOs a local install of k6 is recommended due to performance limitation of docker machine.

## Result visualization

After a run of `k6` a report on the console is displayed.
For deeper insight, is it possible to export run results on an InfluxDB and navigate the results using Grafana.
A base setup is present in `visualization` folder. You can startup the pipeline with docker compose.

After InfluxDB and Grafana are up and running, simply invoke k6 adding an `--out` param:

```
k6 run --out influxdb=http://localhost:8086/k6 k6/scripts/{script_name}.js
```

You can access Grafana from this link: `http://localhost:3000`: use provided dashboard for tests results or explore
metrics.