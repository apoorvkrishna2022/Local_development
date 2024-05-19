
import json

file_path = 'temp2.json'

with open(file_path, 'r') as file:
    # Load the JSON data from the file
    json_data = json.load(file)

# Now you can work with the JSON data as a Python dictionary
test = ""
for i in json_data:
    test = test + ",\'" + str(i['log_context_req_id']) + "\'"

print(test)

