import json
import os
import sys

directory = sys.argv[1]

publicFiles = []
for root, dirs, files in os.walk(directory):
    for file in files:
        if file.lower().endswith(".md"):
            continue
        if file.startswith("."):
            continue
        publicFiles += [os.path.join(root, file)]

filesDict = {str(i): f for i, f in enumerate(publicFiles)}

print(json.dumps(filesDict))
