from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium import webdriver

from webdriver_manager.chrome import ChromeDriverManager


class SeleniumService:
    def __init__(self):
        self.service = Service(ChromeDriverManager().install())
        self.options = Options()
        self.driver = webdriver.Chrome(service=self.service, options=self.options)