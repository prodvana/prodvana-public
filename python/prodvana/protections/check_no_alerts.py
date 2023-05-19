#!/usr/bin/python

# This protection validates that there are no alerts firing on a given PD Service

import os
import sys

import pdpyras  # type: ignore


def check_for_alerts() -> None:
    try:
        api_key = os.environ["PD_API_KEY"]
        service_id = os.environ["PD_SERVICE_ID"]
        session = pdpyras.APISession(api_key)
        incidents = session.list_all("incidents", params={"statuses[]": ["triggered"]})
        incidents_firing = []

        for i in incidents:
            if i["service"]["id"] == service_id:
                incidents_firing.append(i)

        if len(incidents_firing) > 0:
            print("%s alert(s) firing for %s" % (len(incidents_firing), service_id))
            for i in incidents_firing:
                print("%s" % (i["summary"]))
            sys.exit(1)
    except Exception:
        print("Pagerduty protection failed for %s")
        sys.exit(1)

    print("No alerts firing for %s" % (service_id))


if __name__ == "__main__":
    check_for_alerts()
