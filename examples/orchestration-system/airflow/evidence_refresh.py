"""INTENTIONALLY FLAWED starter DAG — candidate review material, not production."""
from airflow import DAG
from airflow.operators.python import PythonOperator
from datetime import datetime
import requests


def refresh():
    # Faults: literal secret, no timeout, no idempotency key, Airflow owns the
    # same per-chemical retry loop that the Temporal workflow is meant to own.
    token = "tidx_replace_me"
    for attempt in range(10):
        requests.post(
            "https://platform.toxindex.com/v1/tools/toxjobs_predict",
            headers={"Authorization": f"Bearer {token}"},
            json={"tool": "admet-ai", "smiles": "CCO"},
        )


with DAG("evidence_refresh", start_date=datetime(2026, 1, 1), schedule="@daily", catchup=True) as dag:
    PythonOperator(task_id="refresh_everything", python_callable=refresh)

