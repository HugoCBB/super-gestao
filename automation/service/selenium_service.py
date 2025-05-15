from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By

import sys
import os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..')))

from selenium import webdriver

from webdriver_manager.chrome import ChromeDriverManager

from automation.data import Data


class SeleniumService:
    def __init__(self, id: int):
        self.data = Data()
        self.data.request(id)
        
        self.response = self.data.get_data()

    def config(self):
        self.service = Service(ChromeDriverManager().install())
        self.options = Options()
        self.driver = webdriver.Chrome(service=self.service, options=self.options)
        self.wait = WebDriverWait(self.driver, timeout=30)
        
    def acessar_site(self):
        self.config()
        try:
            self.driver.get('https://web.whatsapp.com/')
            self.wait.until(EC.presence_of_element_located(By.XPATH , '//*[@id="app"]/div/div[2]/div[2]/div[2]/div/div[2]/div/div[2]/div[1]/div/canvas'))
            
            return {"status":"Site acessado com sucesso"}
        except Exception as e:
            return {"error":e}
        
    def enviar_mensagem(self):
        try:
            data = []
            for item in self.response:
                data.append(item)

            # self.driver.get(f"https://web.whatsapp.com/send?phone={data[0]["numero"]}")
            print(data[0]["numero"])
        except Exception as e:
            return {"error": str(e)}
                    
            
    def status_api(self):
        return self.data.get_status()
    



p1 = SeleniumService(2)
p1.enviar_mensagem()

