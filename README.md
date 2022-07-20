# AirFog 🌁

Common [AirFlow](https://airflow.apache.org) tasks.

## Setup

Create `airfog.json`:

```json
{
  "host": "airflow.k8s.foursquare.com",
  "scheme": "https",
  "username": "ispasic",
  "password": "..."
}
```

## Usage

### `clear-ti` Clear Task Instances

+ Clears Task instances that are _not_ in the `success` state for given `dagID` and latest run.
+ Optionally, `dagRunID` may be used to specify the specific run.
+ Parent DAGs of failed task instances are cleared, too.
+ After the cleanup, the DagRun is re-started, so all cleared tasks will run again.

[Airflow API](https://airflow.apache.org/docs/apache-airflow/stable/stable-rest-api-ref.html#operation/post_clear_task_instances)

```bash
airfog clear-ti --dag=shapes_biweekly_geocode_build
```
