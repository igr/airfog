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

## Run


+ Clear Task instances
```bash
airfog clear-ti --dag=shapes_biweekly_geocode_build
```
