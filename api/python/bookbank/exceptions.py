class InvalidBookbankInstance(Exception):
    """Exception raised when the url given does not hold a bookbank instance."""

    def __init__(self):
        super().__init__("The url given does not hold a bookbank instance.")
    def __str__(self):
        return "The url given does not hold a bookbank instance."

class NotConnectedError(Exception):
    """Exception raised when the Bookbank() class is not connected to a Bookbank instance."""

    def __init__(self):
        super().__init__("Not connected to a Bookbank instance.")

    def __str__(self):
        return "Not connected to a Bookbank instance."

class InvalidEntryId(Exception):
        """Exception raised when the id of the audiobook or podcast requested is invalid or cannot be found on this 
        Bookbank instance."""

        def __init__(self):
            super().__init__("The id of the audiobook or podcast requested is invalid or cannot be found on this Bookbank instance.")

        def __str__(self):
            return "The id of the audiobook or podcast requested is invalid or cannot be found on this Bookbank instance."