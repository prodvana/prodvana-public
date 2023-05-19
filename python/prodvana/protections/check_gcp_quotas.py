#!/usr/bin/python

# This protection validates the provided GCP Service Quotas have headroom.

# Inputs:
# GCP_CREDENTIALS : Contents of a GCP Service Account JSON Keyfile. the Service account needs Cloud Monitoring Viewer permissions.
# GCP_PROJECT     : The GCP Project to check quotas on.
# GCP_REGION      : (Optional) Specific Region to check quotas in (or "global" for only non-regional quotas).
# QUOTA_THRESHOLD : (Optional) 0.0 - 1.0, minimum headroom required. Defaults to 0.05.

# TODO(mike): config file option for more specific / sophisticated options


import json
import os
import sys

from google.cloud import monitoring_v3
from google.oauth2 import service_account  # type: ignore

DEFAULT_QUOTA_THRESHOLD = 0.05

REGIONAL_QUOTA_USAGE_QUERY_MQL = """
fetch consumer_quota
| {{ metric 'serviceruntime.googleapis.com/quota/allocation/usage'
| filter resource.project_id == '{project}' && (resource.location == '{location}')
| align next_older(1d)
| every 1d
| group_by [metric.quota_metric], [value_usage_aggregate: aggregate(value.usage)]
; metric 'serviceruntime.googleapis.com/quota/limit'
| filter resource.project_id == '{project}' && (resource.location == '{location}')
| align next_older(1d)
| every 1d 
| group_by [metric.quota_metric], [value_usage_aggregate: aggregate(value.limit)]
}}
| ratio
"""

# needs to be grouped on location to make sense
ALL_QUOTA_USAGE_QUERY_MQL = """
fetch consumer_quota
| {{ metric 'serviceruntime.googleapis.com/quota/allocation/usage'
| filter resource.project_id == '{project}'
| align next_older(1d)
| every 1d
| group_by [metric.quota_metric, resource.location], [value_usage_aggregate: aggregate(value.usage)]
; metric 'serviceruntime.googleapis.com/quota/limit'
| filter resource.project_id == '{project}'
| align next_older(1d)
| every 1d 
| group_by [metric.quota_metric, resource.location], [value_usage_aggregate: aggregate(value.limit)]
}}
| ratio
"""


def check_gcp_quotas() -> None:
    project = os.environ["GCP_PROJECT"]
    quota_threshold = float(os.environ.get("QUOTA_THRESHOLD", DEFAULT_QUOTA_THRESHOLD))

    try:
        json_acct_info = json.loads(os.environ["GCP_CREDENTIALS"])
        credentials = service_account.Credentials.from_service_account_info(
            json_acct_info
        )

        query_client = monitoring_v3.QueryServiceClient(credentials=credentials)

        failures = []
        region = os.environ.get("GCP_REGION", None)
        if region:
            query = REGIONAL_QUOTA_USAGE_QUERY_MQL.format(
                project=project, location=region
            )
            resp = query_client.query_time_series(
                request=monitoring_v3.QueryTimeSeriesRequest(
                    name="projects/" + project,
                    query=query,
                )
            )
            for ts_data in resp.time_series_data:
                headroom = 1.0 - ts_data.point_data[-1].values[-1].double_value
                if headroom < quota_threshold:
                    failures.append(
                        f"{ts_data.label_values[0].string_value} below threshold ({quota_threshold}): {headroom}"
                    )
        else:
            query = ALL_QUOTA_USAGE_QUERY_MQL.format(project=project)
            resp = query_client.query_time_series(
                request=monitoring_v3.QueryTimeSeriesRequest(
                    name="projects/" + project,
                    query=query,
                )
            )
            for ts_data in resp.time_series_data:
                headroom = 1.0 - ts_data.point_data[-1].values[-1].double_value
                if headroom < quota_threshold:
                    quota_name = ts_data.label_values[0].string_value
                    quota_region = ts_data.label_values[1].string_value
                    failures.append(
                        f"{quota_name} ({quota_region}) below threshold ({quota_threshold}): {headroom}"
                    )

        if len(failures) > 0:
            print("\n".join(failures))
            sys.exit(1)

    except Exception as e:
        raise e
        print(f"Check GCP Quotas protection failed: {e}")
        sys.exit(1)

    print("All quotas are above headroom threshold.")


if __name__ == "__main__":
    check_gcp_quotas()
