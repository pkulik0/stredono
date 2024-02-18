import sys
import os
import hashlib
import json

directory = sys.argv[1]

result = hashlib.sha256()

for root, dirs, files in os.walk(directory):
    for filename in files:
        with open(os.path.join(root, filename), 'rb') as file:
            while True:
                buf = file.read(4096)
                if not buf:
                    break
                result.update(buf)


print(json.dumps({'hash': result.hexdigest()}))