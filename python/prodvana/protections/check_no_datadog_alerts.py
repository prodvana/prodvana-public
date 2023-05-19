import os
import sys

import datadog  # type: ignore


def main() -> None:
    list_of_tags = []
    TAGS = os.getenv("DATADOG_TAG_LIST")
    if TAGS:
        list_of_tags = TAGS.split(" ")
    datadog.initialize(
        api_key=os.getenv("DATADOG_API_KEY"),
        app_key=os.getenv("DATADOG_APP_KEY"),
    )
    print("List of tags checked: %s " % " ".join(list_of_tags))
    monitors = datadog.api.Monitor.get_all(monitor_tags=list_of_tags)
    found_alert = False
    for monitor in monitors:
        if monitor["overall_state"] != "OK" and "subtest" not in monitor["tags"]:
            print(
                "Not OK: %s - https://app.datadoghq.com/monitors/%s"
                % (monitor["name"], monitor["id"])
            )
            found_alert = True
    if found_alert:
        sys.exit(1)


if __name__ == "__main__":
    main()
