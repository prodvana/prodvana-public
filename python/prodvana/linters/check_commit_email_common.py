VALID_DOMAINS = (
    "@prodvana.io",  # company email
    "@foundryinteractive.com",  # contractor email
    "dependabot[bot]@users.noreply.github.com",  # dependabot email
    "snyk-bot@snyk.io",  # snyk bot email
)


def check_email(email: str) -> None:
    assert email.endswith(
        VALID_DOMAINS
    ), f"Invalid commit email {email}, valid domains: {VALID_DOMAINS}"
