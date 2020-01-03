from common import Client, url

client = Client()
client.identify()

print client.spawn('gid1').text