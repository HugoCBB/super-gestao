from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium import webdriver

from webdriver_manager.chrome import ChromeDriverManager

from ..data import Data

class SeleniumService(Data):
    def __init__(self):
        self.service = Service(ChromeDriverManager().install())
        self.options = Options()
        self.driver = webdriver.Chrome(service=self.service, options=self.options)

        self.response = Data.get_data(self)