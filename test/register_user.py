import requests

url:str = "http://localhost:8080/api/user"

payload:dict = {
     "name":"Sonal",
     "email":"sonalvijit@handler.com"
}

response = requests.post(url=url, json=payload)

print("Status Code:", response.status_code)
print("Response JSON:", response.json())

assert response.status_code == 200
assert response.json()["name"] == "Sonal"
assert response.json()["email"] == "sonalvijit@handler.com"

print("Test passed")