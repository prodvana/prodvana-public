#!/usr/bin/python

# This protection validates the provided AWS Quotas have headroom.

# Inputs:
# AWS_ACCESS_KEY_ID: Key ID associated with an IAM User with the ServiceQuotasReadOnlyAccess policy attached.
# AWS_SECRET_ACCESS_KEY: Access Key associated with the above Key ID.
# AWS_REGION: AWS Region to check quotas in.
# QUOTA_THRESHOLD : (Optional) 0.0 - 1.0, minimum headroom required. Defaults to 0.05.

# TODO(mike): config file option for more specific / sophisticated options

import datetime
import os
import sys
from dataclasses import dataclass
from typing import List

from boto3.session import Session
from mypy_boto3_cloudwatch.client import CloudWatchClient
from mypy_boto3_service_quotas.client import ServiceQuotasClient
from mypy_boto3_service_quotas.type_defs import (
    ListServiceQuotasResponseTypeDef,
    ListServicesResponseTypeDef,
    MetricInfoTypeDef,
)

DEFAULT_QUOTA_THRESHOLD = 0.05

SERVICE_QUOTA_CODE = "servicequotas"
CLOUDWATCH_CODE = "monitoring"


@dataclass
class Service:
    name: str
    code: str


@dataclass
class ServiceQuota:
    service: Service
    name: str
    value: float
    usage_metric: MetricInfoTypeDef


def get_latest_quota_usage(client: CloudWatchClient, quota: ServiceQuota) -> float:
    response = client.get_metric_data(
        MetricDataQueries=[
            {
                "Id": "q1",
                "MetricStat": {
                    "Metric": {
                        "Namespace": quota.usage_metric["MetricNamespace"],
                        "MetricName": quota.usage_metric["MetricName"],
                        "Dimensions": [
                            {"Name": name, "Value": val}
                            for name, val in quota.usage_metric[
                                "MetricDimensions"
                            ].items()
                        ],
                    },
                    "Period": 3600,
                    "Stat": quota.usage_metric["MetricStatisticRecommendation"],
                },
                "ReturnData": True,
            },
        ],
        StartTime=datetime.datetime.utcnow() - datetime.timedelta(hours=1),
        EndTime=datetime.datetime.utcnow(),
    )

    # Extract the latest metric data from the response
    result = response["MetricDataResults"][0]
    if len(result["Values"]) > 0:
        usage = result["Values"][-1]
    else:
        usage = 0

    return usage


def get_all_services_with_quotas(client: ServiceQuotasClient) -> List[Service]:
    def handle_response(
        services: List[Service], response: ListServicesResponseTypeDef
    ) -> None:
        for service in response["Services"]:
            services.append(
                Service(
                    name=service["ServiceName"],
                    code=service["ServiceCode"],
                )
            )

    services: List[Service] = []
    response = client.list_services()
    handle_response(services, response)

    while "NextToken" in response:
        response = client.list_services(NextToken=response["NextToken"])
        handle_response(services, response)

    return services


def get_all_service_quotas(
    client: ServiceQuotasClient,
    services: List[Service],
) -> List[ServiceQuota]:
    def handle_response(
        service_quotas: List[ServiceQuota], response: ListServiceQuotasResponseTypeDef
    ) -> None:
        # only collect quotas that have usage metrics
        for quota in [quota for quota in response["Quotas"] if "UsageMetric" in quota]:
            service_quotas.append(
                ServiceQuota(
                    service=svc,
                    name=quota["QuotaName"],
                    value=quota["Value"],
                    usage_metric=quota["UsageMetric"],
                )
            )

    service_quotas: List[ServiceQuota] = []
    for svc in services:
        response = client.list_service_quotas(ServiceCode=svc.code)
        handle_response(service_quotas, response)

        while "NextToken" in response:
            response = client.list_service_quotas(
                ServiceCode=svc.code, NextToken=response["NextToken"]
            )
            handle_response(service_quotas, response)

    return service_quotas


def check_aws_quotas() -> None:
    region = os.environ["AWS_REGION"]
    quota_threshold = float(os.environ.get("QUOTA_THRESHOLD", DEFAULT_QUOTA_THRESHOLD))

    try:
        quota_client = Session().client("service-quotas", region_name=region)
        cw_client = Session().client("cloudwatch", region_name=region)

        services = get_all_services_with_quotas(quota_client)

        # AWS included API rate limits / throttling in the service quotas.
        # when runnining to check, the list_services, list_service_quotas, and get_metric_data
        # calls are enough to result in throttling.. but it's not a hard error or something
        # you'd consider a failure.
        services = [
            svc
            for svc in services
            if svc.code != SERVICE_QUOTA_CODE and svc.code != CLOUDWATCH_CODE
        ]

        quotas = get_all_service_quotas(quota_client, services)

        failures = []
        for quota in [q for q in quotas if q.value > 0]:
            val = get_latest_quota_usage(cw_client, quota)
            headroom = 1.0 - min((val / quota.value), 1)
            if headroom < quota_threshold:
                failures.append(
                    f"{quota.service.name}/{quota.name}  below threshold ({quota_threshold}): {headroom}"
                )

        if len(failures) > 0:
            print("\n".join(failures))
            sys.exit(1)

    except Exception as e:
        raise e
        print(f"Check AWS Quotas protection failed: {e}")
        sys.exit(1)

    print("All quotas are above headroom threshold.")


if __name__ == "__main__":
    check_aws_quotas()
