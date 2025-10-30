import python.bookbank as bookbank

instance = bookbank.Bookbank(log=True) # Default: log=False
instance.connect("http://192.168.1.129",8080)
print(instance.version())
