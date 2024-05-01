import redis
import uuid 

pw = "8RMo89JBm6"
client = redis.StrictRedis(host='localhost', port=6380, password=pw)

with open("audio.wav", "rb") as f:
    file_bytes = f.read()
    id = str(uuid.uuid4())
    for i in range(2000):
        client.set("key-audio" + id + str(i), file_bytes, ex=3600)
        print(i)

print("Number of keys", client.dbsize())

