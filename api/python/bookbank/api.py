import requests
from .utils import is_valid_instance
from .exceptions import (
InvalidBookbankInstance,
NotConnectedError,
InvalidEntryId
)

class Bookbank(object):
    def __init__(self, log=False):
        self.log = log
        self.url = ""
        self.connected = False
    def connect(self,url: str, port: int) -> None:
        if not is_valid_instance(f"{url}:{port}"):
            raise InvalidBookbankInstance
        self.url = f"{url}:{port}"
        self.connected = True
        if not self.log:
            return
        print(f"Successfully connected to bookbank instance {url}:{port}")
    def version(self) -> str:
        if not self.connected:
            raise NotConnectedError
        req = requests.get(f"{self.url}/api/version")
        return req.text
    # def library(self) -> list:
    #     if not self.connected:
    #         raise NotConnectedError
    #     req = requests.get(f"{self.url}/api/books")
    def delete(self, entry: int) -> None:
        if not self.connected:
            raise NotConnectedError
        if entry < 1:
            raise InvalidEntryId
        requests.delete(f"{self.url}/api/delete/{entry}")