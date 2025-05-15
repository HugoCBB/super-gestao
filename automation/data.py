import requests


class Data:
    def request(self, id: int):
        self.url = f"http://localhost:8080/api/v1/task/obter/{id}"
        self.response =  requests.get(self.url)

    def get_data(self):
        data = []
        for item in self.response.json().values():
            data.append({
                "id":item["id"],
                "cliente": item["cliente"],
                "mensagem": item["mensagem"],
                "numero": item["numero"],
                "status": item["status"]
            })
        return data
    
    def get_status(self):
        try:
            return {"status":self.response.status_code}
        except Exception as e:
            return {"error":e}
    
            

if __name__ == "__main__":
    p1 = Data(1)
    p1.get_status()
    print(p1.get_data())

    p2 = Data(2)
    p2.get_status()
    print(p2.get_data())