import requests


def is_valid_instance(url) -> bool:
    try:
        req = requests.get(url + "/api/version")
    except Exception:
        return False
    if req.text != "":
        return True
    else:
        return False