import requests


class Data:
    def __init__(self, id: int):
        self.url = f"http://localhost:8080/api/v1/task/obter/{id}"
        self.response = requests.get(self.url)

    def get_status(self):
        try:
            print(self.response.status_code)
        except Exception as e:
            print(f"Error: {e}")
    
    def get_data(self):
        return self.response.json()
            


p1 = Data(1)
p1.get_status()
print(p1.get_data())

p2 = Data(2)
p2.get_status()
print(p2.get_data())